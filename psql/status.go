package psql

import (
	"context"
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
)

type Status struct {
	ID        string       `db:"id"`
	Name      string       `db:"name"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}

type Paginated[T any] struct {
	TotalCount int
	Data       []*T
}

// create a map type to represent the sort by options
type SortDirection string

const (
	ASC  SortDirection = "ASC"
	DESC SortDirection = "DESC"
)

type SortOptions map[string]SortDirection

type StatusRepository interface {
	FindAll(ctx context.Context) ([]*Status, error)
	PaginateAndCount(ctx context.Context, limit int, offset int, sort SortOptions) (*Paginated[Status], error)
	FindById(ctx context.Context, id string) (*Status, error)
}

// assert that our repository implements the interface

func NewStatusRepository(db *sqlx.DB) StatusRepository {
	return &statusRepository{
		db: db,
	}
}

var _ StatusRepository = &statusRepository{}

type statusRepository struct {
	db *sqlx.DB
}


// FindAll implements StatusRepository.
func (r *statusRepository) FindAll(ctx context.Context) ([]*Status, error) {
	query := `SELECT * FROM status`
	var statuses []*Status
	if err := r.db.SelectContext(ctx, &statuses, query); err != nil {
		return nil, err
	}
	return statuses, nil

}

// PaginateAndCount implements StatusRepository.
func (r *statusRepository) PaginateAndCount(ctx context.Context, limit int, offset int, sort SortOptions) (*Paginated[Status], error) {
  query := `SELECT * FROM status`
  if len(sort) > 0 {
    query += ` ORDER BY `
    for k, v := range sort {
      query += k + " " + string(v) + ","
    }
    query = query[:len(query)-1]
  }
  query += ` LIMIT $1 OFFSET $2`
  var statuses []*Status
  if err := r.db.SelectContext(ctx, &statuses, query, limit, offset); err != nil {
    return nil, err
  }

  count, err := r.count(ctx)
  if err != nil {
    return nil, err
  }

  return &Paginated[Status]{
    TotalCount: count,
    Data:       statuses,
  }, nil
}

func (r *statusRepository) count(ctx context.Context) (int, error) {
  query := `SELECT COUNT(*) FROM status`

  var count int
  if err := r.db.GetContext(ctx, &count, query); err != nil {
    return 0, err
  }
  return count, nil
}

// FindById implements StatusRepository.
func (*statusRepository) FindById(ctx context.Context, id string) (*Status, error) {
	panic("unimplemented")
}
