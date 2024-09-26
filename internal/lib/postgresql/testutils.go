package postgresql

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"os"
	"testing"
)

var (
	DBPool *pgxpool.Pool
)

func Setup() {
	connString := "postgres://postgres:passwordformarkonpostgres@194.190.152.220:5432/testpostgres"
	pool, err := pgxpool.Connect(context.Background(), connString)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	DBPool = pool
}

func Teardown() {
	DBPool.Close()
}

func TestMain(m *testing.M) {
	Setup()
	code := m.Run()
	Teardown()
	os.Exit(code)
}
