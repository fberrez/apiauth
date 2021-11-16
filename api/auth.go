package api

import (
	"fmt"
	"net/http"
)

func (a *API) createToken(w http.ResponseWriter, r *http.Request) {
	token, err := a.auth.CreateToken()
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%v", err)))
		return
	}

	w.Write([]byte(token))
	return
}
