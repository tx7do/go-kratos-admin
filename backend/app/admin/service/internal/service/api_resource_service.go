package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"

	"entgo.io/ent/dialect/sql"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/tx7do/go-utils/trans"
	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	"google.golang.org/protobuf/types/known/emptypb"

	"kratos-admin/app/admin/service/cmd/server/assets"
	"kratos-admin/app/admin/service/internal/data"

	adminV1 "kratos-admin/api/gen/go/admin/service/v1"

	"kratos-admin/pkg/middleware/auth"
)

type ApiResourceService struct {
	adminV1.ApiResourceServiceHTTPServer

	RestServer *http.Server

	log *log.Helper

	repo       *data.ApiResourceRepo
	authorizer *data.Authorizer
}

func NewApiResourceService(
	logger log.Logger,
	repo *data.ApiResourceRepo,
	authorizer *data.Authorizer,
) *ApiResourceService {
	l := log.NewHelper(log.With(logger, "module", "api-resource/service/admin-service"))
	svc := &ApiResourceService{
		log:        l,
		repo:       repo,
		authorizer: authorizer,
	}

	svc.init()

	return svc
}

func (s *ApiResourceService) init() {
	ctx := context.Background()
	if count, _ := s.repo.Count(ctx, []func(s *sql.Selector){}); count == 0 {
		_, _ = s.SyncApiResources(ctx, &emptypb.Empty{})
	}
}

func (s *ApiResourceService) List(ctx context.Context, req *pagination.PagingRequest) (*adminV1.ListApiResourceResponse, error) {
	return s.repo.List(ctx, req)
}

func (s *ApiResourceService) Get(ctx context.Context, req *adminV1.GetApiResourceRequest) (*adminV1.ApiResource, error) {
	return s.repo.Get(ctx, req)
}

func (s *ApiResourceService) Create(ctx context.Context, req *adminV1.CreateApiResourceRequest) (*emptypb.Empty, error) {
	if req.Data == nil {
		return nil, adminV1.ErrorBadRequest("invalid parameter")
	}

	// 获取操作人信息
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	req.Data.CreateBy = trans.Ptr(operator.UserId)

	if err = s.repo.Create(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *ApiResourceService) Update(ctx context.Context, req *adminV1.UpdateApiResourceRequest) (*emptypb.Empty, error) {
	if req.Data == nil {
		return nil, adminV1.ErrorBadRequest("invalid parameter")
	}

	// 获取操作人信息
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	req.Data.UpdateBy = trans.Ptr(operator.UserId)

	if err = s.repo.Update(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *ApiResourceService) Delete(ctx context.Context, req *adminV1.DeleteApiResourceRequest) (*emptypb.Empty, error) {
	if err := s.repo.Delete(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *ApiResourceService) SyncApiResources(ctx context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	_ = s.repo.Truncate(ctx)

	//if err := s.syncWithWalkRoute(ctx); err != nil {
	//	return nil, err
	//}

	if err := s.syncWithOpenAPI(ctx); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *ApiResourceService) syncWithOpenAPI(ctx context.Context) error {
	loader := openapi3.NewLoader()
	doc, err := loader.LoadFromData(assets.OpenApiData)
	if err != nil {
		s.log.Fatalf("加载 OpenAPI 文档失败: %v", err)
		return adminV1.ErrorInternalServerError("load OpenAPI document failed")
	}

	if doc == nil {
		s.log.Fatal("OpenAPI 文档为空")
		return adminV1.ErrorInternalServerError("OpenAPI document is nil")
	}
	if doc.Paths == nil {
		s.log.Fatal("OpenAPI 文档的路径为空")
		return adminV1.ErrorInternalServerError("OpenAPI document paths is nil")
	}

	var count uint32 = 0
	var module string
	var moduleDescription string

	// 遍历所有路径和操作
	for path, pathItem := range doc.Paths.Map() {
		for method, operation := range pathItem.Operations() {

			if len(operation.Tags) > 0 {
				tag := doc.Tags.Get(operation.Tags[0])
				module = tag.Name
				moduleDescription = tag.Description
			}

			count++
			_ = s.repo.Update(ctx, &adminV1.UpdateApiResourceRequest{
				AllowMissing: trans.Ptr(true),
				Data: &adminV1.ApiResource{
					Id:                trans.Ptr(count),
					Path:              trans.Ptr(path),
					Method:            trans.Ptr(method),
					Module:            trans.Ptr(module),
					ModuleDescription: trans.Ptr(moduleDescription),
					Description:       trans.Ptr(operation.Description),
					Operation:         trans.Ptr(operation.OperationID),
				},
			})
		}
	}

	return nil
}

func (s *ApiResourceService) syncWithWalkRoute(ctx context.Context) error {
	if s.RestServer == nil {
		return adminV1.ErrorInternalServerError("rest server is nil")
	}

	var count uint32 = 0
	if err := s.RestServer.WalkRoute(func(info http.RouteInfo) error {
		//log.Infof("Path[%s] Method[%s]", info.Path, info.Method)
		count++
		_ = s.repo.Update(ctx, &adminV1.UpdateApiResourceRequest{
			AllowMissing: trans.Ptr(true),
			Data: &adminV1.ApiResource{
				Id:     trans.Ptr(count),
				Path:   trans.Ptr(info.Path),
				Method: trans.Ptr(info.Method),
			},
		})

		return nil
	}); err != nil {
		s.log.Errorf("failed to walk route: %v", err)
		return adminV1.ErrorInternalServerError("failed to walk route")
	}

	return nil
}

func (s *ApiResourceService) GetWalkRouteData(_ context.Context, _ *emptypb.Empty) (*adminV1.ListApiResourceResponse, error) {
	if s.RestServer == nil {
		return nil, adminV1.ErrorInternalServerError("rest server is nil")
	}

	resp := &adminV1.ListApiResourceResponse{
		Items: []*adminV1.ApiResource{},
	}
	var count uint32 = 0
	if err := s.RestServer.WalkRoute(func(info http.RouteInfo) error {
		//log.Infof("Path[%s] Method[%s]", info.Path, info.Method)
		count++
		resp.Items = append(resp.Items, &adminV1.ApiResource{
			Id:     trans.Ptr(count),
			Path:   trans.Ptr(info.Path),
			Method: trans.Ptr(info.Method),
		})
		return nil
	}); err != nil {
		s.log.Errorf("failed to walk route: %v", err)
		return nil, adminV1.ErrorInternalServerError("failed to walk route")
	}
	resp.Total = count

	return resp, nil
}
