package jwt

import (
	"errors"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang-jwt/jwt/v5"
	"github.com/tx7do/go-utils/trans"

	authn "github.com/tx7do/kratos-authn/engine"

	authenticationV1 "kratos-admin/api/gen/go/authentication/service/v1"
	userV1 "kratos-admin/api/gen/go/user/service/v1"
)

const (
	ClaimFieldUserName  = authn.ClaimFieldSubject // 用户名
	ClaimFieldUserID    = "uid"                   // 用户ID
	ClaimFieldTenantID  = "tid"                   // 租户ID
	ClaimFieldClientID  = "cid"                   // 客户端ID
	ClaimFieldDeviceID  = "did"                   // 设备ID
	ClaimFieldAuthority = "aut"                   // 用户权限
	//ClaimFieldRoleIds   = "rid"                   // 角色ID列表
	ClaimFieldRoleCodes = "roc" // 角色码列表
)

// NewUserTokenPayload 创建用户令牌
func NewUserTokenPayload(user *userV1.User, clientId string) *authenticationV1.UserTokenPayload {
	return &authenticationV1.UserTokenPayload{
		UserId:    user.GetId(),
		TenantId:  user.TenantId,
		Username:  user.Username,
		ClientId:  trans.Ptr(clientId),
		Authority: user.GetAuthority(),
		Roles:     user.Roles,
	}
}

func NewUserTokenAuthClaims(user *userV1.User, clientId string) *authn.AuthClaims {
	return &authn.AuthClaims{
		ClaimFieldUserName:  user.GetUsername(),
		ClaimFieldUserID:    user.GetId(),
		ClaimFieldTenantID:  user.GetTenantId(),
		ClaimFieldClientID:  clientId,
		ClaimFieldAuthority: user.Authority.String(),
		ClaimFieldRoleCodes: user.Roles,
	}
}

func NewUserTokenPayloadWithClaims(claims *authn.AuthClaims) (*authenticationV1.UserTokenPayload, error) {
	payload := &authenticationV1.UserTokenPayload{}

	sub, err := claims.GetSubject()
	if err != nil {
		log.Errorf("GetSubject failed: %v", err)
	}
	if sub != "" {
		payload.Username = trans.Ptr(sub)
	}

	userId, err := claims.GetUint32(ClaimFieldUserID)
	if err != nil {
		log.Errorf("GetUint32 ClaimFieldUserID failed: %v", err)
	}
	if userId != 0 {
		payload.UserId = userId
	}

	tenantId, err := claims.GetUint32(ClaimFieldTenantID)
	if err != nil {
		log.Errorf("GetUint32 ClaimFieldTenantID failed: %v", err)
	}
	if tenantId != 0 {
		payload.TenantId = trans.Ptr(tenantId)
	}

	clientId, err := claims.GetString(ClaimFieldClientID)
	if err != nil {
		log.Errorf("GetString ClaimFieldClientID failed: %v", err)
	}
	if clientId != "" {
		payload.ClientId = trans.Ptr(clientId)
	}

	authority, err := claims.GetString(ClaimFieldAuthority)
	if err != nil {
		log.Errorf("GetString ClaimFieldAuthority failed: %v", err)
	}
	if authority != "" {
		v, ok := userV1.UserAuthority_value[authority]
		if ok {
			payload.Authority = userV1.UserAuthority(v)
		}
	}

	roleCodes, err := claims.GetStrings(ClaimFieldRoleCodes)
	if err != nil {
		log.Errorf("GetStrings ClaimFieldRoleCodes failed: %v", err)
	}
	if roleCodes != nil {
		payload.Roles = roleCodes
	}

	return payload, nil
}

func NewUserTokenPayloadWithJwtMapClaims(claims jwt.MapClaims) (*authenticationV1.UserTokenPayload, error) {
	payload := &authenticationV1.UserTokenPayload{}

	sub, err := claims.GetSubject()
	if err != nil {
		log.Errorf("GetSubject failed: %v", err)
	}
	if sub != "" {
		payload.Username = trans.Ptr(sub)
	}

	userId, _ := claims[ClaimFieldUserID]
	if userId != nil {
		payload.UserId = uint32(userId.(float64))
	}

	tenantId, _ := claims[ClaimFieldTenantID]
	if tenantId != nil {
		payload.TenantId = trans.Ptr(uint32(tenantId.(float64)))
	}

	clientId, _ := claims[ClaimFieldClientID]
	if clientId != nil {
		payload.ClientId = trans.Ptr(clientId.(string))
	}

	authority, _ := claims[ClaimFieldAuthority]
	if authority != nil {
		v, ok := userV1.UserAuthority_value[authority.(string)]
		if ok {
			payload.Authority = userV1.UserAuthority(v)
		}
	}

	roleCodes, _ := claims[ClaimFieldRoleCodes]
	if roleCodes != nil {
		switch itf := roleCodes.(type) {
		case []interface{}:
			for _, rc := range itf {
				payload.Roles = append(payload.Roles, rc.(string))
			}

		case []string:
			payload.Roles = itf

		default:
			return nil, errors.New("invalid roleCodes type")
		}
	}

	return payload, nil
}
