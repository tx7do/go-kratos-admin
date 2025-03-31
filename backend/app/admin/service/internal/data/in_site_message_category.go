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
	timeutil "github.com/tx7do/go-utils/timeutil"
	"github.com/tx7do/go-utils/trans"

	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	systemV1 "kratos-admin/api/gen/go/system/service/v1"
	"kratos-admin/app/admin/service/internal/data/ent"
	"kratos-admin/app/admin/service/internal/data/ent/insitemessagecategory"
)

type InSiteMessageCategoryRepo struct {
	data *Data
	log  *log.Helper
}

func NewInSiteMessageCategoryRepo(data *Data, logger log.Logger) *InSiteMessageCategoryRepo {
	l := log.NewHelper(log.With(logger, "module", "in-site-message-category/repo/admin-service"))
	return &InSiteMessageCategoryRepo{
		data: data,
		log:  l,
	}
}

func (r *InSiteMessageCategoryRepo) convertEntToProto(in *ent.InSiteMessageCategory) *systemV1.InSiteMessageCategory {
	if in == nil {
		return nil
	}
	return &systemV1.InSiteMessageCategory{
		Id:         trans.Ptr(in.ID),
		Name:       in.Name,
		Code:       in.Code,
		SortId:     in.SortID,
		Enable:     in.Enable,
		ParentId:   in.ParentID,
		CreateTime: timeutil.TimeToTimestamppb(in.CreateTime),
		UpdateTime: timeutil.TimeToTimestamppb(in.UpdateTime),
		DeleteTime: timeutil.TimeToTimestamppb(in.DeleteTime),
	}
}

func (r *InSiteMessageCategoryRepo) travelChild(nodes []*systemV1.InSiteMessageCategory, node *systemV1.InSiteMessageCategory) bool {
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

func (r *InSiteMessageCategoryRepo) Count(ctx context.Context, whereCond []func(s *sql.Selector)) (int, error) {
	builder := r.data.db.Client().InSiteMessageCategory.Query()
	if len(whereCond) != 0 {
		builder.Modify(whereCond...)
	}

	count, err := builder.Count(ctx)
	if err != nil {
		r.log.Errorf("query count failed: %s", err.Error())
	}

	return count, err
}

func (r *InSiteMessageCategoryRepo) List(ctx context.Context, req *pagination.PagingRequest) (*systemV1.ListInSiteMessageCategoryResponse, error) {
	builder := r.data.db.Client().InSiteMessageCategory.Query()

	err, whereSelectors, querySelectors := entgo.BuildQuerySelector(
		req.GetQuery(), req.GetOrQuery(),
		req.GetPage(), req.GetPageSize(), req.GetNoPaging(),
		req.GetOrderBy(), insitemessagecategory.FieldCreateTime,
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

	items := make([]*systemV1.InSiteMessageCategory, 0, len(results))
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

	ret := systemV1.ListInSiteMessageCategoryResponse{
		Total: uint32(count),
		Items: items,
	}

	return &ret, err
}

func (r *InSiteMessageCategoryRepo) IsExist(ctx context.Context, id uint32) (bool, error) {
	return r.data.db.Client().InSiteMessageCategory.Query().
		Where(insitemessagecategory.IDEQ(id)).
		Exist(ctx)
}

func (r *InSiteMessageCategoryRepo) Get(ctx context.Context, req *systemV1.GetInSiteMessageCategoryRequest) (*systemV1.InSiteMessageCategory, error) {
	ret, err := r.data.db.Client().InSiteMessageCategory.Get(ctx, req.GetId())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, systemV1.ErrorResourceNotFound("message category not found")
		}

		return nil, err
	}

	return r.convertEntToProto(ret), err
}

func (r *InSiteMessageCategoryRepo) Create(ctx context.Context, req *systemV1.CreateInSiteMessageCategoryRequest) error {
	if req.Data == nil {
		return errors.New("invalid request")
	}

	builder := r.data.db.Client().InSiteMessageCategory.Create().
		SetNillableName(req.Data.Name).
		SetNillableCode(req.Data.Code).
		SetNillableParentID(req.Data.ParentId).
		SetNillableSortID(req.Data.SortId).
		SetNillableEnable(req.Data.Enable).
		SetNillableCreateBy(req.OperatorId).
		SetNillableCreateTime(timeutil.TimestamppbToTime(req.Data.CreateTime))

	if req.Data.CreateTime == nil {
		builder.SetCreateTime(time.Now())
	}

	err := builder.Exec(ctx)
	if err != nil {
		r.log.Errorf("insert one data failed: %s", err.Error())
		return err
	}

	return err
}

func (r *InSiteMessageCategoryRepo) Update(ctx context.Context, req *systemV1.UpdateInSiteMessageCategoryRequest) error {
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
			return r.Create(ctx, &systemV1.CreateInSiteMessageCategoryRequest{Data: req.Data, OperatorId: req.OperatorId})
		}
	}

	if req.UpdateMask != nil {
		req.UpdateMask.Normalize()
		if !req.UpdateMask.IsValid(req.Data) {
			return errors.New("invalid field mask")
		}
		fieldmaskutil.Filter(req.GetData(), req.UpdateMask.GetPaths())
	}

	builder := r.data.db.Client().InSiteMessageCategory.UpdateOneID(req.Data.GetId()).
		SetNillableName(req.Data.Name).
		SetNillableCode(req.Data.Code).
		SetNillableParentID(req.Data.ParentId).
		SetNillableSortID(req.Data.SortId).
		SetNillableEnable(req.Data.Enable).
		SetNillableUpdateBy(req.OperatorId).
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

	err := builder.Exec(ctx)
	if err != nil {
		r.log.Errorf("update one data failed: %s", err.Error())
		return err
	}

	return err
}

func (r *InSiteMessageCategoryRepo) Delete(ctx context.Context, req *systemV1.DeleteInSiteMessageCategoryRequest) (bool, error) {
	err := r.data.db.Client().InSiteMessageCategory.
		DeleteOneID(req.GetId()).
		Exec(ctx)
	return err != nil, err
}
