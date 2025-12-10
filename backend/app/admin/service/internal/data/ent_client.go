package data

import (
	"entgo.io/ent/dialect/sql"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/lib/pq"

	"github.com/go-kratos/kratos/v2/log"

	entCrud "github.com/tx7do/go-crud/entgo"

	conf "github.com/tx7do/kratos-bootstrap/api/gen/go/conf/v1"
	entBootstrap "github.com/tx7do/kratos-bootstrap/database/ent"

	"kratos-admin/app/admin/service/internal/data/ent"
)

// NewEntClient 创建Ent ORM数据库客户端
func NewEntClient(cfg *conf.Bootstrap, logger log.Logger) *entCrud.EntClient[*ent.Client] {
	l := log.NewHelper(log.With(logger, "module", "ent/data/admin-service"))

	return entBootstrap.NewEntClient(cfg, func(drv *sql.Driver) *ent.Client {
		return ent.NewClient(
			ent.Driver(drv),
			ent.Log(func(a ...any) {
				l.Debug(a...)
			}),
		)
	})
}
