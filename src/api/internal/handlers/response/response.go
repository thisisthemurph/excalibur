package response

import (
	"excalibur/internal/handlers/httperror"
	"log"
	"net/http"
)

func Respond(w http.ResponseWriter, i interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	err := ToJSON(i, w)
	if err != nil {
		log.Println("Could not format JSON response: ", err.Error())
	}
}

func ReturnError(w http.ResponseWriter, message error, status int, ) {
	httpErr := httperror.New(status, message.Error())
	Respond(w, httpErr, status)
}