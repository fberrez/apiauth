package api

import (
	"context"
	"encoding/json"
	"net/http"
)

type accountIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (a *API) createAccount(w http.ResponseWriter, r *http.Request) {
	var acc accountIn
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&acc); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := a.accountRepo.Create(context.TODO(), acc.Email, acc.Password, 0); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()
}
