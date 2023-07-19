package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"

	_ "github.com/marianozunino/cc-backend-go/config"
	"github.com/marianozunino/cc-backend-go/graph/generated"
	"github.com/marianozunino/cc-backend-go/resolvers"
	"github.com/marianozunino/cc-backend-go/service"
	"github.com/marianozunino/cc-backend-go/store/postgres"
	"github.com/vektah/gqlparser/v2/gqlerror"
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
	r.Route("/graphql", func(r chi.Router) {
		// Initialize the dataloaders as middleware into our route
		// r.Use(dataloaders.NewMiddleware(db)...)

		schema := generated.NewExecutableSchema(generated.Config{
			Resolvers: &resolvers.Resolver{
				Service: service,
			},
			Directives: generated.DirectiveRoot{},
			Complexity: generated.ComplexityRoot{},
		})

		srv := handler.NewDefaultServer(schema)

		srv.SetErrorPresenter(func(ctx context.Context, e error) *gqlerror.Error {
			err := graphql.DefaultErrorPresenter(ctx, e)

			// check if error is sql.ErrNoRows
			if errors.Is(e, sql.ErrNoRows) {
				return gqlerror.Errorf("not found")
			}

			return err
		})
		srv.Use(extension.FixedComplexityLimit(300))

		r.Handle("/", srv)
	})

	gqlPlayground := playground.Handler("api-gateway", "/graphql")
	r.Get("/", gqlPlayground)

	fmt.Printf("connect to http://localhost%s/ for GraphQL playground", port)
	panic(http.ListenAndServe(port, r))
}
