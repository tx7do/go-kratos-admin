package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/go-utils/trans"

	adminV1 "kratos-admin/api/gen/go/admin/service/v1"
	fileV1 "kratos-admin/api/gen/go/file/service/v1"

	"kratos-admin/pkg/oss"
)

type FileService struct {
	adminV1.FileServiceHTTPServer

	log *log.Helper

	mc *oss.MinIOClient
}

func NewFileService(logger log.Logger, mc *oss.MinIOClient) *FileService {
	l := log.NewHelper(log.With(logger, "module", "file/service/admin-service"))
	return &FileService{
		log: l,
		mc:  mc,
	}
}

func (s *FileService) OssUploadUrl(ctx context.Context, req *fileV1.OssUploadUrlRequest) (*fileV1.OssUploadUrlResponse, error) {
	return s.mc.OssUploadUrl(ctx, req)
}

func (s *FileService) PostUploadFile(ctx context.Context, req *fileV1.UploadFileRequest, file *fileV1.File) (*fileV1.UploadFileResponse, error) {
	if file == nil {
		return nil, fileV1.ErrorUploadFailed("unknown file")
	}

	if req.BucketName == nil {
		req.BucketName = trans.Ptr(s.mc.ContentTypeToBucketName(file.Mime))
	}
	if req.ObjectName == nil {
		req.ObjectName = trans.Ptr(file.FileName)
	}

	downloadUrl, err := s.mc.UploadFile(ctx, req.GetBucketName(), req.GetObjectName(), file.Content)
	return &fileV1.UploadFileResponse{
		Url: downloadUrl,
	}, err
}

func (s *FileService) PutUploadFile(ctx context.Context, req *fileV1.UploadFileRequest, file *fileV1.File) (*fileV1.UploadFileResponse, error) {
	if file == nil {
		return nil, fileV1.ErrorUploadFailed("unknown file")
	}

	if req.BucketName == nil {
		req.BucketName = trans.Ptr(s.mc.ContentTypeToBucketName(file.Mime))
	}
	if req.ObjectName == nil {
		req.ObjectName = trans.Ptr(file.FileName)
	}

	downloadUrl, err := s.mc.UploadFile(ctx, req.GetBucketName(), req.GetObjectName(), file.Content)
	return &fileV1.UploadFileResponse{
		Url: downloadUrl,
	}, err
}
