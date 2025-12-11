package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/go-utils/trans"

	adminV1 "go-wind-admin/api/gen/go/admin/service/v1"
	fileV1 "go-wind-admin/api/gen/go/file/service/v1"

	"go-wind-admin/pkg/oss"
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

func (s *OssService) PostUploadFile(ctx context.Context, req *fileV1.UploadOssFileRequest) (*fileV1.UploadOssFileResponse, error) {
	if req.File == nil {
		return nil, fileV1.ErrorUploadFailed("unknown fileData")
	}

	if req.BucketName == nil {
		req.BucketName = trans.Ptr(s.mc.ContentTypeToBucketName(req.GetMime()))
	}
	if req.ObjectName == nil {
		req.ObjectName = trans.Ptr(req.GetSourceFileName())
	}

	downloadUrl, err := s.mc.UploadFile(ctx, req.GetBucketName(), req.GetObjectName(), req.GetFile())
	return &fileV1.UploadOssFileResponse{
		Url: downloadUrl,
	}, err
}

func (s *OssService) PutUploadFile(ctx context.Context, req *fileV1.UploadOssFileRequest) (*fileV1.UploadOssFileResponse, error) {
	if req.File == nil {
		return nil, fileV1.ErrorUploadFailed("unknown fileData")
	}

	if req.BucketName == nil {
		req.BucketName = trans.Ptr(s.mc.ContentTypeToBucketName(req.GetMime()))
	}
	if req.ObjectName == nil {
		req.ObjectName = trans.Ptr(req.GetSourceFileName())
	}

	downloadUrl, err := s.mc.UploadFile(ctx, req.GetBucketName(), req.GetObjectName(), req.GetFile())
	return &fileV1.UploadOssFileResponse{
		Url: downloadUrl,
	}, err
}
