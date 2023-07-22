package dataloader

import (
	"context"
	"net/http"

	"github.com/marianozunino/cc-backend-go/service"
)

// Middleware stores Loaders as a request-scoped context value.
func Middleware(service service.Service) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			loaders := newLoaders(ctx, service)
			augmentedCtx := context.WithValue(ctx, key, loaders)
			r = r.WithContext(augmentedCtx)
			next.ServeHTTP(w, r)
		})
	}
}
