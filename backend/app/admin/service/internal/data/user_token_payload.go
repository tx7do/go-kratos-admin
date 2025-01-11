package data

import (
	authn "github.com/tx7do/kratos-authn/engine"
)

const (
	ClaimFieldUserID   = "uid"
	ClaimFieldClientID = "cid"
)

// UserTokenPayload 用户JWT令牌载荷
type UserTokenPayload struct {
	UserId   uint32
	UserName string
	ClientId string
}

// NewUserTokenPayload 创建用户令牌
func NewUserTokenPayload(userId uint32, userName string) *UserTokenPayload {
	return &UserTokenPayload{
		UserId:   userId,
		UserName: userName,
	}
}

func NewUserTokenPayloadWithClaims(claims *authn.AuthClaims) (*UserTokenPayload, error) {
	userToken := &UserTokenPayload{}

	if err := userToken.ExtractAuthClaims(claims); err != nil {
		return nil, err
	}

	return userToken, nil
}

// MakeAuthClaims 构建认证声明
func (t *UserTokenPayload) MakeAuthClaims() *authn.AuthClaims {
	return &authn.AuthClaims{
		authn.ClaimFieldSubject: t.UserName,
		ClaimFieldUserID:        t.UserId,
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

	clientId, _ := claims.GetString(ClaimFieldClientID)
	if clientId != "" {
		t.ClientId = clientId
	}

	return nil
}
