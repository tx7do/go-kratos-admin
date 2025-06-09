package data

import (
	"context"
	"errors"
	"kratos-admin/app/admin/service/cmd/server/assets"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/go-utils/trans"

	authzEngine "github.com/tx7do/kratos-authz/engine"
	"github.com/tx7do/kratos-authz/engine/casbin"
	"github.com/tx7do/kratos-authz/engine/noop"
	"github.com/tx7do/kratos-authz/engine/opa"

	conf "github.com/tx7do/kratos-bootstrap/api/gen/go/conf/v1"
	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"

	adminV1 "kratos-admin/api/gen/go/admin/service/v1"
	userV1 "kratos-admin/api/gen/go/user/service/v1"
)

type Authorizer struct {
	log *log.Helper

	roleRepo        *RoleRepo
	apiResourceRepo *ApiResourceRepo

	engine authzEngine.Engine
}

func NewAuthorizer(
	logger log.Logger,
	cfg *conf.Bootstrap,
	roleRepo *RoleRepo,
	apiResourceRepo *ApiResourceRepo,
) *Authorizer {
	a := &Authorizer{
		log:             log.NewHelper(log.With(logger, "module", "authorizer/repo/admin-service")),
		roleRepo:        roleRepo,
		apiResourceRepo: apiResourceRepo,
	}

	a.init(cfg)

	return a
}

func (a *Authorizer) init(cfg *conf.Bootstrap) {
	a.engine = a.newEngine(cfg)

	if err := a.ResetPolicies(context.Background()); err != nil {
		a.log.Errorf("reset policies error: %v", err)
	}
}

func (a *Authorizer) newEngine(cfg *conf.Bootstrap) authzEngine.Engine {
	if cfg.Authz == nil {
		return nil
	}

	ctx := context.Background()

	switch cfg.GetAuthz().GetType() {
	default:
		fallthrough
	case "noop":
		state, err := noop.NewEngine(ctx)
		if err != nil {
			a.log.Errorf("new noop engine error: %v", err)
			return nil
		}
		return state

	case "casbin":
		state, err := casbin.NewEngine(ctx)
		if err != nil {
			a.log.Errorf("init casbin engine error: %v", err)
			return nil
		}
		return state

	case "opa":
		state, err := opa.NewEngine(ctx,
			opa.WithModulesFromString(map[string]string{
				"rbac.rego": string(assets.OpaRbacRego),
			}),
		)
		if err != nil {
			a.log.Errorf("init opa engine error: %v", err)
			return nil
		}

		if err = state.InitModulesFromString(map[string]string{
			"rbac.rego": string(assets.OpaRbacRego),
		}); err != nil {
			a.log.Errorf("init opa modules error: %v", err)
		}

		return state

		//case "zanzibar":
		//	state, err := zanzibar.NewEngine(ctx)
		//	if err != nil {
		//		return nil
		//	}
		//	return state
	}
}

func (a *Authorizer) Engine() authzEngine.Engine {
	return a.engine
}

// ResetPolicies 重置策略
func (a *Authorizer) ResetPolicies(ctx context.Context) error {
	//a.log.Info("*******************reset policies")

	roles, err := a.roleRepo.List(ctx, &pagination.PagingRequest{NoPaging: trans.Ptr(true)})
	if err != nil {
		a.log.Errorf("failed to list roles: %v", err)
		return err
	}

	if roles == nil || len(roles.Items) < 1 {
		a.log.Warnf("no roles found to set policies")
		return nil // No roles to set policies
	}

	apis, err := a.apiResourceRepo.List(ctx, &pagination.PagingRequest{NoPaging: trans.Ptr(true)})
	if err != nil {
		a.log.Errorf("failed to list APIs: %v", err)
		return err
	}

	if apis == nil || len(apis.Items) < 1 {
		a.log.Warnf("no APIs found to set policies for roles")
		return nil // No APIs to set policies
	}

	//a.log.Debugf("roles [%d] apis [%d]", len(roles.Items), len(apis.Items))

	var policies authzEngine.PolicyMap

	switch a.engine.Name() {
	case "casbin":
		if policies, err = a.generateCasbinPolicies(roles, apis); err != nil {
			a.log.Errorf("generate casbin policies error: %v", err)
			return err
		}

	case "opa":
		if policies, err = a.generateOpaPolicies(roles, apis); err != nil {
			a.log.Errorf("generate OPA policies error: %v", err)
			return err
		}

	default:
		a.log.Warnf("unknown engine name: %s", a.engine.Name())
		return errors.New("unknown authz engine name")
	}

	//a.log.Debugf("***************** policy rules len: %v", len(rules))

	if err = a.engine.SetPolicies(context.Background(), policies, nil); err != nil {
		a.log.Errorf("set policies error: %v", err)
		return err
	}

	return nil
}

func (a *Authorizer) generateCasbinPolicies(roles *userV1.ListRoleResponse, apis *adminV1.ListApiResourceResponse) (authzEngine.PolicyMap, error) {
	var rules []casbin.PolicyRule
	apiSet := make(map[uint32]struct{})

	domain := "*"

	for _, role := range roles.Items {
		if role.GetId() == 0 {
			continue // Skip if role or API ID is not set
		}

		for _, apiId := range role.GetApis() {
			apiSet[apiId] = struct{}{}
		}

		for _, api := range apis.Items {
			if api.GetId() == 0 {
				continue // Skip if role or API ID is not set
			}

			if _, exists := apiSet[api.GetId()]; exists {
				rules = append(rules, casbin.PolicyRule{
					PType: "p",
					V0:    role.GetCode(),
					V1:    api.GetPath(),
					V2:    api.GetMethod(),
					V3:    domain,
				})
			}
		}
	}

	policies := authzEngine.PolicyMap{
		"policies": rules,
		"projects": authzEngine.MakeProjects(),
	}

	return policies, nil
}

func (a *Authorizer) generateOpaPolicies(roles *userV1.ListRoleResponse, apis *adminV1.ListApiResourceResponse) (authzEngine.PolicyMap, error) {
	type OpaPolicyPath struct {
		Pattern string `json:"pattern"`
		Method  string `json:"method"`
	}

	policies := make(authzEngine.PolicyMap, len(roles.Items))
	paths := make([]OpaPolicyPath, 0, len(roles.Items)*len(apis.Items))

	//policies["projects"] = authzEngine.MakeProjects("api")

	apiSet := make(map[uint32]struct{})

	for _, role := range roles.Items {
		if role.GetId() == 0 {
			continue // Skip if role or API ID is not set
		}

		paths = paths[:0] // Reset paths for each role

		for _, apiId := range role.GetApis() {
			apiSet[apiId] = struct{}{}
		}

		for _, api := range apis.Items {
			if api.GetId() == 0 {
				continue // Skip if role or API ID is not set
			}

			if _, exists := apiSet[api.GetId()]; exists {
				paths = append(paths, OpaPolicyPath{
					Pattern: api.GetPath(),
					Method:  api.GetMethod(),
				})
			}
		}

		policies[role.GetCode()] = paths
	}

	return policies, nil
}
