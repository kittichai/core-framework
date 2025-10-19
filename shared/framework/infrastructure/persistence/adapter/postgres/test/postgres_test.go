package postgres

import (
	"testing"

	"github.com/kittichai/core-framework/shared/framework/infrastructure/persistence/adapter/postgres"
)

func TestPostgresConnection(t *testing.T) {
	dsn := "host=localhost user=postgres password=postgres dbname=mini_erp sslmode=disable"
	db, err := postgres.NewPostgresDB(dsn)
	if err != nil {
		t.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		t.Fatalf("Failed to ping PostgreSQL: %v", err)
	}
}
