package data

import (
	"context"
	"sort"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/kratos-bootstrap/bootstrap"

	pagination "github.com/tx7do/go-crud/api/gen/go/pagination/v1"
	entCrud "github.com/tx7do/go-crud/entgo"

	"github.com/tx7do/go-utils/copierutil"
	"github.com/tx7do/go-utils/mapper"
	"github.com/tx7do/go-utils/timeutil"

	"go-wind-admin/app/admin/service/internal/data/ent"
	"go-wind-admin/app/admin/service/internal/data/ent/organization"
	"go-wind-admin/app/admin/service/internal/data/ent/predicate"

	userV1 "go-wind-admin/api/gen/go/user/service/v1"
)

type OrganizationRepo struct {
	data *Data
	log  *log.Helper

	mapper          *mapper.CopierMapper[userV1.Organization, ent.Organization]
	typeConverter   *mapper.EnumTypeConverter[userV1.Organization_Type, organization.OrganizationType]
	statusConverter *mapper.EnumTypeConverter[userV1.Organization_Status, organization.Status]

	repository *entCrud.Repository[
		ent.OrganizationQuery, ent.OrganizationSelect,
		ent.OrganizationCreate, ent.OrganizationCreateBulk,
		ent.OrganizationUpdate, ent.OrganizationUpdateOne,
		ent.OrganizationDelete,
		predicate.Organization,
		userV1.Organization, ent.Organization,
	]
}

func NewOrganizationRepo(ctx *bootstrap.Context, data *Data) *OrganizationRepo {
	repo := &OrganizationRepo{
		log:             ctx.NewLoggerHelper("organization/repo/admin-service"),
		data:            data,
		mapper:          mapper.NewCopierMapper[userV1.Organization, ent.Organization](),
		typeConverter:   mapper.NewEnumTypeConverter[userV1.Organization_Type, organization.OrganizationType](userV1.Organization_Type_name, userV1.Organization_Type_value),
		statusConverter: mapper.NewEnumTypeConverter[userV1.Organization_Status, organization.Status](userV1.Organization_Status_name, userV1.Organization_Status_value),
	}

	repo.init()

	return repo
}

func (r *OrganizationRepo) init() {
	r.repository = entCrud.NewRepository[
		ent.OrganizationQuery, ent.OrganizationSelect,
		ent.OrganizationCreate, ent.OrganizationCreateBulk,
		ent.OrganizationUpdate, ent.OrganizationUpdateOne,
		ent.OrganizationDelete,
		predicate.Organization,
		userV1.Organization, ent.Organization,
	](r.mapper)

	r.mapper.AppendConverters(copierutil.NewTimeStringConverterPair())
	r.mapper.AppendConverters(copierutil.NewTimeTimestamppbConverterPair())

	r.mapper.AppendConverters(r.typeConverter.NewConverterPair())
	r.mapper.AppendConverters(r.statusConverter.NewConverterPair())
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

	whereSelectors, _, err := r.repository.BuildListSelectorWithPaging(builder, req)
	if err != nil {
		r.log.Errorf("parse list param error [%s]", err.Error())
		return nil, userV1.ErrorBadRequest("invalid query parameter")
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

			if entCrud.TravelChild(&dtos, dto, func(parent *userV1.Organization, node *userV1.Organization) {
				parent.Children = append(parent.Children, node)
			}) {
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
		Total: uint64(count),
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

	builder := r.data.db.Client().Organization.Query()

	var whereCond []func(s *sql.Selector)
	switch req.QueryBy.(type) {
	default:
	case *userV1.GetOrganizationRequest_Id:
		whereCond = append(whereCond, organization.IDEQ(req.GetId()))
	}

	dto, err := r.repository.Get(ctx, builder, req.GetViewMask(), whereCond...)
	if err != nil {
		return nil, err
	}

	return dto, err
}

// ListOrganizationsByIds 通过多个ID获取组织列表
func (r *OrganizationRepo) ListOrganizationsByIds(ctx context.Context, ids []uint32) ([]*userV1.Organization, error) {
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
		builder.SetID(req.GetData().GetId())
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
		exist, err := r.IsExist(ctx, req.GetId())
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

	builder := r.data.db.Client().Debug().Organization.Update()
	err := r.repository.UpdateX(ctx, builder, req.Data, req.GetUpdateMask(),
		func(dto *userV1.Organization) {
			builder.
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
		},
		func(s *sql.Selector) {
			s.Where(sql.EQ(organization.FieldID, req.GetId()))
		},
	)

	return err
}

func (r *OrganizationRepo) Delete(ctx context.Context, req *userV1.DeleteOrganizationRequest) error {
	if req == nil {
		return userV1.ErrorBadRequest("invalid parameter")
	}

	ids, err := entCrud.QueryAllChildrenIds(ctx, r.data.db, "sys_organizations", req.GetId())
	if err != nil {
		r.log.Errorf("query child organizations failed: %s", err.Error())
		return userV1.ErrorInternalServerError("query child organizations failed")
	}
	ids = append(ids, req.GetId())

	//r.log.Info("organizations ids to delete: ", ids)

	builder := r.data.db.Client().Debug().Organization.Delete()

	_, err = r.repository.Delete(ctx, builder, func(s *sql.Selector) {
		s.Where(sql.In(organization.FieldID, ids))
	})
	if err != nil {
		r.log.Errorf("delete organization failed: %s", err.Error())
		return userV1.ErrorInternalServerError("delete organization failed")
	}

	return nil
}
