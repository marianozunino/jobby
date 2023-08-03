package server

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/marianozunino/cc-backend-go/graph"
	"github.com/marianozunino/cc-backend-go/graph/dataloader"
	"github.com/marianozunino/cc-backend-go/graph/directive"
	"github.com/marianozunino/cc-backend-go/graph/generated"
	"github.com/marianozunino/cc-backend-go/graph/middleware"
	"github.com/marianozunino/cc-backend-go/service"
)

// NewHandler returns a new graphql endpoint handler.
func NewHandler(service service.Service) http.Handler {
	schema := generated.NewExecutableSchema(generated.Config{
		Resolvers: &graph.Resolver{
			Service:     service,
			DataLoaders: dataloader.NewRetriever(),
		},
		Directives: generated.DirectiveRoot{
			Auth:    directive.Auth,
			HasRole: directive.HasRole,
		},
		Complexity: generated.ComplexityRoot{},
	})

	srv := handler.NewDefaultServer(schema)

	srv.SetErrorPresenter(errorPresenter)

	srv.Use(extension.FixedComplexityLimit(300))

	dlMiddleware := middleware.Dataloadewr(service)
	return dlMiddleware(srv)
}

// NewPlaygroundHandler returns a new GraphQL Playground handler.
func NewPlaygroundHandler(endpoint string) http.HandlerFunc {
	return playground.Handler("GraphQL Playground", endpoint)
}
