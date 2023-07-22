package dataloader

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/marianozunino/cc-backend-go/dtos"
	"github.com/marianozunino/cc-backend-go/service"
)

type ctxKey string

const (
	key = ctxKey("dataloaders")
)

//go:generate go run github.com/vektah/dataloaden JobOffersLoader github.com/google/uuid.UUID []*github.com/marianozunino/cc-backend-go/dtos.JobOffer
//go:generate go run github.com/vektah/dataloaden ChildCategoriesLoader github.com/google/uuid.UUID []*github.com/marianozunino/cc-backend-go/dtos.Category
//go:generate go run github.com/vektah/dataloaden ParentCategoryLoader github.com/google/uuid.UUID *github.com/marianozunino/cc-backend-go/dtos.Category
type Loaders struct {
	JobOffersByStatusId        *JobOffersLoader
	ChildCategoriesForParentId *ChildCategoriesLoader
	ParentCategoryForChildId   *ParentCategoryLoader
}

func newLoaders(ctx context.Context, service service.Service) *Loaders {
	return &Loaders{
		// individual loaders will be initialized here
		JobOffersByStatusId:        newJobOffersByStatusID(ctx, service),
		ChildCategoriesForParentId: newCategoriesForParentIDs(ctx, service),
		ParentCategoryForChildId:   newParentCategoryForChildID(ctx, service),
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
	return &retriever{key: key}
}

func newJobOffersByStatusID(ctx context.Context, service service.Service) *JobOffersLoader {
	return NewJobOffersLoader(JobOffersLoaderConfig{
		MaxBatch: 100,
		Wait:     5 * time.Millisecond,
		Fetch: func(statusIDs []uuid.UUID) ([][]*dtos.JobOffer, []error) {
			// Fetch job offers from the service directly
			dbRecords, err := service.JobOffersWithStatus(statusIDs)
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

func newCategoriesForParentIDs(ctx context.Context, service service.Service) *ChildCategoriesLoader {
	return NewChildCategoriesLoader(ChildCategoriesLoaderConfig{
		MaxBatch: 100,
		Wait:     5 * time.Millisecond,
		Fetch: func(parentIDs []uuid.UUID) ([][]*dtos.Category, []error) {
			// Fetch categories from the service directly
			dbRecords, err := service.ChildCategoriesFor(parentIDs)
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
		Fetch: func(childIDs []uuid.UUID) ([]*dtos.Category, []error) {
			// Fetch categories from the service directly
			dbRecords, err := service.ParentCategoriesFor(childIDs)
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
			for _, id := range childIDs {
				finalResults = append(finalResults, results[id.String()])
			}

			// Return the results
			return finalResults, nil
		},
	})
}
