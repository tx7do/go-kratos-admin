package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/go-utils/trans"
	"google.golang.org/genproto/protobuf/field_mask"
	"google.golang.org/protobuf/types/known/emptypb"

	"kratos-admin/app/admin/service/internal/data"

	adminV1 "kratos-admin/api/gen/go/admin/service/v1"
	authenticationV1 "kratos-admin/api/gen/go/authentication/service/v1"
	userV1 "kratos-admin/api/gen/go/user/service/v1"

	"kratos-admin/pkg/middleware/auth"
)

type UserProfileService struct {
	adminV1.UserProfileServiceHTTPServer

	userRepo           *data.UserRepo
	userToken          *data.UserTokenCacheRepo
	roleRepo           *data.RoleRepo
	userCredentialRepo *data.UserCredentialRepo

	log *log.Helper
}

func NewUserProfileService(
	logger log.Logger,
	userRepo *data.UserRepo,
	userToken *data.UserTokenCacheRepo,
	roleRepo *data.RoleRepo,
	userCredentialRepo *data.UserCredentialRepo,
) *UserProfileService {
	l := log.NewHelper(log.With(logger, "module", "user-profile/service/admin-service"))
	return &UserProfileService{
		log:                l,
		userRepo:           userRepo,
		userToken:          userToken,
		roleRepo:           roleRepo,
		userCredentialRepo: userCredentialRepo,
	}
}

func (s *UserProfileService) GetUser(ctx context.Context, _ *emptypb.Empty) (*userV1.User, error) {
	// 获取操作人信息
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return nil, err
	}

	user, err := s.userRepo.Get(ctx, &userV1.GetUserRequest{
		QueryBy: &userV1.GetUserRequest_Id{
			Id: operator.UserId,
		},
	})
	if err != nil {
		s.log.Errorf("查询用户失败[%s]", err.Error())
		return nil, authenticationV1.ErrorNotFound("user not found")
	}

	roleCodes, err := s.roleRepo.GetRoleCodesByRoleIds(ctx, user.GetRoleIds())
	if err != nil {
		s.log.Errorf("get user role codes failed [%s]", err.Error())
	}
	if roleCodes != nil {
		user.Roles = roleCodes
	}

	return user, err
}

func (s *UserProfileService) UpdateUser(ctx context.Context, req *userV1.UpdateUserRequest) (*emptypb.Empty, error) {
	// 获取操作人信息
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return nil, err
	}

	req.Data.Id = trans.Ptr(operator.UserId)

	if err = s.userRepo.Update(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *UserProfileService) ChangePassword(ctx context.Context, req *userV1.ChangePasswordRequest) (*emptypb.Empty, error) {
	// 获取操作人信息
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return nil, err
	}

	err = s.userCredentialRepo.ChangeCredential(ctx, &authenticationV1.ChangeCredentialRequest{
		IdentityType:  authenticationV1.UserCredential_USERNAME,
		Identifier:    operator.GetUsername(),
		OldCredential: req.GetOldPassword(),
		NewCredential: req.GetNewPassword(),
	})
	return &emptypb.Empty{}, err
}

// DeleteAvatar 删除头像
func (s *UserProfileService) DeleteAvatar(ctx context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	// 获取操作人信息
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return nil, err
	}

	if err = s.userRepo.Update(ctx, &userV1.UpdateUserRequest{
		Data: &userV1.User{
			Id:     trans.Ptr(operator.UserId),
			Avatar: trans.Ptr(""),
		},
		UpdateMask: &field_mask.FieldMask{
			Paths: []string{"avatar"},
		},
	}); err != nil {
		s.log.Errorf("delete user avatar failed [%s]", err.Error())
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

// UploadAvatar 上传头像
func (s *UserProfileService) UploadAvatar(ctx context.Context, req *userV1.UploadAvatarRequest) (*userV1.UploadAvatarResponse, error) {
	// 获取操作人信息
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return nil, err
	}

	var avatarURL string
	switch req.GetSource().(type) {
	case *userV1.UploadAvatarRequest_ImageBase64:
	case *userV1.UploadAvatarRequest_ImageUrl:
		avatarURL = req.GetImageUrl()
	default:
		s.log.Errorf("upload avatar failed, invalid avatar source")
		return nil, authenticationV1.ErrorBadRequest("invalid avatar source")
	}

	if err = s.userRepo.Update(ctx, &userV1.UpdateUserRequest{
		Data: &userV1.User{
			Id:     trans.Ptr(operator.UserId),
			Avatar: trans.Ptr(avatarURL),
		},
		UpdateMask: &field_mask.FieldMask{
			Paths: []string{"avatar"},
		},
	}); err != nil {
		s.log.Errorf("delete user avatar failed [%s]", err.Error())
		return nil, err
	}

	return &userV1.UploadAvatarResponse{
		Url: avatarURL,
	}, nil
}

// BindContact 绑定手机号码/邮箱
func (s *UserProfileService) BindContact(context.Context, *userV1.BindContactRequest) (*emptypb.Empty, error) {
	return nil, nil
}

// VerifyContact 验证手机号码/邮箱
func (s *UserProfileService) VerifyContact(context.Context, *userV1.VerifyContactRequest) (*emptypb.Empty, error) {
	return nil, nil
}
