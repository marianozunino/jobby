package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"

	_ "github.com/marianozunino/cc-backend-go/config"
	"github.com/marianozunino/cc-backend-go/graph"
	"github.com/marianozunino/cc-backend-go/service"
	store "github.com/marianozunino/cc-backend-go/store/postgres"
)

const (
	port = ":18080"
)

func main() {
	db := store.NewStore(os.Getenv("DATABASE_URL"))

	service := service.NewService(db)

	r := chi.NewRouter()

	r.Handle("/graphql", graph.NewHandler(service))
	r.Get("/", graph.NewPlaygroundHandler("/graphql"))

	fmt.Printf("connect to http://localhost%s/ for GraphQL playground", port)
	panic(http.ListenAndServe(port, r))
}
