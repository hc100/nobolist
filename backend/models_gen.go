// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package backend

type RegisterUserInput struct {
	Key      string `json:"key"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Token struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	Role         int    `json:"role"`
}
