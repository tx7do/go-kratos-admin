package metadata

import (
	"context"
	"strconv"
	"testing"

	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/stretchr/testify/assert"

	userV1 "kratos-admin/api/gen/go/user/service/v1"
)

func TestFromMetadata(t *testing.T) {
	ctx := context.Background()

	userId := "123"
	tenantId := "456"
	authority := userV1.User_SYS_ADMIN.String()

	ctx = metadata.NewServerContext(ctx, metadata.Metadata{
		mdOperatorId: []string{userId},
		mdTenantId:   []string{tenantId},
		mdAuthority:  []string{authority},
	})

	tUserId, tTenantId, tAuthority := FromOperatorMetadata(ctx)

	assert.NotNil(t, userId)
	assert.Equal(t, uint32(123), *tUserId)

	assert.NotNil(t, tenantId)
	assert.Equal(t, uint32(456), *tTenantId)

	assert.NotNil(t, authority)
	assert.Equal(t, authority, tAuthority.String())
}

func TestNewOperatorMetadataContext(t *testing.T) {
	userId := uint32(123)
	tenantId := uint32(456)
	authority := userV1.User_SYS_ADMIN

	ctx := context.Background()

	ctx = NewOperatorMetadataContext(ctx, &userId, &tenantId, &authority)

	md, ok := metadata.FromClientContext(ctx)
	assert.True(t, ok)

	assert.Equal(t, strconv.Itoa(int(userId)), md.Get(mdOperatorId))
	assert.Equal(t, strconv.Itoa(int(tenantId)), md.Get(mdTenantId))
	assert.Equal(t, authority.String(), md.Get(mdAuthority))
}
