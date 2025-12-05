package data

import (
	"github.com/go-kratos/kratos/v2/log"
	conf "github.com/tx7do/kratos-bootstrap/api/gen/go/conf/v1"
	"gorm.io/gorm"
)

// NewGormClient 创建数据库客户端
func NewGormClient(cfg *conf.Bootstrap, logger log.Logger) *gorm.DB {
	return nil
}
