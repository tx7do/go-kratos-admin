package data

import (
	"github.com/go-kratos/kratos/v2/log"

	gormCrud "github.com/tx7do/go-crud/gorm"

	conf "github.com/tx7do/kratos-bootstrap/api/gen/go/conf/v1"
	gormBootstrap "github.com/tx7do/kratos-bootstrap/database/gorm"
)

// NewGormClient 创建数据库客户端
func NewGormClient(cfg *conf.Bootstrap, logger log.Logger) *gormCrud.Client {
	l := log.NewHelper(log.With(logger, "module", "gorm/data/admin-service"))
	return gormBootstrap.NewGormClient(cfg, l, nil)
}
