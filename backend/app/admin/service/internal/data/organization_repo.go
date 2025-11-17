package data

import (
	"context"
	"sort"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/go-utils/entgo"
	"github.com/tx7do/go-utils/trans"
	"google.golang.org/protobuf/proto"

	"github.com/tx7do/go-utils/copierutil"
	entgoQuery "github.com/tx7do/go-utils/entgo/query"
	entgoUpdate "github.com/tx7do/go-utils/entgo/update"
	"github.com/tx7do/go-utils/fieldmaskutil"
	"github.com/tx7do/go-utils/mapper"
	"github.com/tx7do/go-utils/timeutil"
	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"

	"kratos-admin/app/admin/service/internal/data/ent"
	"kratos-admin/app/admin/service/internal/data/ent/organization"

	userV1 "kratos-admin/api/gen/go/user/service/v1"
)

type OrganizationRepo struct {
	data *Data
	log  *log.Helper

	mapper          *mapper.CopierMapper[userV1.Organization, ent.Organization]
	typeConverter   *mapper.EnumTypeConverter[userV1.Organization_Type, organization.OrganizationType]
	statusConverter *mapper.EnumTypeConverter[userV1.Organization_Status, organization.Status]
}

func NewOrganizationRepo(data *Data, logger log.Logger) *OrganizationRepo {
	repo := &OrganizationRepo{
		log:             log.NewHelper(log.With(logger, "module", "organization/repo/admin-service")),
		data:            data,
		mapper:          mapper.NewCopierMapper[userV1.Organization, ent.Organization](),
		typeConverter:   mapper.NewEnumTypeConverter[userV1.Organization_Type, organization.OrganizationType](userV1.Organization_Type_name, userV1.Organization_Type_value),
		statusConverter: mapper.NewEnumTypeConverter[userV1.Organization_Status, organization.Status](userV1.Organization_Status_name, userV1.Organization_Status_value),
	}

	repo.init()

	return repo
}

func (r *OrganizationRepo) init() {
	r.mapper.AppendConverters(copierutil.NewTimeStringConverterPair())
	r.mapper.AppendConverters(copierutil.NewTimeTimestamppbConverterPair())

	r.mapper.AppendConverters(r.typeConverter.NewConverterPair())
	r.mapper.AppendConverters(r.statusConverter.NewConverterPair())
}

func (r *OrganizationRepo) travelChild(nodes []*userV1.Organization, node *userV1.Organization) bool {
	if nodes == nil {
		return false
	}

	if node.ParentId == nil {
		nodes = append(nodes, node)
		return true
	}

	for _, n := range nodes {
		if node.ParentId == nil {
			continue
		}

		if n.GetId() == node.GetParentId() {
			n.Children = append(n.Children, node)
			return true
		} else {
			if r.travelChild(n.Children, node) {
				return true
			}
		}
	}
	return false
}

func (r *OrganizationRepo) Count(ctx context.Context, whereCond []func(s *sql.Selector)) (int, error) {
	builder := r.data.db.Client().Organization.Query()
	if len(whereCond) != 0 {
		builder.Modify(whereCond...)
	}

	count, err := builder.Count(ctx)
	if err != nil {
		r.log.Errorf("query count failed: %s", err.Error())
		return 0, userV1.ErrorInternalServerError("query count failed")
	}

	return count, nil
}

func (r *OrganizationRepo) List(ctx context.Context, req *pagination.PagingRequest) (*userV1.ListOrganizationResponse, error) {
	if req == nil {
		return nil, userV1.ErrorBadRequest("invalid parameter")
	}

	builder := r.data.db.Client().Organization.Query()

	err, whereSelectors, querySelectors := entgoQuery.BuildQuerySelector(
		req.GetQuery(), req.GetOrQuery(),
		req.GetPage(), req.GetPageSize(), req.GetNoPaging(),
		req.GetOrderBy(), organization.FieldCreatedAt,
		req.GetFieldMask().GetPaths(),
	)
	if err != nil {
		r.log.Errorf("parse list param error [%s]", err.Error())
		return nil, userV1.ErrorBadRequest("invalid query parameter")
	}

	if querySelectors != nil {
		builder.Modify(querySelectors...)
	}

	entities, err := builder.All(ctx)
	if err != nil {
		r.log.Errorf("query list failed: %s", err.Error())
		return nil, userV1.ErrorInternalServerError("query list failed")
	}

	sort.SliceStable(entities, func(i, j int) bool {
		var sortI, sortJ int32
		if entities[i].SortOrder != nil {
			sortI = *entities[i].SortOrder
		}
		if entities[j].SortOrder != nil {
			sortJ = *entities[j].SortOrder
		}
		return sortI < sortJ
	})

	dtos := make([]*userV1.Organization, 0, len(entities))
	for _, entity := range entities {
		if entity.ParentID == nil {
			dto := r.mapper.ToDTO(entity)
			dtos = append(dtos, dto)
		}
	}
	for _, entity := range entities {
		if entity.ParentID != nil {
			dto := r.mapper.ToDTO(entity)

			if r.travelChild(dtos, dto) {
				continue
			}

			dtos = append(dtos, dto)
		}
	}

	count, err := r.Count(ctx, whereSelectors)
	if err != nil {
		return nil, err
	}

	return &userV1.ListOrganizationResponse{
		Total: uint32(count),
		Items: dtos,
	}, err
}

func (r *OrganizationRepo) IsExist(ctx context.Context, id uint32) (bool, error) {
	exist, err := r.data.db.Client().Organization.Query().
		Where(organization.IDEQ(id)).
		Exist(ctx)
	if err != nil {
		r.log.Errorf("query exist failed: %s", err.Error())
		return false, userV1.ErrorInternalServerError("query exist failed")
	}
	return exist, nil
}

func (r *OrganizationRepo) Get(ctx context.Context, req *userV1.GetOrganizationRequest) (*userV1.Organization, error) {
	if req == nil {
		return nil, userV1.ErrorBadRequest("invalid parameter")
	}

	entity, err := r.data.db.Client().Organization.Get(ctx, req.GetId())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, userV1.ErrorOrganizationNotFound("organization not found")
		}

		r.log.Errorf("query one data failed: %s", err.Error())

		return nil, userV1.ErrorInternalServerError("query data failed")
	}

	return r.mapper.ToDTO(entity), nil
}

// GetOrganizationsByIds 通过多个ID获取组织列表
func (r *OrganizationRepo) GetOrganizationsByIds(ctx context.Context, ids []uint32) ([]*userV1.Organization, error) {
	if len(ids) == 0 {
		return []*userV1.Organization{}, nil
	}

	entities, err := r.data.db.Client().Organization.Query().
		Where(organization.IDIn(ids...)).
		All(ctx)
	if err != nil {
		r.log.Errorf("query organization by ids failed: %s", err.Error())
		return nil, userV1.ErrorInternalServerError("query organization by ids failed")
	}

	dtos := make([]*userV1.Organization, 0, len(entities))
	for _, entity := range entities {
		dto := r.mapper.ToDTO(entity)
		dtos = append(dtos, dto)
	}

	return dtos, nil
}

func (r *OrganizationRepo) Create(ctx context.Context, req *userV1.CreateOrganizationRequest) error {
	if req == nil || req.Data == nil {
		return userV1.ErrorBadRequest("invalid parameter")
	}

	builder := r.data.db.Client().Organization.Create().
		SetNillableName(req.Data.Name).
		SetNillableParentID(req.Data.ParentId).
		SetNillableSortOrder(req.Data.SortOrder).
		SetNillableRemark(req.Data.Remark).
		SetNillableStatus(r.statusConverter.ToEntity(req.Data.Status)).
		SetNillableOrganizationType(r.typeConverter.ToEntity(req.Data.OrganizationType)).
		SetNillableIsLegalEntity(req.Data.IsLegalEntity).
		SetNillableBusinessScope(req.Data.BusinessScope).
		SetNillableCreditCode(req.Data.CreditCode).
		SetNillableAddress(req.Data.Address).
		SetNillableManagerID(req.Data.ManagerId).
		SetNillableCreatedBy(req.Data.CreatedBy).
		SetNillableCreatedAt(timeutil.TimestamppbToTime(req.Data.CreatedAt))

	if req.Data.TenantId == nil {
		builder.SetTenantID(req.Data.GetTenantId())
	}
	if req.Data.CreatedAt == nil {
		builder.SetCreatedAt(time.Now())
	}

	if req.Data.Id != nil {
		builder.SetID(req.Data.GetId())
	}

	if err := builder.Exec(ctx); err != nil {
		r.log.Errorf("insert one data failed: %s", err.Error())
		return userV1.ErrorInternalServerError("insert data failed")
	}

	return nil
}

func (r *OrganizationRepo) Update(ctx context.Context, req *userV1.UpdateOrganizationRequest) error {
	if req == nil || req.Data == nil {
		return userV1.ErrorBadRequest("invalid parameter")
	}

	// 如果不存在则创建
	if req.GetAllowMissing() {
		exist, err := r.IsExist(ctx, req.GetData().GetId())
		if err != nil {
			return err
		}
		if !exist {
			createReq := &userV1.CreateOrganizationRequest{Data: req.Data}
			createReq.Data.CreatedBy = createReq.Data.UpdatedBy
			createReq.Data.UpdatedBy = nil
			return r.Create(ctx, createReq)
		}
	}

	if err := fieldmaskutil.FilterByFieldMask(trans.Ptr(proto.Message(req.GetData())), req.UpdateMask); err != nil {
		r.log.Errorf("invalid field mask [%v], error: %s", req.UpdateMask, err.Error())
		return userV1.ErrorBadRequest("invalid field mask")
	}

	builder := r.data.db.Client().Organization.
		UpdateOneID(req.Data.GetId()).
		SetNillableName(req.Data.Name).
		SetNillableParentID(req.Data.ParentId).
		SetNillableSortOrder(req.Data.SortOrder).
		SetNillableRemark(req.Data.Remark).
		SetNillableStatus(r.statusConverter.ToEntity(req.Data.Status)).
		SetNillableOrganizationType(r.typeConverter.ToEntity(req.Data.OrganizationType)).
		SetNillableIsLegalEntity(req.Data.IsLegalEntity).
		SetNillableBusinessScope(req.Data.BusinessScope).
		SetNillableCreditCode(req.Data.CreditCode).
		SetNillableAddress(req.Data.Address).
		SetNillableManagerID(req.Data.ManagerId).
		SetNillableUpdatedBy(req.Data.UpdatedBy).
		SetNillableUpdatedAt(timeutil.TimestamppbToTime(req.Data.UpdatedAt))

	if req.Data.UpdatedAt == nil {
		builder.SetUpdatedAt(time.Now())
	}

	entgoUpdate.ApplyNilFieldMask(proto.Message(req.GetData()), req.UpdateMask, builder)

	if err := builder.Exec(ctx); err != nil {
		r.log.Errorf("update one data failed: %s", err.Error())
		return userV1.ErrorInternalServerError("update data failed")
	}

	return nil
}

func (r *OrganizationRepo) Delete(ctx context.Context, req *userV1.DeleteOrganizationRequest) error {
	if req == nil {
		return userV1.ErrorBadRequest("invalid parameter")
	}

	ids, err := entgo.QueryAllChildrenIds(ctx, r.data.db, "sys_organizations", req.GetId())
	if err != nil {
		r.log.Errorf("query child organizations failed: %s", err.Error())
		return userV1.ErrorInternalServerError("query child organizations failed")
	}
	ids = append(ids, req.GetId())

	//r.log.Info("organizations ids to delete: ", ids)

	if _, err = r.data.db.Client().Organization.Delete().
		Where(organization.IDIn(ids...)).
		Exec(ctx); err != nil {
		r.log.Errorf("delete organizations failed: %s", err.Error())
		return userV1.ErrorInternalServerError("delete organizations failed")
	}

	return nil
}
