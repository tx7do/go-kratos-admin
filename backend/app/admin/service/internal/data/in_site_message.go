package data

import (
	"context"
	"errors"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"
	entgo "github.com/tx7do/go-utils/entgo/query"
	entgoUpdate "github.com/tx7do/go-utils/entgo/update"
	"github.com/tx7do/go-utils/fieldmaskutil"
	timeutil "github.com/tx7do/go-utils/timeutil"
	"github.com/tx7do/go-utils/trans"
	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"

	"kratos-admin/app/admin/service/internal/data/ent"
	"kratos-admin/app/admin/service/internal/data/ent/insitemessage"

	systemV1 "kratos-admin/api/gen/go/system/service/v1"
)

type InSiteMessageRepo struct {
	data *Data
	log  *log.Helper
}

func NewInSiteMessageRepo(data *Data, logger log.Logger) *InSiteMessageRepo {
	l := log.NewHelper(log.With(logger, "module", "in-site-message/repo/admin-service"))
	return &InSiteMessageRepo{
		data: data,
		log:  l,
	}
}

func (r *InSiteMessageRepo) toProtoStatus(in *insitemessage.Status) *systemV1.MessageStatus {
	if in == nil {
		return nil
	}

	switch *in {
	case insitemessage.StatusUnknown:
		return trans.Ptr(systemV1.MessageStatus_MessageStatus_Unknown)

	case insitemessage.StatusDraft:
		return trans.Ptr(systemV1.MessageStatus_MessageStatus_Draft)

	case insitemessage.StatusPublished:
		return trans.Ptr(systemV1.MessageStatus_MessageStatus_Published)

	case insitemessage.StatusScheduled:
		return trans.Ptr(systemV1.MessageStatus_MessageStatus_Scheduled)

	case insitemessage.StatusRevoked:
		return trans.Ptr(systemV1.MessageStatus_MessageStatus_Revoked)

	case insitemessage.StatusArchived:
		return trans.Ptr(systemV1.MessageStatus_MessageStatus_Archived)

	default:
		return nil
	}
}

func (r *InSiteMessageRepo) toEntStatus(in *systemV1.MessageStatus) *insitemessage.Status {
	if in == nil {
		return nil
	}

	switch *in {
	case systemV1.MessageStatus_MessageStatus_Unknown:
		return trans.Ptr(insitemessage.StatusUnknown)

	case systemV1.MessageStatus_MessageStatus_Draft:
		return trans.Ptr(insitemessage.StatusDraft)

	case systemV1.MessageStatus_MessageStatus_Published:
		return trans.Ptr(insitemessage.StatusPublished)

	case systemV1.MessageStatus_MessageStatus_Scheduled:
		return trans.Ptr(insitemessage.StatusScheduled)

	case systemV1.MessageStatus_MessageStatus_Revoked:
		return trans.Ptr(insitemessage.StatusRevoked)

	case systemV1.MessageStatus_MessageStatus_Archived:
		return trans.Ptr(insitemessage.StatusArchived)

	default:
		return nil
	}
}

func (r *InSiteMessageRepo) convertEntToProto(in *ent.InSiteMessage) *systemV1.InSiteMessage {
	if in == nil {
		return nil
	}
	return &systemV1.InSiteMessage{
		Id:         trans.Ptr(in.ID),
		Title:      in.Title,
		Content:    in.Content,
		CategoryId: in.CategoryID,
		Status:     r.toProtoStatus(in.Status),
		CreateTime: timeutil.TimeToTimestamppb(in.CreateTime),
		UpdateTime: timeutil.TimeToTimestamppb(in.UpdateTime),
		DeleteTime: timeutil.TimeToTimestamppb(in.DeleteTime),
	}
}

func (r *InSiteMessageRepo) Count(ctx context.Context, whereCond []func(s *sql.Selector)) (int, error) {
	builder := r.data.db.Client().InSiteMessage.Query()
	if len(whereCond) != 0 {
		builder.Modify(whereCond...)
	}

	count, err := builder.Count(ctx)
	if err != nil {
		r.log.Errorf("query count failed: %s", err.Error())
	}

	return count, err
}

func (r *InSiteMessageRepo) List(ctx context.Context, req *pagination.PagingRequest) (*systemV1.ListInSiteMessageResponse, error) {
	builder := r.data.db.Client().InSiteMessage.Query()

	err, whereSelectors, querySelectors := entgo.BuildQuerySelector(
		req.GetQuery(), req.GetOrQuery(),
		req.GetPage(), req.GetPageSize(), req.GetNoPaging(),
		req.GetOrderBy(), insitemessage.FieldCreateTime,
		req.GetFieldMask().GetPaths(),
	)
	if err != nil {
		r.log.Errorf("解析条件发生错误[%s]", err.Error())
		return nil, err
	}

	if querySelectors != nil {
		builder.Modify(querySelectors...)
	}

	results, err := builder.All(ctx)
	if err != nil {
		return nil, err
	}

	items := make([]*systemV1.InSiteMessage, 0, len(results))
	for _, res := range results {
		item := r.convertEntToProto(res)
		items = append(items, item)
	}

	count, err := r.Count(ctx, whereSelectors)
	if err != nil {
		return nil, err
	}

	return &systemV1.ListInSiteMessageResponse{
		Total: uint32(count),
		Items: items,
	}, err
}

func (r *InSiteMessageRepo) IsExist(ctx context.Context, id uint32) (bool, error) {
	return r.data.db.Client().InSiteMessage.Query().
		Where(insitemessage.IDEQ(id)).
		Exist(ctx)
}

func (r *InSiteMessageRepo) Get(ctx context.Context, req *systemV1.GetInSiteMessageRequest) (*systemV1.InSiteMessage, error) {
	ret, err := r.data.db.Client().InSiteMessage.Get(ctx, req.GetId())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, systemV1.ErrorResourceNotFound("message not found")
		}
		return nil, err
	}

	return r.convertEntToProto(ret), err
}

func (r *InSiteMessageRepo) Create(ctx context.Context, req *systemV1.CreateInSiteMessageRequest) error {
	if req.Data == nil {
		return errors.New("invalid request")
	}

	builder := r.data.db.Client().InSiteMessage.Create().
		SetNillableTitle(req.Data.Title).
		SetNillableContent(req.Data.Content).
		SetNillableCategoryID(req.Data.CategoryId).
		SetNillableStatus(r.toEntStatus(req.Data.Status)).
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

func (r *InSiteMessageRepo) Update(ctx context.Context, req *systemV1.UpdateInSiteMessageRequest) error {
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
			return r.Create(ctx, &systemV1.CreateInSiteMessageRequest{Data: req.Data, OperatorId: req.OperatorId})
		}
	}

	if req.UpdateMask != nil {
		req.UpdateMask.Normalize()
		if !req.UpdateMask.IsValid(req.Data) {
			return errors.New("invalid field mask")
		}
		fieldmaskutil.Filter(req.GetData(), req.UpdateMask.GetPaths())
	}

	builder := r.data.db.Client().InSiteMessage.UpdateOneID(req.Data.GetId()).
		SetNillableTitle(req.Data.Title).
		SetNillableContent(req.Data.Content).
		SetNillableCategoryID(req.Data.CategoryId).
		SetNillableStatus(r.toEntStatus(req.Data.Status)).
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

func (r *InSiteMessageRepo) Delete(ctx context.Context, req *systemV1.DeleteInSiteMessageRequest) (bool, error) {
	err := r.data.db.Client().InSiteMessage.
		DeleteOneID(req.GetId()).
		Exec(ctx)
	return err != nil, err
}
