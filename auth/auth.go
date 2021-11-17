package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/juju/errors"
)

// Auth contains all tools needed to perform an authentication.
type Auth struct {
	// jwt is an instance used to generate, sign and parse jwt
	jwt *JWTSettings
	// rdb is a redis client
	rdb *redis.Client
}

// New initializes a new `Auth` instance.
func New(jwtSett *JWTSettings, rdb *redis.Client) *Auth {
	return &Auth{
		jwt: jwtSett,
		rdb: rdb,
	}
}

// GetTokens is used to generate a new token with the given email.
// The generated token is signed and then cached.
func (a *Auth) GetToken(ctx context.Context, email string) (string, error) {
	token := a.jwt.createJWT(email)
	signed, err := a.jwt.sign(token)
	if err != nil {
		return "", errors.Annotatef(err, "sign token")
	}

	err = a.cacheToken(ctx, token.Email, signed)
	if err != nil {
		return "", errors.Annotatef(err, "cache token")
	}

	return "", err
}

// ParseToken parses the signed token.
func (a *Auth) ParseToken(signedToken string) (*JWT, error) {
	return a.jwt.parse(signedToken)
}

// DelToken delete the cached token corresponding to the given email address.
func (a *Auth) DelToken(ctx context.Context, email string) error {
	return a.rdb.Del(ctx, email).Err()
}

// cacheToken sets in redis a new entry where the key is token id and the value, to signed token.
func (a *Auth) cacheToken(ctx context.Context, email, signedToken string) error {
	duration, err := time.ParseDuration(fmt.Sprintf("%ds", a.jwt.ExpiresInSeconds))
	if err != nil {
		return errors.Annotatef(err, "parse duration:")
	}

	err = a.rdb.Set(ctx, email, signedToken, duration).Err()
	if err != nil {
		return errors.Annotatef(err, "set token:")
	}

	return nil
}
