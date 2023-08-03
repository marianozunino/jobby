package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"

	_ "github.com/marianozunino/cc-backend-go/config"
	"github.com/marianozunino/cc-backend-go/graph/middleware"
	"github.com/marianozunino/cc-backend-go/graph/server"
	"github.com/marianozunino/cc-backend-go/service"
	store "github.com/marianozunino/cc-backend-go/store/postgres"
)

const (
	port = ":18080"
)

//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate ./ent/schema --template ./ent/nullsetter.tmpl
//go:generate go run github.com/99designs/gqlgen generate
func main() {
	db := store.NewStore(os.Getenv("DATABASE_URL"))

	service := service.NewService(db)

	r := chi.NewRouter()

	r.Use(middleware.Auth(service))

	r.Handle("/graphql", server.NewHandler(service))
	r.Get("/", server.NewPlaygroundHandler("/graphql"))

	fmt.Printf("connect to http://localhost%s/ for GraphQL playground\n", port)
	panic(http.ListenAndServe(port, r))
}
