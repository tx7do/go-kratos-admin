package data

import (
	"context"
	"sort"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"

	"github.com/tx7do/go-utils/copierutil"
	entgo "github.com/tx7do/go-utils/entgo/query"
	entgoUpdate "github.com/tx7do/go-utils/entgo/update"
	"github.com/tx7do/go-utils/fieldmaskutil"
	"github.com/tx7do/go-utils/mapper"
	"github.com/tx7do/go-utils/timeutil"
	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"

	"kratos-admin/app/admin/service/internal/data/ent"
	"kratos-admin/app/admin/service/internal/data/ent/organization"

	userV1 "kratos-admin/api/gen/go/user/service/v1"
)

var (
	OrganizationStatusNameMap = map[int32]string{
		int32(userV1.OrganizationStatus_ORGANIZATION_STATUS_ON):  string(organization.StatusORGANIZATION_STATUS_ON),
		int32(userV1.OrganizationStatus_ORGANIZATION_STATUS_OFF): string(organization.StatusORGANIZATION_STATUS_OFF),
	}

	OrganizationStatusValueMap = map[string]int32{
		string(organization.StatusORGANIZATION_STATUS_ON):  int32(userV1.OrganizationStatus_ORGANIZATION_STATUS_ON),
		string(organization.StatusORGANIZATION_STATUS_OFF): int32(userV1.OrganizationStatus_ORGANIZATION_STATUS_OFF),
	}

	OrganizationTypeNameMap = map[int32]string{
		int32(userV1.OrganizationType_ORGANIZATION_TYPE_GROUP):      string(organization.OrganizationTypeORGANIZATION_TYPE_GROUP),
		int32(userV1.OrganizationType_ORGANIZATION_TYPE_SUBSIDIARY): string(organization.OrganizationTypeORGANIZATION_TYPE_SUBSIDIARY),
		int32(userV1.OrganizationType_ORGANIZATION_TYPE_FILIALE):    string(organization.OrganizationTypeORGANIZATION_TYPE_FILIALE),
		int32(userV1.OrganizationType_ORGANIZATION_TYPE_DIVISION):   string(organization.OrganizationTypeORGANIZATION_TYPE_DIVISION),
	}

	OrganizationTypeValueMap = map[string]int32{
		string(organization.OrganizationTypeORGANIZATION_TYPE_GROUP):      int32(userV1.OrganizationType_ORGANIZATION_TYPE_GROUP),
		string(organization.OrganizationTypeORGANIZATION_TYPE_SUBSIDIARY): int32(userV1.OrganizationType_ORGANIZATION_TYPE_SUBSIDIARY),
		string(organization.OrganizationTypeORGANIZATION_TYPE_FILIALE):    int32(userV1.OrganizationType_ORGANIZATION_TYPE_FILIALE),
		string(organization.OrganizationTypeORGANIZATION_TYPE_DIVISION):   int32(userV1.OrganizationType_ORGANIZATION_TYPE_DIVISION),
	}
)

type OrganizationRepo struct {
	data *Data
	log  *log.Helper

	mapper          *mapper.CopierMapper[userV1.Organization, ent.Organization]
	typeConverter   *mapper.EnumTypeConverter[userV1.OrganizationType, organization.OrganizationType]
	statusConverter *mapper.EnumTypeConverter[userV1.OrganizationStatus, organization.Status]
}

func NewOrganizationRepo(data *Data, logger log.Logger) *OrganizationRepo {
	repo := &OrganizationRepo{
		log:             log.NewHelper(log.With(logger, "module", "organization/repo/admin-service")),
		data:            data,
		mapper:          mapper.NewCopierMapper[userV1.Organization, ent.Organization](),
		typeConverter:   mapper.NewEnumTypeConverter[userV1.OrganizationType, organization.OrganizationType](OrganizationTypeNameMap, OrganizationTypeValueMap),
		statusConverter: mapper.NewEnumTypeConverter[userV1.OrganizationStatus, organization.Status](OrganizationStatusNameMap, OrganizationStatusValueMap),
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

	err, whereSelectors, querySelectors := entgo.BuildQuerySelector(
		req.GetQuery(), req.GetOrQuery(),
		req.GetPage(), req.GetPageSize(), req.GetNoPaging(),
		req.GetOrderBy(), organization.FieldCreateTime,
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
		if entities[i].SortID != nil {
			sortI = *entities[i].SortID
		}
		if entities[j].SortID != nil {
			sortJ = *entities[j].SortID
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
		SetNillableSortID(req.Data.SortId).
		SetNillableRemark(req.Data.Remark).
		SetNillableStatus(r.statusConverter.ToEntity(req.Data.Status)).
		SetNillableOrganizationType(r.typeConverter.ToEntity(req.Data.OrganizationType)).
		SetNillableIsLegalEntity(req.Data.IsLegalEntity).
		SetNillableBusinessScope(req.Data.BusinessScope).
		SetNillableCreditCode(req.Data.CreditCode).
		SetNillableAddress(req.Data.Address).
		SetNillableManagerID(req.Data.ManagerId).
		SetNillableTenantID(req.Data.TenantId).
		SetNillableCreateBy(req.Data.CreateBy).
		SetNillableCreateTime(timeutil.TimestamppbToTime(req.Data.CreateTime))

	if req.Data.CreateTime == nil {
		builder.SetCreateTime(time.Now())
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
			createReq.Data.CreateBy = createReq.Data.UpdateBy
			createReq.Data.UpdateBy = nil
			return r.Create(ctx, createReq)
		}
	}

	if req.UpdateMask != nil {
		req.UpdateMask.Normalize()
		if !req.UpdateMask.IsValid(req.Data) {
			r.log.Errorf("invalid field mask [%v]", req.UpdateMask)
			return userV1.ErrorBadRequest("invalid field mask")
		}
		fieldmaskutil.Filter(req.GetData(), req.UpdateMask.GetPaths())
	}

	builder := r.data.db.Client().Organization.
		UpdateOneID(req.Data.GetId()).
		SetNillableName(req.Data.Name).
		SetNillableParentID(req.Data.ParentId).
		SetNillableSortID(req.Data.SortId).
		SetNillableRemark(req.Data.Remark).
		SetNillableStatus(r.statusConverter.ToEntity(req.Data.Status)).
		SetNillableOrganizationType(r.typeConverter.ToEntity(req.Data.OrganizationType)).
		SetNillableIsLegalEntity(req.Data.IsLegalEntity).
		SetNillableBusinessScope(req.Data.BusinessScope).
		SetNillableCreditCode(req.Data.CreditCode).
		SetNillableAddress(req.Data.Address).
		SetNillableManagerID(req.Data.ManagerId).
		SetNillableUpdateBy(req.Data.UpdateBy).
		SetNillableUpdateTime(timeutil.TimestamppbToTime(req.Data.UpdateTime))

	if req.Data.UpdateTime == nil {
		builder.SetUpdateTime(time.Now())
	}

	if req.UpdateMask != nil {
		nilPaths := fieldmaskutil.NilValuePaths(req.Data, req.GetUpdateMask().GetPaths())
		nilUpdater := entgoUpdate.BuildSetNullUpdater(nilPaths)
		if nilUpdater != nil {
			builder.Modify(nilUpdater)
		}
	}

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

	if err := r.data.db.Client().Organization.DeleteOneID(req.GetId()).Exec(ctx); err != nil {
		if ent.IsNotFound(err) {
			return userV1.ErrorNotFound("organization not found")
		}

		r.log.Errorf("delete one data failed: %s", err.Error())

		return userV1.ErrorInternalServerError("delete failed")
	}

	return nil
}
