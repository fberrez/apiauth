package api

import (
	"net/http"

	"github.com/fberrez/apiauth/auth"
	"github.com/gorilla/mux"
)

type API struct {
	auth   *auth.Auth
	router *mux.Router
}

func New(authSecret string, authExpireInSeconds int) *API {
	router := mux.NewRouter()

	// Initializes dependencies
	auth := auth.New(authSecret, authExpireInSeconds)

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
