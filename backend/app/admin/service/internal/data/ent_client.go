package data

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/tx7do/kratos-bootstrap/bootstrap"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/lib/pq"

	"github.com/go-kratos/kratos/v2/log"

	entCrud "github.com/tx7do/go-crud/entgo"

	entBootstrap "github.com/tx7do/kratos-bootstrap/database/ent"

	"go-wind-admin/app/admin/service/internal/data/ent"
	"go-wind-admin/app/admin/service/internal/data/ent/migrate"
)

// NewEntClient 创建Ent ORM数据库客户端
func NewEntClient(ctx *bootstrap.Context) *entCrud.EntClient[*ent.Client] {
	l := log.NewHelper(log.With(ctx.Logger, "module", "ent/data/admin-service"))

	return entBootstrap.NewEntClient(ctx.Config, func(drv *sql.Driver) *ent.Client {
		client := ent.NewClient(
			ent.Driver(drv),
			ent.Log(func(a ...any) {
				l.Debug(a...)
			}),
		)
		if client == nil {
			l.Fatalf("failed creating ent client")
			return nil
		}

		// 运行数据库迁移工具
		if ctx.Config.Data.Database.GetMigrate() {
			if err := client.Schema.Create(context.Background(), migrate.WithForeignKeys(true)); err != nil {
				l.Fatalf("failed creating schema resources: %v", err)
			}
		}

		return client
	})
}
