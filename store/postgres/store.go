package postgres

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/marianozunino/cc-backend-go/store"
	"github.com/rs/zerolog"
	sqldblogger "github.com/simukti/sqldb-logger"
	"github.com/simukti/sqldb-logger/logadapter/zerologadapter"
)

type Store struct {
	store.StatusStore
	store.JobOfferStore
	store.MessageStore
}

// assert that Store implements store.Store
var _ store.Store = &Store{}

func NewStore(databaseURL string) store.Store {

	db, err := sqlx.Open("postgres", databaseURL)
	if err != nil {
		panic(fmt.Sprintf("error opening database: %q", err))
	}

	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()
	loggerOptions := []sqldblogger.Option{
		sqldblogger.WithSQLQueryFieldname("sql"),
		sqldblogger.WithWrapResult(false),
		sqldblogger.WithLogDriverErrorSkip(false),
		sqldblogger.WithPreparerLevel(sqldblogger.LevelDebug),
		sqldblogger.WithQueryerLevel(sqldblogger.LevelInfo),
		sqldblogger.WithExecerLevel(sqldblogger.LevelInfo),
	}
	sqlDB := sqldblogger.OpenDriver(databaseURL, db.Driver(), zerologadapter.New(logger), loggerOptions...)
	db = sqlx.NewDb(sqlDB, "postgres")

	if err != nil {
		panic(fmt.Sprintf("error opening database: %q", err))
	}

	if err := db.Ping(); err != nil {
		panic("error connecting to database: " + err.Error())
	}

	return &Store{
		StatusStore:   &StatusStore{db},
		JobOfferStore: &JobOfferStore{db},
		MessageStore:  &MessageStore{db},
	}
}
