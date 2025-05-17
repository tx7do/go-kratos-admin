package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/go-utils/trans"
	"google.golang.org/protobuf/types/known/emptypb"

	"kratos-admin/app/admin/service/internal/data"

	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	adminV1 "kratos-admin/api/gen/go/admin/service/v1"
	fileV1 "kratos-admin/api/gen/go/file/service/v1"

	"kratos-admin/pkg/middleware/auth"
)

type FileService struct {
	adminV1.FileServiceHTTPServer

	log *log.Helper

	uc *data.FileRepo
}

func NewFileService(uc *data.FileRepo, logger log.Logger) *FileService {
	l := log.NewHelper(log.With(logger, "module", "file/service/admin-service"))
	return &FileService{
		log: l,
		uc:  uc,
	}
}

func (s *FileService) List(ctx context.Context, req *pagination.PagingRequest) (*fileV1.ListFileResponse, error) {
	return s.uc.List(ctx, req)
}

func (s *FileService) Get(ctx context.Context, req *fileV1.GetFileRequest) (*fileV1.File, error) {
	return s.uc.Get(ctx, req)
}

func (s *FileService) Create(ctx context.Context, req *fileV1.CreateFileRequest) (*emptypb.Empty, error) {
	if req.Data == nil {
		return nil, adminV1.ErrorBadRequest("错误的参数")
	}

	// 获取操作人信息
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	req.Data.CreateBy = trans.Ptr(operator.UserId)

	if err = s.uc.Create(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *FileService) Update(ctx context.Context, req *fileV1.UpdateFileRequest) (*emptypb.Empty, error) {
	if req.Data == nil {
		return nil, adminV1.ErrorBadRequest("错误的参数")
	}

	// 获取操作人信息
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	req.Data.UpdateBy = trans.Ptr(operator.UserId)

	if err = s.uc.Update(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *FileService) Delete(ctx context.Context, req *fileV1.DeleteFileRequest) (*emptypb.Empty, error) {
	if err := s.uc.Delete(ctx, req); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
