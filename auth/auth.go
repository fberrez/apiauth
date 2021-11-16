package auth

import (
	"time"

	"github.com/golang-jwt/jwt"
)

// Auth contains authentication settings.
type Auth struct {
	secret          string
	expireInSeconds int
}

// New creates a new `Auth` instance with the given arguments.
func New(secret string, expireInSeconds int) *Auth {
	return &Auth{
		secret:          secret,
		expireInSeconds: expireInSeconds,
	}
}

// CreateToken creates a token with the given configuration. It returns the signed
// token as a string.
func (a *Auth) CreateToken() (string, error) {
	// Sets time fields
	exp := time.Now().Add(time.Second * time.Duration(a.expireInSeconds)).Unix()
	iat := time.Now().Unix()

	// Initializes token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": exp,
		"iat": iat,
	})

	// Signs and gets token as a string
	tokenString, err := token.SignedString([]byte(a.secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
