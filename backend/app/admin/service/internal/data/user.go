package data

import (
	"context"
	"errors"
	"time"

	"encoding/base64"

	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"

	"github.com/tx7do/go-utils/crypto"
	entgo "github.com/tx7do/go-utils/entgo/query"
	entgoUpdate "github.com/tx7do/go-utils/entgo/update"
	"github.com/tx7do/go-utils/fieldmaskutil"
	"github.com/tx7do/go-utils/timeutil"
	"github.com/tx7do/go-utils/trans"
	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"

	"kratos-admin/app/admin/service/internal/data/ent"
	"kratos-admin/app/admin/service/internal/data/ent/user"

	userV1 "kratos-admin/api/gen/go/user/service/v1"

	"kratos-admin/pkg/middleware/auth"
)

type UserRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) *UserRepo {
	l := log.NewHelper(log.With(logger, "module", "user/repo/admin-service"))
	return &UserRepo{
		data: data,
		log:  l,
	}
}

func (r *UserRepo) toEntAuthority(authority *userV1.UserAuthority) *user.Authority {
	if authority == nil {
		return nil
	}
	find, ok := userV1.UserAuthority_name[int32(*authority)]
	if !ok {
		return nil
	}
	return (*user.Authority)(trans.Ptr(find))
}

func (r *UserRepo) toProtoAuthority(authority *user.Authority) *userV1.UserAuthority {
	if authority == nil {
		return nil
	}
	find, ok := userV1.UserAuthority_value[string(*authority)]
	if !ok {
		return nil
	}
	return (*userV1.UserAuthority)(trans.Ptr(find))
}

func (r *UserRepo) toEntGender(gender *userV1.UserGender) *user.Gender {
	if gender == nil {
		return nil
	}
	find, ok := userV1.UserGender_name[int32(*gender)]
	if !ok {
		return nil
	}
	return (*user.Gender)(trans.Ptr(find))
}

func (r *UserRepo) toProtoGender(gender *user.Gender) *userV1.UserGender {
	if gender == nil {
		return nil
	}
	find, ok := userV1.UserGender_value[string(*gender)]
	if !ok {
		return nil
	}
	return (*userV1.UserGender)(trans.Ptr(find))
}

func (r *UserRepo) toEntStatus(status *userV1.UserStatus) *user.Status {
	if status == nil {
		return nil
	}
	find, ok := userV1.UserStatus_name[int32(*status)]
	if !ok {
		return nil
	}
	return (*user.Status)(trans.Ptr(find))
}

func (r *UserRepo) toProtoStatus(status *user.Status) *userV1.UserStatus {
	if status == nil {
		return nil
	}
	find, ok := userV1.UserStatus_value[string(*status)]
	if !ok {
		return nil
	}
	return (*userV1.UserStatus)(trans.Ptr(find))
}

func (r *UserRepo) convertEntToProto(in *ent.User) *userV1.User {
	if in == nil {
		return nil
	}

	return &userV1.User{
		Id:            trans.Ptr(in.ID),
		RoleId:        in.RoleID,
		WorkId:        in.WorkID,
		OrgId:         in.OrgID,
		PositionId:    in.PositionID,
		TenantId:      in.TenantID,
		UserName:      in.Username,
		NickName:      in.NickName,
		RealName:      in.RealName,
		Email:         in.Email,
		Avatar:        in.Avatar,
		Telephone:     in.Telephone,
		Mobile:        in.Mobile,
		Address:       in.Address,
		Region:        in.Region,
		Description:   in.Description,
		Remark:        in.Remark,
		LastLoginTime: in.LastLoginTime,
		LastLoginIp:   in.LastLoginIP,
		CreateBy:      in.CreateBy,
		UpdateBy:      in.UpdateBy,
		Gender:        r.toProtoGender(in.Gender),
		Authority:     r.toProtoAuthority(in.Authority),
		Status:        r.toProtoStatus(in.Status),
		CreateTime:    timeutil.TimeToTimeString(in.CreateTime),
		UpdateTime:    timeutil.TimeToTimeString(in.UpdateTime),
		DeleteTime:    timeutil.TimeToTimeString(in.DeleteTime),
	}
}

func (r *UserRepo) Count(ctx context.Context, whereCond []func(s *sql.Selector)) (int, error) {
	builder := r.data.db.Client().User.Query()
	if len(whereCond) != 0 {
		builder.Modify(whereCond...)
	}

	count, err := builder.Count(ctx)
	if err != nil {
		r.log.Errorf("query count failed: %s", err.Error())
	}

	return count, err
}

func (r *UserRepo) List(ctx context.Context, req *pagination.PagingRequest) (*userV1.ListUserResponse, error) {
	builder := r.data.db.Client().User.Query()

	err, whereSelectors, querySelectors := entgo.BuildQuerySelector(
		req.GetQuery(), req.GetOrQuery(),
		req.GetPage(), req.GetPageSize(), req.GetNoPaging(),
		req.GetOrderBy(), user.FieldCreateTime,
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
		r.log.Errorf("query list failed: %s", err.Error())
		return nil, err
	}

	items := make([]*userV1.User, 0, len(results))
	for _, res := range results {
		item := r.convertEntToProto(res)
		items = append(items, item)
	}

	count, err := r.Count(ctx, whereSelectors)
	if err != nil {
		return nil, err
	}

	return &userV1.ListUserResponse{
		Total: uint32(count),
		Items: items,
	}, nil
}

func (r *UserRepo) IsExist(ctx context.Context, id uint32) (bool, error) {
	return r.data.db.Client().User.Query().
		Where(user.IDEQ(id)).
		Exist(ctx)
}

func (r *UserRepo) Get(ctx context.Context, userId uint32) (*userV1.User, error) {
	ret, err := r.data.db.Client().User.Get(ctx, userId)
	if err != nil {
		r.log.Errorf("query one data failed: %s", err.Error())

		if ent.IsNotFound(err) {
			return nil, userV1.ErrorUserNotFound("user not found")
		}

		return nil, err
	}

	u := r.convertEntToProto(ret)

	return u, err
}

func (r *UserRepo) Create(ctx context.Context, req *userV1.CreateUserRequest, operator *auth.UserTokenPayload) error {
	if req.Data == nil {
		return errors.New("invalid request")
	}

	builder := r.data.db.Client().User.Create().
		SetNillableUsername(req.Data.UserName).
		SetNillableNickName(req.Data.NickName).
		SetNillableEmail(req.Data.Email).
		SetNillableRealName(req.Data.RealName).
		SetNillableEmail(req.Data.Email).
		SetNillableTelephone(req.Data.Telephone).
		SetNillableMobile(req.Data.Mobile).
		SetNillableAvatar(req.Data.Avatar).
		SetNillableRegion(req.Data.Region).
		SetNillableAddress(req.Data.Address).
		SetNillableDescription(req.Data.Description).
		SetNillableRemark(req.Data.Remark).
		SetNillableLastLoginTime(req.Data.LastLoginTime).
		SetNillableLastLoginIP(req.Data.LastLoginIp).
		SetNillableStatus(r.toEntStatus(req.Data.Status)).
		SetNillableGender(r.toEntGender(req.Data.Gender)).
		SetNillableAuthority(r.toEntAuthority(req.Data.Authority)).
		SetNillableOrgID(req.Data.OrgId).
		SetNillableRoleID(req.Data.RoleId).
		SetNillableWorkID(req.Data.WorkId).
		SetNillablePositionID(req.Data.PositionId).
		SetNillableTenantID(req.Data.TenantId).
		SetNillableCreateTime(timeutil.StringTimeToTime(req.Data.CreateTime))

	if len(req.GetPassword()) > 0 {
		cryptoPassword, err := crypto.HashPassword(req.GetPassword())
		if err == nil {
			builder.SetPassword(cryptoPassword)
		}
	}

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

func (r *UserRepo) Update(ctx context.Context, req *userV1.UpdateUserRequest, operator *auth.UserTokenPayload) error {
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
			return r.Create(ctx, &userV1.CreateUserRequest{Data: req.Data}, operator)
		}
	}

	if req.UpdateMask != nil {
		req.UpdateMask.Normalize()
		if !req.UpdateMask.IsValid(req.Data) {
			return errors.New("invalid field mask")
		}
		fieldmaskutil.Filter(req.GetData(), req.UpdateMask.GetPaths())
	}

	builder := r.data.db.Client().User.
		UpdateOneID(req.Data.GetId()).
		SetNillableNickName(req.Data.NickName).
		SetNillableEmail(req.Data.Email).
		SetNillableRealName(req.Data.RealName).
		SetNillableEmail(req.Data.Email).
		SetNillableTelephone(req.Data.Telephone).
		SetNillableMobile(req.Data.Mobile).
		SetNillableAvatar(req.Data.Avatar).
		SetNillableRegion(req.Data.Region).
		SetNillableAddress(req.Data.Address).
		SetNillableDescription(req.Data.Description).
		SetNillableRemark(req.Data.Remark).
		SetNillableLastLoginTime(req.Data.LastLoginTime).
		SetNillableLastLoginIP(req.Data.LastLoginIp).
		SetNillableStatus(r.toEntStatus(req.Data.Status)).
		SetNillableGender(r.toEntGender(req.Data.Gender)).
		SetNillableAuthority(r.toEntAuthority(req.Data.Authority)).
		SetNillableOrgID(req.Data.OrgId).
		SetNillableRoleID(req.Data.RoleId).
		SetNillableWorkID(req.Data.WorkId).
		SetNillablePositionID(req.Data.PositionId).
		SetNillableUpdateTime(timeutil.StringTimeToTime(req.Data.UpdateTime))

	if req.Data.UpdateTime == nil {
		builder.SetUpdateTime(time.Now())
	}

	if operator != nil {
		builder.SetUpdateBy(operator.UserId)
	}

	if len(req.GetPassword()) > 0 {
		cryptoPassword, err := crypto.HashPassword(req.GetPassword())
		if err == nil {
			builder.SetPassword(cryptoPassword)
		}
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

func (r *UserRepo) Delete(ctx context.Context, userId uint32) (bool, error) {
	err := r.data.db.Client().User.
		DeleteOneID(userId).
		Exec(ctx)
	if err != nil {
		r.log.Errorf("delete one data failed: %s", err.Error())
	}

	return err == nil, err
}

func (r *UserRepo) GetUserByUserName(ctx context.Context, userName string) (*userV1.User, error) {
	ret, err := r.data.db.Client().User.Query().
		Where(user.UsernameEQ(userName)).
		Only(ctx)
	if err != nil {
		r.log.Errorf("query user data failed: %s", err.Error())

		if ent.IsNotFound(err) {
			return nil, userV1.ErrorUserNotFound("user not found")
		}

		return nil, err
	}

	u := r.convertEntToProto(ret)
	return u, err
}

func (r *UserRepo) VerifyPassword(ctx context.Context, req *userV1.VerifyPasswordRequest) (*userV1.VerifyPasswordResponse, error) {
	ret, err := r.data.db.Client().User.
		Query().
		Select(user.FieldID, user.FieldPassword).
		Where(user.UsernameEQ(req.GetUserName())).
		Only(ctx)
	if err != nil {
		return &userV1.VerifyPasswordResponse{
			Result: userV1.VerifyPasswordResult_ACCOUNT_NOT_EXISTS,
		}, userV1.ErrorUserNotFound("用户未找到")
	}

	// 解密密码
	bytesPass, err := base64.StdEncoding.DecodeString(req.GetPassword())
	plainPassword, _ := crypto.AesDecrypt(bytesPass, crypto.DefaultAESKey, nil)

	// 校验密码
	bMatched := crypto.VerifyPassword(string(plainPassword), *ret.Password)

	if !bMatched {
		return &userV1.VerifyPasswordResponse{
			Result: userV1.VerifyPasswordResult_WRONG_PASSWORD,
		}, userV1.ErrorIncorrectPassword("密码错误")
	}

	return &userV1.VerifyPasswordResponse{
		Result: userV1.VerifyPasswordResult_SUCCESS,
	}, nil
}

func (r *UserRepo) UserExists(ctx context.Context, req *userV1.UserExistsRequest) (*userV1.UserExistsResponse, error) {
	count, _ := r.data.db.Client().User.
		Query().
		Select(user.FieldID).
		Where(user.UsernameEQ(req.GetUserName())).
		Count(ctx)
	return &userV1.UserExistsResponse{
		Exist: count > 0,
	}, nil
}
