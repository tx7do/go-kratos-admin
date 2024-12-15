package data

import (
	"strings"
	"testing"

	"github.com/go-kratos/kratos/v2/log"

	entgoUpdate "github.com/tx7do/go-utils/entgo/update"
	"github.com/tx7do/go-utils/fieldmaskutil"
	"github.com/tx7do/go-utils/trans"

	"google.golang.org/genproto/protobuf/field_mask"

	systemV1 "kratos-admin/api/gen/go/system/service/v1"
)

func TestMenuMetaFieldMask(t *testing.T) {
	updateMenuReq := &systemV1.UpdateMenuRequest{
		Menu: &systemV1.Menu{
			Meta: &systemV1.RouteMeta{
				Title: trans.Ptr("标题1"),
				Order: trans.Ptr(int32(1)),
			},
		},
		UpdateMask: &field_mask.FieldMask{
			Paths: []string{"id", "meta", "meta.order", "meta.title"},
		},
	}
	var metaPaths []string
	for _, v := range updateMenuReq.UpdateMask.GetPaths() {
		if strings.HasPrefix(v, "meta.") {
			metaPaths = append(metaPaths, strings.SplitAfter(v, "meta.")[1])
		}
	}
	updateMenuReq.UpdateMask.Normalize()
	if !updateMenuReq.UpdateMask.IsValid(updateMenuReq.Menu) {
		// Return an error.
		panic("invalid field mask")
	}
	fieldmaskutil.Filter(updateMenuReq.GetMenu(), updateMenuReq.UpdateMask.GetPaths())

	fieldmaskutil.Filter(updateMenuReq.GetMenu().Meta, metaPaths)

	nilPaths := fieldmaskutil.NilValuePaths(updateMenuReq.GetMenu().Meta, metaPaths)
	keyValues := entgoUpdate.ExtractJsonFieldKeyValues(updateMenuReq.GetMenu().Meta, metaPaths, false)

	log.Infof("UPDATE: [%v] [%v] [%v] [%v]", updateMenuReq.Menu, updateMenuReq.Menu.Meta, nilPaths, keyValues)
}
