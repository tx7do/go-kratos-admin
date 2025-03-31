package server

import (
	"context"
	"io"
	"strings"

	"github.com/go-kratos/kratos/v2/transport/http"

	"kratos-admin/app/admin/service/internal/service"

	fileV1 "kratos-admin/api/gen/go/file/service/v1"
)

func registerUEditorUploadHandler(srv *http.Server, svc *service.UEditorService) {
	r := srv.Route("/")
	r.POST("admin/v1/ueditor", _UEditorService_UploadFile_HTTP_Handler(svc))
	r.PUT("admin/v1/ueditor", _UEditorService_UploadFile_HTTP_Handler(svc))
}

const OperationUEditorServiceUploadFile = "/admin.service.v1.UEditorService/UploadFile"

func _UEditorService_UploadFile_HTTP_Handler(svc *service.UEditorService) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		http.SetOperation(ctx, OperationUEditorServiceUploadFile)

		var in fileV1.UEditorUploadRequest
		var err error

		var fileData *fileV1.FileData

		file, header, err := ctx.Request().FormFile("file")
		if err == nil {
			defer file.Close()

			b := new(strings.Builder)
			_, err = io.Copy(b, file)

			fileData = &fileV1.FileData{
				FileName: header.Filename,
				Mime:     header.Header.Get("Content-Type"),
				Content:  []byte(b.String()),
			}
		}

		if err = ctx.BindQuery(&in); err != nil {
			return err
		}

		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return svc.UploadFile(ctx, req.(*fileV1.UEditorUploadRequest), fileData)
		})

		// 逻辑处理，取数据
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}

		reply := out.(*fileV1.UEditorUploadResponse)

		return ctx.Result(200, reply)
	}
}
