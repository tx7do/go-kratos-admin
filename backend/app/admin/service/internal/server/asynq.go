package server

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	conf "github.com/tx7do/kratos-bootstrap/api/gen/go/conf/v1"
	"github.com/tx7do/kratos-transport/transport/asynq"
	"google.golang.org/protobuf/types/known/emptypb"

	"kratos-admin/app/admin/service/internal/service"

	"kratos-admin/pkg/task"
)

// NewAsynqServer creates a new asynq server.
func NewAsynqServer(cfg *conf.Bootstrap, _ log.Logger, svc *service.TaskService) *asynq.Server {
	if cfg == nil || cfg.Server == nil || cfg.Server.Asynq == nil {
		return nil
	}

	srv := asynq.NewServer(
		asynq.WithCodec(cfg.Server.Asynq.GetCodec()),
		asynq.WithRedisURI(cfg.Server.Asynq.GetUri()),
		asynq.WithLocation(cfg.Server.Asynq.GetLocation()),
		asynq.WithGracefullyShutdown(cfg.Server.Asynq.GetEnableGracefullyShutdown()),
		asynq.WithShutdownTimeout(cfg.Server.Asynq.GetShutdownTimeout().AsDuration()),
	)

	svc.Server = srv

	var err error

	// 注册任务
	if err = asynq.RegisterSubscriber(srv, task.BackupTaskType, svc.AsyncBackup); err != nil {
		log.Error(err)
	}

	// 启动所有的任务
	_, _ = svc.StartAllTask(context.Background(), &emptypb.Empty{})

	return srv
}
