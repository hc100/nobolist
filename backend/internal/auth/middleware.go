package auth

import (
	"context"
	"net/http"

	"github.com/hc100/nobolist/backend/ent"
	"github.com/hc100/nobolist/backend/ent/user"
	"github.com/hc100/nobolist/backend/jwt"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

func Middleware(client *ent.Client) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")

			// Allow unauthenticated users in
			if header == "" {
				next.ServeHTTP(w, r)
				return
			}

			//validate jwt token
			tokenStr := header
			email, err := jwt.ParseToken(tokenStr)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			// create user and check if user exists in db
			users, err := client.User.
				Query().
				Where(user.EmailEQ(email)).
				All(context.Background())
			u := users[0]

			// put it in context
			ctx := context.WithValue(r.Context(), userCtxKey, &u)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *ent.User {
	raw, _ := ctx.Value(userCtxKey).(*ent.User)
	return raw
}
