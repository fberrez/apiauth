package api

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/fberrez/apiauth/auth"
	"github.com/fberrez/apiauth/backend"
	"github.com/fberrez/apiauth/models"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
)

type API struct {
	auth        *auth.Auth
	router      *chi.Mux
	accountRepo *models.AccountRepo
}

func New(ctx context.Context, authSecret string, authExpireInSeconds int) (*API, error) {
	r := chi.NewRouter()

	// Declares middlewares
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	// Initializes dependencies
	jwt := auth.NewJWTSettings(100, "myapp", "apiAuthentication", []string{"authenticatedUsers"}, []byte("mysecret"), []byte("myverify"))
	redis, err := backend.NewRedis("redis:6379", "", ctx)
	if err != nil {
		return nil, err
	}

	postgres, err := backend.NewPostgres(ctx, "postgres://postgres:postgres@postgres:5432/apiauth")
	if err != nil {
		return nil, err
	}

	// Defines API
	api := &API{
		auth:        auth.New(jwt, redis),
		accountRepo: models.NewAccountRepo(postgres),
	}

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(jwt.TokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Route("/accounts", func(r chi.Router) {
			r.Post("/", api.createAccount)

			// /accounts/{id}
			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", notImplemented)
				r.Put("/", notImplemented)
				r.Delete("/", notImplemented)
			})
		})
	})

	// Defines routes

	r.Post("/login", api.Login)

	api.router = r

	return api, nil
}

func (a *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}

// notImplemented is a naive controller used to return a 501 response for routes in development.
func notImplemented(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
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
