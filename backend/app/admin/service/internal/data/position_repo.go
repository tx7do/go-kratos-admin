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
	"go-wind-admin/app/admin/service/internal/data/ent/position"
	"go-wind-admin/app/admin/service/internal/data/ent/predicate"

	userV1 "go-wind-admin/api/gen/go/user/service/v1"
)

type PositionRepo struct {
	data *Data
	log  *log.Helper

	mapper          *mapper.CopierMapper[userV1.Position, ent.Position]
	statusConverter *mapper.EnumTypeConverter[userV1.Position_Status, position.Status]

	repository *entCrud.Repository[
		ent.PositionQuery, ent.PositionSelect,
		ent.PositionCreate, ent.PositionCreateBulk,
		ent.PositionUpdate, ent.PositionUpdateOne,
		ent.PositionDelete,
		predicate.Position,
		userV1.Position, ent.Position,
	]
}

func NewPositionRepo(ctx *bootstrap.Context, data *Data) *PositionRepo {
	repo := &PositionRepo{
		log:             ctx.NewLoggerHelper("position/repo/admin-service"),
		data:            data,
		mapper:          mapper.NewCopierMapper[userV1.Position, ent.Position](),
		statusConverter: mapper.NewEnumTypeConverter[userV1.Position_Status, position.Status](userV1.Position_Status_name, userV1.Position_Status_value),
	}

	repo.init()

	return repo
}

func (r *PositionRepo) init() {
	r.repository = entCrud.NewRepository[
		ent.PositionQuery, ent.PositionSelect,
		ent.PositionCreate, ent.PositionCreateBulk,
		ent.PositionUpdate, ent.PositionUpdateOne,
		ent.PositionDelete,
		predicate.Position,
		userV1.Position, ent.Position,
	](r.mapper)

	r.mapper.AppendConverters(copierutil.NewTimeStringConverterPair())
	r.mapper.AppendConverters(copierutil.NewTimeTimestamppbConverterPair())

	r.mapper.AppendConverters(r.statusConverter.NewConverterPair())
}

func (r *PositionRepo) Count(ctx context.Context, whereCond []func(s *sql.Selector)) (int, error) {
	builder := r.data.db.Client().Position.Query()
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

func (r *PositionRepo) List(ctx context.Context, req *pagination.PagingRequest) (*userV1.ListPositionResponse, error) {
	if req == nil {
		return nil, userV1.ErrorBadRequest("invalid parameter")
	}

	builder := r.data.db.Client().Position.Query()

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

	dtos := make([]*userV1.Position, 0, len(entities))
	for _, entity := range entities {
		if entity.ParentID == nil {
			dto := r.mapper.ToDTO(entity)
			dtos = append(dtos, dto)
		}
	}
	for _, entity := range entities {
		if entity.ParentID != nil {
			dto := r.mapper.ToDTO(entity)

			if entCrud.TravelChild(&dtos, dto, func(parent *userV1.Position, node *userV1.Position) {
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

	return &userV1.ListPositionResponse{
		Total: uint64(count),
		Items: dtos,
	}, err
}

func (r *PositionRepo) IsExist(ctx context.Context, id uint32) (bool, error) {
	exist, err := r.data.db.Client().Position.Query().
		Where(position.IDEQ(id)).
		Exist(ctx)
	if err != nil {
		r.log.Errorf("query exist failed: %s", err.Error())
		return false, userV1.ErrorInternalServerError("query exist failed")
	}
	return exist, nil
}

func (r *PositionRepo) Get(ctx context.Context, req *userV1.GetPositionRequest) (*userV1.Position, error) {
	if req == nil {
		return nil, userV1.ErrorBadRequest("invalid parameter")
	}

	builder := r.data.db.Client().Position.Query()

	var whereCond []func(s *sql.Selector)
	switch req.QueryBy.(type) {
	default:
	case *userV1.GetPositionRequest_Id:
		whereCond = append(whereCond, position.IDEQ(req.GetId()))
	}

	dto, err := r.repository.Get(ctx, builder, req.GetViewMask(), whereCond...)
	if err != nil {
		return nil, err
	}

	return dto, err
}

// ListPositionByIds 通过多个ID获取职位信息
func (r *PositionRepo) ListPositionByIds(ctx context.Context, ids []uint32) ([]*userV1.Position, error) {
	if len(ids) == 0 {
		return []*userV1.Position{}, nil
	}

	entities, err := r.data.db.Client().Position.Query().
		Where(position.IDIn(ids...)).
		All(ctx)
	if err != nil {
		r.log.Errorf("query position by ids failed: %s", err.Error())
		return nil, userV1.ErrorInternalServerError("query position by ids failed")
	}

	dtos := make([]*userV1.Position, 0, len(entities))
	for _, entity := range entities {
		dto := r.mapper.ToDTO(entity)
		dtos = append(dtos, dto)
	}

	return dtos, nil
}

func (r *PositionRepo) Create(ctx context.Context, req *userV1.CreatePositionRequest) error {
	if req == nil || req.Data == nil {
		return userV1.ErrorBadRequest("invalid parameter")
	}

	builder := r.data.db.Client().Position.Create().
		SetNillableName(req.Data.Name).
		SetNillableParentID(req.Data.ParentId).
		SetNillableSortOrder(req.Data.SortOrder).
		SetNillableCode(req.Data.Code).
		SetNillableStatus(r.statusConverter.ToEntity(req.Data.Status)).
		SetNillableRemark(req.Data.Remark).
		SetNillableQuota(req.Data.Quota).
		SetNillableDescription(req.Data.Description).
		SetOrganizationID(req.Data.GetOrganizationId()).
		SetDepartmentID(req.Data.GetDepartmentId()).
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

func (r *PositionRepo) Update(ctx context.Context, req *userV1.UpdatePositionRequest) error {
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
			createReq := &userV1.CreatePositionRequest{Data: req.Data}
			createReq.Data.CreatedBy = createReq.Data.UpdatedBy
			createReq.Data.UpdatedBy = nil
			return r.Create(ctx, createReq)
		}
	}

	builder := r.data.db.Client().Debug().Position.Update()
	err := r.repository.UpdateX(ctx, builder, req.Data, req.GetUpdateMask(),
		func(dto *userV1.Position) {
			builder.
				SetNillableName(req.Data.Name).
				SetNillableParentID(req.Data.ParentId).
				SetNillableSortOrder(req.Data.SortOrder).
				SetNillableCode(req.Data.Code).
				SetNillableStatus(r.statusConverter.ToEntity(req.Data.Status)).
				SetNillableRemark(req.Data.Remark).
				SetNillableQuota(req.Data.Quota).
				SetNillableDescription(req.Data.Description).
				SetNillableUpdatedBy(req.Data.UpdatedBy).
				SetNillableUpdatedAt(timeutil.TimestamppbToTime(req.Data.UpdatedAt))

			if req.Data.UpdatedAt == nil {
				builder.SetUpdatedAt(time.Now())
			}

			if req.Data.OrganizationId == nil {
				builder.SetOrganizationID(req.Data.GetOrganizationId())
			}

			if req.Data.DepartmentId == nil {
				builder.SetDepartmentID(req.Data.GetDepartmentId())
			}
		},
		func(s *sql.Selector) {
			s.Where(sql.EQ(position.FieldID, req.GetId()))
		},
	)

	return err
}

func (r *PositionRepo) Delete(ctx context.Context, req *userV1.DeletePositionRequest) error {
	if req == nil {
		return userV1.ErrorBadRequest("invalid parameter")
	}

	ids, err := entCrud.QueryAllChildrenIds(ctx, r.data.db, "sys_positions", req.GetId())
	if err != nil {
		r.log.Errorf("query child positions failed: %s", err.Error())
		return userV1.ErrorInternalServerError("query child positions failed")
	}
	ids = append(ids, req.GetId())

	//r.log.Infof("child positions to delete: %+v", ids)

	builder := r.data.db.Client().Debug().Position.Delete()

	_, err = r.repository.Delete(ctx, builder, func(s *sql.Selector) {
		s.Where(sql.In(position.FieldID, ids))
	})
	if err != nil {
		r.log.Errorf("delete position failed: %s", err.Error())
		return userV1.ErrorInternalServerError("delete position failed")
	}

	return nil
}
