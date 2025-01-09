package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"

	"kratos-admin/app/admin/service/internal/data"

	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	adminV1 "kratos-admin/api/gen/go/admin/service/v1"
	systemV1 "kratos-admin/api/gen/go/system/service/v1"
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
	err := s.uc.Create(ctx, req)
	if err != nil {

		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *DictService) UpdateDict(ctx context.Context, req *systemV1.UpdateDictRequest) (*emptypb.Empty, error) {
	err := s.uc.Update(ctx, req)
	if err != nil {

		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *DictService) DeleteDict(ctx context.Context, req *systemV1.DeleteDictRequest) (*emptypb.Empty, error) {
	_, err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
