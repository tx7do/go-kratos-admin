package data

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/lib/pq"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/go-utils/entgo"

	"kratos-admin/app/admin/service/internal/data/ent"
	"kratos-admin/app/admin/service/internal/data/ent/migrate"

	conf "github.com/tx7do/kratos-bootstrap/api/gen/go/conf/v1"
)

// NewEntClient 创建Ent ORM数据库客户端
func NewEntClient(cfg *conf.Bootstrap, logger log.Logger) *entgo.EntClient[*ent.Client] {
	l := log.NewHelper(log.With(logger, "module", "ent/data/admin-service"))

	drv, err := entgo.CreateDriver(
		cfg.Data.Database.GetDriver(),
		cfg.Data.Database.GetSource(),
		cfg.Data.Database.GetEnableTrace(),
		cfg.Data.Database.GetEnableMetrics(),
	)
	if err != nil {
		l.Fatalf("failed opening connection to db: %v", err)
		return nil
	}

	client := ent.NewClient(
		ent.Driver(drv),
		ent.Log(func(a ...any) {
			l.Debug(a...)
		}),
	)

	// 运行数据库迁移工具
	if cfg.Data.Database.GetMigrate() {
		if err = client.Schema.Create(context.Background(), migrate.WithForeignKeys(true)); err != nil {
			l.Fatalf("failed creating schema resources: %v", err)
		}
	}

	cli := entgo.NewEntClient(client, drv)

	cli.SetConnectionOption(
		int(cfg.Data.Database.GetMaxIdleConnections()),
		int(cfg.Data.Database.GetMaxOpenConnections()),
		cfg.Data.Database.GetConnectionMaxLifetime().AsDuration(),
	)

	return cli
}

// queryAllChildrenIDs 使用CTE递归查询所有子节点ID
func queryAllChildrenIDs(ctx context.Context, entClient *entgo.EntClient[*ent.Client], tableName string, parentID uint32) ([]uint32, error) {
	var query string
	switch entClient.Driver().Dialect() {
	case dialect.MySQL:
		query = fmt.Sprintf(`
			WITH RECURSIVE all_descendants AS (
				SELECT 
					id,
					parent_id,
					name,
					1 AS depth
				FROM %s
				WHERE parent_id = ?
				
				UNION ALL
				
				SELECT 
					p.id,
					p.parent_id,
					p.name,
					ad.depth + 1 AS depth
				FROM %s p
				INNER JOIN all_descendants ad 
					ON p.parent_id = ad.id
			)
			SELECT id FROM all_descendants;
		`, tableName, tableName)

	case dialect.Postgres:
		query = fmt.Sprintf(`
        WITH RECURSIVE all_descendants AS (
            SELECT * FROM %s WHERE parent_id = $1
            UNION ALL
            SELECT p.* FROM %s p
            INNER JOIN all_descendants ad ON p.parent_id = ad.id
        )
        SELECT id FROM all_descendants;
    `, tableName, tableName)
	}

	rows := &sql.Rows{}
	if err := entClient.Query(ctx, query, []any{parentID}, rows); err != nil {
		return nil, errors.New("query child nodes failed: " + err.Error())
	}
	defer rows.Close()

	childIDs := make([]uint32, 0)
	for rows.Next() {
		var id uint32

		if err := rows.Scan(&id); err != nil {
			log.Errorf("scan child node failed: %s", err.Error())
			return nil, errors.New("scan child node failed")
		}

		childIDs = append(childIDs, id)
	}

	return childIDs, nil
}
