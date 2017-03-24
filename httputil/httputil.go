package httputil

import "net/http"

// RespondWithError will return the specified status code and the message associated with the error
func RespondWithError(w http.ResponseWriter, err error, status int) {
	errorString := err.Error()
	RespondWithErrorMessage(w, errorString, status)
}

// RespondWithErrorMessage will return the specified status code and message
func RespondWithErrorMessage(w http.ResponseWriter, errMessage string, status int) {
	w.WriteHeader(status)
	w.Write([]byte(errMessage))
}

// RespondWithValue will return "200 OK" with the specified value
func RespondWithValue(w http.ResponseWriter, value string) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(value))
}
