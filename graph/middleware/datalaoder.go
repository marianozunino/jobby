package middleware

import (
	"context"
	"net/http"

	"github.com/marianozunino/cc-backend-go/graph/dataloader"
	"github.com/marianozunino/cc-backend-go/service"
)

// Middleware stores Loaders as a request-scoped context value.
func Dataloadewr(service service.Service) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			loaders := dataloader.NewLoaders(ctx, service)
			augmentedCtx := context.WithValue(ctx, dataloader.Key, loaders)
			r = r.WithContext(augmentedCtx)
			next.ServeHTTP(w, r)
		})
	}
}
