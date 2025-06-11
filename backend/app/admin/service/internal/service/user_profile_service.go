package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/go-utils/trans"
	"google.golang.org/protobuf/types/known/emptypb"

	"kratos-admin/app/admin/service/internal/data"

	adminV1 "kratos-admin/api/gen/go/admin/service/v1"
	authenticationV1 "kratos-admin/api/gen/go/authentication/service/v1"
	userV1 "kratos-admin/api/gen/go/user/service/v1"

	"kratos-admin/pkg/middleware/auth"
)

type UserProfileService struct {
	adminV1.UserProfileServiceHTTPServer

	userRepo  *data.UserRepo
	userToken *data.UserTokenCacheRepo
	roleRepo  *data.RoleRepo

	log *log.Helper
}

func NewUserProfileService(
	logger log.Logger,
	userRepo *data.UserRepo,
	userToken *data.UserTokenCacheRepo,
	roleRepo *data.RoleRepo,
) *UserProfileService {
	l := log.NewHelper(log.With(logger, "module", "user-profile/service/admin-service"))
	return &UserProfileService{
		log:       l,
		userRepo:  userRepo,
		userToken: userToken,
		roleRepo:  roleRepo,
	}
}

func (s *UserProfileService) GetUser(ctx context.Context, _ *emptypb.Empty) (*userV1.User, error) {
	// 获取操作人信息
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return nil, err
	}

	user, err := s.userRepo.Get(ctx, operator.UserId)
	if err != nil {
		s.log.Errorf("查询用户失败[%s]", err.Error())
		return nil, authenticationV1.ErrorNotFound("user not found")
	}

	//role, err := s.roleRepo.Get(ctx, user.GetRoleId())
	//if err == nil && role != nil {
	//	user.Roles = append(user.Roles, role.GetCode())
	//}

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
