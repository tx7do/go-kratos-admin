package data

import (
	"strings"
	"testing"

	"github.com/go-kratos/kratos/v2/log"

	entgoUpdate "github.com/tx7do/go-utils/entgo/update"
	"github.com/tx7do/go-utils/fieldmaskutil"
	"github.com/tx7do/go-utils/trans"

	"google.golang.org/genproto/protobuf/field_mask"

	adminV1 "kratos-admin/api/gen/go/admin/service/v1"
)

func TestMenuMetaFieldMask(t *testing.T) {
	updateMenuReq := &adminV1.UpdateMenuRequest{
		Data: &adminV1.Menu{
			Meta: &adminV1.RouteMeta{
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
	if !updateMenuReq.UpdateMask.IsValid(updateMenuReq.Data) {
		// Return an error.
		panic("invalid field mask")
	}
	fieldmaskutil.Filter(updateMenuReq.GetData(), updateMenuReq.UpdateMask.GetPaths())

	fieldmaskutil.Filter(updateMenuReq.GetData().Meta, metaPaths)

	nilPaths := fieldmaskutil.NilValuePaths(updateMenuReq.GetData().Meta, metaPaths)
	keyValues := entgoUpdate.ExtractJsonFieldKeyValues(updateMenuReq.GetData().Meta, metaPaths, false)

	log.Infof("UPDATE: [%v] [%v] [%v] [%v]", updateMenuReq.Data, updateMenuReq.Data.Meta, nilPaths, keyValues)
}
