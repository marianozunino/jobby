package server

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"

	"github.com/marianozunino/cc-backend-go/ent"
)

func errorPresenter(ctx context.Context, e error) *gqlerror.Error {
	err := graphql.DefaultErrorPresenter(ctx, e)

	// check if error is sql.ErrNoRows
	if errors.Is(e, sql.ErrNoRows) {
		return &gqlerror.Error{
			Message: "Entity not found",
			Extensions: map[string]interface{}{
				"code": "NOT_FOUND",
			},
			Path: graphql.GetPath(ctx),
		}
	}

	if ent.IsConstraintError(e) {
		return &gqlerror.Error{
			Message: "Entity already exists",
			Extensions: map[string]interface{}{
				"code": "CONFLICT",
			},
			Path: graphql.GetPath(ctx),
		}
	}

	fmt.Printf("Error: %v\n", e)

	return err
}
