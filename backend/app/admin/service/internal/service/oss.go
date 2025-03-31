package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/go-utils/trans"

	adminV1 "kratos-admin/api/gen/go/admin/service/v1"
	fileV1 "kratos-admin/api/gen/go/file/service/v1"

	"kratos-admin/pkg/oss"
)

type OssService struct {
	adminV1.OssServiceHTTPServer

	log *log.Helper

	mc *oss.MinIOClient
}

func NewOssService(logger log.Logger, mc *oss.MinIOClient) *OssService {
	l := log.NewHelper(log.With(logger, "module", "oss/service/admin-service"))
	return &OssService{
		log: l,
		mc:  mc,
	}
}

func (s *OssService) OssUploadUrl(ctx context.Context, req *fileV1.OssUploadUrlRequest) (*fileV1.OssUploadUrlResponse, error) {
	return s.mc.OssUploadUrl(ctx, req)
}

func (s *OssService) PostUploadFile(ctx context.Context, req *fileV1.UploadOssFileRequest, fileData *fileV1.FileData) (*fileV1.UploadOssFileResponse, error) {
	if fileData == nil {
		return nil, fileV1.ErrorUploadFailed("unknown fileData")
	}

	if req.BucketName == nil {
		req.BucketName = trans.Ptr(s.mc.ContentTypeToBucketName(fileData.Mime))
	}
	if req.ObjectName == nil {
		req.ObjectName = trans.Ptr(fileData.FileName)
	}

	downloadUrl, err := s.mc.UploadFile(ctx, req.GetBucketName(), req.GetObjectName(), fileData.Content)
	return &fileV1.UploadOssFileResponse{
		Url: downloadUrl,
	}, err
}

func (s *OssService) PutUploadFile(ctx context.Context, req *fileV1.UploadOssFileRequest, fileData *fileV1.FileData) (*fileV1.UploadOssFileResponse, error) {
	if fileData == nil {
		return nil, fileV1.ErrorUploadFailed("unknown fileData")
	}

	if req.BucketName == nil {
		req.BucketName = trans.Ptr(s.mc.ContentTypeToBucketName(fileData.Mime))
	}
	if req.ObjectName == nil {
		req.ObjectName = trans.Ptr(fileData.FileName)
	}

	downloadUrl, err := s.mc.UploadFile(ctx, req.GetBucketName(), req.GetObjectName(), fileData.Content)
	return &fileV1.UploadOssFileResponse{
		Url: downloadUrl,
	}, err
}
