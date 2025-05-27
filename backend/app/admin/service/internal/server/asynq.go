package server

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	conf "github.com/tx7do/kratos-bootstrap/api/gen/go/conf/v1"
	"github.com/tx7do/kratos-transport/transport/asynq"

	"kratos-admin/app/admin/service/internal/service"

	"kratos-admin/pkg/task"
)

// NewAsynqServer creates a new asynq server.
func NewAsynqServer(cfg *conf.Bootstrap, _ log.Logger, svc *service.TaskService) *asynq.Server {
	if cfg == nil || cfg.Server == nil || cfg.Server.Asynq == nil {
		return nil
	}

	srv := asynq.NewServer(
		asynq.WithRedisURI(cfg.Server.Asynq.GetUri()),
		asynq.WithLocation(cfg.Server.Asynq.GetLocation()),
		asynq.WithEnableKeepAlive(false),
		asynq.WithGracefullyShutdown(true),
		asynq.WithShutdownTimeout(3*time.Second),
	)

	svc.Server = srv

	var err error

	// 注册任务
	if err = asynq.RegisterSubscriber(srv, task.BackupTaskType, svc.AsyncBackup); err != nil {
		log.Error(err)
	}

	// 启动所有的任务
	_, _ = svc.StartAllTask(context.Background())

	return srv
}
