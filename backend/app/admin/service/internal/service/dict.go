package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/go-utils/trans"
	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	"google.golang.org/protobuf/types/known/emptypb"

	"kratos-admin/app/admin/service/internal/data"

	adminV1 "kratos-admin/api/gen/go/admin/service/v1"

	"kratos-admin/pkg/middleware/auth"
)

type DictService struct {
	adminV1.DictServiceHTTPServer

	log *log.Helper

	repo *data.DictRepo
}

func NewDictService(logger log.Logger, repo *data.DictRepo) *DictService {
	l := log.NewHelper(log.With(logger, "module", "dict/service/admin-service"))
	return &DictService{
		log:  l,
		repo: repo,
	}
}

func (s *DictService) List(ctx context.Context, req *pagination.PagingRequest) (*adminV1.ListDictResponse, error) {
	return s.repo.List(ctx, req)
}

func (s *DictService) Get(ctx context.Context, req *adminV1.GetDictRequest) (*adminV1.Dict, error) {
	return s.repo.Get(ctx, req)
}

func (s *DictService) Create(ctx context.Context, req *adminV1.CreateDictRequest) (*emptypb.Empty, error) {
	if req.Data == nil {
		return nil, adminV1.ErrorBadRequest("invalid parameter")
	}

	// 获取操作人信息
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	req.Data.CreateBy = trans.Ptr(operator.UserId)

	if err = s.repo.Create(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *DictService) Update(ctx context.Context, req *adminV1.UpdateDictRequest) (*emptypb.Empty, error) {
	if req.Data == nil {
		return nil, adminV1.ErrorBadRequest("invalid parameter")
	}

	// 获取操作人信息
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	req.Data.UpdateBy = trans.Ptr(operator.UserId)

	if err = s.repo.Update(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *DictService) Delete(ctx context.Context, req *adminV1.DeleteDictRequest) (*emptypb.Empty, error) {
	if err := s.repo.Delete(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
