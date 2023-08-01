package postgres

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/marianozunino/cc-backend-go/dtos"
	"github.com/marianozunino/cc-backend-go/ent"
	"github.com/marianozunino/cc-backend-go/ent/degreelevel"
	"github.com/marianozunino/cc-backend-go/store"
)

// assert that DegreeLevelStore implements store.DegreeLevelStore
var _ store.DegreeLevelStore = &DegreeLevelStore{}

type DegreeLevelStore struct {
	*ent.Client
}

func (d DegreeLevelStore) DegreeLevel(ctx context.Context, id uuid.UUID) (*ent.DegreeLevel, error) {
	return d.Client.DegreeLevel.Query().Where(degreelevel.IDEQ(id), degreelevel.DeletedAtIsNil()).Only(ctx)
}

func (d DegreeLevelStore) PaginatedDegreeLevels(ctx context.Context, orderBy *dtos.DegreeLevelAggregationInput, take *int, skip *int, where *dtos.DegreeLevelWhereInput) (ent.DegreeLevels, error) {
	query := d.Client.DegreeLevel.Query().Where(degreelevel.DeletedAtIsNil())

	order := []degreelevel.OrderOption{}

	if orderBy != nil {
		if orderBy.ID != nil {
			if *orderBy.ID == dtos.SortOrderAsc {
				order = append(order, ent.Asc(degreelevel.FieldID))
			} else {
				order = append(order, ent.Desc(degreelevel.FieldID))
			}
		}
		if orderBy.Name != nil {
			if *orderBy.Name == dtos.SortOrderAsc {
				order = append(order, ent.Asc(degreelevel.FieldName))
			} else {
				order = append(order, ent.Desc(degreelevel.FieldName))
			}
		}
	}

	if take != nil {
		query = query.Limit(*take)
	}
	if skip != nil {
		query = query.Offset(*skip)
	}

	if len(order) > 0 {
		query = query.Order(order...)
	}

	return query.All(ctx)
}

func (d DegreeLevelStore) CountDegreeLevels(ctx context.Context) (int, error) {
	return d.Client.DegreeLevel.Query().Where(degreelevel.DeletedAtIsNil()).Count(ctx)
}

func (d DegreeLevelStore) CreateDegreeLevel(ctx context.Context, degreeLevel *ent.DegreeLevel) (*ent.DegreeLevel, error) {
	return d.Client.DegreeLevel.Create().SetName(degreeLevel.Name).Save(ctx)
}

func (d DegreeLevelStore) UpdateDegreeLevel(ctx context.Context, degreeLevel *ent.DegreeLevel) (*ent.DegreeLevel, error) {
	return d.Client.DegreeLevel.UpdateOne(degreeLevel).SetName(degreeLevel.Name).SetUpdatedAt(time.Now()).Save(ctx)
}

func (d DegreeLevelStore) DeleteDegreeLevel(ctx context.Context, id uuid.UUID) (*ent.DegreeLevel, error) {
	return d.Client.DegreeLevel.UpdateOneID(id).SetDeletedAt(time.Now()).Save(ctx)
}
