package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/go-utils/trans"
	"google.golang.org/protobuf/types/known/emptypb"

	"kratos-admin/app/admin/service/internal/data"
	"kratos-admin/app/admin/service/internal/middleware/auth"

	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	adminV1 "kratos-admin/api/gen/go/admin/service/v1"
	userV1 "kratos-admin/api/gen/go/user/service/v1"
)

type UserService struct {
	adminV1.UserServiceHTTPServer

	log *log.Helper

	userRepo *data.UserRepo
	roleRepo *data.RoleRepo
}

func NewUserService(
	logger log.Logger,
	userRepo *data.UserRepo,
	roleRepo *data.RoleRepo,
) *UserService {
	l := log.NewHelper(log.With(logger, "module", "user/service/admin-service"))
	return &UserService{
		log:      l,
		userRepo: userRepo,
		roleRepo: roleRepo,
	}
}

func (s *UserService) ListUser(ctx context.Context, req *pagination.PagingRequest) (*userV1.ListUserResponse, error) {
	return s.userRepo.ListUser(ctx, req)
}

func (s *UserService) GetUser(ctx context.Context, req *userV1.GetUserRequest) (*userV1.User, error) {
	user, err := s.userRepo.GetUser(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	role, err := s.roleRepo.GetRole(ctx, user.GetRoleId())
	if err == nil && role != nil {
		user.Roles = append(user.Roles, role.GetCode())
	}

	return user, nil
}

func (s *UserService) GetUserByUserName(ctx context.Context, req *userV1.GetUserByUserNameRequest) (*userV1.User, error) {
	user, err := s.userRepo.GetUserByUserName(ctx, req.GetUserName())
	if err != nil {
		return nil, err
	}

	role, err := s.roleRepo.GetRole(ctx, user.GetRoleId())
	if err == nil && role != nil {
		user.Roles = append(user.Roles, role.GetCode())
	}

	return user, nil
}

func (s *UserService) CreateUser(ctx context.Context, req *userV1.CreateUserRequest) (*emptypb.Empty, error) {
	authInfo, err := auth.FromContext(ctx)
	if err != nil {
		s.log.Errorf("用户认证失败[%s]", err.Error())
		return nil, adminV1.ErrorAccessForbidden("用户认证失败")
	}

	if req.Data == nil {
		return nil, adminV1.ErrorBadRequest("错误的参数")
	}

	req.OperatorId = trans.Ptr(authInfo.UserId)
	//req.Data.CreateBy = trans.Ptr(authInfo.UserId)
	if req.Data.Authority == nil {
		req.Data.Authority = userV1.UserAuthority_CUSTOMER_USER.Enum()
	}

	// 获取操作者的用户信息
	operator, err := s.userRepo.GetUser(ctx, req.GetOperatorId())
	if err != nil {
		return nil, err
	}

	// 校验操作者的权限
	if operator.GetAuthority() != userV1.UserAuthority_SYS_ADMIN && operator.GetAuthority() != userV1.UserAuthority_SYS_MANAGER {
		return nil, adminV1.ErrorAccessForbidden("权限不够")
	}

	if req.Data.Authority != nil {
		if operator.GetAuthority() >= req.Data.GetAuthority() {
			return nil, adminV1.ErrorAccessForbidden("不能够创建同级用户或者比自己权限高的用户")
		}
	}

	// 创建用户
	err = s.userRepo.CreateUser(ctx, req)

	return &emptypb.Empty{}, nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *userV1.UpdateUserRequest) (*emptypb.Empty, error) {
	authInfo, err := auth.FromContext(ctx)
	if err != nil {
		s.log.Errorf("用户认证失败[%s]", err.Error())
		return nil, adminV1.ErrorAccessForbidden("用户认证失败")
	}

	if req.Data == nil {
		return nil, adminV1.ErrorBadRequest("错误的参数")
	}

	req.OperatorId = trans.Ptr(authInfo.UserId)

	// 获取操作者的用户信息
	operator, err := s.userRepo.GetUser(ctx, req.GetOperatorId())
	if err != nil {
		return nil, err
	}

	// 校验操作者的权限
	if operator.GetAuthority() != userV1.UserAuthority_SYS_ADMIN && operator.GetAuthority() != userV1.UserAuthority_SYS_MANAGER {
		return nil, adminV1.ErrorAccessForbidden("权限不够")
	}

	if req.Data.Authority != nil {
		if operator.GetAuthority() >= req.Data.GetAuthority() {
			return nil, adminV1.ErrorAccessForbidden("不能够赋权同级用户或者比自己权限高的用户")
		}
	}

	// 更新用户
	err = s.userRepo.UpdateUser(ctx, req)

	return &emptypb.Empty{}, nil
}

func (s *UserService) DeleteUser(ctx context.Context, req *userV1.DeleteUserRequest) (*emptypb.Empty, error) {
	authInfo, err := auth.FromContext(ctx)
	if err != nil {
		s.log.Errorf("用户认证失败[%s]", err.Error())
		return nil, adminV1.ErrorAccessForbidden("用户认证失败")
	}

	req.OperatorId = trans.Ptr(authInfo.UserId)

	// 获取操作者的用户信息
	operator, err := s.userRepo.GetUser(ctx, req.GetOperatorId())
	if err != nil {
		return nil, err
	}

	// 校验操作者的权限
	if operator.GetAuthority() != userV1.UserAuthority_SYS_ADMIN && operator.GetAuthority() != userV1.UserAuthority_SYS_MANAGER {
		return nil, adminV1.ErrorAccessForbidden("权限不够")
	}

	// 获取将被删除的用户信息
	user, err := s.userRepo.GetUser(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	// 不能删除超级管理员
	if user.GetAuthority() == userV1.UserAuthority_SYS_ADMIN {
		return nil, adminV1.ErrorAccessForbidden("闹哪样？不能删除超级管理员！")
	}

	if operator.GetAuthority() == user.GetAuthority() {
		return nil, adminV1.ErrorAccessForbidden("不能删除同级用户！")
	}

	// 删除用户
	_, err = s.userRepo.DeleteUser(ctx, req.GetId())

	return &emptypb.Empty{}, err
}

func (s *UserService) UserExists(ctx context.Context, req *userV1.UserExistsRequest) (*userV1.UserExistsResponse, error) {
	return s.userRepo.UserExists(ctx, req)
}

func (s *UserService) VerifyPassword(ctx context.Context, req *userV1.VerifyPasswordRequest) (*userV1.VerifyPasswordResponse, error) {
	return s.userRepo.VerifyPassword(ctx, req)
}
