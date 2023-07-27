package postgres

import (
	"fmt"
	"os"

	"entgo.io/ent/dialect"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/marianozunino/cc-backend-go/ent"
	"github.com/marianozunino/cc-backend-go/store"
	"github.com/rs/zerolog"
	sqldblogger "github.com/simukti/sqldb-logger"
	"github.com/simukti/sqldb-logger/logadapter/zerologadapter"
)

type Store struct {
	store.StatusStore
	store.JobOfferStore
	store.MessageStore
	store.CategoryStore
	store.DegreeLevelStore
}

// assert that Store implements store.Store
var _ store.Store = &Store{}

func NewStore(databaseURL string) store.Store {
	client, err := ent.Open(dialect.Postgres, os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

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
		StatusStore:      &StatusStore{client},
		JobOfferStore:    &JobOfferStore{db},
		MessageStore:     &MessageStore{db},
		CategoryStore:    &CategoryStore{client},
		DegreeLevelStore: &DegreeLevelStore{db},
	}
}
