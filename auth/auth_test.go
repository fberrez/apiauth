package auth

import (
	"testing"
)

func TestCreateToken(t *testing.T) {
	secret := "mysecret"
	expireInSeconds := 100
	auth := New(secret, expireInSeconds)
	if auth.secret != secret {
		t.Fatalf("secret not valid, expected '%s', got '%s'", secret, auth.secret)
	}

	if auth.expireInSeconds != expireInSeconds {
		t.Fatalf("expireInSeconds not valid, expected '%d', got '%d'", expireInSeconds, auth.expireInSeconds)
	}

	token, err := auth.CreateToken()
	if err != nil {
		t.Fatalf("%v", err)
	} else {
		t.Logf("%s", token)
	}
}
