package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/marianozunino/jobby/ent"
	"github.com/marianozunino/jobby/ent/user"
	"github.com/marianozunino/jobby/store"
)

// assert that StatusStore implements store.StatusStore
var _ store.UserStore = &UserStore{}

type UserStore struct {
	*ent.Client
}

// User implements store.UserStore.
func (u *UserStore) User(ctx context.Context, id uuid.UUID) (*ent.User, error) {
	return u.Client.User.Query().Where(user.IDEQ(id)).Only(ctx)
}

// UserByEmail implements store.UserStore.
func (u *UserStore) UserByEmail(ctx context.Context, email string) (*ent.User, error) {
	return u.Client.User.Query().Where(user.EmailContainsFold(email)).Only(ctx)
}
