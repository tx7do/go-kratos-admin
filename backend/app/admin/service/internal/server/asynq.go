package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/kratos-transport/transport/asynq"

	conf "github.com/tx7do/kratos-bootstrap/api/gen/go/conf/v1"

	"kratos-admin/app/admin/service/internal/service"

	"kratos-admin/pkg/task"
)

func NewAsynqServer(cfg *conf.Bootstrap, _ log.Logger, svc *service.TaskService) *asynq.Server {
	srv := asynq.NewServer(
		asynq.WithAddress(cfg.Server.Asynq.GetEndpoint()),
		asynq.WithRedisPassword(cfg.Server.Asynq.GetPassword()),
		asynq.WithRedisDatabase(int(cfg.Server.Asynq.GetDb())),
		asynq.WithLocation(cfg.Server.Asynq.Location),
	)

	svc.Server = srv

	var err error
	if err = asynq.RegisterSubscriber(srv, task.BackupTaskType, svc.AsyncBackup); err != nil {
		log.Error(err)
	}

	svc.StartAllPeriodicTask()
	svc.StartAllDelayTask()

	return srv
}
