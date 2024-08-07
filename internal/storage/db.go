package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"shortener/internal/app"
	"shortener/internal/config"
	"shortener/internal/db"
	"shortener/internal/logger"
	"strconv"
	"strings"
)

type DBStorage struct {
	DB *sql.DB
}

func CreateDBStorage() {
	res, err := db.Connect()
	if err != nil {
		logger.Log.Fatal(err.Error())
	}

	db.Migrate(res)

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

func ReplaceSQL(stmt, pattern string, len int) string {
	pattern += ","
	stmt = fmt.Sprintf(stmt, strings.Repeat(pattern, len))
	n := 0
	for strings.IndexByte(stmt, '?') != -1 {
		n++
		param := "$" + strconv.Itoa(n)
		stmt = strings.Replace(stmt, "?", param, 1)
	}
	return strings.TrimSuffix(stmt, ",")
}

func (d DBStorage) BatchAdd(inputURLs []BatchInputParams) ([]BatchOutputParams, error) {
	var output []BatchOutputParams
	vals := []interface{}{}
	for _, row := range inputURLs {
		genURL, err := app.CreateShortURL(row.URL)
		if err != nil {
			return nil, err
		}
		vals = append(vals, genURL, row.URL)
		output = append(output, BatchOutputParams{CorrelationId: row.CorrelationId, ShortURL: genURL})
	}

	stmt, _ := d.DB.Prepare(ReplaceSQL("INSERT INTO url_list(short_url, url) VALUES %s", "(?, ?)", len(inputURLs)) +
		" ON CONFLICT(short_url) DO UPDATE SET url = EXCLUDED.url, updated_at = NOW()")
	_, err := stmt.Exec(vals...)
	if err != nil {
		return nil, err
	}
	return output, nil
}
