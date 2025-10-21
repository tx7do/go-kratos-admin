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
	"kratos-admin/app/admin/service/internal/data/ent/notificationmessagecategory"

	internalMessageV1 "kratos-admin/api/gen/go/internal_message/service/v1"
)

type NotificationMessageCategoryRepo struct {
	data *Data
	log  *log.Helper

	mapper *mapper.CopierMapper[internalMessageV1.NotificationMessageCategory, ent.NotificationMessageCategory]
}

func NewNotificationMessageCategoryRepo(data *Data, logger log.Logger) *NotificationMessageCategoryRepo {
	repo := &NotificationMessageCategoryRepo{
		log:    log.NewHelper(log.With(logger, "module", "notification-message-category/repo/admin-service")),
		data:   data,
		mapper: mapper.NewCopierMapper[internalMessageV1.NotificationMessageCategory, ent.NotificationMessageCategory](),
	}

	repo.init()

	return repo
}

func (r *NotificationMessageCategoryRepo) init() {
	r.mapper.AppendConverters(copierutil.NewTimeStringConverterPair())
	r.mapper.AppendConverters(copierutil.NewTimeTimestamppbConverterPair())
}

func (r *NotificationMessageCategoryRepo) travelChild(nodes []*internalMessageV1.NotificationMessageCategory, node *internalMessageV1.NotificationMessageCategory) bool {
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

func (r *NotificationMessageCategoryRepo) Count(ctx context.Context, whereCond []func(s *sql.Selector)) (int, error) {
	builder := r.data.db.Client().NotificationMessageCategory.Query()
	if len(whereCond) != 0 {
		builder.Modify(whereCond...)
	}

	count, err := builder.Count(ctx)
	if err != nil {
		r.log.Errorf("query count failed: %s", err.Error())
		return 0, internalMessageV1.ErrorInternalServerError("query count failed")
	}

	return count, nil
}

func (r *NotificationMessageCategoryRepo) List(ctx context.Context, req *pagination.PagingRequest) (*internalMessageV1.ListNotificationMessageCategoryResponse, error) {
	if req == nil {
		return nil, internalMessageV1.ErrorBadRequest("invalid parameter")
	}

	builder := r.data.db.Client().NotificationMessageCategory.Query()

	err, whereSelectors, querySelectors := entgo.BuildQuerySelector(
		req.GetQuery(), req.GetOrQuery(),
		req.GetPage(), req.GetPageSize(), req.GetNoPaging(),
		req.GetOrderBy(), notificationmessagecategory.FieldCreateTime,
		req.GetFieldMask().GetPaths(),
	)
	if err != nil {
		r.log.Errorf("parse list param error [%s]", err.Error())
		return nil, internalMessageV1.ErrorBadRequest("invalid query parameter")
	}

	if querySelectors != nil {
		builder.Modify(querySelectors...)
	}

	entities, err := builder.All(ctx)
	if err != nil {
		r.log.Errorf("query list failed: %s", err.Error())
		return nil, internalMessageV1.ErrorInternalServerError("query list failed")
	}

	sort.SliceStable(entities, func(i, j int) bool {
		if entities[j].ParentID == nil {
			return true
		}
		if entities[i].ParentID == nil {
			return true
		}
		return *entities[i].ParentID < *entities[j].ParentID
	})

	dtos := make([]*internalMessageV1.NotificationMessageCategory, 0, len(entities))
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

	return &internalMessageV1.ListNotificationMessageCategoryResponse{
		Total: uint32(count),
		Items: dtos,
	}, err
}

func (r *NotificationMessageCategoryRepo) IsExist(ctx context.Context, id uint32) (bool, error) {
	exist, err := r.data.db.Client().NotificationMessageCategory.Query().
		Where(notificationmessagecategory.IDEQ(id)).
		Exist(ctx)
	if err != nil {
		r.log.Errorf("query exist failed: %s", err.Error())
		return false, internalMessageV1.ErrorInternalServerError("query exist failed")
	}
	return exist, nil
}

func (r *NotificationMessageCategoryRepo) Get(ctx context.Context, req *internalMessageV1.GetNotificationMessageCategoryRequest) (*internalMessageV1.NotificationMessageCategory, error) {
	if req == nil {
		return nil, internalMessageV1.ErrorBadRequest("invalid parameter")
	}

	entity, err := r.data.db.Client().NotificationMessageCategory.Get(ctx, req.GetId())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, internalMessageV1.ErrorNotFound("message category not found")
		}

		r.log.Errorf("query one data failed: %s", err.Error())

		return nil, internalMessageV1.ErrorInternalServerError("query data failed")
	}

	return r.mapper.ToDTO(entity), nil
}

func (r *NotificationMessageCategoryRepo) Create(ctx context.Context, req *internalMessageV1.CreateNotificationMessageCategoryRequest) error {
	if req == nil || req.Data == nil {
		return internalMessageV1.ErrorBadRequest("invalid parameter")
	}

	builder := r.data.db.Client().NotificationMessageCategory.Create().
		SetNillableName(req.Data.Name).
		SetNillableCode(req.Data.Code).
		SetNillableParentID(req.Data.ParentId).
		SetNillableSortID(req.Data.SortId).
		SetNillableEnable(req.Data.Enable).
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
		return internalMessageV1.ErrorInternalServerError("insert data failed")
	}

	return nil
}

func (r *NotificationMessageCategoryRepo) Update(ctx context.Context, req *internalMessageV1.UpdateNotificationMessageCategoryRequest) error {
	if req == nil || req.Data == nil {
		return internalMessageV1.ErrorBadRequest("invalid parameter")
	}

	// 如果不存在则创建
	if req.GetAllowMissing() {
		exist, err := r.IsExist(ctx, req.GetData().GetId())
		if err != nil {
			return err
		}
		if !exist {
			createReq := &internalMessageV1.CreateNotificationMessageCategoryRequest{Data: req.Data}
			createReq.Data.CreateBy = createReq.Data.UpdateBy
			createReq.Data.UpdateBy = nil
			return r.Create(ctx, createReq)
		}
	}

	if req.UpdateMask != nil {
		req.UpdateMask.Normalize()
		if !req.UpdateMask.IsValid(req.Data) {
			r.log.Errorf("invalid field mask [%v]", req.UpdateMask)
			return internalMessageV1.ErrorBadRequest("invalid field mask")
		}
		fieldmaskutil.Filter(req.GetData(), req.UpdateMask.GetPaths())
	}

	builder := r.data.db.Client().NotificationMessageCategory.UpdateOneID(req.Data.GetId()).
		SetNillableName(req.Data.Name).
		SetNillableCode(req.Data.Code).
		SetNillableParentID(req.Data.ParentId).
		SetNillableSortID(req.Data.SortId).
		SetNillableEnable(req.Data.Enable).
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
		return internalMessageV1.ErrorInternalServerError("update data failed")
	}

	return nil
}

func (r *NotificationMessageCategoryRepo) Delete(ctx context.Context, req *internalMessageV1.DeleteNotificationMessageCategoryRequest) error {
	if req == nil {
		return internalMessageV1.ErrorBadRequest("invalid parameter")
	}

	ids, err := queryAllChildrenIDs(ctx, r.data.db, "notification_message_categories", req.GetId())
	if err != nil {
		r.log.Errorf("query child notification message categories failed: %s", err.Error())
		return internalMessageV1.ErrorInternalServerError("query child notification message categories failed")
	}
	ids = append(ids, req.GetId())

	//r.log.Info("notification message category ids to delete: ", ids)

	if _, err = r.data.db.Client().NotificationMessageCategory.Delete().
		Where(notificationmessagecategory.IDIn(ids...)).
		Exec(ctx); err != nil {
		r.log.Errorf("delete notification message categories failed: %s", err.Error())
		return internalMessageV1.ErrorInternalServerError("delete notification message categories failed")
	}

	return nil
}
