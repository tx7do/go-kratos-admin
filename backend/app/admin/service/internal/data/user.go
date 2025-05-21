package data

import (
	"context"
	"encoding/base64"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"

	"github.com/tx7do/go-utils/crypto"
	entgo "github.com/tx7do/go-utils/entgo/query"
	entgoUpdate "github.com/tx7do/go-utils/entgo/update"
	"github.com/tx7do/go-utils/fieldmaskutil"
	"github.com/tx7do/go-utils/timeutil"
	"github.com/tx7do/go-utils/trans"
	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"

	"kratos-admin/app/admin/service/internal/data/ent"
	_ "kratos-admin/app/admin/service/internal/data/ent/runtime"
	"kratos-admin/app/admin/service/internal/data/ent/user"

	userV1 "kratos-admin/api/gen/go/user/service/v1"
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

func (r *UserRepo) toEntAuthority(in *userV1.UserAuthority) *user.Authority {
	if in == nil {
		return nil
	}

	find, ok := userV1.UserAuthority_name[int32(*in)]
	if !ok {
		return nil
	}

	return (*user.Authority)(trans.Ptr(find))
}

func (r *UserRepo) toProtoAuthority(in *user.Authority) *userV1.UserAuthority {
	if in == nil {
		return nil
	}

	find, ok := userV1.UserAuthority_value[string(*in)]
	if !ok {
		return nil
	}

	return (*userV1.UserAuthority)(trans.Ptr(find))
}

func (r *UserRepo) toEntGender(in *userV1.UserGender) *user.Gender {
	if in == nil {
		return nil
	}

	find, ok := userV1.UserGender_name[int32(*in)]
	if !ok {
		return nil
	}
	return (*user.Gender)(trans.Ptr(find))
}

func (r *UserRepo) toProtoGender(in *user.Gender) *userV1.UserGender {
	if in == nil {
		return nil
	}

	find, ok := userV1.UserGender_value[string(*in)]
	if !ok {
		return nil
	}
	return (*userV1.UserGender)(trans.Ptr(find))
}

func (r *UserRepo) toEntStatus(in *userV1.UserStatus) *user.Status {
	if in == nil {
		return nil
	}

	find, ok := userV1.UserStatus_name[int32(*in)]
	if !ok {
		return nil
	}

	return (*user.Status)(trans.Ptr(find))
}

func (r *UserRepo) toProtoStatus(in *user.Status) *userV1.UserStatus {
	if in == nil {
		return nil
	}

	find, ok := userV1.UserStatus_value[string(*in)]
	if !ok {
		return nil
	}
	return (*userV1.UserStatus)(trans.Ptr(find))
}

func (r *UserRepo) toEnt(in *userV1.User) *ent.User {
	if in == nil {
		return nil
	}

	var out ent.User
	_ = copier.Copy(&out, in)

	out.Gender = r.toEntGender(in.Gender)
	out.Authority = r.toEntAuthority(in.Authority)
	out.Status = r.toEntStatus(in.Status)
	out.CreateTime = timeutil.StringTimeToTime(in.CreateTime)
	out.UpdateTime = timeutil.StringTimeToTime(in.UpdateTime)
	out.DeleteTime = timeutil.StringTimeToTime(in.DeleteTime)

	return &out
}

func (r *UserRepo) toProto(in *ent.User) *userV1.User {
	if in == nil {
		return nil
	}

	var out userV1.User
	_ = copier.Copy(&out, in)

	out.Gender = r.toProtoGender(in.Gender)
	out.Authority = r.toProtoAuthority(in.Authority)
	out.Status = r.toProtoStatus(in.Status)
	out.CreateTime = timeutil.TimeToTimeString(in.CreateTime)
	out.UpdateTime = timeutil.TimeToTimeString(in.UpdateTime)
	out.DeleteTime = timeutil.TimeToTimeString(in.DeleteTime)

	return &out
}

func (r *UserRepo) Count(ctx context.Context, whereCond []func(s *sql.Selector)) (int, error) {
	builder := r.data.db.Client().User.Query()
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

func (r *UserRepo) List(ctx context.Context, req *pagination.PagingRequest) (*userV1.ListUserResponse, error) {
	if req == nil {
		return nil, userV1.ErrorBadRequest("invalid parameter")
	}

	builder := r.data.db.Client().User.Query()

	err, whereSelectors, querySelectors := entgo.BuildQuerySelector(
		req.GetQuery(), req.GetOrQuery(),
		req.GetPage(), req.GetPageSize(), req.GetNoPaging(),
		req.GetOrderBy(), user.FieldCreateTime,
		req.GetFieldMask().GetPaths(),
	)
	if err != nil {
		r.log.Errorf("parse list param error [%s]", err.Error())
		return nil, userV1.ErrorBadRequest("invalid query parameter")
	}

	if querySelectors != nil {
		builder.Modify(querySelectors...)
	}

	results, err := builder.All(ctx)
	if err != nil {
		r.log.Errorf("query list failed: %s", err.Error())
		return nil, userV1.ErrorInternalServerError("query list failed")
	}

	items := make([]*userV1.User, 0, len(results))
	for _, res := range results {
		item := r.toProto(res)
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
	exist, err := r.data.db.Client().User.Query().
		Where(user.IDEQ(id)).
		Exist(ctx)
	if err != nil {
		r.log.Errorf("query exist failed: %s", err.Error())
		return false, userV1.ErrorInternalServerError("query exist failed")
	}
	return exist, nil
}

func (r *UserRepo) Get(ctx context.Context, userId uint32) (*userV1.User, error) {
	if userId == 0 {
		return nil, userV1.ErrorBadRequest("invalid parameter")
	}

	ret, err := r.data.db.Client().User.Get(ctx, userId)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, userV1.ErrorUserNotFound("user not found")
		}

		r.log.Errorf("query one data failed: %s", err.Error())

		return nil, userV1.ErrorInternalServerError("query data failed")
	}

	return r.toProto(ret), nil
}

func (r *UserRepo) Create(ctx context.Context, req *userV1.CreateUserRequest) error {
	if req == nil || req.Data == nil {
		return userV1.ErrorBadRequest("invalid parameter")
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
		SetNillableCreateBy(req.Data.CreateBy).
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

	if req.Data.Id != nil {
		builder.SetID(req.Data.GetId())
	}

	if err := builder.Exec(ctx); err != nil {
		r.log.Errorf("insert one data failed: %s", err.Error())
		return userV1.ErrorInternalServerError("insert data failed")
	}

	return nil
}

func (r *UserRepo) Update(ctx context.Context, req *userV1.UpdateUserRequest) error {
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
			createReq := &userV1.CreateUserRequest{Data: req.Data}
			createReq.Data.CreateBy = createReq.Data.UpdateBy
			createReq.Data.UpdateBy = nil
			return r.Create(ctx, createReq)
		}
	}

	if req.UpdateMask != nil {
		req.UpdateMask.Normalize()
		if !req.UpdateMask.IsValid(req.Data) {
			return userV1.ErrorBadRequest("invalid field mask")
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
		SetNillableUpdateBy(req.Data.UpdateBy).
		SetNillableUpdateTime(timeutil.StringTimeToTime(req.Data.UpdateTime))

	if req.Data.UpdateTime == nil {
		builder.SetUpdateTime(time.Now())
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
		return userV1.ErrorInternalServerError("update data failed")
	}

	return nil
}

func (r *UserRepo) Delete(ctx context.Context, userId uint32) error {
	if err := r.data.db.Client().User.DeleteOneID(userId).Exec(ctx); err != nil {
		if ent.IsNotFound(err) {
			return userV1.ErrorNotFound("user not found")
		}

		r.log.Errorf("delete one data failed: %s", err.Error())

		return userV1.ErrorInternalServerError("delete failed")
	}

	return nil
}

func (r *UserRepo) GetUserByUserName(ctx context.Context, userName string) (*userV1.User, error) {
	if userName == "" {
		return nil, userV1.ErrorBadRequest("invalid parameter")
	}

	ret, err := r.data.db.Client().User.Query().
		Where(user.UsernameEQ(userName)).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, userV1.ErrorUserNotFound("user not found")
		}

		r.log.Errorf("query user data failed: %s", err.Error())

		return nil, userV1.ErrorInternalServerError("query data failed")
	}

	return r.toProto(ret), nil
}

func (r *UserRepo) VerifyPassword(ctx context.Context, req *userV1.VerifyPasswordRequest) (*userV1.VerifyPasswordResponse, error) {
	ret, err := r.data.db.Client().User.
		Query().
		Select(user.FieldID, user.FieldPassword).
		Where(user.UsernameEQ(req.GetUserName())).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, userV1.ErrorUserNotFound("user not found")
		}

		r.log.Errorf("query user data failed: %s", err.Error())

		return nil, userV1.ErrorInternalServerError("query data failed")
	}

	// 解密密码
	bytesPass, err := base64.StdEncoding.DecodeString(req.GetPassword())
	plainPassword, _ := crypto.AesDecrypt(bytesPass, crypto.DefaultAESKey, nil)

	// 校验密码
	bMatched := crypto.VerifyPassword(string(plainPassword), *ret.Password)

	if !bMatched {
		return &userV1.VerifyPasswordResponse{
			Result: userV1.VerifyPasswordResult_WRONG_PASSWORD,
		}, userV1.ErrorIncorrectPassword("Incorrect Password")
	}

	return &userV1.VerifyPasswordResponse{
		Result: userV1.VerifyPasswordResult_SUCCESS,
	}, nil
}

func (r *UserRepo) UserExists(ctx context.Context, req *userV1.UserExistsRequest) (*userV1.UserExistsResponse, error) {
	exist, err := r.data.db.Client().User.Query().
		Where(user.UsernameEQ(req.GetUserName())).
		Exist(ctx)
	if err != nil {
		r.log.Errorf("query exist failed: %s", err.Error())
		return &userV1.UserExistsResponse{
			Exist: false,
		}, userV1.ErrorInternalServerError("query exist failed")
	}

	return &userV1.UserExistsResponse{
		Exist: exist,
	}, nil
}
