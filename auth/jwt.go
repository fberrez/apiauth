package auth

import (
	"time"

	"github.com/go-chi/jwtauth/v5"
	"github.com/google/uuid"
	"github.com/juju/errors"
)

// JWTSettings represents the common data used to create tokens.
type JWTSettings struct {
	// ExpiresInSeconds determines the number of seconds to wait before a token can be considered as invalid.
	// It is used to calculate the `exp` claim.
	ExpiresInSeconds int
	// Issuer `iss` identifies the principal that issued the jwt.
	Issuer string
	// Subject `sub` identifies the principal that is the subject of the jwt.
	Subject string
	// Audience `aud` identifies the recipients that the jwt is intended for.
	Audience []string
	// TokenAuth is an instance of JWTAuth used to describe the algorithm and the secret key
	// used to sign a token (and verify its validity)
	TokenAuth *jwtauth.JWTAuth
}

// Algorithm used to sign the token.
const algorithm = "HS256"

// NewJWTSettings creates new settings used to generate jwt.
func NewJWTSettings(expiresInSeconds int, iss string, sub string, aud []string, signKey []byte, verifyKey []byte) *JWTSettings {
	return &JWTSettings{
		ExpiresInSeconds: expiresInSeconds,
		Issuer:           iss,
		Subject:          sub,
		Audience:         aud,
		TokenAuth:        jwtauth.New(algorithm, signKey, verifyKey),
	}
}

// createJWT creates a new jwt with the given username and settings.
func (j *JWTSettings) createJWT(username string) map[string]interface{} {
	return map[string]interface{}{
		"username": username,
		"exp":      time.Now().Add(time.Duration(j.ExpiresInSeconds) * time.Second).Unix(),
		"iat":      time.Now().Unix(),
		"nbf":      time.Now().Unix(),
		"iss":      j.Issuer,
		"sub":      j.Subject,
		"jti":      uuid.NewString(),
		"aud":      j.Audience,
	}
}

// sign signs the given token and returns the signed token as a ready-to-send string.
func (j *JWTSettings) sign(jwtToSign map[string]interface{}) (string, error) {
	_, tokenString, err := j.TokenAuth.Encode(jwtToSign)
	if err != nil {
		return "", errors.Annotatef(err, "sign token")
	}

	return tokenString, nil
}
