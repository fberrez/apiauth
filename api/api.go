package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/fberrez/apiauth/auth"
	"github.com/fberrez/apiauth/backend"
	"github.com/fberrez/apiauth/models"
	"github.com/gorilla/mux"
)

type API struct {
	auth        *auth.Auth
	router      *mux.Router
	accountRepo *models.AccountRepo
}

func New(ctx context.Context, authSecret string, authExpireInSeconds int) (*API, error) {
	router := mux.NewRouter()

	// Initializes dependencies
	jwt := auth.NewJWTSettings(100, "myapp", "apiAuthentication", []string{"authenticatedUsers"}, []byte("mysecret"))
	redis, err := backend.NewRedis("127.0.0.1:6379", "", ctx)
	if err != nil {
		return nil, err
	}

	postgres, err := backend.NewPostgres(ctx, "postgres://postgres:postgres@localhost:5432/apiauth")
	if err != nil {
		return nil, err
	}

	// Defines API
	api := &API{
		auth:        auth.New(jwt, redis),
		accountRepo: models.NewAccountRepo(postgres),
	}

	// Defines routes
	router.HandleFunc("/auth/token", api.createToken).Methods("GET")
	router.HandleFunc("/accounts", api.createAccount).Methods("POST")

	api.router = router

	return api, nil
}

func (a *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
