// Package db работа с db
package db

import (
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
)

// Connect соединение с db, DatabaseDSN - параметры подключения
func Connect(DatabaseDSN string) (*sql.DB, error) {
	db, err := sql.Open("pgx", DatabaseDSN)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
