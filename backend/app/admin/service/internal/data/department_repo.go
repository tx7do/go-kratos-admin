package data

import (
	"context"
	"sort"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"

	pagination "github.com/tx7do/go-crud/api/gen/go/pagination/v1"
	entCrud "github.com/tx7do/go-crud/entgo"

	"github.com/tx7do/go-utils/copierutil"
	"github.com/tx7do/go-utils/mapper"
	"github.com/tx7do/go-utils/timeutil"

	"kratos-admin/app/admin/service/internal/data/ent"
	"kratos-admin/app/admin/service/internal/data/ent/department"
	"kratos-admin/app/admin/service/internal/data/ent/predicate"

	userV1 "kratos-admin/api/gen/go/user/service/v1"
)

type DepartmentRepo struct {
	data *Data
	log  *log.Helper

	mapper          *mapper.CopierMapper[userV1.Department, ent.Department]
	statusConverter *mapper.EnumTypeConverter[userV1.Department_Status, department.Status]

	repository *entCrud.Repository[
		ent.DepartmentQuery, ent.DepartmentSelect,
		ent.DepartmentCreate, ent.DepartmentCreateBulk,
		ent.DepartmentUpdate, ent.DepartmentUpdateOne,
		ent.DepartmentDelete,
		predicate.Department,
		userV1.Department, ent.Department,
	]
}

func NewDepartmentRepo(data *Data, logger log.Logger) *DepartmentRepo {
	repo := &DepartmentRepo{
		log:             log.NewHelper(log.With(logger, "module", "department/repo/admin-service")),
		data:            data,
		mapper:          mapper.NewCopierMapper[userV1.Department, ent.Department](),
		statusConverter: mapper.NewEnumTypeConverter[userV1.Department_Status, department.Status](userV1.Department_Status_name, userV1.Department_Status_value),
	}

	repo.init()

	return repo
}

func (r *DepartmentRepo) init() {
	r.repository = entCrud.NewRepository[
		ent.DepartmentQuery, ent.DepartmentSelect,
		ent.DepartmentCreate, ent.DepartmentCreateBulk,
		ent.DepartmentUpdate, ent.DepartmentUpdateOne,
		ent.DepartmentDelete,
		predicate.Department,
		userV1.Department, ent.Department,
	](r.mapper)

	r.mapper.AppendConverters(copierutil.NewTimeStringConverterPair())
	r.mapper.AppendConverters(copierutil.NewTimeTimestamppbConverterPair())

	r.mapper.AppendConverters(r.statusConverter.NewConverterPair())
}

func (r *DepartmentRepo) Count(ctx context.Context, whereCond []func(s *sql.Selector)) (int, error) {
	builder := r.data.db.Client().Department.Query()
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

func (r *DepartmentRepo) List(ctx context.Context, req *pagination.PagingRequest) (*userV1.ListDepartmentResponse, error) {
	if req == nil {
		return nil, userV1.ErrorBadRequest("invalid parameter")
	}

	builder := r.data.db.Client().Department.Query()

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

	dtos := make([]*userV1.Department, 0, len(entities))
	for _, entity := range entities {
		if entity.ParentID == nil {
			dto := r.mapper.ToDTO(entity)
			dtos = append(dtos, dto)
		}
	}
	for _, entity := range entities {
		if entity.ParentID != nil {
			dto := r.mapper.ToDTO(entity)

			if entCrud.TravelChild(&dtos, dto, func(parent *userV1.Department, node *userV1.Department) {
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

	ret := userV1.ListDepartmentResponse{
		Total: uint64(count),
		Items: dtos,
	}

	return &ret, err
}

func (r *DepartmentRepo) IsExist(ctx context.Context, id uint32) (bool, error) {
	exist, err := r.data.db.Client().Department.Query().
		Where(department.IDEQ(id)).
		Exist(ctx)
	if err != nil {
		r.log.Errorf("query exist failed: %s", err.Error())
		return false, userV1.ErrorInternalServerError("query exist failed")
	}
	return exist, nil
}

func (r *DepartmentRepo) Get(ctx context.Context, req *userV1.GetDepartmentRequest) (*userV1.Department, error) {
	if req == nil {
		return nil, userV1.ErrorBadRequest("invalid parameter")
	}

	builder := r.data.db.Client().Department.Query()

	var whereCond []func(s *sql.Selector)
	switch req.QueryBy.(type) {
	default:
	case *userV1.GetDepartmentRequest_Id:
		whereCond = append(whereCond, department.IDEQ(req.GetId()))
	}

	dto, err := r.repository.Get(ctx, builder, req.GetViewMask(), whereCond...)
	if err != nil {
		return nil, err
	}

	return dto, err
}

// ListDepartmentsByIds 通过多个ID获取部门信息列表
func (r *DepartmentRepo) ListDepartmentsByIds(ctx context.Context, ids []uint32) ([]*userV1.Department, error) {
	if len(ids) == 0 {
		return []*userV1.Department{}, nil
	}

	entities, err := r.data.db.Client().Department.Query().
		Where(department.IDIn(ids...)).
		All(ctx)
	if err != nil {
		r.log.Errorf("query department by ids failed: %s", err.Error())
		return nil, userV1.ErrorInternalServerError("query department by ids failed")
	}

	dtos := make([]*userV1.Department, 0, len(entities))
	for _, entity := range entities {
		dto := r.mapper.ToDTO(entity)
		dtos = append(dtos, dto)
	}

	return dtos, nil
}

func (r *DepartmentRepo) Create(ctx context.Context, req *userV1.CreateDepartmentRequest) error {
	if req == nil || req.Data == nil {
		return userV1.ErrorBadRequest("invalid parameter")
	}

	builder := r.data.db.Client().Department.Create().
		SetNillableName(req.Data.Name).
		SetNillableParentID(req.Data.ParentId).
		SetNillableSortOrder(req.Data.SortOrder).
		SetNillableRemark(req.Data.Remark).
		SetNillableStatus(r.statusConverter.ToEntity(req.Data.Status)).
		SetOrganizationID(req.Data.GetOrganizationId()).
		SetNillableManagerID(req.Data.ManagerId).
		SetNillableDescription(req.Data.Description).
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

func (r *DepartmentRepo) Update(ctx context.Context, req *userV1.UpdateDepartmentRequest) error {
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
			createReq := &userV1.CreateDepartmentRequest{Data: req.Data}
			createReq.Data.CreatedBy = createReq.Data.UpdatedBy
			createReq.Data.UpdatedBy = nil
			return r.Create(ctx, createReq)
		}
	}

	builder := r.data.db.Client().Debug().Department.Update()
	err := r.repository.UpdateX(ctx, builder, req.Data, req.GetUpdateMask(),
		func(dto *userV1.Department) {
			builder.
				SetNillableName(req.Data.Name).
				SetNillableParentID(req.Data.ParentId).
				SetNillableSortOrder(req.Data.SortOrder).
				SetNillableRemark(req.Data.Remark).
				SetNillableStatus(r.statusConverter.ToEntity(req.Data.Status)).
				SetNillableDescription(req.Data.Description).
				SetNillableUpdatedBy(req.Data.UpdatedBy).
				SetNillableUpdatedAt(timeutil.TimestamppbToTime(req.Data.UpdatedAt))

			if req.Data.UpdatedAt == nil {
				builder.SetUpdatedAt(time.Now())
			}

			if req.Data.OrganizationId == nil {
				builder.SetOrganizationID(req.Data.GetOrganizationId())
			}
		},
		func(s *sql.Selector) {
			s.Where(sql.EQ(department.FieldID, req.Data.GetId()))
		},
	)

	return err
}

func (r *DepartmentRepo) Delete(ctx context.Context, req *userV1.DeleteDepartmentRequest) error {
	if req == nil {
		return userV1.ErrorBadRequest("invalid parameter")
	}

	ids, err := entCrud.QueryAllChildrenIds(ctx, r.data.db, "sys_departments", req.GetId())
	if err != nil {
		r.log.Errorf("query child departments failed: %s", err.Error())
		return userV1.ErrorInternalServerError("query child departments failed")
	}
	ids = append(ids, req.GetId())

	//r.log.Info("department ids to delete: ", ids)

	builder := r.data.db.Client().Debug().Department.Delete()

	_, err = r.repository.Delete(ctx, builder, func(s *sql.Selector) {
		s.Where(sql.In(department.FieldID, ids))
	})
	if err != nil {
		r.log.Errorf("delete department failed: %s", err.Error())
		return userV1.ErrorInternalServerError("delete department failed")
	}

	return nil
}
