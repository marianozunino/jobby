package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.35

import (
	"context"

	"github.com/google/uuid"
	"github.com/marianozunino/cc-backend-go/dtos"
	"github.com/marianozunino/cc-backend-go/graph/generated"
)

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input dtos.AuthInput) (*dtos.Auth, error) {
	panic("not implemented")
}

// CreateStatus is the resolver for the createStatus field.
func (r *mutationResolver) CreateStatus(ctx context.Context, input dtos.StatusCreateInput) (*dtos.Status, error) {
	return r.Service.CreateStatus(ctx, input)
}

// DeleteStatus is the resolver for the deleteStatus field.
func (r *mutationResolver) DeleteStatus(ctx context.Context, id string) (*dtos.Status, error) {
	return r.Service.DeleteStatus(ctx, id)
}

// UpdateStatus is the resolver for the updateStatus field.
func (r *mutationResolver) UpdateStatus(ctx context.Context, id string, input dtos.StatusUpdateInput) (*dtos.Status, error) {
	return r.Service.UpdateStatus(ctx, id, input)
}

// SendMessages implements generated.MutationResolver.
func (r *mutationResolver) SendMessage(ctx context.Context, input dtos.MessageCreateInput) (*dtos.Message, error) {
	return r.Service.SendMessage(input)
}

// DeleteMessage implements generated.MutationResolver.
func (r *mutationResolver) DeleteMessage(ctx context.Context, id string) (*dtos.Message, error) {
	return r.Service.DeleteMessage(id)
}

// UpdateMessage implements generated.MutationResolver.
func (r *mutationResolver) UpdateMessage(ctx context.Context, id string, input dtos.MessageUpdateInput) (*dtos.Message, error) {
	return r.Service.UpdateMessage(id, input)
}

// CreateCategory is the resolver for the createCategory field.
func (r *mutationResolver) CreateCategory(ctx context.Context, input dtos.CategoryCreateInput) (*dtos.Category, error) {
	return r.Service.CreateCategory(input)
}

// DeleteCategory is the resolver for the deleteCategory field.
func (r *mutationResolver) DeleteCategory(ctx context.Context, id uuid.UUID) (*dtos.Category, error) {
	return r.Service.DeleteCategory(id)
}

// UpdateCategory is the resolver for the updateCategory field.
func (r *mutationResolver) UpdateCategory(ctx context.Context, id uuid.UUID, input dtos.CategoryUpdateInput) (*dtos.Category, error) {
	return r.Service.UpdateCategory(id, input)
}

// CreateDegreeLevel is the resolver for the createDegreeLevel field.
func (r *mutationResolver) CreateDegreeLevel(ctx context.Context, input dtos.DegreeLevelCreateInput) (*dtos.DegreeLevel, error) {
	return r.Service.CreateDegreeLevel(input)
}

// DeleteDegreeLevel is the resolver for the deleteDegreeLevel field.
func (r *mutationResolver) DeleteDegreeLevel(ctx context.Context, id uuid.UUID) (*dtos.DegreeLevel, error) {
	return r.Service.DeleteDegreeLevel(id)
}

// UpdateDegreeLevel is the resolver for the updateDegreeLevel field.
func (r *mutationResolver) UpdateDegreeLevel(ctx context.Context, id uuid.UUID, input dtos.DegreeLevelUpdateInput) (*dtos.DegreeLevel, error) {
	return r.Service.UpdateDegreeLevel(id, input)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
