package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tx7do/go-utils/trans"
	authn "github.com/tx7do/kratos-authn/engine"

	userV1 "kratos-admin/api/gen/go/user/service/v1"
)

func TestNewUserTokenPayload(t *testing.T) {
	clientId := "client_123"
	user := userV1.User{
		Id:        trans.Ptr(uint32(1)),
		TenantId:  trans.Ptr(uint32(2)),
		UserName:  trans.Ptr("test_user"),
		Authority: trans.Ptr(userV1.UserAuthority_TENANT_ADMIN),
	}

	payload := NewUserTokenPayload(&user, clientId)

	assert.Equal(t, user.GetTenantId(), payload.GetTenantId())
	assert.Equal(t, user.GetId(), payload.GetUserId())
	assert.Equal(t, user.GetUserName(), payload.GetUsername())
	assert.Equal(t, user.GetAuthority(), payload.GetAuthority())
	assert.Equal(t, clientId, payload.ClientId)
}

func TestMakeAuthClaims(t *testing.T) {
	clientId := "client_123"
	user := userV1.User{
		Id:        trans.Ptr(uint32(1)),
		TenantId:  trans.Ptr(uint32(2)),
		UserName:  trans.Ptr("test_user"),
		Authority: trans.Ptr(userV1.UserAuthority_TENANT_ADMIN),
	}

	claims := NewUserTokenAuthClaims(&user, clientId)

	assert.Equal(t, user.GetUserName(), (*claims)[authn.ClaimFieldSubject])
	assert.Equal(t, user.GetId(), (*claims)[ClaimFieldUserID])
	assert.Equal(t, user.GetTenantId(), (*claims)[ClaimFieldTenantID])
	assert.Equal(t, clientId, (*claims)[ClaimFieldClientID])
	assert.Equal(t, user.GetAuthority().String(), (*claims)[ClaimFieldAuthority])
}

func TestExtractAuthClaims(t *testing.T) {
	clientId := "client_123"
	user := userV1.User{
		Id:        trans.Ptr(uint32(1)),
		TenantId:  trans.Ptr(uint32(2)),
		UserName:  trans.Ptr("test_user"),
		Authority: trans.Ptr(userV1.UserAuthority_TENANT_ADMIN),
	}

	claims := &authn.AuthClaims{
		authn.ClaimFieldSubject: user.GetUserName(),
		ClaimFieldUserID:        user.GetId(),
		ClaimFieldTenantID:      user.GetTenantId(),
		ClaimFieldClientID:      clientId,
		ClaimFieldAuthority:     user.GetAuthority().String(),
	}

	payload, err := NewUserTokenPayloadWithClaims(claims)
	assert.NoError(t, err)
	assert.Equal(t, user.GetUserName(), payload.GetUsername())
	assert.Equal(t, user.GetId(), payload.GetUserId())
	assert.Equal(t, user.GetTenantId(), payload.GetTenantId())
	assert.Equal(t, clientId, payload.GetClientId())
	assert.Equal(t, user.GetAuthority(), payload.GetAuthority())
}
