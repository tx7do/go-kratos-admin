package metadata

import (
	"context"
	"strconv"

	"github.com/go-kratos/kratos/v2/metadata"

	userV1 "go-wind-admin/api/gen/go/user/service/v1"
)

const (
	mdOperatorId = "x-md-operator-id"
	mdAuthority  = "x-md-authority"
	mdTenantId   = "x-md-tenant-id"
)

func FromOperatorMetadata(ctx context.Context) (userId *uint32, tenantId *uint32, authority *userV1.User_Authority) {
	md, ok := metadata.FromServerContext(ctx)
	if !ok {
		return
	}

	if id := md.Get(mdOperatorId); id != "" {
		if i, err := strconv.Atoi(id); err == nil {
			userId = new(uint32)
			*userId = uint32(i)
		}
	}

	if id := md.Get(mdTenantId); id != "" {
		if i, err := strconv.Atoi(id); err == nil {
			tenantId = new(uint32)
			*tenantId = uint32(i)
		}
	}

	if authorityStr := md.Get(mdAuthority); authorityStr != "" {
		a, ok := userV1.User_Authority_value[authorityStr]
		if ok {
			authority = new(userV1.User_Authority)
			*authority = userV1.User_Authority(a)
		}
	}

	return
}

func NewOperatorMetadataContext(ctx context.Context, userId *uint32, tenantId *uint32, authority *userV1.User_Authority) context.Context {
	if userId != nil {
		ctx = metadata.AppendToClientContext(ctx, mdOperatorId, strconv.Itoa(int(*userId)))
	}
	if tenantId != nil {
		ctx = metadata.AppendToClientContext(ctx, mdTenantId, strconv.Itoa(int(*tenantId)))
	}
	if authority != nil {
		ctx = metadata.AppendToClientContext(ctx, mdAuthority, authority.String())
	}
	return ctx
}
