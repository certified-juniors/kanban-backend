package workers

import (
	"context"
	"fmt"
	"kanban/internal/domain/models"
	taskservice "kanban/internal/domain/usecases/services/task_service"
	"kanban/internal/lib/workerpool"
	"log/slog"
)

type TaskJob struct {
	Task models.TaskFilters
}

func (job TaskJob) Process() error {
	//TODO add controller
	return nil
}

type TaskJobProcessor struct {
	taskService taskservice.TaskService
	log         *slog.Logger
}

func NewTaskJobProcessor(taskService taskservice.TaskService, log *slog.Logger) TaskJobProcessor {
	return TaskJobProcessor{
		taskService: taskService,
		log:         log,
	}
}

func (p TaskJobProcessor) Process(job workerpool.Job) error {
	const op = "CronJobs.Workers.TaskJobProcessor.Process"
	log := p.log.With("op", op).With("job", job)

	taskJob, ok := job.(TaskJob)
	if !ok {
		return fmt.Errorf("invalid job type")
	}

	ctx := context.Background()
	_, _, err := p.taskService.GetTaskList(ctx, taskJob.Task, 1, 20)
	if err != nil {
		log.Error("failed to task", "error", err)
		return err
	}

	return nil
}
