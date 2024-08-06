package storage

import (
	"database/sql"
	"errors"
	"shortener/internal/app"
	"shortener/internal/config"
	"shortener/internal/db"
	"shortener/internal/logger"
)

type DBStorage struct {
	DB *sql.DB
}

func CreateDBStorage() {
	res, err := db.Connect()
	if err != nil {
		logger.Log.Fatal(err.Error())
	}
	Source = DBStorage{
		DB: res,
	}
}

func (d DBStorage) Add(inputURL string) (string, error) {
	genURL, err := app.CreateShortURL(inputURL)
	if err != nil {
		return "", err
	}

	insertDynStmt := "INSERT INTO url_list (short_url, url) VALUES ($1, $2)"
	_, err = d.DB.Exec(insertDynStmt, genURL, inputURL)
	if err != nil {
		return "", err
	}
	return genURL, nil
}

func (d DBStorage) Find(key string) (string, error) {
	row := d.DB.QueryRow("SELECT url FROM url_list WHERE short_url = $1", config.Cfg.BaseURL+key)

	var originalURL string
	if err := row.Scan(&originalURL); err != nil {
		return "", errors.New("ключ не найден")
	}
	return originalURL, nil
}
