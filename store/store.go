package store

import (
	"context"

	"github.com/google/uuid"
	"github.com/marianozunino/jobby/dtos"
	"github.com/marianozunino/jobby/ent"
)

type ErrNotFound struct {
	Err error
}

func (r *ErrNotFound) Error() string {
	return r.Err.Error()
}

type StatusStore interface {
	Status(ctx context.Context, id uuid.UUID) (*ent.Status, error)
	Statuses(ctx context.Context) (ent.StatusSlice, error)
	CreateStatus(ctx context.Context, status *ent.Status) (*ent.Status, error)
	UpdateStatus(ctx context.Context, status *ent.Status) (*ent.Status, error)
	DeleteStatus(ctx context.Context, id uuid.UUID) (*ent.Status, error)
	PaginatedStatuses(ctx context.Context, orderBy *dtos.StatusAggregationInput, take *int, skip *int) (ent.StatusSlice, error)
	CountStatuses(ctx context.Context) (int, error)
}

type JobOfferStore interface {
	JobOffersWithStatus(ctx context.Context, statusID []uuid.UUID) (ent.JobOffers, error)
}

type MessageStore interface {
	Message(ctx context.Context, id uuid.UUID) (*ent.ContactUsMessage, error)
	PaginatedMessages(ctx context.Context, orderBy *dtos.MessageAggregationInput, take *int, skip *int) (ent.ContactUsMessages, error)
	CountMessages(ctx context.Context) (int, error)
	SendMessage(ctx context.Context, message *ent.ContactUsMessage) (*ent.ContactUsMessage, error)
	UpdateMessage(ctx context.Context, message *ent.ContactUsMessage) (*ent.ContactUsMessage, error)
	DeleteMessage(ctx context.Context, id uuid.UUID) (*ent.ContactUsMessage, error)
}

type CategoryStore interface {
	Category(ctx context.Context, id uuid.UUID) (*ent.Category, error)
	PaginatedCategories(ctx context.Context, orderBy *dtos.CategoryAggregationInput, take *int, skip *int, where *dtos.CategoryWhereInput) (ent.Categories, error)
	CountCategories(ctx context.Context) (int, error)
	CreateCategory(ctx context.Context, category *ent.Category) (*ent.Category, error)
	UpdateCategory(ctx context.Context, category *ent.Category) (*ent.Category, error)
	DeleteCategory(ctx context.Context, id uuid.UUID) (*ent.Category, error)
	ChildCategoriesFor(ctx context.Context, parentIDs []uuid.UUID) (ent.Categories, error)
	ParentCategoriesFor(ctx context.Context, childIDs []uuid.UUID) (ent.Categories, error)
	IsSlugTaken(ctx context.Context, slug string) (bool, error)
}

type PostCategoryStore interface {
	PostCategory(ctx context.Context, id uuid.UUID) (*ent.PostCategory, error)
	PaginatedPostCategories(ctx context.Context, orderBy *dtos.PostCategoryAggregationInput, take *int, skip *int, where *dtos.PostCategoryWhereInput) (ent.PostCategories, error)
	CountPostCategories(ctx context.Context) (int, error)
	CreatePostCategory(ctx context.Context, category *ent.PostCategory) (*ent.PostCategory, error)
	UpdatePostCategory(ctx context.Context, category *ent.PostCategory) (*ent.PostCategory, error)
	DeletePostCategory(ctx context.Context, id uuid.UUID) (*ent.PostCategory, error)
}

type DegreeLevelStore interface {
	DegreeLevel(ctx context.Context, id uuid.UUID) (*ent.DegreeLevel, error)
	PaginatedDegreeLevels(ctx context.Context, orderBy *dtos.DegreeLevelAggregationInput, take *int, skip *int, where *dtos.DegreeLevelWhereInput) (ent.DegreeLevels, error)
	CountDegreeLevels(ctx context.Context) (int, error)
	CreateDegreeLevel(ctx context.Context, degreeLevel *ent.DegreeLevel) (*ent.DegreeLevel, error)
	UpdateDegreeLevel(ctx context.Context, degreeLevel *ent.DegreeLevel) (*ent.DegreeLevel, error)
	DeleteDegreeLevel(ctx context.Context, id uuid.UUID) (*ent.DegreeLevel, error)
}

type PostStore interface {
	Post(ctx context.Context, id uuid.UUID) (*ent.Post, error)
	PaginatedPosts(ctx context.Context, orderBy *dtos.PostAggregationInput, take *int, skip *int, where *dtos.PostWhereInput) (ent.Posts, error)
	CountPosts(ctx context.Context) (int, error)
	CreatePost(ctx context.Context, post *ent.Post) (*ent.Post, error)
	UpdatePost(ctx context.Context, post *ent.Post) (*ent.Post, error)
	DeletePost(ctx context.Context, id uuid.UUID) (*ent.Post, error)
	IsPostSlugTaken(ctx context.Context, slug string) (bool, error)
	PublishPost(ctx context.Context, id uuid.UUID) (*ent.Post, error)
	PostCategoriesFor(ctx context.Context, postIDs []uuid.UUID) (map[uuid.UUID]ent.PostCategories, error)
}

type UserStore interface {
	User(ctx context.Context, id uuid.UUID) (*ent.User, error)
	UserByEmail(ctx context.Context, email string) (*ent.User, error)
	PostAuthorFor(ctx context.Context, postIDs []uuid.UUID) (map[uuid.UUID]*ent.User, error)
}

type Store interface {
	StatusStore
	JobOfferStore
	MessageStore
	CategoryStore
	DegreeLevelStore
	PostCategoryStore
	PostStore
	UserStore
	Close() error
}
