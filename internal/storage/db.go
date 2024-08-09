package storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"shortener/internal/app"
	"shortener/internal/config"
	"shortener/internal/db"
	"shortener/internal/logger"
	"strconv"
	"strings"
	"time"
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

func (d DBStorage) Add(ctx context.Context, inputURL string) (string, error) {
	genURL, err := app.CreateShortURL(inputURL)
	if err != nil {
		return "", err
	}

	baseShortURL := ""
	childCtx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	err = d.DB.QueryRowContext(childCtx, "INSERT INTO url_list (short_url, url) VALUES ($1, $2)  ON CONFLICT(url) DO UPDATE SET updated_at = NOW() RETURNING short_url", genURL, inputURL).Scan(&baseShortURL)
	if err != nil {
		return "", err
	} else if baseShortURL != genURL {
		return baseShortURL, errors.New("URL уже присутствует в базе")
	}
	return genURL, nil
}

func (d DBStorage) Find(ctx context.Context, key string) (string, error) {
	childCtx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	row := d.DB.QueryRowContext(childCtx, "SELECT url FROM url_list WHERE short_url = $1", config.Cfg.BaseURL+key)

	var originalURL string
	if err := row.Scan(&originalURL); err != nil {
		return "", errors.New("ключ не найден")
	}
	return originalURL, nil
}

func replaceSQL(stmt, pattern string, len int) string {
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

func (d DBStorage) BatchAdd(ctx context.Context, inputURLs []BatchInputParams) ([]BatchOutputParams, error) {
	var output []BatchOutputParams
	vals := []interface{}{}
	for _, row := range inputURLs {
		genURL, err := app.CreateShortURL(row.URL)
		if err != nil {
			return nil, err
		}
		vals = append(vals, genURL, row.URL)
		output = append(output, BatchOutputParams{CorrelationID: row.CorrelationID, ShortURL: genURL})
	}

	childCtx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	stmt, _ := d.DB.Prepare(replaceSQL("INSERT INTO url_list(short_url, url) VALUES %s", "(?, ?)", len(inputURLs)) +
		" ON CONFLICT(short_url) DO UPDATE SET url = EXCLUDED.url, updated_at = NOW()")
	_, err := stmt.ExecContext(childCtx, vals...)
	if err != nil {
		return nil, err
	}
	return output, nil
}
