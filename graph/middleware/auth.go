package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/marianozunino/cc-backend-go/ent"
	"github.com/marianozunino/cc-backend-go/service"
)

// A private key for context that only this package can access. This is important
// to prevent collisions between different context uses
var userCtxKey = &contextKey{"user"}

const bearerPrefix = "bearer: "

type contextKey struct {
	name string
}

// Middleware decodes the share session cookie and packs the session into context
func Auth(service service.Service) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			auth := r.Header.Get("Authorization")

			if auth == "" {
				next.ServeHTTP(w, r)
				return
			}

			if len(auth) < len(bearerPrefix) || strings.ToLower(auth[:len(bearerPrefix)]) != bearerPrefix {
				http.Error(w, "Invalid Authorization header", http.StatusForbidden)
				return
			}

			token := auth[len(bearerPrefix):]

			user, err := service.DecodeToken(r.Context(), token)

			if err != nil {
				http.Error(w, "Invalid Authorization Token", http.StatusForbidden)
				return
			}

			// put it in context
			ctx := context.WithValue(r.Context(), userCtxKey, user)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *ent.User {
	raw, _ := ctx.Value(userCtxKey).(*ent.User)
	return raw
}
