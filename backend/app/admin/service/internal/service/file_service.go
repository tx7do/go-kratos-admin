package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	pagination "github.com/tx7do/go-crud/api/gen/go/pagination/v1"
	"github.com/tx7do/go-utils/trans"
	"github.com/tx7do/kratos-bootstrap/bootstrap"
	"google.golang.org/protobuf/types/known/emptypb"

	"go-wind-admin/app/admin/service/internal/data"

	adminV1 "go-wind-admin/api/gen/go/admin/service/v1"
	fileV1 "go-wind-admin/api/gen/go/file/service/v1"

	"go-wind-admin/pkg/middleware/auth"
)

type FileService struct {
	adminV1.FileServiceHTTPServer

	log *log.Helper

	fileRepo *data.FileRepo
}

func NewFileService(ctx *bootstrap.Context, repo *data.FileRepo) *FileService {
	l := log.NewHelper(log.With(ctx.Logger, "module", "file/service/admin-service"))
	return &FileService{
		log:      l,
		fileRepo: repo,
	}
}

func (s *FileService) List(ctx context.Context, req *pagination.PagingRequest) (*fileV1.ListFileResponse, error) {
	return s.fileRepo.List(ctx, req)
}

func (s *FileService) Get(ctx context.Context, req *fileV1.GetFileRequest) (*fileV1.File, error) {
	return s.fileRepo.Get(ctx, req)
}

func (s *FileService) Create(ctx context.Context, req *fileV1.CreateFileRequest) (*emptypb.Empty, error) {
	if req.Data == nil {
		return nil, adminV1.ErrorBadRequest("invalid parameter")
	}

	// 获取操作人信息
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return nil, err
	}

	req.Data.CreatedBy = trans.Ptr(operator.UserId)

	if err = s.fileRepo.Create(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *FileService) Update(ctx context.Context, req *fileV1.UpdateFileRequest) (*emptypb.Empty, error) {
	if req.Data == nil {
		return nil, adminV1.ErrorBadRequest("invalid parameter")
	}

	// 获取操作人信息
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return nil, err
	}

	req.Data.UpdatedBy = trans.Ptr(operator.UserId)

	if err = s.fileRepo.Update(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *FileService) Delete(ctx context.Context, req *fileV1.DeleteFileRequest) (*emptypb.Empty, error) {
	if err := s.fileRepo.Delete(ctx, req); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
