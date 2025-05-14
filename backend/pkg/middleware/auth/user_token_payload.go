package auth

import (
	authn "github.com/tx7do/kratos-authn/engine"
	userV1 "kratos-admin/api/gen/go/user/service/v1"
)

const (
	ClaimFieldUserID    = "uid"
	ClaimFieldTenantID  = "tid"
	ClaimFieldClientID  = "cid"
	ClaimFieldDeviceID  = "did"
	ClaimFieldAuthority = "aut"
)

// UserTokenPayload 用户JWT令牌载荷
type UserTokenPayload struct {
	UserId    uint32
	TenantId  uint32
	UserName  string
	ClientId  string
	Authority string
}

// NewUserTokenPayload 创建用户令牌
func NewUserTokenPayload(tenantId uint32, userId uint32, userName string, authority userV1.UserAuthority, clientId string) *UserTokenPayload {
	return &UserTokenPayload{
		UserId:    userId,
		TenantId:  tenantId,
		UserName:  userName,
		ClientId:  clientId,
		Authority: authority.String(),
	}
}

func NewUserTokenPayloadWithClaims(claims *authn.AuthClaims) (*UserTokenPayload, error) {
	payload := &UserTokenPayload{}

	if err := payload.ExtractAuthClaims(claims); err != nil {
		return nil, err
	}

	return payload, nil
}

// MakeAuthClaims 构建认证声明
func (t *UserTokenPayload) MakeAuthClaims() *authn.AuthClaims {
	return &authn.AuthClaims{
		authn.ClaimFieldSubject: t.UserName,
		ClaimFieldUserID:        t.UserId,
		ClaimFieldTenantID:      t.TenantId,
		ClaimFieldClientID:      t.ClientId,
		ClaimFieldAuthority:     t.Authority,
	}
}

// ExtractAuthClaims 解析认证声明
func (t *UserTokenPayload) ExtractAuthClaims(claims *authn.AuthClaims) error {
	sub, _ := claims.GetSubject()
	if sub != "" {
		t.UserName = sub
	}

	userId, _ := claims.GetUint32(ClaimFieldUserID)
	if userId != 0 {
		t.UserId = userId
	}

	tenantId, _ := claims.GetUint32(ClaimFieldTenantID)
	if userId != 0 {
		t.TenantId = tenantId
	}

	clientId, _ := claims.GetString(ClaimFieldClientID)
	if clientId != "" {
		t.ClientId = clientId
	}

	return nil
}

func (t *UserTokenPayload) GetAuthority() userV1.UserAuthority {
	authority, ok := userV1.UserAuthority_value[t.Authority]
	if !ok {
		return userV1.UserAuthority_GUEST_USER
	}
	return userV1.UserAuthority(authority)
}
