package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func Connect(connStr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func RunMigrations(db *sql.DB, migrateDir string) error {
	fmt.Println("Migrations runner: manuel veya golang-migrate kullanın")
	return nil
}
