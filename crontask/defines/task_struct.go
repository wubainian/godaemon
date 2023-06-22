package defines

import "github.com/go-co-op/gocron"

type TaskItem struct {
	Name      string
	Task      func(params []string)
	CronLabel string
	Params    []string
}

type CronTask interface {
	StartCron(scheduler *gocron.Scheduler)
}
