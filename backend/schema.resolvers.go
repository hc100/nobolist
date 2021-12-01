package backend

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/hc100/nobolist/backend/ent"
	"github.com/hc100/nobolist/backend/ent/user"
)

func (r *mutationResolver) CreateUser(ctx context.Context, email string) (*ent.User, error) {
	u, err := r.client.User.
		Create().
		SetActive(false).
		SetEmail(email).
		SetEmailAuthenticationKey(randStringRunes(20)).
		SetEmailAuthenticationKeyCreatedAt(time.Now()).
		SetEmailAuthenticationStatus(false).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (r *queryResolver) Users(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.UserOrder) (*ent.UserConnection, error) {
	return r.client.User.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithUserOrder(orderBy),
		)
}

func (r *queryResolver) ExistUserEmail(ctx context.Context, email string) (bool, error) {
	users, err := r.client.User.
		Query().
		Where(user.EmailEQ(email)).
		All(ctx)
	if err != nil {
		return true, err
	}

	if len(users) > 0 {
		return true, nil
	}

	return false, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
