package auth

import (
	authn "github.com/tx7do/kratos-authn/engine"
)

const (
	ClaimFieldUserID   = "uid"
	ClaimFieldTenantID = "tid"
	ClaimFieldClientID = "cid"
	ClaimFieldDeviceID = "did"
)

// UserTokenPayload 用户JWT令牌载荷
type UserTokenPayload struct {
	UserId   uint32
	TenantId uint32
	UserName string
	ClientId string
}

// NewUserTokenPayload 创建用户令牌
func NewUserTokenPayload(tenantId uint32, userId uint32, userName string, clientId string) *UserTokenPayload {
	return &UserTokenPayload{
		UserId:   userId,
		TenantId: tenantId,
		UserName: userName,
		ClientId: clientId,
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
