package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	"google.golang.org/protobuf/types/known/emptypb"

	"kratos-admin/app/admin/service/internal/data"

	adminV1 "kratos-admin/api/gen/go/admin/service/v1"
	systemV1 "kratos-admin/api/gen/go/system/service/v1"

	"kratos-admin/pkg/middleware/auth"
)

type DictService struct {
	adminV1.DictServiceHTTPServer

	log *log.Helper

	uc *data.DictRepo
}

func NewDictService(uc *data.DictRepo, logger log.Logger) *DictService {
	l := log.NewHelper(log.With(logger, "module", "dict/service/admin-service"))
	return &DictService{
		log: l,
		uc:  uc,
	}
}

func (s *DictService) ListDict(ctx context.Context, req *pagination.PagingRequest) (*systemV1.ListDictResponse, error) {
	return s.uc.List(ctx, req)
}

func (s *DictService) GetDict(ctx context.Context, req *systemV1.GetDictRequest) (*systemV1.Dict, error) {
	return s.uc.Get(ctx, req)
}

func (s *DictService) CreateDict(ctx context.Context, req *systemV1.CreateDictRequest) (*emptypb.Empty, error) {
	if req.Data == nil {
		return nil, adminV1.ErrorBadRequest("错误的参数")
	}

	// 获取操作人信息
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	if err = s.uc.Create(ctx, req, operator); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *DictService) UpdateDict(ctx context.Context, req *systemV1.UpdateDictRequest) (*emptypb.Empty, error) {
	if req.Data == nil {
		return nil, adminV1.ErrorBadRequest("错误的参数")
	}

	// 获取操作人信息
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	if err = s.uc.Update(ctx, req, operator); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *DictService) DeleteDict(ctx context.Context, req *systemV1.DeleteDictRequest) (*emptypb.Empty, error) {
	if _, err := s.uc.Delete(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
