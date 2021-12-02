package backend

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/hc100/nobolist/backend/ent"
	"github.com/hc100/nobolist/backend/ent/user"
	"github.com/hc100/nobolist/backend/lib/util/sendmail"
	"golang.org/x/crypto/bcrypt"
)

func (r *mutationResolver) CreateUser(ctx context.Context, email string) (*ent.User, error) {
	u, err := r.client.User.
		Query().
		Where(
			user.EmailEQ(email),
			user.EmailAuthenticationStatusEQ(false),
		).
		First(ctx)
	if ent.IsNotFound(err) {
		u, err = r.client.User.
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
	} else if err != nil {
		return nil, err
	} else {
		u, err = u.Update().
			SetEmailAuthenticationKey(randStringRunes(20)).
			SetEmailAuthenticationKeyCreatedAt(time.Now()).
			Save(ctx)
		if err != nil {
			return nil, err
		}
	}

	err = sendmail.UserRegistration(u.Email, u.EmailAuthenticationKey)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (r *mutationResolver) RegisterUser(ctx context.Context, input RegisterUserInput) (*ent.User, error) {
	u, err := r.client.User.
		Query().
		Where(
			user.ActiveEQ(false),
			user.EmailAuthenticationKeyCreatedAtGT(time.Now().AddDate(0, 0, -1)),
			user.EmailAuthenticationKeyEQ(input.Key),
		).
		First(ctx)
	if err != nil {
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	u, err = u.Update().
		SetActive(true).
		SetName(input.Name).
		SetName(input.Name).
		SetPassword(string(hashedPassword)).
		SetEmailAuthenticationStatus(true).
		SetUpdatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	err = sendmail.UserRegistrationComplete(u.Email)
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
		Where(
			user.EmailEQ(email),
			user.EmailAuthenticationStatusEQ(true),
		).
		All(ctx)
	if err != nil {
		return true, err
	}

	if len(users) > 0 {
		return true, nil
	}

	return false, nil
}

func (r *queryResolver) IsValidRegistrationKey(ctx context.Context, key string) (*ent.User, error) {
	u, err := r.client.User.
		Query().
		Where(
			user.ActiveEQ(false),
			user.EmailAuthenticationKeyCreatedAtGT(time.Now().AddDate(0, 0, -1)),
			user.EmailAuthenticationKeyEQ(key),
		).
		First(ctx)
	if err != nil {
		return nil, err
	}

	return u, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
