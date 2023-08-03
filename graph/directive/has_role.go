package directive

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/marianozunino/cc-backend-go/dtos"
	"github.com/marianozunino/cc-backend-go/graph/middleware"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func HasRole(ctx context.Context, obj interface{}, next graphql.Resolver, role dtos.Role) (interface{}, error) {
	if user := middleware.ForContext(ctx); user != nil {
		if user.Role.String() == role.String() {
			return next(ctx)
		}
	}
	return nil, &gqlerror.Error{
		Message: "Access Denied",
	}

}
