package sl

import (
	"kanban/internal/config"
	"kanban/internal/lib/logger/handlers/slogpretty"
	"log/slog"
	"os"
	"path/filepath"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}

func SetupLogger(env string, logger config.Logger) *slog.Logger {
	var log *slog.Logger

	logFile, err := os.OpenFile(logger.Path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	log = slog.New(
		slog.NewJSONHandler(logFile, &slog.HandlerOptions{Level: slog.LevelDebug}),
	)

	switch env {
	case envLocal:
		log = setupPrettySlog()
	case envDev:
		logDir := filepath.Dir(logger.Path)
		if err := os.MkdirAll(logDir, 0755); err != nil {
			panic(err)
		}
	case envProd:
		logDir := filepath.Dir(logger.Path)
		if err := os.MkdirAll(logDir, 0755); err != nil {
			panic(err)
		}
	default:
		log = setupPrettySlog()
	}
	return log
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
