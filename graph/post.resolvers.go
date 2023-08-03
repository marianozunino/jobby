package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.35

import (
	"context"

	"github.com/google/uuid"
	"github.com/marianozunino/jobby/dtos"
	"github.com/marianozunino/jobby/graph/generated"
	"github.com/marianozunino/jobby/graph/middleware"
)

// CreatePost is the resolver for the createPost field.
func (r *mutationResolver) CreatePost(ctx context.Context, input dtos.PostCreateInput) (*dtos.Post, error) {
	user := middleware.ForContext(ctx)

	return r.Service.CreatePost(ctx, input, user.ID)
}

// DeletePost is the resolver for the deletePost field.
func (r *mutationResolver) DeletePost(ctx context.Context, id uuid.UUID) (*dtos.Post, error) {
	return r.Service.DeletePost(ctx, id)
}

// UpdatePost is the resolver for the updatePost field.
func (r *mutationResolver) UpdatePost(ctx context.Context, id uuid.UUID, input dtos.PostUpdateInput) (*dtos.Post, error) {
	return r.Service.UpdatePost(ctx, id, input)
}

// PublishPost is the resolver for the publishPost field.
func (r *mutationResolver) PublishPost(ctx context.Context, id uuid.UUID) (*dtos.Post, error) {
	return r.Service.PublishPost(ctx, id)
}

// Categories is the resolver for the categories field.
func (r *postResolver) Categories(ctx context.Context, obj *dtos.Post) ([]*dtos.PostCategory, error) {
	return r.DataLoaders.Retrieve(ctx).PostCategoriesForPostId.Load(obj.ID)
}

// Post is the resolver for the post field.
func (r *queryResolver) Post(ctx context.Context, id uuid.UUID) (*dtos.Post, error) {
	return r.Service.GetPost(ctx, id)
}

// Posts is the resolver for the posts field.
func (r *queryResolver) Posts(ctx context.Context, orderBy *dtos.PostAggregationInput, take *int, skip *int, where *dtos.PostWhereInput) (*dtos.PaginatedPostResponse, error) {
	return r.Service.PaginatedPosts(ctx, orderBy, take, skip, where)
}

// Post returns generated.PostResolver implementation.
func (r *Resolver) Post() generated.PostResolver { return &postResolver{r} }

type postResolver struct{ *Resolver }
