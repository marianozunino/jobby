package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/marianozunino/jobby/config"
	_ "github.com/marianozunino/jobby/config"
	"github.com/marianozunino/jobby/graph/middleware"
	"github.com/marianozunino/jobby/graph/server"
	"github.com/marianozunino/jobby/service"
	store "github.com/marianozunino/jobby/store/postgres"
)

//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate ./ent/schema --template ./ent/nullsetter.tmpl
//go:generate go run github.com/99designs/gqlgen generate
func main() {
	db := store.NewStore()

	service := service.NewService(db)

	r := chi.NewRouter()

	r.Use(middleware.Auth(service))

	r.Handle("/graphql", server.NewHandler(service))
	r.Get("/", server.NewPlaygroundHandler("/graphql"))

	fmt.Printf("connect to http://localhost%s/ for GraphQL playground\n", config.Port)

	panic(http.ListenAndServe(config.Port, r))
}
