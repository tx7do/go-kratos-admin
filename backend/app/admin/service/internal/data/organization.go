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

	"kratos-admin/app/admin/service/internal/data/ent"
	"kratos-admin/app/admin/service/internal/data/ent/organization"

	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	userV1 "kratos-admin/api/gen/go/user/service/v1"
)

type OrganizationRepo struct {
	data *Data
	log  *log.Helper
}

func NewOrganizationRepo(data *Data, logger log.Logger) *OrganizationRepo {
	l := log.NewHelper(log.With(logger, "module", "organization/repo/admin-service"))
	return &OrganizationRepo{
		data: data,
		log:  l,
	}
}

func (r *OrganizationRepo) convertEntToProto(in *ent.Organization) *userV1.Organization {
	if in == nil {
		return nil
	}
	return &userV1.Organization{
		Id:         trans.Ptr(in.ID),
		Name:       in.Name,
		Remark:     in.Remark,
		SortId:     in.SortID,
		ParentId:   in.ParentID,
		Status:     (*string)(in.Status),
		CreateTime: timeutil.TimeToTimestamppb(in.CreateTime),
		UpdateTime: timeutil.TimeToTimestamppb(in.UpdateTime),
		DeleteTime: timeutil.TimeToTimestamppb(in.DeleteTime),
	}
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
	}

	return count, err
}

func (r *OrganizationRepo) List(ctx context.Context, req *pagination.PagingRequest) (*userV1.ListOrganizationResponse, error) {
	builder := r.data.db.Client().Organization.Query()

	err, whereSelectors, querySelectors := entgo.BuildQuerySelector(
		req.GetQuery(), req.GetOrQuery(),
		req.GetPage(), req.GetPageSize(), req.GetNoPaging(),
		req.GetOrderBy(), organization.FieldCreateTime,
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

	items := make([]*userV1.Organization, 0, len(results))
	for _, m := range results {
		if m.ParentID == nil {
			item := r.convertEntToProto(m)
			items = append(items, item)
		}
	}
	for _, m := range results {
		if m.ParentID != nil {
			item := r.convertEntToProto(m)
			r.travelChild(items, item)
		}
	}

	count, err := r.Count(ctx, whereSelectors)
	if err != nil {
		return nil, err
	}

	ret := userV1.ListOrganizationResponse{
		Total: int32(count),
		Items: items,
	}

	return &ret, err
}

func (r *OrganizationRepo) IsExist(ctx context.Context, id uint32) (bool, error) {
	return r.data.db.Client().Organization.Query().
		Where(organization.IDEQ(id)).
		Exist(ctx)
}

func (r *OrganizationRepo) Get(ctx context.Context, req *userV1.GetOrganizationRequest) (*userV1.Organization, error) {
	ret, err := r.data.db.Client().Organization.Get(ctx, req.GetId())
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	return r.convertEntToProto(ret), err
}

func (r *OrganizationRepo) Create(ctx context.Context, req *userV1.CreateOrganizationRequest) error {
	if req.Data == nil {
		return errors.New("invalid request")
	}

	builder := r.data.db.Client().Organization.Create().
		SetNillableName(req.Data.Name).
		SetNillableParentID(req.Data.ParentId).
		SetNillableSortID(req.Data.SortId).
		SetNillableRemark(req.Data.Remark).
		SetNillableStatus((*organization.Status)(req.Data.Status)).
		SetNillableCreateBy(req.OperatorId)

	if req.Data.CreateTime == nil {
		builder.SetCreateTime(time.Now())
	} else {
		builder.SetCreateTime(*timeutil.TimestamppbToTime(req.Data.CreateTime))
	}

	err := builder.Exec(ctx)
	if err != nil {
		r.log.Errorf("insert one data failed: %s", err.Error())
		return err
	}

	return err
}

func (r *OrganizationRepo) Update(ctx context.Context, req *userV1.UpdateOrganizationRequest) error {
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
			return r.Create(ctx, &userV1.CreateOrganizationRequest{Data: req.Data, OperatorId: req.OperatorId})
		}
	}

	if req.UpdateMask != nil {
		req.UpdateMask.Normalize()
		if !req.UpdateMask.IsValid(req.Data) {
			return errors.New("invalid field mask")
		}
		fieldmaskutil.Filter(req.GetData(), req.UpdateMask.GetPaths())
	}

	builder := r.data.db.Client().Organization.
		UpdateOneID(req.Data.GetId()).
		SetNillableName(req.Data.Name).
		SetNillableParentID(req.Data.ParentId).
		SetNillableSortID(req.Data.SortId).
		SetNillableRemark(req.Data.Remark).
		SetNillableStatus((*organization.Status)(req.Data.Status)).
		SetNillableUpdateBy(req.OperatorId)

	if req.Data.UpdateTime == nil {
		builder.SetUpdateTime(time.Now())
	} else {
		builder.SetUpdateTime(*timeutil.TimestamppbToTime(req.Data.UpdateTime))
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

func (r *OrganizationRepo) Delete(ctx context.Context, req *userV1.DeleteOrganizationRequest) (bool, error) {
	err := r.data.db.Client().Organization.
		DeleteOneID(req.GetId()).
		Exec(ctx)
	return err != nil, err
}
