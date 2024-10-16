package db

import (
	"database/sql"
	"github.com/pressly/goose/v3"
)

// Migrate запуск миграций, db - ресурс подключение к db
func Migrate(db *sql.DB) error {
	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	if err := goose.Up(db, "internal/migrations"); err != nil {
		return err
	}
	return nil
}
