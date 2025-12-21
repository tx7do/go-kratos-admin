package data

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/kratos-bootstrap/bootstrap"

	authzEngine "github.com/tx7do/kratos-authz/engine"
	"github.com/tx7do/kratos-authz/engine/casbin"
	"github.com/tx7do/kratos-authz/engine/noop"
	"github.com/tx7do/kratos-authz/engine/opa"

	pagination "github.com/tx7do/go-crud/api/gen/go/pagination/v1"
	"github.com/tx7do/go-utils/trans"
	conf "github.com/tx7do/kratos-bootstrap/api/gen/go/conf/v1"

	"go-wind-admin/app/admin/service/cmd/server/assets"

	adminV1 "go-wind-admin/api/gen/go/admin/service/v1"
	userV1 "go-wind-admin/api/gen/go/user/service/v1"
)

// Authorizer 权限管理器
type Authorizer struct {
	log *log.Helper

	roleRepo        *RoleRepo
	apiResourceRepo *ApiResourceRepo

	engine authzEngine.Engine
}

func NewAuthorizer(
	ctx *bootstrap.Context,
	roleRepo *RoleRepo,
	apiResourceRepo *ApiResourceRepo,
) *Authorizer {
	a := &Authorizer{
		log:             log.NewHelper(log.With(ctx.Logger, "module", "authorizer/repo/admin-service")),
		roleRepo:        roleRepo,
		apiResourceRepo: apiResourceRepo,
	}

	a.init(ctx.Config)

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
		state, err := casbin.NewEngine(ctx, casbin.WithStringModel(string(assets.OpaRbacRego)))
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
	//a.log.Debugf("Generating policies for engine: %s", a.engine.Name())

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

	case "noop":
		return nil

	default:
		err = fmt.Errorf("unknown engine name: %s", a.engine.Name())
		a.log.Warnf(err.Error())
		return err
	}

	//a.log.Debugf("***************** policy rules len: %v", len(policies))

	if err = a.engine.SetPolicies(ctx, policies, nil); err != nil {
		a.log.Errorf("set policies error: %v", err)
		return err
	}

	a.log.Infof("reloaded policy rules [%d] successfully for engine: %s", len(policies), a.engine.Name())

	return nil
}

// generateCasbinPolicies 生成 Casbin 策略
func (a *Authorizer) generateCasbinPolicies(roles *userV1.ListRoleResponse, apis *adminV1.ListApiResourceResponse) (authzEngine.PolicyMap, error) {
	var rules []casbin.PolicyRule

	domain := "*"

	for _, role := range roles.Items {
		if role.GetId() == 0 {
			continue // Skip if role or API ID is not set
		}

		apiSet := make(map[uint32]struct{}) // API set for current role

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

// generateOpaPolicies 生成 OPA 策略
func (a *Authorizer) generateOpaPolicies(roles *userV1.ListRoleResponse, apis *adminV1.ListApiResourceResponse) (authzEngine.PolicyMap, error) {
	type OpaPolicyPath struct {
		Pattern string `json:"pattern"`
		Method  string `json:"method"`
	}

	policies := make(authzEngine.PolicyMap, len(roles.Items))

	for _, role := range roles.Items {
		if role.GetId() == 0 {
			continue // Skip if role or API ID is not set
		}

		paths := make([]OpaPolicyPath, 0, len(apis.Items))
		apiSet := make(map[uint32]struct{}) // API set for current role

		//a.log.Debugf("Processing Role: [%s] with APIs: [%v]", role.GetCode(), role.GetApis())

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

				//a.log.Debugf("OPA Policy - Role: [%s], Path: [%s], Method: [%s]", role.GetCode(), api.GetPath(), api.GetMethod())
			}
		}

		policies[role.GetCode()] = paths
	}

	return policies, nil
}
