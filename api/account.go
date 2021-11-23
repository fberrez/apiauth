package api

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/fberrez/apiauth/models"
	"github.com/go-chi/render"
)

type AccountIn struct {
	Account *models.Account
}

type AccountOut struct {
	Account *models.Account
}

func (a *API) createAccount(w http.ResponseWriter, r *http.Request) {
	data := &AccountIn{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrBadRequest(err))
		return
	}

	account := data.Account
	fmt.Printf("%#v\n", account)
	render.Status(r, http.StatusCreated)
	render.Render(w, r, NewAccountOut(account))
}

func (a *AccountIn) Bind(r *http.Request) error {
	if a.Account == nil {
		return errors.New("Missing Required Account Fields")
	}

	a.Account.Username = strings.Trim(a.Account.Username, " ")
	a.Account.Password = strings.Trim(a.Account.Password, " ")
	if a.Account.Username == "" {
		return errors.New("Missing Required Field: Username")
	}

	if a.Account.Password == "" {
		return errors.New("Missing Required Field: Password")
	}

	return nil
}

func (a *AccountOut) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// NewAccountOut creates a new instance of AccountOut, used as api response.
func NewAccountOut(account *models.Account) *AccountOut {
	resp := &AccountOut{Account: account}
	return resp
}
