package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() (*sql.DB, error) {
	// konfigurasi database
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/crud_go")
	if err != nil {
		return nil, fmt.Errorf("error opening database %w", err)
	}
	// tes connection
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database %w", err)
	}
	return db, nil
}
