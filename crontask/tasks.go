package crontask

import (
	"time"

	"github.com/go-co-op/gocron"
	"github.com/wubainian/godaemon/crontask/defines"
	"github.com/wubainian/godaemon/global"
	"go.uber.org/zap"
)

var (
	gCrontask []defines.CronTask
)

func AddCronTask(task defines.CronTask) {
	gCrontask = append(gCrontask, task)
}

func StartTasks() {
	scheduler := gocron.NewScheduler(time.Local)
	scheduler.SingletonModeAll() //单例模式，在当前运行完成之前，不会重新安排长时间运行的作业

	//启动zps定时器
	for _, task := range gCrontask {
		task.StartCron(scheduler)
	}

	scheduler.StartAsync()

	//打印job信息
	jobs := scheduler.Jobs()
	for idx, job := range jobs {
		global.GVA_LOG.Info("globals->StartTasks:", zap.Any("idx", idx),
			zap.Any("Tags", job.Tags()),
			zap.Any("IsRunning", job.IsRunning()),
			zap.Any("NextRun", job.NextRun()))
	}
}
