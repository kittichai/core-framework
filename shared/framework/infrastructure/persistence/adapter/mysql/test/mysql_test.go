package mysql

import (
	"testing"

	"github.com/kittichai/core-framework/shared/framework/infrastructure/persistence/adapter/mysql"
)

func TestMySQLConnection(t *testing.T) {
	dsn := "user:password@tcp(localhost:3306)/testdb?charset=utf8mb4&parseTime=True"
	db, err := mysql.NewMySQLDB(dsn)
	if err != nil {
		t.Fatalf("Failed to connect to MySQL: %v", err)
		return
	}
	defer db.CloseDB()

	if err := db.Ping(); err != nil {
		t.Fatalf("Failed to ping MySQL: %v", err)
		return
	}
}
