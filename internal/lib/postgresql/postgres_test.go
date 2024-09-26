package postgresql

import (
	"context"
	"io/ioutil"
	"log/slog"
	"salepoint/internal/lib/logger/handlers/slogpretty"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"salepoint/internal/config"
)

// TestPostgresInitialInsertions тестирует выполнение миграций и инициализацию таблиц
func TestPostgresInitialInsertions(t *testing.T) {
	const op = "TestPostgresAutoMigrate"
	log := setupPrettySlog()

	cfg := &config.PostgresConfig{
		URL:         "postgres://postgres:passwordformarkonpostgres@194.190.152.220:5432/testpostgres",
		AutoMigrate: true,
		Migrations:  "../../../migrations/postgresql/002_initial_insertions.sql",
	}

	p, err := New(log, cfg)
	if err != nil {
		t.Fatalf("Failed to create Postgres instance: %v", err)
	}
	defer p.Close()

	err = p.AutoMigrate(log, cfg.Migrations)
	assert.NoError(t, err, "AutoMigrate should not return an error")

	migrationScript, err := ioutil.ReadFile(cfg.Migrations)
	if err != nil {
		t.Fatalf("%s: failed to read migration file: %v", op, err)
	}

	sqlCommands := strings.Split(string(migrationScript), ";")

	for _, cmd := range sqlCommands {
		cmd = strings.TrimSpace(cmd)
		if cmd == "" {
			continue
		}
		if _, err := p.DB.Exec(context.Background(), cmd); err != nil {
			t.Errorf("%s: failed to execute migration command: %v", op, err)
			return
		}
	}
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
