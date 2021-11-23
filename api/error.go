package api

import (
	"net/http"

	"github.com/go-chi/render"
)

// ErrOut renderer type for handling all sorts of errors.
type ErrOut struct {
	// Err is the low-level runtime error
	Err error `json:"-"`
	// HTTPStatusCode is the http response status code
	HTTPStatusCode int `json:"-"`

	// StatusText is the user-level status message
	StatusText string `json:"status"`
	// ErrorText is the application-level error message
	ErrorText string `json:"error,omitempty"`
}

// Render is the function called to render the error.
func (e *ErrOut) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

// ErrBadRequest is called when a bad request error occured in a controller.
func ErrBadRequest(err error) render.Renderer {
	return &ErrOut{
		Err:            err,
		HTTPStatusCode: http.StatusBadRequest,
		StatusText:     "Invalid Request.",
		ErrorText:      err.Error(),
	}
}
