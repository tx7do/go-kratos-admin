package data

import (
	"context"
	"errors"
	"sort"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"

	entgo "github.com/tx7do/go-utils/entgo/query"
	entgoUpdate "github.com/tx7do/go-utils/entgo/update"
	"github.com/tx7do/go-utils/fieldmaskutil"
	"github.com/tx7do/go-utils/timeutil"
	"github.com/tx7do/go-utils/trans"
	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"

	"kratos-admin/app/admin/service/internal/data/ent"
	"kratos-admin/app/admin/service/internal/data/ent/notificationmessagecategory"

	internalMessageV1 "kratos-admin/api/gen/go/internal_message/service/v1"

	"kratos-admin/pkg/middleware/auth"
)

type NotificationMessageCategoryRepo struct {
	data *Data
	log  *log.Helper
}

func NewNotificationMessageCategoryRepo(data *Data, logger log.Logger) *NotificationMessageCategoryRepo {
	l := log.NewHelper(log.With(logger, "module", "notification-message-category/repo/admin-service"))
	return &NotificationMessageCategoryRepo{
		data: data,
		log:  l,
	}
}

func (r *NotificationMessageCategoryRepo) convertEntToProto(in *ent.NotificationMessageCategory) *internalMessageV1.NotificationMessageCategory {
	if in == nil {
		return nil
	}
	return &internalMessageV1.NotificationMessageCategory{
		Id:         trans.Ptr(in.ID),
		Name:       in.Name,
		Code:       in.Code,
		SortId:     in.SortID,
		Enable:     in.Enable,
		ParentId:   in.ParentID,
		CreateBy:   in.CreateBy,
		UpdateBy:   in.UpdateBy,
		CreateTime: timeutil.TimeToTimeString(in.CreateTime),
		UpdateTime: timeutil.TimeToTimeString(in.UpdateTime),
		DeleteTime: timeutil.TimeToTimeString(in.DeleteTime),
	}
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
	}

	return count, err
}

func (r *NotificationMessageCategoryRepo) List(ctx context.Context, req *pagination.PagingRequest) (*internalMessageV1.ListNotificationMessageCategoryResponse, error) {
	builder := r.data.db.Client().NotificationMessageCategory.Query()

	err, whereSelectors, querySelectors := entgo.BuildQuerySelector(
		req.GetQuery(), req.GetOrQuery(),
		req.GetPage(), req.GetPageSize(), req.GetNoPaging(),
		req.GetOrderBy(), notificationmessagecategory.FieldCreateTime,
		req.GetFieldMask().GetPaths(),
	)
	if err != nil {
		r.log.Errorf("解析SELECT条件发生错误[%s]", err.Error())
		return nil, err
	}

	if querySelectors != nil {
		builder.Modify(querySelectors...)
	}

	results, err := builder.All(ctx)
	if err != nil {
		return nil, err
	}

	sort.SliceStable(results, func(i, j int) bool {
		if results[j].ParentID == nil {
			return true
		}
		if results[i].ParentID == nil {
			return true
		}
		return *results[i].ParentID < *results[j].ParentID
	})

	items := make([]*internalMessageV1.NotificationMessageCategory, 0, len(results))
	for _, m := range results {
		if m.ParentID == nil {
			item := r.convertEntToProto(m)
			items = append(items, item)
		}
	}
	for _, m := range results {
		if m.ParentID != nil {
			item := r.convertEntToProto(m)

			if r.travelChild(items, item) {
				continue
			}

			items = append(items, item)
		}
	}

	count, err := r.Count(ctx, whereSelectors)
	if err != nil {
		return nil, err
	}

	ret := internalMessageV1.ListNotificationMessageCategoryResponse{
		Total: uint32(count),
		Items: items,
	}

	return &ret, err
}

func (r *NotificationMessageCategoryRepo) IsExist(ctx context.Context, id uint32) (bool, error) {
	return r.data.db.Client().NotificationMessageCategory.Query().
		Where(notificationmessagecategory.IDEQ(id)).
		Exist(ctx)
}

func (r *NotificationMessageCategoryRepo) Get(ctx context.Context, req *internalMessageV1.GetNotificationMessageCategoryRequest) (*internalMessageV1.NotificationMessageCategory, error) {
	ret, err := r.data.db.Client().NotificationMessageCategory.Get(ctx, req.GetId())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, internalMessageV1.ErrorResourceNotFound("message category not found")
		}

		return nil, err
	}

	return r.convertEntToProto(ret), err
}

func (r *NotificationMessageCategoryRepo) Create(ctx context.Context, req *internalMessageV1.CreateNotificationMessageCategoryRequest, operator *auth.UserTokenPayload) error {
	if req.Data == nil {
		return errors.New("invalid request")
	}

	builder := r.data.db.Client().NotificationMessageCategory.Create().
		SetNillableName(req.Data.Name).
		SetNillableCode(req.Data.Code).
		SetNillableParentID(req.Data.ParentId).
		SetNillableSortID(req.Data.SortId).
		SetNillableEnable(req.Data.Enable).
		SetNillableCreateTime(timeutil.StringTimeToTime(req.Data.CreateTime))

	if req.Data.CreateTime == nil {
		builder.SetCreateTime(time.Now())
	}

	if operator != nil {
		builder.SetCreateBy(operator.UserId)
	}

	if req.Data.Id != nil {
		builder.SetID(req.Data.GetId())
	}

	if err := builder.Exec(ctx); err != nil {
		r.log.Errorf("insert one data failed: %s", err.Error())
		return err
	}

	return nil
}

func (r *NotificationMessageCategoryRepo) Update(ctx context.Context, req *internalMessageV1.UpdateNotificationMessageCategoryRequest, operator *auth.UserTokenPayload) error {
	if req.Data == nil {
		return errors.New("invalid request")
	}

	// 如果不存在则创建
	if req.GetAllowMissing() {
		exist, err := r.IsExist(ctx, req.GetData().GetId())
		if err != nil {
			return err
		}
		if !exist {
			return r.Create(ctx, &internalMessageV1.CreateNotificationMessageCategoryRequest{Data: req.Data}, operator)
		}
	}

	if req.UpdateMask != nil {
		req.UpdateMask.Normalize()
		if !req.UpdateMask.IsValid(req.Data) {
			return errors.New("invalid field mask")
		}
		fieldmaskutil.Filter(req.GetData(), req.UpdateMask.GetPaths())
	}

	builder := r.data.db.Client().NotificationMessageCategory.UpdateOneID(req.Data.GetId()).
		SetNillableName(req.Data.Name).
		SetNillableCode(req.Data.Code).
		SetNillableParentID(req.Data.ParentId).
		SetNillableSortID(req.Data.SortId).
		SetNillableEnable(req.Data.Enable).
		SetNillableUpdateTime(timeutil.StringTimeToTime(req.Data.UpdateTime))

	if req.Data.UpdateTime == nil {
		builder.SetUpdateTime(time.Now())
	}

	if operator != nil {
		builder.SetUpdateBy(operator.UserId)
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
		return err
	}

	return nil
}

func (r *NotificationMessageCategoryRepo) Delete(ctx context.Context, req *internalMessageV1.DeleteNotificationMessageCategoryRequest) (bool, error) {
	err := r.data.db.Client().NotificationMessageCategory.
		DeleteOneID(req.GetId()).
		Exec(ctx)
	return err != nil, err
}
