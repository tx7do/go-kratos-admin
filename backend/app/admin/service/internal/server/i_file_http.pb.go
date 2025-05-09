package server

import (
	"context"
	"io"
	"strings"

	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/tx7do/go-utils/trans"

	"kratos-admin/app/admin/service/internal/service"

	fileV1 "kratos-admin/api/gen/go/file/service/v1"
)

func registerFileUploadHandler(srv *http.Server, svc *service.OssService) {
	r := srv.Route("/")
	r.POST("admin/v1/file:upload", _FileService_PostUploadFile_HTTP_Handler(svc))
	r.PUT("admin/v1/file:upload", _FileService_PutUploadFile_HTTP_Handler(svc))
}

const OperationFileServicePostUploadFile = "/admin.service.v1.FileService/PostUploadFile"
const OperationFileServicePutUploadFile = "/admin.service.v1.FileService/PutUploadFile"

func _FileService_PostUploadFile_HTTP_Handler(svc *service.OssService) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		http.SetOperation(ctx, OperationFileServicePostUploadFile)

		var in fileV1.UploadOssFileRequest
		var err error

		file, header, err := ctx.Request().FormFile("file")
		if err == nil {
			defer file.Close()

			b := new(strings.Builder)
			_, err = io.Copy(b, file)

			in.SourceFileName = trans.Ptr(header.Filename)
			in.Mime = trans.Ptr(header.Header.Get("Content-Type"))
			in.File = []byte(b.String())
		}

		if err = ctx.BindQuery(&in); err != nil {
			return err
		}

		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			aReq := req.(*fileV1.UploadOssFileRequest)

			var resp *fileV1.UploadOssFileResponse
			resp, err = svc.PostUploadFile(ctx, aReq)
			in.File = nil

			return resp, err
		})

		// 逻辑处理，取数据
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}

		reply := out.(*fileV1.UploadOssFileResponse)

		return ctx.Result(200, reply)
	}
}

func _FileService_PutUploadFile_HTTP_Handler(svc *service.OssService) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		http.SetOperation(ctx, OperationFileServicePutUploadFile)

		var in fileV1.UploadOssFileRequest
		var err error

		file, header, err := ctx.Request().FormFile("file")
		if err == nil {
			defer file.Close()

			b := new(strings.Builder)
			_, err = io.Copy(b, file)

			in.SourceFileName = trans.Ptr(header.Filename)
			in.Mime = trans.Ptr(header.Header.Get("Content-Type"))
			in.File = []byte(b.String())
		}

		if err = ctx.BindQuery(&in); err != nil {
			return err
		}

		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			aReq := req.(*fileV1.UploadOssFileRequest)

			var resp *fileV1.UploadOssFileResponse
			resp, err = svc.PutUploadFile(ctx, aReq)
			in.File = nil

			return resp, err
		})

		// 逻辑处理，取数据
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}

		reply := out.(*fileV1.UploadOssFileResponse)

		return ctx.Result(200, reply)
	}
}
