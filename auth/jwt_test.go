package auth

import (
	"reflect"
	"testing"
)

func TestNewJWTSettings(t *testing.T) {
	expiresInSeconds := 10
	iss := "myapp"
	sub := "api"
	aud := []string{"registredUsers"}
	key := []byte("mysecretkey")
	settings := NewJWTSettings(expiresInSeconds, iss, sub, aud, key)
	if settings.ExpiresInSeconds != expiresInSeconds {
		t.Fatalf("ExpiresInSeconds not valid, got '%d', expected '%d'", settings.ExpiresInSeconds, expiresInSeconds)
	}

	if settings.Issuer != iss {
		t.Fatalf("Issuer not valid, got '%s', expected '%s'", settings.Issuer, iss)
	}

	if settings.Subject != sub {
		t.Fatalf("Subject not valid, got '%s', expected '%s'", settings.Subject, sub)
	}

	if !reflect.DeepEqual(settings.Audience, aud) {
		t.Fatalf("Audience not valid, got '%v', expected '%v'", settings.Audience, aud)
	}

	if !reflect.DeepEqual(settings.Key, key) {
		t.Fatalf("Key not valid, got '%s', expected '%s'", string(settings.Key), key)
	}
}

func TestCreateJWT(t *testing.T) {
	expiresInSeconds := 10
	iss := "myapp"
	sub := "api"
	aud := []string{"registredUsers"}
	key := []byte("mysecretkey")
	settings := NewJWTSettings(expiresInSeconds, iss, sub, aud, key)

	email := "user@host.com"
	token := settings.CreateJWT(email)
	if token.Email == email {
		t.Fatalf("Email not valid, got '%s', expected '%s'", token.Email, email)
	}
}

func TestSign(t *testing.T) {
	expiresInSeconds := 10
	iss := "myapp"
	sub := "api"
	aud := []string{"registredUsers"}
	key := []byte("mysecretkey")
	settings := NewJWTSettings(expiresInSeconds, iss, sub, aud, key)

	email := "user@host.com"
	token := settings.CreateJWT(email)
	_, err := settings.Sign(token)
	if err != nil {
		t.Fatalf("%v", err)
	}
}

func TestParse(t *testing.T) {
	expiresInSeconds := 10
	iss := "myapp"
	sub := "api"
	aud := []string{"registredUsers"}
	key := []byte("mysecretkey")
	settings := NewJWTSettings(expiresInSeconds, iss, sub, aud, key)

	email := "user@host.com"
	token := settings.CreateJWT(email)
	signed, err := settings.Sign(token)
	if err != nil {
		t.Fatalf("%v", err)
	}

	parsedToken, err := settings.Parse(signed)
	if err != nil {
		t.Fatalf("%v", err)
	}

	t.Logf("%+v", parsedToken)
}
