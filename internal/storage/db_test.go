package storage

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"reflect"
	"shortener/internal/authenticate"
	"shortener/internal/config"
	"shortener/internal/logger"
	"testing"
)

func TestDBStorage_Add(t *testing.T) {

	conf := config.NewConfig()
	loggerRes, _ := logger.NewLogger(zap.InfoLevel)
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	d := DBStorage{
		config: conf,
		logger: loggerRes,
		DB:     db,
	}

	userID := uuid.MustParse("00000000-0000-0000-0000-000000000000")
	ctx := context.WithValue(context.Background(), authenticate.ContextUserID, userID)
	mock.ExpectExec(addQuery).WithArgs("http://localhost/xxxxxx", "https://ya.ru", userID)

	// now we execute our method
	_, _ = d.Add(ctx, "https://ya.ru")
}

func TestDBStorage_Find(t *testing.T) {
	conf := config.NewConfig()
	loggerRes, _ := logger.NewLogger(zap.InfoLevel)
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	d := DBStorage{
		config: conf,
		logger: loggerRes,
		DB:     db,
	}

	userID := uuid.MustParse("00000000-0000-0000-0000-000000000000")
	ctx := context.WithValue(context.Background(), authenticate.ContextUserID, userID)
	mock.ExpectExec(findByShortQuery).WithArgs("xxxxxx")

	// now we execute our method
	_, _ = d.Find(ctx, "https://ya.ru")
}

func TestDBStorage_FindByUser(t *testing.T) {
	conf := config.NewConfig()
	loggerRes, _ := logger.NewLogger(zap.InfoLevel)
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	d := DBStorage{
		config: conf,
		logger: loggerRes,
		DB:     db,
	}

	userID := uuid.MustParse("00000000-0000-0000-0000-000000000000")
	ctx := context.WithValue(context.Background(), authenticate.ContextUserID, userID)
	mock.ExpectExec(findByUserQuery).WithArgs(userID)

	// now we execute our method
	_, _ = d.FindByUser(ctx)
}

func TestDBStorage_Remove(t *testing.T) {
	conf := config.NewConfig()
	loggerRes, _ := logger.NewLogger(zap.InfoLevel)
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	d := DBStorage{
		config: conf,
		logger: loggerRes,
		DB:     db,
	}
	shorts := []string{"xxxxxx"}

	userID := uuid.MustParse("00000000-0000-0000-0000-000000000000")
	ctx := context.WithValue(context.Background(), authenticate.ContextUserID, userID)
	mock.ExpectExec(removeQuery).WithArgs(userID, d.getFormatShorts(shorts))

	// now we execute our method
	_ = d.Remove(ctx, userID, shorts)
}

func TestDBStorage_CreateDBStorage(t *testing.T) {
	conf := config.NewConfig()
	loggerRes, _ := logger.NewLogger(zap.InfoLevel)
	db, _, _ := sqlmock.New()

	store := CreateDBStorage(conf, loggerRes, db)
	assert.Equal(t, reflect.TypeOf(store).String() == "storage.DBStorage", true)
}

func TestDBStorage_replaceSQL(t *testing.T) {
	inputURLs := []BatchInputParams{
		{
			CorrelationID: "b9253cb9-03e9-4850-a3cb-16e84e9f8a37",
			URL:           "http://lenta.ru",
		},
	}
	replaceSQL("INSERT INTO url_list(short_url, url, user_id) VALUES %s", "(?, ?, ?)", len(inputURLs))
}

func TestDBStorage_getUserOrNull(t *testing.T) {
	res := getUserOrNull(context.Background())
	assert.Equal(t, res, sql.NullString{})

	userID := uuid.MustParse("00000000-0000-0000-0000-000000000000")
	ctx := context.WithValue(context.Background(), authenticate.ContextUserID, userID)
	res = getUserOrNull(ctx)
	assert.NotEqual(t, res, sql.NullString{})
}
