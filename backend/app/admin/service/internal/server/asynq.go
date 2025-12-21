package server

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/kratos-bootstrap/bootstrap"
	"github.com/tx7do/kratos-transport/transport/asynq"
	"google.golang.org/protobuf/types/known/emptypb"

	"go-wind-admin/app/admin/service/internal/service"

	"go-wind-admin/pkg/task"
)

// NewAsynqServer creates a new asynq server.
func NewAsynqServer(ctx *bootstrap.Context, svc *service.TaskService) *asynq.Server {
	if ctx.Config == nil || ctx.Config.Server == nil || ctx.Config.Server.Asynq == nil {
		return nil
	}

	srv := asynq.NewServer(
		asynq.WithCodec(ctx.Config.Server.Asynq.GetCodec()),
		asynq.WithRedisURI(ctx.Config.Server.Asynq.GetUri()),
		asynq.WithLocation(ctx.Config.Server.Asynq.GetLocation()),
		asynq.WithGracefullyShutdown(ctx.Config.Server.Asynq.GetEnableGracefullyShutdown()),
		asynq.WithShutdownTimeout(ctx.Config.Server.Asynq.GetShutdownTimeout().AsDuration()),
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
