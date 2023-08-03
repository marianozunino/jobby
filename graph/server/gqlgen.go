package server

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/marianozunino/jobby/config"
	"github.com/marianozunino/jobby/graph"
	"github.com/marianozunino/jobby/graph/dataloader"
	"github.com/marianozunino/jobby/graph/directive"
	"github.com/marianozunino/jobby/graph/generated"
	"github.com/marianozunino/jobby/graph/middleware"
	"github.com/marianozunino/jobby/service"
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

	if config.InstrospectionEnabled {
		srv.AddTransport(transport.Options{})
		srv.AddTransport(transport.POST{})
		srv.Use(extension.Introspection{})
	}

	srv.SetErrorPresenter(errorPresenter)

	srv.Use(extension.FixedComplexityLimit(300))

	dlMiddleware := middleware.Dataloadewr(service)
	return dlMiddleware(srv)
}

// NewPlaygroundHandler returns a new GraphQL Playground handler.
func NewPlaygroundHandler(endpoint string) http.HandlerFunc {
	return playground.Handler("GraphQL Playground", endpoint)
}
