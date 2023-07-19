//go:generate go run github.com/99designs/gqlgen generate
package resolvers

import "github.com/marianozunino/cc-backend-go/service"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	service.Service
}
