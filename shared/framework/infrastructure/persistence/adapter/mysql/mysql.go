package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLDB struct {
	*sql.DB
}

func NewMySQLDB(dsn string) (*MySQLDB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &MySQLDB{DB: db}, nil
}

func (db *MySQLDB) CloseDB() error {
	return db.Close()
}
