package dataloader

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/marianozunino/jobby/dtos"
	"github.com/marianozunino/jobby/service"
)

type ctxKey string

const (
	Key = ctxKey("dataloaders")
)

//go:generate go run github.com/vektah/dataloaden JobOffersLoader github.com/google/uuid.UUID []*github.com/marianozunino/jobby/dtos.JobOffer
//go:generate go run github.com/vektah/dataloaden ChildCategoriesLoader github.com/google/uuid.UUID []*github.com/marianozunino/jobby/dtos.Category
//go:generate go run github.com/vektah/dataloaden ParentCategoryLoader github.com/google/uuid.UUID *github.com/marianozunino/jobby/dtos.Category
//go:generate go run github.com/vektah/dataloaden PostCategoriesLoader github.com/google/uuid.UUID []*github.com/marianozunino/jobby/dtos.PostCategory
//go:generate go run github.com/vektah/dataloaden PostAuthorLoader github.com/google/uuid.UUID *github.com/marianozunino/jobby/dtos.User
//go:generate go run github.com/vektah/dataloaden JobOffersForCategoryLoader github.com/google/uuid.UUID []*github.com/marianozunino/jobby/dtos.JobOffer
//go:generate go run github.com/vektah/dataloaden StatusForJobOfferLoader github.com/google/uuid.UUID *github.com/marianozunino/jobby/dtos.Status
//go:generate go run github.com/vektah/dataloaden CategoriesForJobOfferLoader github.com/google/uuid.UUID []*github.com/marianozunino/jobby/dtos.Category
type Loaders struct {
	JobOffersByStatusId        *JobOffersLoader
	ChildCategoriesForParentId *ChildCategoriesLoader
	ParentCategoryForChildId   *ParentCategoryLoader
	PostCategoriesForPostId    *PostCategoriesLoader
	PostAuthorForPostId        *PostAuthorLoader
	JobOffersForCategoryId     *JobOffersForCategoryLoader
	StatusForJobOfferId        *StatusForJobOfferLoader
	CategoriesForJobOfferId    *CategoriesForJobOfferLoader
}

func NewLoaders(ctx context.Context, service service.Service) *Loaders {
	return &Loaders{
		// individual loaders will be initialized here
		JobOffersByStatusId:        newJobOffersByStatusID(ctx, service),
		ChildCategoriesForParentId: newChildCategoriesForParentID(ctx, service),
		ParentCategoryForChildId:   newParentCategoryForChildID(ctx, service),
		PostCategoriesForPostId:    newPostCategoriesForPostID(ctx, service),
		PostAuthorForPostId:        newPostAuthorForPostID(ctx, service),
		JobOffersForCategoryId:     newJobOffersForCategoryID(ctx, service),
		StatusForJobOfferId:        newStatusForJobOfferID(ctx, service),
		CategoriesForJobOfferId:    newCategoriesForJobOfferID(ctx, service),
	}
}

// Retriever retrieves dataloaders from the request context.
type Retriever interface {
	Retrieve(context.Context) *Loaders
}

type retriever struct {
	key ctxKey
}

func (r *retriever) Retrieve(ctx context.Context) *Loaders {
	return ctx.Value(r.key).(*Loaders)
}

// NewRetriever instantiates a new implementation of Retriever.
func NewRetriever() Retriever {
	return &retriever{key: Key}
}

func newJobOffersByStatusID(ctx context.Context, service service.Service) *JobOffersLoader {
	return NewJobOffersLoader(JobOffersLoaderConfig{
		MaxBatch: 100,
		Wait:     5 * time.Millisecond,
		Fetch: func(statusIDs []uuid.UUID) ([][]*dtos.JobOffer, []error) {
			// Fetch job offers from the service directly
			dbRecords, err := service.JobOffersWithStatus(ctx, statusIDs)
			if err != nil {
				return nil, []error{err}
			}

			// Prepare a map to hold the results
			results := make(map[string][]*dtos.JobOffer)

			// Group the records based on their StatusID
			for _, record := range dbRecords {
				results[record.StatusID.String()] = append(results[record.StatusID.String()], record)
			}

			// Prepare the final results slice using the original order of statusIDs
			var finalResults [][]*dtos.JobOffer
			for _, id := range statusIDs {
				finalResults = append(finalResults, results[id.String()])
			}

			// Return the results
			return finalResults, nil
		},
	})
}

func newChildCategoriesForParentID(ctx context.Context, service service.Service) *ChildCategoriesLoader {
	return NewChildCategoriesLoader(ChildCategoriesLoaderConfig{
		MaxBatch: 100,
		Wait:     5 * time.Millisecond,
		Fetch: func(parentIDs []uuid.UUID) ([][]*dtos.Category, []error) {
			// Fetch categories from the service directly
			dbRecords, err := service.ChildCategoriesFor(ctx, parentIDs)
			if err != nil {
				return nil, []error{err}
			}

			// Prepare a map to hold the results
			results := make(map[string][]*dtos.Category)

			// Group the records based on their ParentID
			for _, record := range dbRecords {
				results[record.ParentID.String()] = append(results[record.ParentID.String()], record)
			}

			// Prepare the final results slice using the original order of parentIDs
			var finalResults [][]*dtos.Category
			for _, id := range parentIDs {
				finalResults = append(finalResults, results[id.String()])
			}

			// Return the results
			return finalResults, nil
		},
	})
}

func newParentCategoryForChildID(ctx context.Context, service service.Service) *ParentCategoryLoader {
	return NewParentCategoryLoader(ParentCategoryLoaderConfig{
		MaxBatch: 100,
		Wait:     5 * time.Millisecond,
		Fetch: func(parentIDs []uuid.UUID) ([]*dtos.Category, []error) {
			// Fetch categories from the service directly
			dbRecords, err := service.ParentCategoriesFor(ctx, parentIDs)
			if err != nil {
				return nil, []error{err}
			}

			// Prepare a map to hold the results
			results := make(map[string]*dtos.Category)

			// Group the records based on their ChildID
			for _, record := range dbRecords {
				results[record.ID.String()] = record
			}

			// Prepare the final results slice using the original order of childIDs
			var finalResults []*dtos.Category
			for _, id := range parentIDs {
				finalResults = append(finalResults, results[id.String()])
			}

			// Return the results
			return finalResults, nil
		},
	})
}

func newPostCategoriesForPostID(ctx context.Context, service service.Service) *PostCategoriesLoader {
	return NewPostCategoriesLoader(PostCategoriesLoaderConfig{
		MaxBatch: 100,
		Wait:     5 * time.Millisecond,
		Fetch: func(postIDs []uuid.UUID) ([][]*dtos.PostCategory, []error) {
			// Fetch post categories from the service directly
			mappedDtos, err := service.PostCategoriesFor(ctx, postIDs)
			if err != nil {
				return nil, []error{err}
			}

			// Prepare the final results slice using the original order of postIDs
			var finalResults [][]*dtos.PostCategory
			for _, id := range postIDs {
				finalResults = append(finalResults, mappedDtos[id])
			}

			// Return the results
			return finalResults, nil
		},
	})
}

func newPostAuthorForPostID(ctx context.Context, service service.Service) *PostAuthorLoader {
	return NewPostAuthorLoader(PostAuthorLoaderConfig{
		MaxBatch: 100,
		Wait:     5 * time.Millisecond,
		Fetch: func(postIDs []uuid.UUID) ([]*dtos.User, []error) {
			// Fetch post author from the service directly
			mappedDtos, err := service.PostAuthorFor(ctx, postIDs)

			if err != nil {
				return nil, []error{err}
			}

			// Prepare the final results slice using the original order of postIDs
			var finalResults []*dtos.User

			for _, id := range postIDs {
				finalResults = append(finalResults, mappedDtos[id])
			}

			// Return the results
			return finalResults, nil
		},
	})
}

func newJobOffersForCategoryID(ctx context.Context, service service.Service) *JobOffersForCategoryLoader {
	return NewJobOffersForCategoryLoader(JobOffersForCategoryLoaderConfig{
		MaxBatch: 100,
		Wait:     5 * time.Millisecond,
		Fetch: func(categoryIDs []uuid.UUID) ([][]*dtos.JobOffer, []error) {
			// Fetch job offers from the service directly
			mappedDtos, err := service.JobOffersForCategories(ctx, categoryIDs)

			if err != nil {
				return nil, []error{err}
			}

			// Prepare the final results slice using the original order of categoryIDs
			var finalResults [][]*dtos.JobOffer

			for _, id := range categoryIDs {
				finalResults = append(finalResults, mappedDtos[id])
			}

			// Return the results
			return finalResults, nil
		},
	})
}

func newStatusForJobOfferID(ctx context.Context, service service.Service) *StatusForJobOfferLoader {
	return NewStatusForJobOfferLoader(StatusForJobOfferLoaderConfig{
		MaxBatch: 100,
		Wait:     5 * time.Millisecond,
		Fetch: func(jobOfferIDs []uuid.UUID) ([]*dtos.Status, []error) {
			// Fetch status from the service directly
			mappedDtos, err := service.StatusForJobOffers(ctx, jobOfferIDs)

			if err != nil {
				return nil, []error{err}
			}

			// Prepare the final results slice using the original order of jobOfferIDs
			var finalResults []*dtos.Status

			for _, id := range jobOfferIDs {
				finalResults = append(finalResults, mappedDtos[id])
			}

			// Return the results
			return finalResults, nil
		},
	})
}

func newCategoriesForJobOfferID(ctx context.Context, service service.Service) *CategoriesForJobOfferLoader {
	return NewCategoriesForJobOfferLoader(CategoriesForJobOfferLoaderConfig{
		MaxBatch: 100,
		Wait:     5 * time.Millisecond,
		Fetch: func(jobOfferIDs []uuid.UUID) ([][]*dtos.Category, []error) {
			// Fetch categories from the service directly
			mappedDtos, err := service.CategoriesForJobOffers(ctx, jobOfferIDs)

			if err != nil {
				return nil, []error{err}
			}

			// Prepare the final results slice using the original order of jobOfferIDs
			var finalResults [][]*dtos.Category

			for _, id := range jobOfferIDs {
				finalResults = append(finalResults, mappedDtos[id])
			}

			// Return the results
			return finalResults, nil
		},
	})
}
