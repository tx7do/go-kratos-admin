package auth

import (
	"context"
	"reflect"
	"strconv"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"

	authnEngine "github.com/tx7do/kratos-authn/engine"
	authn "github.com/tx7do/kratos-authn/middleware"

	authzEngine "github.com/tx7do/kratos-authz/engine"
	authz "github.com/tx7do/kratos-authz/middleware"

	"github.com/tx7do/go-utils/stringutil"
	"github.com/tx7do/go-utils/trans"
	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"

	authenticationV1 "kratos-admin/api/gen/go/authentication/service/v1"

	"kratos-admin/pkg/entgo/viewer"
	"kratos-admin/pkg/jwt"
	"kratos-admin/pkg/metadata"
)

var defaultAction = authzEngine.Action("ANY")

// Server 衔接认证和权鉴
func Server(opts ...Option) middleware.Middleware {
	op := options{
		injectOperatorId: false,
		injectTenantId:   false,
		enableAuthz:      true,
		injectEnt:        true,
		injectMetadata:   true,
	}
	for _, o := range opts {
		o(&op)
	}

	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			tr, ok := transport.FromServerContext(ctx)
			if !ok {
				return nil, ErrWrongContext
			}

			authnClaims, ok := authn.FromContext(ctx)
			if !ok {
				return nil, ErrWrongContext
			}

			tokenPayload, err := jwt.NewUserTokenPayloadWithClaims(authnClaims)
			if err != nil {
				return nil, ErrExtractUserInfoFailed
			}

			// 校验访问令牌是否存在
			if op.isExistAccessToken != nil {
				if !op.isExistAccessToken(ctx, tokenPayload.UserId) {
					return nil, ErrAccessTokenExpired
				}
			}

			if op.injectOperatorId {
				if err = setRequestOperationId(req, tokenPayload); err != nil {
					return nil, err
				}
			}
			if op.injectTenantId {
				if err = setRequestTenantId(req, tokenPayload); err != nil {
					return nil, err
				}

				if err = ensurePagingRequestTenantId(req, tokenPayload); err != nil {
					return nil, err
				}
			}

			if op.injectEnt {
				ctx = viewer.NewContext(ctx, viewer.UserViewer{
					Role:     tokenPayload.GetAuthority(),
					TenantId: tokenPayload.TenantId,
				})
			}

			if op.injectMetadata {
				ctx = metadata.NewOperatorMetadataContext(ctx,
					trans.Ptr(tokenPayload.UserId),
					tokenPayload.TenantId,
					trans.Ptr(tokenPayload.GetAuthority()),
				)
			}

			if op.enableAuthz {
				ctx, err = processAuthz(ctx, tr, tokenPayload)
				if err != nil {
					return nil, err
				}
			}

			return handler(ctx, req)
		}
	}
}

func processAuthz(
	ctx context.Context,
	tr transport.Transporter,
	tokenPayload *authenticationV1.UserTokenPayload,
) (context.Context, error) {
	//var sub string
	//if sub, err = tokenPayload.GetSubject(); err != nil {
	//	return nil, ErrExtractSubjectFailed
	//}

	path := authzEngine.Resource(tr.Operation())
	action := defaultAction

	var htr *http.Transport
	var ok bool
	if htr, ok = tr.(*http.Transport); ok {
		path = authzEngine.Resource(htr.PathTemplate())
		action = authzEngine.Action(htr.Request().Method)
	}

	//log.Infof("**************** PATH[%s] ACTION[%s]", path, action)

	authzClaims := authzEngine.AuthClaims{
		//Subject:  (*authzEngine.Subject)(&sub),
		Subjects: trans.Ptr(tokenPayload.GetRoles()),
		Action:   trans.Ptr(action),
		Resource: trans.Ptr(path),
	}

	ctx = authz.NewContext(ctx, &authzClaims)

	return ctx, nil
}

func FromContext(ctx context.Context) (*authenticationV1.UserTokenPayload, error) {
	claims, ok := authnEngine.AuthClaimsFromContext(ctx)
	if !ok {
		return nil, ErrMissingJwtToken
	}

	return jwt.NewUserTokenPayloadWithClaims(claims)
}

func setRequestOperationId(req interface{}, payload *authenticationV1.UserTokenPayload) error {
	if req == nil {
		return ErrInvalidRequest
	}

	v := reflect.ValueOf(req).Elem()
	field := v.FieldByName("OperatorId")
	if field.IsValid() && field.Kind() == reflect.Ptr {
		field.Set(reflect.ValueOf(&payload.UserId))
	}

	return nil
}

func setRequestTenantId(req interface{}, payload *authenticationV1.UserTokenPayload) error {
	if req == nil {
		return ErrInvalidRequest
	}

	v := reflect.ValueOf(req).Elem()
	field := v.FieldByName("TenantId")
	if field.IsValid() && field.Kind() == reflect.Ptr {
		field.Set(reflect.ValueOf(&payload.TenantId))
	}

	return nil
}

func ensurePagingRequestTenantId(req interface{}, payload *authenticationV1.UserTokenPayload) error {
	if paging, ok := req.(*pagination.PagingRequest); ok && payload.GetTenantId() > 0 {
		if paging.Query != nil {
			newStr := stringutil.ReplaceJSONField("tenantId|tenant_id", strconv.Itoa(int(payload.GetTenantId())), paging.GetQuery())
			paging.Query = trans.Ptr(newStr)
		}
		if paging.OrQuery != nil {
			newStr := stringutil.ReplaceJSONField("tenantId|tenant_id", strconv.Itoa(int(payload.GetTenantId())), paging.GetOrQuery())
			paging.OrQuery = trans.Ptr(newStr)
		}
	}
	return nil
}
