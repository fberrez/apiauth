package api

import (
	"context"
	"net/http"

	"github.com/fberrez/apiauth/auth"
	"github.com/fberrez/apiauth/backend"
	"github.com/gorilla/mux"
)

type API struct {
	auth   *auth.Auth
	router *mux.Router
}

func New(authSecret string, authExpireInSeconds int) *API {
	router := mux.NewRouter()

	// Initializes dependencies
	jwt := auth.NewJWTSettings(100, "myapp", "apiAuthentication", []string{"authenticatedUsers"}, []byte("mysecret"))
	redis, err := backend.NewRedis("127.0.0.1:6379", "", context.Background())
	if err != nil {
		panic(err)
	}
	auth := auth.New(jwt, redis)

	// Defines API
	api := &API{
		auth: auth,
	}

	// Defines routes
	router.HandleFunc("/auth/token", api.createToken).Methods("GET")

	api.router = router

	return api
}

func (a *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}
