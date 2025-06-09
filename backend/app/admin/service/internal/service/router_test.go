package service

import (
	"fmt"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"

	"kratos-admin/app/admin/service/cmd/server/assets"
)

func TestMenuListToQueryString(t *testing.T) {
	type args struct {
		menus []uint32
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Single menu ID",
			args: args{menus: []uint32{1}},
			want: `{"id__in":"[\"1\"]","status":"ON","type__not":"BUTTON"}`,
		},
		{
			name: "Multiple menu IDs",
			args: args{menus: []uint32{1, 2, 3}},
			want: `{"id__in":"[\"1\", \"2\", \"3\"]","status":"ON","type__not":"BUTTON"}`,
		},
		{
			name: "No menu IDs",
			args: args{menus: []uint32{}},
			want: `{"id__in":"[]","status":"ON","type__not":"BUTTON"}`,
		},
		{
			name: "Large menu IDs",
			args: args{menus: []uint32{1234567890, 987654321}},
			want: `{"id__in":"[\"1234567890\", \"987654321\"]","status":"ON","type__not":"BUTTON"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &RouterService{}
			if got := s.menuListToQueryString(tt.args.menus, false); got != tt.want {
				t.Errorf("menuListToQueryString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOpenAPI(t *testing.T) {
	// 加载 OpenAPI V3 文档
	loader := openapi3.NewLoader()
	doc, err := loader.LoadFromData(assets.OpenApiData)
	if err != nil {
		t.Fatalf("加载 OpenAPI 文档失败: %v", err)
	}

	if doc == nil {
		t.Fatal("OpenAPI 文档为空")
	}
	if doc.Paths == nil {
		t.Fatal("OpenAPI 文档的路径为空")
	}

	// 遍历所有路径和操作
	for path, pathItem := range doc.Paths.Map() {
		fmt.Printf("路径: %s\n", path)
		for method, operation := range pathItem.Operations() {
			fmt.Printf("  方法: %s\n", method)
			if operation.Summary != "" {
				fmt.Printf("    摘要: %s\n", operation.Summary)
			}
			if operation.Description != "" {
				fmt.Printf("    描述: %s\n", operation.Description)
			}
			if len(operation.Tags) > 0 {
				tag := doc.Tags.Get(operation.Tags[0])
				fmt.Printf("    服务: %s %s\n", tag.Name, tag.Description)
			}
		}
	}
}
