package cronjobs

import (
	"context"
	"github.com/robfig/cron/v3"
	"kanban/internal/config"
	"kanban/internal/cronjobs/workers"
	"kanban/internal/domain/models"
	taskservice "kanban/internal/domain/usecases/services/task_service"
	"kanban/internal/lib/workerpool"
	"log/slog"
)

type CronJob struct {
	cron           *cron.Cron
	taskService    taskservice.TaskService
	logger         *slog.Logger
	taskWorkerPool *workerpool.WorkerPool
}

// NewCronJob создает новый экземпляр CronJob
func NewCronJob(
	taskService taskservice.TaskService,
	logger *slog.Logger, workerCount int) *CronJob {
	return &CronJob{
		cron:           cron.New(),
		taskService:    taskService,
		logger:         logger,
		taskWorkerPool: workerpool.NewWorkerPool(workerCount),
	}
}

// Start запускает все cron job
func (c *CronJob) Start(cfg *config.Cronjob) {
	const op = "CronJobs.CronJob.Start"
	log := c.logger.With("op", op)

	log.Info("Starting cron job")
	c.taskWorkerPool.Start(workers.NewTaskJobProcessor(c.taskService, c.logger))

	if cfg.TaskJob {
		_, err := c.cron.AddFunc(cfg.TaskSpec, func() {
			err := c.processWaitingTasks()
			if err != nil {
				log.Error("failed to process waiting checks", "error", err)
			}
		})

		//TODO: можно добавлять новые кронджобы

		if err != nil {
			log.Error(op, "failed to start cron job")
			panic(err)
		}
		log.Info("InitProJob started")
	}
	c.cron.Start()
}

// Stop останавливает все cron job
func (c *CronJob) Stop() {
	const op = "CronJobs.CronJob.Stop"
	log := c.logger.With("op", op)

	log.Info("Stopping cron job")
	c.cron.Stop()
	log.Info("Stopping worker pool")
	c.taskWorkerPool.Stop()
}

func (c *CronJob) processWaitingTasks() error {
	const op = "CronJobs.CronJob.processWaitingTasks"
	log := c.logger.With("op", op)
	ctx := context.Background()
	taskFilter := models.TaskFilters{}
	tasks, _, err := c.taskService.GetTaskList(ctx, taskFilter, 0, 20)
	if err != nil {
		return err
	}

	log.Info("Found waiting tasks")
	for _, task := range tasks {
		log.With("task", task).Info("log info")
		job := workers.TaskJob{Task: taskFilter}
		c.taskWorkerPool.AddJob(job)
	}

	return nil
}
