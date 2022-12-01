package errorhandler

import (
	"errors"
	"net/http"
)

func makeErrorMap() map[error]int {
	return map[error]int{
		ErrorNotFound:       http.StatusNotFound,
		ErrorAlreadyExists:  http.StatusBadRequest,
		ErrorJSONParse:      http.StatusBadRequest,
		ErrorInternalServer: http.StatusInternalServerError,
	}
}

var (
	// ErrorNotFound 404
	ErrorNotFound = errors.New("the specific item could not be found")

	// ErrorAlreadyExists 400
	ErrorAlreadyExists = errors.New("the specified item already exists")

	// ErrorJSONParse 400
	ErrorJSONParse = errors.New("could not process the given data")

	// ErrorInternalServer 500
	ErrorInternalServer = errors.New("there has been an internal server error")
)
