package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi"

	_ "github.com/marianozunino/cc-backend-go/config"
	"github.com/marianozunino/cc-backend-go/graph"
	"github.com/marianozunino/cc-backend-go/service"
	"github.com/marianozunino/cc-backend-go/store/postgres"
)

const (
	port = ":18080"
)

func main() {
	db := postgres.NewStore(os.Getenv("DATABASE_URL"))

	service := service.NewService(db)

	r := chi.NewRouter()

	// The base path that users would use is POST /graphql which is fairly
	// idiomatic.

	r.Handle("/graphql", graph.NewHandler(service))

	r.Get("/", graph.NewPlaygroundHandler("/graphql"))

	// fmt.Printf("connect to http://localhost%s/ for GraphQL playground", port)
	// client, err := ent.Open(dialect.Postgres, os.Getenv("DATABASE_URL"))
	// if err != nil {
	// 	panic(err)
	// }
	// defer client.Close()
	// ctx := context.Background()
	// data, err := client.Debug().Category.Query().WithParentCategory().Limit(10).Order(ent.Desc(category.FieldID)).All(ctx)
	// if err != nil {
	// 	panic(err)
	// }
	// for _, d := range data {
	// 	fmt.Printf("%v\n\n", d.Edges.ParentCategory)
	// }

	panic(http.ListenAndServe(port, r))
}
