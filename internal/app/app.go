package app

import (
	"context"
	"fmt"
	"kanban/internal/config"
	"kanban/internal/cronjobs"
	taskspostgresrepo "kanban/internal/data/task_repo/task_postgres_repo"
	taskservice "kanban/internal/domain/usecases/services/task_service"
	httpserver "kanban/internal/http-server"
	"kanban/internal/lib/logger/sl"
	"kanban/internal/lib/postgresql"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type App struct{}

func New() App {
	return App{}
}

func (a *App) Run() {
	cfg := config.MustLoad()
	log := sl.SetupLogger(cfg.Env, cfg.Logger)

	log.With("config", cfg).Info("Application start!")

	postgres, err := postgresql.New(log, &cfg.PostgresConfig)
	if err != nil {
		log.Error("failed to create postgres instance", sl.Err(err))
		panic(err)
	}

	taskRepository := taskspostgresrepo.NewTaskPostgresRepository(postgres)

	taskService := taskservice.NewTaskService(
		log,
		taskRepository)

	cron := cronjobs.NewCronJob(taskService, log, 10) // TODO: кол-во воркеров убрать в конфиги
	cron.Start(&cfg.Cronjob)
	defer cron.Stop()

	router := httpserver.New(log, cfg, taskService)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	srv := &http.Server{
		Addr:         cfg.HTTPServer.Address,
		Handler:      router,
		ReadTimeout:  time.Duration(cfg.HTTPServer.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.HTTPServer.WriteTimeout) * time.Second,
		IdleTimeout:  time.Duration(cfg.HTTPServer.IdleTimeout) * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Error(fmt.Sprintf("failed to start server because: %v", err))
		}
	}()

	<-done
	log.Info("stopping server")

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.GracefulTimeout)*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Error("failed to stop server", sl.Err(err))
		return
	}

	// TODO: close storages
	postgres.Close()

	log.Info("Application stopped!")
}
