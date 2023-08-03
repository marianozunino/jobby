package directive

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/marianozunino/jobby/graph/middleware"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func Auth(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	if user := middleware.ForContext(ctx); user != nil {
		return next(ctx)
	}
	return nil, &gqlerror.Error{
		Message: "Access Denied",
	}
}
