package graph

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/vektah/gqlparser/v2/gqlerror"

	"github.com/marianozunino/cc-backend-go/graph/dataloader"
	"github.com/marianozunino/cc-backend-go/graph/directive"
	"github.com/marianozunino/cc-backend-go/graph/generated"
	"github.com/marianozunino/cc-backend-go/graph/middleware"
	"github.com/marianozunino/cc-backend-go/service"
)

// NewHandler returns a new graphql endpoint handler.
func NewHandler(service service.Service) http.Handler {
	schema := generated.NewExecutableSchema(generated.Config{
		Resolvers: &Resolver{
			Service:     service,
			DataLoaders: dataloader.NewRetriever(),
		},
		Directives: generated.DirectiveRoot{
			Auth: directive.Auth,
		},
		Complexity: generated.ComplexityRoot{},
	})

	srv := handler.NewDefaultServer(schema)

	srv.SetErrorPresenter(func(ctx context.Context, e error) *gqlerror.Error {
		err := graphql.DefaultErrorPresenter(ctx, e)
		fmt.Print("err: ", err)

		// check if error is sql.ErrNoRows
		if errors.Is(e, sql.ErrNoRows) {
			return gqlerror.Errorf("not found")
		}

		return err
	})

	srv.Use(extension.FixedComplexityLimit(300))

	dlMiddleware := middleware.Dataloadewr(service)
	return dlMiddleware(srv)
}

// NewPlaygroundHandler returns a new GraphQL Playground handler.
func NewPlaygroundHandler(endpoint string) http.HandlerFunc {
	return playground.Handler("GraphQL Playground", endpoint)
}
