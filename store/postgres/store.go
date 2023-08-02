package postgres

import (
	"os"
	"reflect"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	_ "github.com/lib/pq"
	"github.com/marianozunino/cc-backend-go/dtos"
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
	store.PostCategoryStore
	store.PostStore
	store.UserStore
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
		&PostCategoryStore{client},
		&PostStore{client},
		&UserStore{client},
	}
}

func (s Store) Close() error {
	return s.client.Close()
}

func createRawOrderOptions[K func(*sql.Selector)](orderBy interface{}) []interface{} {
	var order []interface{}

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
