package jwt

import (
	"github.com/golang-jwt/jwt/v5"

	authn "github.com/tx7do/kratos-authn/engine"

	authenticationV1 "kratos-admin/api/gen/go/authentication/service/v1"
	userV1 "kratos-admin/api/gen/go/user/service/v1"
)

const (
	ClaimFieldUserID    = "uid"
	ClaimFieldTenantID  = "tid"
	ClaimFieldClientID  = "cid"
	ClaimFieldDeviceID  = "did"
	ClaimFieldAuthority = "aut"
)

// NewUserTokenPayload 创建用户令牌
func NewUserTokenPayload(user *userV1.User, clientId string) *authenticationV1.UserTokenPayload {
	return &authenticationV1.UserTokenPayload{
		UserId:    user.GetId(),
		TenantId:  user.GetTenantId(),
		Username:  user.GetUsername(),
		ClientId:  clientId,
		Authority: user.GetAuthority(),
	}
}

func NewUserTokenAuthClaims(user *userV1.User, clientId string) *authn.AuthClaims {
	return &authn.AuthClaims{
		authn.ClaimFieldSubject: user.GetUsername(),
		ClaimFieldUserID:        user.GetId(),
		ClaimFieldTenantID:      user.GetTenantId(),
		ClaimFieldClientID:      clientId,
		ClaimFieldAuthority:     user.Authority.String(),
	}
}

func NewUserTokenPayloadWithClaims(claims *authn.AuthClaims) (*authenticationV1.UserTokenPayload, error) {
	payload := &authenticationV1.UserTokenPayload{}

	sub, _ := claims.GetSubject()
	if sub != "" {
		payload.Username = sub
	}

	userId, _ := claims.GetUint32(ClaimFieldUserID)
	if userId != 0 {
		payload.UserId = userId
	}

	tenantId, _ := claims.GetUint32(ClaimFieldTenantID)
	if userId != 0 {
		payload.TenantId = tenantId
	}

	clientId, _ := claims.GetString(ClaimFieldClientID)
	if clientId != "" {
		payload.ClientId = clientId
	}

	authority, _ := claims.GetString(ClaimFieldAuthority)
	if authority != "" {
		v, ok := userV1.UserAuthority_value[authority]
		if ok {
			payload.Authority = userV1.UserAuthority(v)
		}
	}

	return payload, nil
}

func NewUserTokenPayloadWithJwtMapClaims(claims jwt.MapClaims) (*authenticationV1.UserTokenPayload, error) {
	payload := &authenticationV1.UserTokenPayload{}

	sub, _ := claims.GetSubject()
	if sub != "" {
		payload.Username = sub
	}

	userId, _ := claims[ClaimFieldUserID]
	if userId != nil {
		payload.UserId = uint32(userId.(float64))
	}

	tenantId, _ := claims[ClaimFieldTenantID]
	if userId != nil {
		payload.TenantId = uint32(tenantId.(float64))
	}

	clientId, _ := claims[ClaimFieldClientID]
	if clientId != nil {
		payload.ClientId = clientId.(string)
	}

	authority, _ := claims[ClaimFieldAuthority]
	if authority != nil {
		v, ok := userV1.UserAuthority_value[authority.(string)]
		if ok {
			payload.Authority = userV1.UserAuthority(v)
		}
	}

	return payload, nil
}
