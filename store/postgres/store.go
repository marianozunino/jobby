package postgres

import (
	"os"

	"entgo.io/ent/dialect"
	_ "github.com/lib/pq"
	"github.com/marianozunino/cc-backend-go/ent"
	"github.com/marianozunino/cc-backend-go/store"
)

type Store struct {
	client *ent.Client
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

	if os.Getenv("DEBUG") == "true" {
		client = client.Debug()
	}

	if err != nil {
		panic(err)
	}

	return &Store{
		client,
		&StatusStore{client},
		&JobOfferStore{client},
		&MessageStore{client},
		&CategoryStore{client},
		&DegreeLevelStore{client},
	}
}

func (s Store) Close() error {
	return s.client.Close()
}
