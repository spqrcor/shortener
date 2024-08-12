package db

import (
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"shortener/internal/config"
	"shortener/internal/logger"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("pgx", config.Cfg.DatabaseDSN)
	if err != nil {
		logger.Log.Fatal(err.Error())
		return nil, err
	}
	if err := db.Ping(); err != nil {
		logger.Log.Error(err.Error())
		return nil, err
	}
	return db, nil
}
