package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/juju/errors"
)

// JWT represents the structure of a generated token.
type JWT struct {
	// Email is the email of the user that owns the token.
	Email string `json: "email"`
	jwt.RegisteredClaims
}

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
	// Key is the key used to sign the build jwt
	Key []byte
}

// NewJWTSettings creates new settings used to generate jwt.
func NewJWTSettings(expiresInSeconds int, iss string, sub string, aud []string, key []byte) *JWTSettings {
	return &JWTSettings{
		ExpiresInSeconds: expiresInSeconds,
		Issuer:           iss,
		Subject:          sub,
		Audience:         aud,
		Key:              key,
	}
}

// CreateJWT creates a new jwt with the given email and settings.
func (j *JWTSettings) CreateJWT(email string) JWT {
	return JWT{
		email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(j.ExpiresInSeconds) * time.Second)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    j.Issuer,
			Subject:   j.Subject,
			ID:        uuid.NewString(),
			Audience:  j.Audience,
		},
	}
}

// Sign signs the given token and returns the signed token.
func (j *JWTSettings) Sign(jwtToSign JWT) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtToSign)
	return token.SignedString(j.Key)
}

// Parse parses the given signed token and returns a new struct containing the token data.
func (j *JWTSettings) Parse(signedToken string) (*JWT, error) {
	token, err := jwt.ParseWithClaims(signedToken, &JWT{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.Errorf("Unexpected signing method, got '%v', expected 'HMAC'", token.Header["alg"])
		}

		return j.Key, nil
	})

	if err != nil {
		return nil, errors.Annotatef(err, "parse token:")
	}

	if claims, ok := token.Claims.(*JWT); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.Annotatef(err, "token not valid")
}
