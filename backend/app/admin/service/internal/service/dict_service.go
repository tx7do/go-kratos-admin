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

	dictMainRepo *data.DictMainRepo
	dictItemRepo *data.DictItemRepo
}

func NewDictService(
	logger log.Logger,
	dictMainRepo *data.DictMainRepo,
	dictItemRepo *data.DictItemRepo,
) *DictService {
	l := log.NewHelper(log.With(logger, "module", "dict/service/admin-service"))
	return &DictService{
		log:          l,
		dictMainRepo: dictMainRepo,
		dictItemRepo: dictItemRepo,
	}
}

func (s *DictService) ListDictMain(ctx context.Context, req *pagination.PagingRequest) (*adminV1.ListDictMainResponse, error) {
	return s.dictMainRepo.List(ctx, req)
}

func (s *DictService) GetDictMain(ctx context.Context, req *adminV1.GetDictMainRequest) (*adminV1.DictMain, error) {
	return s.dictMainRepo.Get(ctx, req)
}

func (s *DictService) CreateDictMain(ctx context.Context, req *adminV1.CreateDictMainRequest) (*emptypb.Empty, error) {
	if req.Data == nil {
		return nil, adminV1.ErrorBadRequest("invalid parameter")
	}

	// 获取操作人信息
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	req.Data.CreateBy = trans.Ptr(operator.UserId)

	if err = s.dictMainRepo.Create(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *DictService) UpdateDictMain(ctx context.Context, req *adminV1.UpdateDictMainRequest) (*emptypb.Empty, error) {
	if req.Data == nil {
		return nil, adminV1.ErrorBadRequest("invalid parameter")
	}

	// 获取操作人信息
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	req.Data.UpdateBy = trans.Ptr(operator.UserId)

	if err = s.dictMainRepo.Update(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *DictService) DeleteDictMain(ctx context.Context, req *adminV1.BatchDeleteDictRequest) (*emptypb.Empty, error) {
	if err := s.dictMainRepo.BatchDelete(ctx, req.GetIds()); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *DictService) ListDictItem(ctx context.Context, req *pagination.PagingRequest) (*adminV1.ListDictItemResponse, error) {
	return s.dictItemRepo.List(ctx, req)
}

func (s *DictService) CreateDictItem(ctx context.Context, req *adminV1.CreateDictItemRequest) (*emptypb.Empty, error) {
	if req.Data == nil {
		return nil, adminV1.ErrorBadRequest("invalid parameter")
	}

	// 获取操作人信息
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	req.Data.CreateBy = trans.Ptr(operator.UserId)

	if err = s.dictItemRepo.Create(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *DictService) UpdateDictItem(ctx context.Context, req *adminV1.UpdateDictItemRequest) (*emptypb.Empty, error) {
	if req.Data == nil {
		return nil, adminV1.ErrorBadRequest("invalid parameter")
	}

	// 获取操作人信息
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	req.Data.UpdateBy = trans.Ptr(operator.UserId)

	if err = s.dictItemRepo.Update(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *DictService) DeleteDictItem(ctx context.Context, req *adminV1.BatchDeleteDictRequest) (*emptypb.Empty, error) {
	if err := s.dictItemRepo.BatchDelete(ctx, req.GetIds()); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
