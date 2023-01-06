// Package errorhandler assists with normalizing errors
package errorhandler

import (
	"excalibur/internal/handler/response"
	"fmt"
	"net/http"
)

// ErrorHandler structure representing possible errors
type ErrorHandler struct {
	em map[error]int
}

// New creates a new error handler
func New() ErrorHandler {
	return ErrorHandler{
		em: makeErrorMap(),
	}
}

// GetStatus returns the http status code (-1 if not an appropriate error) associated with the given error
func (e ErrorHandler) GetStatus(err error) (int, error) {
	status, prs := e.em[err]
	if !prs {
		return -1, fmt.Errorf("the given error does not exist in the error map")
	}

	return status, nil
}

// TODO: Should this be moved to the `httperror` package and this package moved somewhere better?

// HandleAPIError writes an error response to the http.ResponseWriter if there is an error
func (e ErrorHandler) HandleAPIError(w http.ResponseWriter, err error) int {
	if err == nil {
		return http.StatusOK
	}

	status, _ := e.GetStatus(err)
	response.ReturnError(w, err, status)

	return status
}
