package storage

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"shortener/internal/app"
	"shortener/internal/authenticate"
	"shortener/internal/config"
	"shortener/internal/db"
	"strconv"
	"strings"
	"time"
)

// sql запросы
const (
	addQuery         = "INSERT INTO url_list (short_url, url, user_id) VALUES ($1, $2, $3) ON CONFLICT(url) DO UPDATE SET updated_at = NOW() RETURNING short_url" // добавление записи
	findByShortQuery = "SELECT url, deleted_at FROM url_list WHERE short_url = $1"                                                                                // поиск записи
	findByUserQuery  = "SELECT short_url, url FROM url_list WHERE user_id = $1 AND deleted_at IS NULL"                                                            // поиск записей пользователя
	removeQuery      = "UPDATE url_list SET deleted_at = NOW() WHERE user_id = $1 AND deleted_at IS NULL AND short_url= ANY($2)"                                  // удаление записи
)

// DBStorage тип db хранилища
type DBStorage struct {
	config config.Config
	logger *zap.Logger
	DB     *sql.DB
}

// ErrURLExists URL существует
var ErrURLExists = fmt.Errorf("url exists")

// ErrUserNotExists Пользователь не найден
var ErrUserNotExists = fmt.Errorf("user not exists")

// ErrShortIsRemoved шорткей был удален
var ErrShortIsRemoved = fmt.Errorf("short is removed")

// ErrKeyNotFound шорткей не найден
var ErrKeyNotFound = fmt.Errorf("key not found")

// CreateDBStorage создание db хранилища, config - конфиг, logger - логгер
func CreateDBStorage(config config.Config, logger *zap.Logger) Storage {
	res, err := db.Connect(config.DatabaseDSN)
	if err != nil {
		logger.Fatal(err.Error())
	}
	if err := db.Migrate(res); err != nil {
		logger.Fatal(err.Error())
	}

	return DBStorage{
		config: config,
		logger: logger,
		DB:     res,
	}
}

// Add добавление, ctx - контекст, inputURL - входящий url
func (d DBStorage) Add(ctx context.Context, inputURL string) (string, error) {
	if err := app.ValidateURL(inputURL); err != nil {
		return "", err
	}

	baseShortURL := ""
	genURL := app.GenerateShortURL(d.config.ShortStringLength, d.config.BaseURL)
	childCtx, cancel := context.WithTimeout(ctx, time.Second*d.config.QueryTimeOut)
	defer cancel()
	err := d.DB.QueryRowContext(childCtx, addQuery, genURL, inputURL, ctx.Value(authenticate.ContextUserID)).Scan(&baseShortURL)
	if err != nil {
		return "", err
	} else if baseShortURL != genURL {
		return baseShortURL, ErrURLExists
	}
	return genURL, nil
}

// Find поиск, ctx - контекст, key - шорткей
func (d DBStorage) Find(ctx context.Context, key string) (string, error) {
	childCtx, cancel := context.WithTimeout(ctx, time.Second*d.config.QueryTimeOut)
	defer cancel()
	row := d.DB.QueryRowContext(childCtx, findByShortQuery, d.config.BaseURL+key)

	var originalURL string
	var deletedAt sql.NullTime
	if err := row.Scan(&originalURL, &deletedAt); err != nil {
		return "", ErrKeyNotFound
	}
	if deletedAt.Valid {
		return originalURL, ErrShortIsRemoved
	}
	return originalURL, nil
}

// replaceSQL хелпер для multiple insert, stmt - запрос, pattern - паттерн, len - длина
func replaceSQL(stmt string, pattern string, len int) string {
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

// getUserOrNull хелпер для получения null значения пользователя, ctx - контекст
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

// BatchAdd групповое добавление, ctx - контекст, inputURLs массив данных
func (d DBStorage) BatchAdd(ctx context.Context, inputURLs []BatchInputParams) ([]BatchOutputParams, error) {
	var output []BatchOutputParams
	UserID := getUserOrNull(ctx)

	vals := []interface{}{}
	for _, row := range inputURLs {
		err := app.ValidateURL(row.URL)
		if err != nil {
			return nil, err
		}

		genURL := app.GenerateShortURL(d.config.ShortStringLength, d.config.BaseURL)
		vals = append(vals, genURL, row.URL, UserID)
		output = append(output, BatchOutputParams{CorrelationID: row.CorrelationID, ShortURL: genURL})
	}

	childCtx, cancel := context.WithTimeout(ctx, time.Second*d.config.QueryTimeOut)
	defer cancel()
	stmt, _ := d.DB.PrepareContext(childCtx, replaceSQL("INSERT INTO url_list(short_url, url, user_id) VALUES %s", "(?, ?, ?)", len(inputURLs))+
		" ON CONFLICT(short_url) DO UPDATE SET url = EXCLUDED.url, updated_at = NOW(), deleted_at = NULL")
	_, err := stmt.ExecContext(childCtx, vals...)
	if err != nil {
		return nil, err
	}
	return output, nil
}

// FindByUser поиск по пользователю, ctx - контекст
func (d DBStorage) FindByUser(ctx context.Context) ([]FindByUserOutputParams, error) {
	var output []FindByUserOutputParams
	UserID, ok := ctx.Value(authenticate.ContextUserID).(uuid.UUID)
	if !ok {
		return output, ErrUserNotExists
	}

	childCtx, cancel := context.WithTimeout(ctx, time.Second*d.config.QueryTimeOut)
	defer cancel()
	rows, err := d.DB.QueryContext(childCtx, findByUserQuery, UserID)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			d.logger.Error(err.Error())
		}
		if err := rows.Err(); err != nil {
			d.logger.Error(err.Error())
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

// getFormatShorts форматирование шорткеев, shorts - массив даных
func (d DBStorage) getFormatShorts(shorts []string) []string {
	for i, short := range shorts {
		shorts[i] = d.config.BaseURL + "/" + short
	}
	return shorts
}

// Remove удаление, ctx - контекст, UserID - guid пользователя, shorts - массив шорткеев
func (d DBStorage) Remove(ctx context.Context, UserID uuid.UUID, shorts []string) error {
	childCtx, cancel := context.WithTimeout(ctx, time.Second*d.config.QueryTimeOut)
	defer cancel()

	_, err := d.DB.ExecContext(childCtx, removeQuery, UserID, d.getFormatShorts(shorts))
	if err != nil {
		return err
	}
	return nil
}
