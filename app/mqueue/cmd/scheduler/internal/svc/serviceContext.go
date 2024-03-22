package svc

import (
	"github.com/hibiken/asynq"
	"golodge/app/mqueue/cmd/scheduler/internal/config"
)

type ServiceContext struct {
	Config config.Config

	Scheduler *asynq.Scheduler
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Scheduler:newScheduler(c),
	}
}

