package storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"shortener/internal/app"
	"shortener/internal/authenticate"
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

var ErrURLExists = fmt.Errorf("url exists")
var ErrUserNotExists = fmt.Errorf("user not exists")

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
	err := app.ValidateURL(inputURL)
	if err != nil {
		return "", err
	}

	baseShortURL := ""
	genURL := app.GenerateShortURL()
	childCtx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	err = d.DB.QueryRowContext(childCtx, "INSERT INTO url_list (short_url, url, user_id) VALUES ($1, $2, $3)  "+
		"ON CONFLICT(url) DO UPDATE SET updated_at = NOW() RETURNING short_url", genURL, inputURL, ctx.Value(authenticate.ContextUserID)).Scan(&baseShortURL)
	if err != nil {
		return "", err
	} else if baseShortURL != genURL {
		return baseShortURL, ErrURLExists
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

func getUserOrNull(ctx context.Context) sql.NullString {
	id, ok := ctx.Value(authenticate.ContextUserID).(uuid.UUID)
	if ok {
		return sql.NullString{
			String: id.String(),
			Valid:  true,
		}
	}
	return sql.NullString{}
}

func (d DBStorage) BatchAdd(ctx context.Context, inputURLs []BatchInputParams) ([]BatchOutputParams, error) {
	var output []BatchOutputParams
	UserID := getUserOrNull(ctx)

	vals := []interface{}{}
	for _, row := range inputURLs {
		err := app.ValidateURL(row.URL)
		if err != nil {
			return nil, err
		}

		genURL := app.GenerateShortURL()
		vals = append(vals, genURL, row.URL, UserID)
		output = append(output, BatchOutputParams{CorrelationID: row.CorrelationID, ShortURL: genURL})
	}

	childCtx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	stmt, _ := d.DB.Prepare(replaceSQL("INSERT INTO url_list(short_url, url, user_id) VALUES %s", "(?, ?, ?)", len(inputURLs)) +
		" ON CONFLICT(short_url) DO UPDATE SET url = EXCLUDED.url, updated_at = NOW()")
	_, err := stmt.ExecContext(childCtx, vals...)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (d DBStorage) FindByUser(ctx context.Context) ([]FindByUserOutputParams, error) {
	var output []FindByUserOutputParams
	UserID, ok := ctx.Value(authenticate.ContextUserID).(uuid.UUID)
	if !ok {
		return output, ErrUserNotExists
	}

	childCtx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	rows, err := d.DB.QueryContext(childCtx, "SELECT short_url, url FROM url_list WHERE user_id = $1", UserID)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			logger.Log.Error(err.Error())
		}
		if err := rows.Err(); err != nil {
			logger.Log.Error(err.Error())
		}
	}()

	for rows.Next() {
		var s FindByUserOutputParams
		if err = rows.Scan(&s.ShortURL, &s.OriginalURL); err != nil {
			return nil, err
		}
		output = append(output, s)
	}
	return output, nil
}
