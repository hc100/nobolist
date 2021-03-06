package backend

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/hc100/nobolist/backend/ent"
	"github.com/hc100/nobolist/backend/ent/user"
	"github.com/hc100/nobolist/backend/internal/auth"
	"github.com/hc100/nobolist/backend/jwt"
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
			SetEmailAuthenticationKey(randStringRunes(40)).
			SetEmailAuthenticationKeyCreatedAt(time.Now()).
			SetEmailAuthenticationStatus(false).
			SetResetPasswordKey(randStringRunes(40)).
			SetResetPasswordKeyCreatedAt(time.Now()).
			Save(ctx)
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	} else {
		u, err = u.Update().
			SetEmailAuthenticationKey(randStringRunes(40)).
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

func (r *mutationResolver) UpdateUser(ctx context.Context, input UpdateUserInput) (*ent.User, error) {
	u := auth.ForContext(ctx)
	if u == nil {
		return nil, fmt.Errorf("Access denied")
	}

	layout := "2006-01-02"
	birthday, err := time.Parse(layout, input.Birthday)

	u, err = u.Update().
		SetName(input.Name).
		SetHeight(input.Height).
		SetHeightDisplay(user.HeightDisplay(input.HeightDisplay)).
		SetWeight(input.Weight).
		SetWeightDisplay(user.WeightDisplay(input.WeightDisplay)).
		SetWingspan(input.Wingspan).
		SetWingspanDisplay(user.WingspanDisplay(input.WingspanDisplay)).
		SetBirthday(birthday).
		SetBirthdayDisplay(user.BirthdayDisplay(input.BirthdayDisplay)).
		SetGender(user.Gender(input.Gender)).
		SetGenderDisplay(user.GenderDisplay(input.GenderDisplay)).
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

func (r *mutationResolver) Login(ctx context.Context, email string, password string) (*Token, error) {
	u, err := r.client.User.
		Query().
		Where(
			user.EmailEQ(email),
			user.ActiveEQ(true),
		).
		First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, fmt.Errorf("???????????????????????????????????????????????????")
		} else {
			return nil, err
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("???????????????????????????????????????")
	}

	at, err := jwt.CreateAccessToken(u.Email)
	if err != nil {
		return nil, err
	}

	rt, err := jwt.CreateRefreshToken(u.Email)
	if err != nil {
		return nil, err
	}

	token := Token{
		AccessToken:  at,
		RefreshToken: rt,
		Role:         u.Role,
	}

	return &token, nil
}

func (r *mutationResolver) ResetPasswordRequest(ctx context.Context, email string) (bool, error) {
	u, err := r.client.User.
		Query().
		Where(
			user.ActiveEQ(true),
			user.EmailEQ(email),
		).
		First(ctx)
	if err != nil {
		return false, err
	}

	u, err = u.Update().
		SetResetPasswordKey(randStringRunes(40)).
		SetResetPasswordKeyCreatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return false, err
	}

	err = sendmail.ResetPasswordRequest(u.Email, u.ResetPasswordKey)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *mutationResolver) ResetPassword(ctx context.Context, key string, password string) (*ent.User, error) {
	u, err := r.client.User.
		Query().
		Where(
			user.ActiveEQ(true),
			user.ResetPasswordKeyEQ(key),
			user.ResetPasswordKeyCreatedAtGT(time.Now().AddDate(0, 0, -1)),
		).
		First(ctx)
	if err != nil {
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	u, err = u.Update().
		SetPassword(string(hashedPassword)).
		SetResetPasswordKey(randStringRunes(40)).
		SetResetPasswordKeyCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	err = sendmail.ResetPassword(u.Email)
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

func (r *queryResolver) IsValidResetPasswordKey(ctx context.Context, key string) (*ent.User, error) {
	u, err := r.client.User.
		Query().
		Where(
			user.ActiveEQ(true),
			user.ResetPasswordKeyCreatedAtGT(time.Now().AddDate(0, 0, -1)),
			user.ResetPasswordKeyEQ(key),
		).
		First(ctx)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (r *queryResolver) Myself(ctx context.Context) (*Myself, error) {
	u := auth.ForContext(ctx)
	if u == nil {
		return nil, fmt.Errorf("Access denied")
	}

	m := &Myself{
		ID:              u.ID,
		Name:            u.Name,
		Email:           u.Email,
		Height:          u.Height,
		HeightDisplay:   u.HeightDisplay.String(),
		Weight:          u.Weight,
		WeightDisplay:   u.WeightDisplay.String(),
		Wingspan:        u.Wingspan,
		WingspanDisplay: u.WingspanDisplay.String(),
		Birthday:        u.Birthday.Format("2006-01-02"),
		BirthdayDisplay: u.BirthdayDisplay.String(),
		Gender:          u.Gender.String(),
		GenderDisplay:   u.GenderDisplay.String(),
	}

	return m, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
