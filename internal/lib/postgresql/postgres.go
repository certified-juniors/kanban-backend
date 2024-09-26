package postgresql

import (
	"context"
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"io/ioutil"
	"kanban/internal/config"
	"kanban/internal/lib/logger/sl"
	"log/slog"
	"strings"
	"time"
)

type Postgres struct {
	DB  *pgxpool.Pool
	log *slog.Logger
}

// New создает новый экземпляр Postgres
func New(log *slog.Logger, cfg *config.PostgresConfig) (*Postgres, error) {
	const op = "Postgres.New"

	connString := fmt.Sprintf(cfg.URL)
	poolConfig, err := pgxpool.ParseConfig(connString)
	if err != nil {
		log.Error("Error parsing pool config", slog.String("operation", op), sl.Err(err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	dbPool, err := pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		log.Error("Error connecting to database", slog.String("operation", op), sl.Err(err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	p := &Postgres{
		DB:  dbPool,
		log: log,
	}

	if cfg.AutoMigrate {
		if err = p.AutoMigrate(cfg.Migrations); err != nil {
			p.Close()
			log.Error("Error during auto migration", slog.String("operation", op), sl.Err(err))
			return nil, fmt.Errorf("%s: %w", op, err)
		}
	}

	log.Info("Postgres intance created successfully")
	return p, nil
}

// Close закрывает подключение к базе данных
func (p *Postgres) Close() {
	p.DB.Close()
	p.log.Info("Postgres connection closed")
}

// AutoMigrate выполняет автоматическую миграцию базы данных
func (p *Postgres) AutoMigrate(migrationsPath string) error {
	const op = "Postgres.AutoMigrate"

	migrationScript, err := ioutil.ReadFile(migrationsPath)
	if err != nil {
		p.log.Error("Failed to read migration file", slog.String("operation", op), sl.Err(err))
		return fmt.Errorf("%s: failed to read migration file: %w", op, err)
	}

	sqlCommands := strings.Split(string(migrationScript), ";")

	for _, cmd := range sqlCommands {
		cmd = strings.TrimSpace(cmd)
		if cmd == "" {
			continue
		}
		if _, err = p.Exec(context.Background(), cmd); err != nil {
			p.log.Error("Failed to execute migration command", slog.String("operation", op), slog.String("command", cmd), sl.Err(err))
			return fmt.Errorf("%s: failed to execute migration command: %w", op, err)
		}
	}

	p.log.Info("Auto migration completed successfully")
	return nil
}

// Exec обертка для выполнения SQL команд с логированием
func (p *Postgres) Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error) {
	start := time.Now()
	p.log.Info("Executing query", slog.String("query", sql), slog.Any("args", args))

	result, err := p.DB.Exec(ctx, sql, args...)
	duration := time.Since(start)

	if err != nil {
		p.log.Error("Error executing query",
			slog.String("query", formatSQLQuery(sql)),
			slog.Any("args", args),
			slog.Duration("duration", duration),
			sl.Err(err))
	} else {
		p.log.Info("Query executed successfully",
			slog.String("query", formatSQLQuery(sql)),
			slog.Any("args", args),
			slog.Duration("duration", duration))
	}
	return result, err
}

// Query обертка для выполнения SQL запросов с логированием
func (p *Postgres) Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
	start := time.Now()
	p.log.Info("Executing query", slog.String("query", formatSQLQuery(sql)), slog.Any("args", args))

	rows, err := p.DB.Query(ctx, sql, args...)
	duration := time.Since(start)

	if err != nil {
		p.log.Error("Error executing query",
			slog.String("query", formatSQLQuery(sql)),
			slog.Any("args", args),
			slog.Duration("duration", duration),
			sl.Err(err))
	} else {
		p.log.Info("Query executed successfully",
			slog.String("query", formatSQLQuery(sql)),
			slog.Any("args", args),
			slog.Duration("duration", duration))
	}
	return rows, err
}

// QueryRow обертка для выполнения SQL запросов с логированием
func (p *Postgres) QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row {
	start := time.Now()
	p.log.Info("Executing query", slog.String("query", formatSQLQuery(sql)), slog.Any("args", args))

	row := p.DB.QueryRow(ctx, sql, args...)
	duration := time.Since(start)

	p.log.Info("Query executed", slog.String("query", formatSQLQuery(sql)), slog.Any("args", args), slog.Duration("duration", duration))
	return row
}

// Begin начинает новую транзакцию
func (p *Postgres) Begin(ctx context.Context) (pgx.Tx, error) {
	const op = "Postgres.Begin"
	p.log.Info("Starting new transaction", slog.String("operation", op))
	tx, err := p.DB.Begin(ctx)
	if err != nil {
		p.log.Error("Failed to begin transaction", slog.String("operation", op), sl.Err(err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return tx, nil
}

// Commit фиксирует транзакцию
func (p *Postgres) Commit(ctx context.Context, tx pgx.Tx) error {
	const op = "Postgres.Commit"
	start := time.Now()
	p.log.Info("Committing transaction", slog.String("operation", op))
	err := tx.Commit(ctx)
	duration := time.Since(start)

	if err != nil {
		p.log.Error("Error committing transaction", slog.String("operation", op), slog.Duration("duration", duration), sl.Err(err))
		return fmt.Errorf("%s: %w", op, err)
	}

	p.log.Info("Transaction committed successfully", slog.String("operation", op), slog.Duration("duration", duration))
	return nil
}

// Rollback откатывает транзакцию
func (p *Postgres) Rollback(ctx context.Context, tx pgx.Tx) error {
	const op = "Postgres.Rollback"
	start := time.Now()
	p.log.Info("Rolling back transaction", slog.String("operation", op))
	err := tx.Rollback(ctx)
	duration := time.Since(start)

	if err != nil && err != pgx.ErrTxClosed {
		p.log.Error("Error rolling back transaction", slog.String("operation", op), slog.Duration("duration", duration), sl.Err(err))
		return fmt.Errorf("%s: %w", op, err)
	}

	p.log.Info("Transaction rolled back successfully", slog.String("operation", op), slog.Duration("duration", duration))
	return nil
}

// ExecInTx выполняет SQL команду в рамках транзакции
func (p *Postgres) ExecInTx(ctx context.Context, tx pgx.Tx, sql string, args ...interface{}) (pgconn.CommandTag, error) {
	start := time.Now()
	p.log.Info("Executing query in transaction", slog.String("query", sql), slog.Any("args", args))

	result, err := tx.Exec(ctx, sql, args...)
	duration := time.Since(start)

	if err != nil {
		p.log.Error("Error executing query in transaction",
			slog.String("query", formatSQLQuery(sql)),
			slog.Any("args", args),
			slog.Duration("duration", duration),
			sl.Err(err))
	} else {
		p.log.Info("Query executed successfully in transaction",
			slog.String("query", formatSQLQuery(sql)),
			slog.Any("args", args),
			slog.Duration("duration", duration))
	}
	return result, err
}

// QueryInTx выполняет SQL запрос в рамках транзакции
func (p *Postgres) QueryInTx(ctx context.Context, tx pgx.Tx, sql string, args ...interface{}) (pgx.Rows, error) {
	start := time.Now()
	p.log.Info("Executing query in transaction", slog.String("query", formatSQLQuery(sql)), slog.Any("args", args))

	rows, err := tx.Query(ctx, sql, args...)
	duration := time.Since(start)

	if err != nil {
		p.log.Error("Error executing query in transaction",
			slog.String("query", formatSQLQuery(sql)),
			slog.Any("args", args),
			slog.Duration("duration", duration),
			sl.Err(err))
	} else {
		p.log.Info("Query executed successfully in transaction",
			slog.String("query", formatSQLQuery(sql)),
			slog.Any("args", args),
			slog.Duration("duration", duration))
	}
	return rows, err
}

// QueryRowInTx выполняет SQL запрос с возвратом одной строки в рамках транзакции
func (p *Postgres) QueryRowInTx(ctx context.Context, tx pgx.Tx, sql string, args ...interface{}) pgx.Row {
	start := time.Now()
	p.log.Info("Executing query in transaction", slog.String("query", formatSQLQuery(sql)), slog.Any("args", args))

	row := tx.QueryRow(ctx, sql, args...)
	duration := time.Since(start)

	p.log.Info("Query executed in transaction", slog.String("query", formatSQLQuery(sql)), slog.Any("args", args), slog.Duration("duration", duration))
	return row
}

// formatSQLQuery форматирует SQL запрос для лучшей читаемости в логах
func formatSQLQuery(query string) string {
	query = strings.ReplaceAll(query, "\n", " ")
	return strings.ReplaceAll(query, "\r", " ")
}
