package storage

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

	err = db.Ping()
	if err != nil {
		logger.Log.Fatal(err.Error())
		return nil, err
	}

	defer func() {
		err = db.Close()
	}()

	return db, nil
}
