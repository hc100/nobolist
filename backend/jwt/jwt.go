package jwt

import (
	"fmt"
	"os"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// ErrorExpiredRefreshToken error string
const ErrorExpiredRefreshToken string = "Expired refresh token"

// secret key being used to sign tokens
var (
	SecretKey = []byte(os.Getenv("JWT_SECRET"))
)

// CreateAccessToken create JWT
func CreateAccessToken(email string) (string, error) {
	var err error
	atClaims := jwt.MapClaims{}
	atClaims["email"] = email
	// atClaims["exp"] = time.Now().Add(time.Minute * 10).Unix()
	atClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString(SecretKey)
	if err != nil {
		return "", err
	}
	return token, nil
}

// CreateRefreshToken create JWT
func CreateRefreshToken(email string) (string, error) {
	var err error
	atClaims := jwt.MapClaims{}
	atClaims["refresh"] = "true"
	atClaims["email"] = email
	atClaims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString(SecretKey)
	if err != nil {
		return "", err
	}
	return token, nil
}

// ParseToken parses a jwt token and returns the email in it's claims
func ParseToken(tokenStr string) (string, error) {

	t, err := stripBearerPrefixFromTokenString(tokenStr)
	if err != nil {
		return "", err
	}

	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if err != nil && token == nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("Invalid Token")
	}

	if token.Valid && claims["email"] != nil {
		email := claims["email"].(string)
		return email, nil
	}

	if claims["refresh"] != nil {
		refresh := claims["refresh"].(string)
		if refresh == "true" {
			return "", fmt.Errorf(ErrorExpiredRefreshToken)
		}
	}

	return "", fmt.Errorf("Invalid Token")
}

// Strips 'Bearer ' prefix from bearer token string
func stripBearerPrefixFromTokenString(tok string) (string, error) {
	// Should be a bearer token
	if len(tok) > 6 && strings.ToUpper(tok[0:7]) == "BEARER " {
		return tok[7:], nil
	} else {
		return "", fmt.Errorf("Bad Authorization header format")
	}

	return tok, nil
}
