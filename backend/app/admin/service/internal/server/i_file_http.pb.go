package server

import (
	"context"
	"io"
	"strings"

	"github.com/go-kratos/kratos/v2/transport/http"

	"kratos-admin/app/admin/service/internal/service"

	fileV1 "kratos-admin/api/gen/go/file/service/v1"
)

func registerFileUploadHandler(srv *http.Server, svc *service.FileService) {
	r := srv.Route("/")
	r.POST("admin/v1/file:upload", _FileService_PostUploadFile_HTTP_Handler(svc))
	r.PUT("admin/v1/file:upload", _FileService_PutUploadFile_HTTP_Handler(svc))
}

const OperationFileServicePostUploadFile = "/admin.service.v1.FileService/PostUploadFile"
const OperationFileServicePutUploadFile = "/admin.service.v1.FileService/PutUploadFile"

func _FileService_PostUploadFile_HTTP_Handler(svc *service.FileService) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		http.SetOperation(ctx, OperationFileServicePostUploadFile)

		var in fileV1.UploadFileRequest
		var err error

		var aFile *fileV1.File

		file, header, err := ctx.Request().FormFile("file")
		if err == nil {
			defer file.Close()

			b := new(strings.Builder)
			_, err = io.Copy(b, file)

			aFile = &fileV1.File{
				FileName: header.Filename,
				Mime:     header.Header.Get("Content-Type"),
				Content:  []byte(b.String()),
			}
		}

		if err = ctx.BindQuery(&in); err != nil {
			return err
		}

		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return svc.PostUploadFile(ctx, req.(*fileV1.UploadFileRequest), aFile)
		})

		// 逻辑处理，取数据
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}

		reply := out.(*fileV1.UploadFileResponse)

		return ctx.Result(200, reply)
	}
}

func _FileService_PutUploadFile_HTTP_Handler(svc *service.FileService) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		http.SetOperation(ctx, OperationFileServicePutUploadFile)

		var in fileV1.UploadFileRequest
		var err error

		var aFile *fileV1.File

		file, header, err := ctx.Request().FormFile("file")
		if err == nil {
			defer file.Close()

			b := new(strings.Builder)
			_, err = io.Copy(b, file)

			aFile = &fileV1.File{
				FileName: header.Filename,
				Mime:     header.Header.Get("Content-Type"),
				Content:  []byte(b.String()),
			}
		}

		if err = ctx.BindQuery(&in); err != nil {
			return err
		}

		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return svc.PutUploadFile(ctx, req.(*fileV1.UploadFileRequest), aFile)
		})

		// 逻辑处理，取数据
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}

		reply := out.(*fileV1.UploadFileResponse)

		return ctx.Result(200, reply)
	}
}
