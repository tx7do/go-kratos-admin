package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/go-utils/trans"
	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	"google.golang.org/protobuf/types/known/emptypb"

	"kratos-admin/app/admin/service/internal/data"

	adminV1 "kratos-admin/api/gen/go/admin/service/v1"
	dictV1 "kratos-admin/api/gen/go/dict/service/v1"

	"kratos-admin/pkg/middleware/auth"
)

type DictService struct {
	adminV1.DictServiceHTTPServer

	log *log.Helper

	dictTypeRepo  *data.DictTypeRepo
	dictEntryRepo *data.DictEntryRepo
}

func NewDictService(
	logger log.Logger,
	dictTypeRepo *data.DictTypeRepo,
	dictEntryRepo *data.DictEntryRepo,
) *DictService {
	l := log.NewHelper(log.With(logger, "module", "dict/service/admin-service"))
	return &DictService{
		log:           l,
		dictTypeRepo:  dictTypeRepo,
		dictEntryRepo: dictEntryRepo,
	}
}

func (s *DictService) ListDictType(ctx context.Context, req *pagination.PagingRequest) (*dictV1.ListDictTypeResponse, error) {
	return s.dictTypeRepo.List(ctx, req)
}

func (s *DictService) GetDictType(ctx context.Context, req *dictV1.GetDictTypeRequest) (*dictV1.DictType, error) {
	return s.dictTypeRepo.Get(ctx, req)
}

func (s *DictService) CreateDictType(ctx context.Context, req *dictV1.CreateDictTypeRequest) (*emptypb.Empty, error) {
	if req.Data == nil {
		return nil, adminV1.ErrorBadRequest("invalid parameter")
	}

	// 获取操作人信息
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return nil, err
	}

	req.Data.CreatedBy = trans.Ptr(operator.UserId)

	if err = s.dictTypeRepo.Create(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *DictService) UpdateDictType(ctx context.Context, req *dictV1.UpdateDictTypeRequest) (*emptypb.Empty, error) {
	if req.Data == nil {
		return nil, adminV1.ErrorBadRequest("invalid parameter")
	}

	// 获取操作人信息
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return nil, err
	}

	req.Data.UpdatedBy = trans.Ptr(operator.UserId)

	if err = s.dictTypeRepo.Update(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *DictService) DeleteDictType(ctx context.Context, req *dictV1.BatchDeleteDictRequest) (*emptypb.Empty, error) {
	if err := s.dictTypeRepo.BatchDelete(ctx, req.GetIds()); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *DictService) ListDictEntry(ctx context.Context, req *pagination.PagingRequest) (*dictV1.ListDictEntryResponse, error) {
	return s.dictEntryRepo.List(ctx, req)
}

func (s *DictService) CreateDictEntry(ctx context.Context, req *dictV1.CreateDictEntryRequest) (*emptypb.Empty, error) {
	if req.Data == nil {
		return nil, adminV1.ErrorBadRequest("invalid parameter")
	}

	// 获取操作人信息
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return nil, err
	}

	req.Data.CreatedBy = trans.Ptr(operator.UserId)

	if err = s.dictEntryRepo.Create(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *DictService) UpdateDictEntry(ctx context.Context, req *dictV1.UpdateDictEntryRequest) (*emptypb.Empty, error) {
	if req.Data == nil {
		return nil, adminV1.ErrorBadRequest("invalid parameter")
	}

	// 获取操作人信息
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return nil, err
	}

	req.Data.UpdatedBy = trans.Ptr(operator.UserId)

	if err = s.dictEntryRepo.Update(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *DictService) DeleteDictEntry(ctx context.Context, req *dictV1.BatchDeleteDictRequest) (*emptypb.Empty, error) {
	if err := s.dictEntryRepo.BatchDelete(ctx, req.GetIds()); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
