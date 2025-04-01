package oss

import (
	"bytes"
	"context"
	"mime"
	"net/url"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/tx7do/go-utils/trans"

	conf "github.com/tx7do/kratos-bootstrap/api/gen/go/conf/v1"
	ossMinio "github.com/tx7do/kratos-bootstrap/oss/minio"

	fileV1 "kratos-admin/api/gen/go/file/service/v1"
)

const (
	defaultExpiryTime = time.Minute * 60 // 1小时
)

type MinIOClient struct {
	mc   *minio.Client
	conf *conf.OSS
	log  *log.Helper
}

func NewMinIoClient(cfg *conf.Bootstrap, logger log.Logger) *MinIOClient {
	l := log.NewHelper(log.With(logger, "module", "minio/data/admin-service"))
	return &MinIOClient{
		log:  l,
		conf: cfg.Oss,
		mc:   ossMinio.NewClient(cfg.Oss),
	}
}

// ContentTypeToBucketName 根据文件类型获取bucket名称
func (c *MinIOClient) ContentTypeToBucketName(contentType string) string {
	h := strings.Split(contentType, "/")
	if len(h) != 2 {
		return "images"
	}

	var bucketName string
	switch h[0] {
	case "image":
		bucketName = "images"
		break
	case "video":
		bucketName = "videos"
		break
	case "audio":
		bucketName = "audios"
		break
	case "application", "text":
		bucketName = "docs"
		break
	default:
		bucketName = "files"
		break
	}

	return bucketName
}

// contentTypeToFileSuffix 根据文件类型获取文件后缀
func (c *MinIOClient) contentTypeToFileSuffix(contentType string) string {
	fileSuffix := ""

	switch contentType {
	case "text/plain":
		fileSuffix = ".txt"
		break

	case "image/jpeg":
		fileSuffix = ".jpg"
		break

	case "image/png":
		fileSuffix = ".png"
		break

	default:
		extensions, _ := mime.ExtensionsByType(contentType)
		if len(extensions) > 0 {
			fileSuffix = extensions[0]
		}
	}

	return fileSuffix
}

// jointObjectName 拼接objectName
func (c *MinIOClient) jointObjectName(contentType string, filePath, fileName *string) (string, string) {
	fileSuffix := c.contentTypeToFileSuffix(contentType)

	var _fileName string
	if fileName == nil {
		_fileName = uuid.New().String() + fileSuffix
	} else {
		_fileName = *fileName
	}

	var objectName string
	if filePath != nil {
		objectName = *filePath + "/" + _fileName
	} else {
		objectName = _fileName
	}

	return objectName, _fileName
}

// OssUploadUrl 获取上传地址
func (c *MinIOClient) OssUploadUrl(ctx context.Context, req *fileV1.OssUploadUrlRequest) (*fileV1.OssUploadUrlResponse, error) {
	var bucketName string
	if req.BucketName != nil {
		bucketName = req.GetBucketName()
	} else {
		bucketName = c.ContentTypeToBucketName(req.GetContentType())
	}

	objectName, _ := c.jointObjectName(req.GetContentType(), req.FilePath, req.FileName)

	expiry := defaultExpiryTime

	var uploadUrl string
	var downloadUrl string
	var formData map[string]string

	var err error
	var presignedURL *url.URL

	switch req.GetMethod() {
	case fileV1.UploadMethod_Put:
		presignedURL, err = c.mc.PresignedPutObject(ctx, bucketName, objectName, expiry)
		if err != nil {
			return nil, err
		}

		uploadUrl = presignedURL.String()
		uploadUrl = strings.Replace(uploadUrl, c.conf.Minio.Endpoint, c.conf.Minio.UploadHost, -1)

		downloadUrl = presignedURL.Host + "/" + bucketName + "/" + objectName
		downloadUrl = strings.Replace(downloadUrl, c.conf.Minio.Endpoint, c.conf.Minio.DownloadHost, -1)
		if !strings.HasPrefix(downloadUrl, presignedURL.Scheme) {
			downloadUrl = presignedURL.Scheme + "://" + downloadUrl
		}

	case fileV1.UploadMethod_Post:
		policy := minio.NewPostPolicy()
		_ = policy.SetBucket(bucketName)
		_ = policy.SetKey(objectName)
		_ = policy.SetExpires(time.Now().UTC().Add(expiry))
		_ = policy.SetContentType(req.GetContentType())

		presignedURL, formData, err = c.mc.PresignedPostPolicy(ctx, policy)
		if err != nil {
			return nil, err
		}

		uploadUrl = presignedURL.String()
		uploadUrl = strings.Replace(uploadUrl, c.conf.Minio.Endpoint, c.conf.Minio.UploadHost, -1)

		downloadUrl = presignedURL.Host + "/" + bucketName + "/" + objectName
		downloadUrl = strings.Replace(downloadUrl, c.conf.Minio.Endpoint, c.conf.Minio.DownloadHost, -1)
		if !strings.HasPrefix(downloadUrl, presignedURL.Scheme) {
			downloadUrl = presignedURL.Scheme + "://" + downloadUrl
		}
	}

	return &fileV1.OssUploadUrlResponse{
		UploadUrl:   uploadUrl,
		DownloadUrl: downloadUrl,
		ObjectName:  objectName,
		BucketName:  trans.Ptr(bucketName),
		FormData:    formData,
	}, nil
}

// ListFile 获取文件夹下面的文件列表
func (c *MinIOClient) ListFile(ctx context.Context, req *fileV1.ListOssFileRequest) (*fileV1.ListOssFileResponse, error) {
	resp := &fileV1.ListOssFileResponse{
		Files: make([]string, 0),
	}
	for object := range c.mc.ListObjects(ctx,
		req.GetBucketName(),
		minio.ListObjectsOptions{
			Prefix:    req.GetFolder(),
			Recursive: req.GetRecursive(),
		},
	) {
		//fmt.Printf("%+v\n", object)
		resp.Files = append(resp.Files, object.Key)
	}
	return resp, nil
}

// ListFileForUEditor 获取文件夹下面的文件列表
func (c *MinIOClient) ListFileForUEditor(ctx context.Context, bucketName string, folder string) (*fileV1.UEditorResponse, error) {
	resp := &fileV1.UEditorResponse{
		State: trans.Ptr("SUCCESS"),
		List:  make([]*fileV1.UEditorResponse_Item, 0),
	}
	for object := range c.mc.ListObjects(ctx,
		bucketName,
		minio.ListObjectsOptions{
			Prefix:    folder,
			Recursive: true,
		},
	) {
		//fmt.Printf("%+v\n", object)
		resp.List = append(resp.List, &fileV1.UEditorResponse_Item{
			Url:   "/" + bucketName + "/" + folder + object.Key,
			Mtime: object.LastModified.Unix(),
		})
	}

	resp.Start = trans.Ptr(int32(0))
	resp.Total = trans.Ptr(int32(len(resp.List)))

	return resp, nil
}

// DeleteFile 删除一个文件
func (c *MinIOClient) DeleteFile(ctx context.Context, req *fileV1.DeleteOssFileRequest) (*fileV1.DeleteOssFileResponse, error) {
	err := c.mc.RemoveObject(ctx, req.GetBucketName(), req.GetObjectName(), minio.RemoveObjectOptions{})
	if err != nil {
		return nil, err
	}

	return &fileV1.DeleteOssFileResponse{}, nil
}

func (c *MinIOClient) UploadFile(ctx context.Context, bucketName string, objectName string, file []byte) (string, error) {
	reader := bytes.NewReader(file)

	_, err := c.mc.PutObject(
		ctx,
		bucketName,
		objectName,
		reader, reader.Size(),
		minio.PutObjectOptions{},
	)
	if err != nil {
		return "", err
	}

	downloadUrl := "/" + bucketName + "/" + objectName

	return downloadUrl, nil
}
