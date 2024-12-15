package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/kratos-transport/transport/asynq"

	"kratos-admin/app/admin/service/internal/data"

	"kratos-admin/pkg/task"
)

type TaskService struct {
	log *log.Helper

	Server *asynq.Server

	userRepo *data.UserRepo
}

func NewTaskService(
	logger log.Logger,
	userRepo *data.UserRepo,
) *TaskService {
	l := log.NewHelper(log.With(logger, "module", "task/service/admin-service"))
	return &TaskService{
		log:      l,
		userRepo: userRepo,
	}
}

// StartAllPeriodicTask 启动所有的定时任务
func (s *TaskService) StartAllPeriodicTask() {
	_, _ = s.Server.NewPeriodicTask("*/1 * * * ?", task.BackupTaskType, task.BackupTaskData{Name: "test"})
}

// StartAllDelayTask 启动所有的延迟任务
func (s *TaskService) StartAllDelayTask() {

}

func (s *TaskService) AsyncBackup(taskType string, taskData *task.BackupTaskData) error {
	s.log.Infof("AsyncBackup [%s] [%+v]", taskType, taskData)
	return nil
}
