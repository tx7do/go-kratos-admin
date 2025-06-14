package data

import (
	"context"
	"encoding/base64"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"
	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"

	"github.com/tx7do/go-utils/copierutil"
	"github.com/tx7do/go-utils/crypto"
	entgo "github.com/tx7do/go-utils/entgo/query"
	entgoUpdate "github.com/tx7do/go-utils/entgo/update"
	"github.com/tx7do/go-utils/fieldmaskutil"
	"github.com/tx7do/go-utils/mapper"
	"github.com/tx7do/go-utils/timeutil"
	"github.com/tx7do/go-utils/trans"

	"kratos-admin/app/admin/service/internal/data/ent"
	"kratos-admin/app/admin/service/internal/data/ent/usercredential"

	authenticationV1 "kratos-admin/api/gen/go/authentication/service/v1"
)

type UserCredentialRepo struct {
	data *Data
	log  *log.Helper

	mapper                  *mapper.CopierMapper[ent.UserCredential, authenticationV1.UserCredential]
	statusConverter         *mapper.EnumTypeConverter[usercredential.Status, authenticationV1.UserCredentialStatus]
	identityTypeConverter   *mapper.EnumTypeConverter[usercredential.IdentityType, authenticationV1.IdentityType]
	credentialTypeConverter *mapper.EnumTypeConverter[usercredential.CredentialType, authenticationV1.CredentialType]
}

func NewUserCredentialRepo(data *Data, logger log.Logger) *UserCredentialRepo {
	repo := &UserCredentialRepo{
		log:                     log.NewHelper(log.With(logger, "module", "user-credentials/repo/admin-service")),
		data:                    data,
		mapper:                  mapper.NewCopierMapper[ent.UserCredential, authenticationV1.UserCredential](),
		statusConverter:         mapper.NewEnumTypeConverter[usercredential.Status, authenticationV1.UserCredentialStatus](authenticationV1.UserCredentialStatus_name, authenticationV1.UserCredentialStatus_value),
		identityTypeConverter:   mapper.NewEnumTypeConverter[usercredential.IdentityType, authenticationV1.IdentityType](authenticationV1.IdentityType_name, authenticationV1.IdentityType_value),
		credentialTypeConverter: mapper.NewEnumTypeConverter[usercredential.CredentialType, authenticationV1.CredentialType](authenticationV1.CredentialType_name, authenticationV1.CredentialType_value),
	}

	repo.init()

	return repo
}

func (r *UserCredentialRepo) init() {
	r.mapper.AppendConverters(copierutil.NewTimeStringConverterPair())
	r.mapper.AppendConverters(copierutil.NewTimeTimestamppbConverterPair())
	r.mapper.AppendConverters(r.statusConverter.NewConverterPair())
	r.mapper.AppendConverters(r.identityTypeConverter.NewConverterPair())
	r.mapper.AppendConverters(r.credentialTypeConverter.NewConverterPair())
}

func (r *UserCredentialRepo) IsExist(ctx context.Context, id uint32) (bool, error) {
	exist, err := r.data.db.Client().UserCredential.Query().
		Where(usercredential.IDEQ(id)).
		Exist(ctx)
	if err != nil {
		r.log.Errorf("query exist failed: %s", err.Error())
		return false, authenticationV1.ErrorInternalServerError("query exist failed")
	}
	return exist, nil
}

func (r *UserCredentialRepo) Count(ctx context.Context, whereCond []func(s *sql.Selector)) (int, error) {
	builder := r.data.db.Client().UserCredential.Query()
	if len(whereCond) != 0 {
		builder.Modify(whereCond...)
	}

	count, err := builder.Count(ctx)
	if err != nil {
		r.log.Errorf("query count failed: %s", err.Error())
		return 0, authenticationV1.ErrorInternalServerError("query count failed")
	}

	return count, nil
}

func (r *UserCredentialRepo) List(ctx context.Context, req *pagination.PagingRequest) (*authenticationV1.ListUserCredentialResponse, error) {
	if req == nil {
		return nil, authenticationV1.ErrorBadRequest("invalid parameter")
	}

	builder := r.data.db.Client().UserCredential.Query()

	err, whereSelectors, querySelectors := entgo.BuildQuerySelector(
		req.GetQuery(), req.GetOrQuery(),
		req.GetPage(), req.GetPageSize(), req.GetNoPaging(),
		req.GetOrderBy(), usercredential.FieldCreateTime,
		req.GetFieldMask().GetPaths(),
	)
	if err != nil {
		r.log.Errorf("parse list param error [%s]", err.Error())
		return nil, authenticationV1.ErrorBadRequest("invalid query parameter")
	}

	if querySelectors != nil {
		builder.Modify(querySelectors...)
	}

	results, err := builder.All(ctx)
	if err != nil {
		r.log.Errorf("query list failed: %s", err.Error())
		return nil, authenticationV1.ErrorInternalServerError("query list failed")
	}

	models := make([]*authenticationV1.UserCredential, 0, len(results))
	for _, dto := range results {
		model := r.mapper.ToModel(dto)
		models = append(models, model)
	}

	count, err := r.Count(ctx, whereSelectors)
	if err != nil {
		return nil, err
	}

	return &authenticationV1.ListUserCredentialResponse{
		Total: uint32(count),
		Items: models,
	}, nil
}

func (r *UserCredentialRepo) Create(ctx context.Context, req *authenticationV1.CreateUserCredentialRequest) error {
	if req == nil || req.Data == nil {
		return authenticationV1.ErrorBadRequest("invalid request")
	}

	var err error

	if req.Data.Credential != nil {
		var newCredential string
		newCredential, err = r.prepareCredential(r.credentialTypeConverter.ToDto(req.Data.CredentialType), req.Data.GetCredential())
		if err != nil {
			r.log.Errorf("prepare new credential failed: %s", err.Error())
			return authenticationV1.ErrorBadRequest("prepare new credential failed")
		}
		req.Data.Credential = trans.Ptr(newCredential)
	}

	builder := r.data.db.Client().UserCredential.Create()
	builder.
		SetUserID(req.Data.GetUserId()).
		SetNillableTenantID(req.Data.TenantId).
		SetNillableIdentityType(r.identityTypeConverter.ToDto(req.Data.IdentityType)).
		SetNillableIdentifier(req.Data.Identifier).
		SetNillableCredentialType(r.credentialTypeConverter.ToDto(req.Data.CredentialType)).
		SetNillableCredential(req.Data.Credential).
		SetNillableIsPrimary(req.Data.IsPrimary).
		SetNillableStatus(r.statusConverter.ToDto(req.Data.Status)).
		SetNillableExtraInfo(req.Data.ExtraInfo).
		SetNillableCreateTime(timeutil.TimestamppbToTime(req.Data.CreateTime))

	if req.Data.CreateTime == nil {
		builder.SetCreateTime(time.Now())
	}

	if err = builder.Exec(ctx); err != nil {
		r.log.Errorf("insert one data failed: %s [%v]", err.Error(), req.Data)
		return authenticationV1.ErrorInternalServerError("insert data failed")
	}

	return nil
}

func (r *UserCredentialRepo) Update(ctx context.Context, req *authenticationV1.UpdateUserCredentialRequest) error {
	if req == nil || req.Data == nil {
		return authenticationV1.ErrorBadRequest("invalid request")
	}

	// 如果不存在则创建
	if req.GetAllowMissing() {
		exist, err := r.IsExist(ctx, req.GetData().GetId())
		if err != nil {
			return err
		}
		if !exist {
			err = r.Create(ctx, &authenticationV1.CreateUserCredentialRequest{Data: req.Data})
			return err
		}
	}

	var err error

	if req.Data.Credential != nil {
		var newCredential string
		newCredential, err = r.prepareCredential(r.credentialTypeConverter.ToDto(req.Data.CredentialType), req.Data.GetCredential())
		if err != nil {
			r.log.Errorf("prepare new credential failed: %s", err.Error())
			return authenticationV1.ErrorBadRequest("prepare new credential failed")
		}
		req.Data.Credential = trans.Ptr(newCredential)
	}

	if req.UpdateMask != nil {
		req.UpdateMask.Normalize()
		if !req.UpdateMask.IsValid(req.Data) {
			r.log.Errorf("invalid field mask [%v]", req.UpdateMask)
			return authenticationV1.ErrorBadRequest("invalid field mask")
		}
		fieldmaskutil.Filter(req.GetData(), req.UpdateMask.GetPaths())
	}

	builder := r.data.db.Client().UserCredential.UpdateOneID(req.Data.Id).
		SetNillableIdentityType(r.identityTypeConverter.ToDto(req.Data.IdentityType)).
		SetNillableIdentifier(req.Data.Identifier).
		SetNillableCredentialType(r.credentialTypeConverter.ToDto(req.Data.CredentialType)).
		SetNillableCredential(req.Data.Credential).
		SetNillableIsPrimary(req.Data.IsPrimary).
		SetNillableStatus(r.statusConverter.ToDto(req.Data.Status)).
		SetNillableExtraInfo(req.Data.ExtraInfo).
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

	if err = builder.Exec(ctx); err != nil {
		r.log.Errorf("update one data failed: %s", err.Error())
		return authenticationV1.ErrorInternalServerError("update data failed")
	}

	return nil
}

func (r *UserCredentialRepo) Delete(ctx context.Context, id uint32) error {
	builder := r.data.db.Client().UserCredential.Delete()
	builder.Where(usercredential.IDEQ(id))
	if affected, err := builder.Exec(ctx); err != nil {
		if ent.IsNotFound(err) {
			return authenticationV1.ErrorNotFound("user credential not found")
		}

		r.log.Errorf("delete one data failed: %s", err.Error())

		return authenticationV1.ErrorInternalServerError("delete one data failed")
	} else {
		if affected == 0 {
			return authenticationV1.ErrorNotFound("user credential not found")
		} else {
			return nil
		}
	}
}

func (r *UserCredentialRepo) DeleteByUserId(ctx context.Context, userId uint32) error {
	builder := r.data.db.Client().UserCredential.Delete()
	builder.Where(usercredential.UserIDEQ(userId))
	if affected, err := builder.Exec(ctx); err != nil {
		if ent.IsNotFound(err) {
			return authenticationV1.ErrorNotFound("user credential not found")
		}

		r.log.Errorf("delete one data failed: %s", err.Error())

		return authenticationV1.ErrorInternalServerError("delete one data failed")
	} else {
		if affected == 0 {
			return authenticationV1.ErrorNotFound("user credential not found")
		} else {
			return nil
		}
	}
}

func (r *UserCredentialRepo) DeleteByIdentifier(ctx context.Context, identityType authenticationV1.IdentityType, identifier string) error {
	builder := r.data.db.Client().UserCredential.Delete()
	builder.Where(
		usercredential.IdentityTypeEQ(*r.identityTypeConverter.ToDto(&identityType)),
		usercredential.IdentifierEQ(identifier),
	)
	if affected, err := builder.Exec(ctx); err != nil {
		if ent.IsNotFound(err) {
			return authenticationV1.ErrorNotFound("user credential not found")
		}

		r.log.Errorf("delete one data failed: %s", err.Error())

		return authenticationV1.ErrorInternalServerError("delete one data failed")
	} else {
		if affected == 0 {
			return authenticationV1.ErrorNotFound("user credential not found")
		} else {
			return nil
		}
	}
}

func (r *UserCredentialRepo) Get(ctx context.Context, req *authenticationV1.GetUserCredentialRequest) (*authenticationV1.UserCredential, error) {
	builder := r.data.db.Client().UserCredential.Query()

	builder.Where(
		usercredential.IDEQ(req.GetId()),
	)

	ret, err := builder.Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, authenticationV1.ErrorNotFound("user credential not found")
		}

		r.log.Errorf("query one data failed: %s", err.Error())

		return nil, authenticationV1.ErrorInternalServerError("query data failed")
	}

	return r.mapper.ToModel(ret), nil
}

func (r *UserCredentialRepo) GetByIdentifier(ctx context.Context, req *authenticationV1.GetUserCredentialByIdentifierRequest) (*authenticationV1.UserCredential, error) {
	builder := r.data.db.Client().UserCredential.Query()

	builder.Where(
		usercredential.IdentityTypeEQ(*r.identityTypeConverter.ToDto(trans.Ptr(req.GetIdentityType()))),
		usercredential.IdentifierEQ(req.GetIdentifier()),
	)

	ret, err := builder.Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, authenticationV1.ErrorNotFound("user credential not found")
		}

		r.log.Errorf("query one data failed: %s", err.Error())

		return nil, authenticationV1.ErrorInternalServerError("query data failed")
	}

	return r.mapper.ToModel(ret), nil
}

func (r *UserCredentialRepo) VerifyCredential(ctx context.Context, req *authenticationV1.VerifyCredentialRequest) (*authenticationV1.VerifyCredentialResponse, error) {
	if req.GetNeedDecrypt() {
		// 解密密码
		bytesPass, err := base64.StdEncoding.DecodeString(req.GetCredential())
		if err != nil {
			r.log.Errorf("decode base64 credential failed: %s", err.Error())
			return &authenticationV1.VerifyCredentialResponse{
				Success: false,
			}, authenticationV1.ErrorBadRequest("invalid credential format")
		}
		plainPassword, err := crypto.AesDecrypt(bytesPass, crypto.DefaultAESKey, nil)
		if err != nil {
			r.log.Errorf("decrypt credential failed: %s", err.Error())
			return &authenticationV1.VerifyCredentialResponse{
				Success: false,
			}, authenticationV1.ErrorBadRequest("decrypt credential failed")
		}

		req.Credential = string(plainPassword)
	}

	ret, err := r.data.db.Client().UserCredential.Query().
		Select(usercredential.FieldCredentialType, usercredential.FieldCredential, usercredential.FieldStatus).
		Where(
			usercredential.IdentityTypeEQ(*r.identityTypeConverter.ToDto(trans.Ptr(req.GetIdentityType()))),
			usercredential.IdentifierEQ(req.GetIdentifier()),
		).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return &authenticationV1.VerifyCredentialResponse{
				Success: false,
			}, authenticationV1.ErrorUserNotFound("user not found")
		}

		r.log.Errorf("query one data failed: %s", err.Error())

		return &authenticationV1.VerifyCredentialResponse{
			Success: false,
		}, authenticationV1.ErrorServiceUnavailable("db error")
	}

	if *ret.Status != usercredential.StatusEnabled {
		return &authenticationV1.VerifyCredentialResponse{
			Success: false,
		}, authenticationV1.ErrorUserFreeze("account has freeze")
	}

	if !r.verifyCredential(ret.CredentialType, req.GetCredential(), *ret.Credential) {
		return &authenticationV1.VerifyCredentialResponse{
			Success: false,
		}, authenticationV1.ErrorIncorrectPassword("incorrect password")
	}

	return &authenticationV1.VerifyCredentialResponse{
		Success: true,
	}, nil
}

func (r *UserCredentialRepo) verifyCredential(credentialType *usercredential.CredentialType, plainCredential, targetCredential string) bool {
	if credentialType == nil || plainCredential == "" {
		return false
	}

	switch *credentialType {
	case usercredential.CredentialTypePasswordHash:
		return crypto.VerifyPassword(plainCredential, targetCredential)
	default:
		return plainCredential == targetCredential
	}
}

func (r *UserCredentialRepo) prepareCredential(credentialType *usercredential.CredentialType, plainCredential string) (string, error) {
	var newCredential string
	switch *credentialType {
	case usercredential.CredentialTypePasswordHash:
		var err error
		// 加密明文密码
		newCredential, err = crypto.HashPassword(plainCredential)
		if err != nil {
			r.log.Errorf("hash new password failed: %s", err.Error())
			return "", authenticationV1.ErrorBadRequest("hash new password failed")
		}

	default:
		newCredential = plainCredential
	}

	return newCredential, nil
}

func (r *UserCredentialRepo) ChangeCredential(ctx context.Context, req *authenticationV1.ChangeCredentialRequest) error {
	if req.GetNeedDecrypt() {
		// 解密密码
		bytesPass, _ := base64.StdEncoding.DecodeString(req.GetOldCredential())
		plainPassword, _ := crypto.AesDecrypt(bytesPass, crypto.DefaultAESKey, nil)
		req.OldCredential = string(plainPassword)

		bytesPass, _ = base64.StdEncoding.DecodeString(req.GetNewCredential())
		plainPassword, _ = crypto.AesDecrypt(bytesPass, crypto.DefaultAESKey, nil)
		req.NewCredential = string(plainPassword)
	}

	ret, err := r.data.db.Client().UserCredential.
		Query().
		Select(
			usercredential.FieldCredentialType,
			usercredential.FieldCredential,
		).
		Where(
			usercredential.IdentityTypeEQ(*r.identityTypeConverter.ToDto(trans.Ptr(req.GetIdentityType()))),
			usercredential.IdentifierEQ(req.GetIdentifier()),
		).
		Only(ctx)
	if err != nil {
		r.log.Errorf("query one data failed: %s", err.Error())
		return authenticationV1.ErrorInternalServerError("query one data failed")
	}

	if ret.CredentialType == nil {
		return authenticationV1.ErrorNotFound("user credential not found")
	}

	// 验证旧认证信息
	if !r.verifyCredential(ret.CredentialType, req.GetOldCredential(), *ret.Credential) {
		return authenticationV1.ErrorBadRequest("invalid old password")
	}

	var newCredential string
	newCredential, err = r.prepareCredential(ret.CredentialType, req.GetOldCredential())
	if err != nil {
		r.log.Errorf("prepare new credential failed: %s", err.Error())
		return authenticationV1.ErrorBadRequest("prepare new credential failed")
	}

	if newCredential == "" {
		return authenticationV1.ErrorBadRequest("new credential cannot be empty")
	}

	builder := r.data.db.Client().UserCredential.Update()
	builder.Where(
		usercredential.IdentityTypeEQ(*r.identityTypeConverter.ToDto(trans.Ptr(req.GetIdentityType()))),
		usercredential.IdentifierEQ(req.GetIdentifier()),
	)
	builder.
		SetCredential(newCredential).
		SetUpdateTime(time.Now())
	if err = builder.Exec(ctx); err != nil {
		r.log.Errorf("update one data failed: %s", err.Error())
		return authenticationV1.ErrorInternalServerError("update data failed")
	}

	return nil
}

func (r *UserCredentialRepo) ResetCredential(ctx context.Context, req *authenticationV1.ResetCredentialRequest) error {
	if req.GetNeedDecrypt() {
		// 解密密码
		bytesPass, _ := base64.StdEncoding.DecodeString(req.GetNewCredential())
		plainPassword, _ := crypto.AesDecrypt(bytesPass, crypto.DefaultAESKey, nil)
		req.NewCredential = string(plainPassword)
	}

	ret, err := r.data.db.Client().UserCredential.
		Query().
		Select(
			usercredential.FieldCredentialType,
		).
		Where(
			usercredential.IdentityTypeEQ(*r.identityTypeConverter.ToDto(trans.Ptr(req.GetIdentityType()))),
			usercredential.IdentifierEQ(req.GetIdentifier()),
		).
		Only(ctx)
	if err != nil {
		r.log.Errorf("query one data failed: %s", err.Error())
		return authenticationV1.ErrorInternalServerError("query one data failed")
	}

	if ret.CredentialType == nil {
		return authenticationV1.ErrorNotFound("user credential not found")
	}

	var newCredential string
	newCredential, err = r.prepareCredential(ret.CredentialType, req.GetNewCredential())
	if err != nil {
		r.log.Errorf("prepare new credential failed: %s", err.Error())
		return authenticationV1.ErrorBadRequest("prepare new credential failed")
	}

	if newCredential == "" {
		return authenticationV1.ErrorBadRequest("new credential cannot be empty")
	}

	builder := r.data.db.Client().UserCredential.Update()
	builder.Where(
		usercredential.IdentityTypeEQ(*r.identityTypeConverter.ToDto(trans.Ptr(req.GetIdentityType()))),
		usercredential.IdentifierEQ(req.GetIdentifier()),
	)
	builder.
		SetCredential(newCredential).
		SetUpdateTime(time.Now())
	if err = builder.Exec(ctx); err != nil {
		r.log.Errorf("update one data failed: %s", err.Error())
		return authenticationV1.ErrorInternalServerError("update data failed")
	}

	return nil
}
