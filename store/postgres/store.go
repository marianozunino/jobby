package postgres

import (
	"reflect"

	"entgo.io/ent/dialect"
	_ "github.com/lib/pq"
	"github.com/marianozunino/jobby/config"
	"github.com/marianozunino/jobby/dtos"
	"github.com/marianozunino/jobby/ent"
	"github.com/marianozunino/jobby/store"
)

type Store struct {
	client *ent.Client
	store.StatusStore
	store.JobOfferStore
	store.MessageStore
	store.CategoryStore
	store.DegreeLevelStore
	store.PostCategoryStore
	store.PostStore
	store.UserStore
}

// assert that Store implements store.Store
var _ store.Store = &Store{}

func NewStore() store.Store {
	client, err := ent.Open(dialect.Postgres, config.DatabaseURL)

	if err != nil {
		panic(err)
	}

	if config.EntDebug {
		client = client.Debug()
	}

	return &Store{
		client,
		&StatusStore{client},
		&JobOfferStore{client},
		&MessageStore{client},
		&CategoryStore{client},
		&DegreeLevelStore{client},
		&PostCategoryStore{client},
		&PostStore{client},
		&UserStore{client},
	}
}

func (s Store) Close() error {
	return s.client.Close()
}

func createRawOrderOptions[K []ent.OrderFunc](orderBy interface{}) []ent.OrderFunc {
	var order []ent.OrderFunc

	// Use reflection to get the value of the `orderBy` struct
	val := reflect.ValueOf(orderBy)
	if val.Kind() != reflect.Struct {
		// Handle invalid input
		return order
	}

	// Use reflection to get the type of the `orderBy` struct
	typ := val.Type()

	// Loop through each field in the struct
	for i := 0; i < typ.NumField(); i++ {
		fieldValue := val.Field(i)
		fieldName := typ.Field(i).Name

		// Check if the field value is a pointer to a SortOrder (dtos.SortOrderAsc/dtos.SortOrderDesc)
		if fieldValue.Kind() == reflect.Ptr {
			// Get the dereferenced value of the pointer
			orderByValue := fieldValue.Elem()
			if orderByValue.Kind() == reflect.Int && orderByValue.Type() == reflect.TypeOf(dtos.SortOrderAsc) {
				// Append Asc order if the SortOrder is dtos.SortOrderAsc
				// order = append(order, ent.Asc(fieldName))
				order = append(order, ent.Asc(fieldName))
			} else if orderByValue.Kind() == reflect.Int && orderByValue.Type() == reflect.TypeOf(dtos.SortOrderDesc) {
				// Append Desc order if the SortOrder is dtos.SortOrderDesc
				order = append(order, ent.Desc(fieldName))
			}
		}
	}

	return order
}
