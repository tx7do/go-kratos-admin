package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/go-utils/trans"
	"google.golang.org/protobuf/types/known/emptypb"

	"kratos-admin/app/admin/service/internal/data"

	adminV1 "kratos-admin/api/gen/go/admin/service/v1"
	userV1 "kratos-admin/api/gen/go/user/service/v1"

	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"

	"kratos-admin/pkg/middleware/auth"
)

type UserService struct {
	adminV1.UserServiceHTTPServer

	uc  *data.UserRepo
	log *log.Helper
}

func NewUserService(logger log.Logger, uc *data.UserRepo) *UserService {
	l := log.NewHelper(log.With(logger, "module", "user/service/admin-service"))
	return &UserService{
		log: l,
		uc:  uc,
	}
}

func (s *UserService) ListUser(ctx context.Context, req *pagination.PagingRequest) (*userV1.ListUserResponse, error) {
	return s.uc.List(ctx, req)
}

func (s *UserService) GetUser(ctx context.Context, req *userV1.GetUserRequest) (*userV1.User, error) {
	return s.uc.Get(ctx, req)
}

func (s *UserService) GetUserByUserName(ctx context.Context, req *userV1.GetUserByUserNameRequest) (*userV1.User, error) {
	return s.uc.GetUserByUserName(ctx, req.GetUserName())
}

func (s *UserService) CreateUser(ctx context.Context, req *userV1.CreateUserRequest) (*emptypb.Empty, error) {
	authInfo, err := auth.FromContext(ctx)
	if err != nil {
		s.log.Errorf("用户认证失败[%s]", err.Error())
		return nil, adminV1.ErrorAccessForbidden("用户认证失败")
	}

	if req.User == nil {
		return nil, adminV1.ErrorBadRequest("错误的参数")
	}

	req.OperatorId = trans.Ptr(authInfo.UserId)
	req.User.CreatorId = trans.Ptr(authInfo.UserId)
	if req.User.Authority == nil {
		req.User.Authority = userV1.UserAuthority_CUSTOMER_USER.Enum()
	}

	err = s.uc.Create(ctx, req)
	return &emptypb.Empty{}, nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *userV1.UpdateUserRequest) (*emptypb.Empty, error) {
	authInfo, err := auth.FromContext(ctx)
	if err != nil {
		s.log.Errorf("用户认证失败[%s]", err.Error())
		return nil, adminV1.ErrorAccessForbidden("用户认证失败")
	}

	if req.User == nil {
		return nil, adminV1.ErrorBadRequest("错误的参数")
	}

	req.OperatorId = trans.Ptr(authInfo.UserId)

	err = s.uc.Update(ctx, req)
	return &emptypb.Empty{}, nil
}

func (s *UserService) DeleteUser(ctx context.Context, req *userV1.DeleteUserRequest) (*emptypb.Empty, error) {
	authInfo, err := auth.FromContext(ctx)
	if err != nil {
		s.log.Errorf("用户认证失败[%s]", err.Error())
		return nil, adminV1.ErrorAccessForbidden("用户认证失败")
	}

	req.OperatorId = trans.Ptr(authInfo.UserId)

	_, err = s.uc.Delete(ctx, req)

	return &emptypb.Empty{}, err
}

func (s *UserService) UserExists(ctx context.Context, req *userV1.UserExistsRequest) (*userV1.UserExistsResponse, error) {
	return s.uc.UserExists(ctx, req)
}
