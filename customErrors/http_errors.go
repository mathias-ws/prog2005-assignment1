package customErrors

import (
	"net/http"
)

// HttpSearchParameters http error message for when the parameters or its values are wrong.
func HttpSearchParameters(w http.ResponseWriter) {
	http.Error(w, "Search must contain the valid amount of search parameter(s) "+
		"with a valid value. See the documentation.", http.StatusBadRequest)
}

// HttpErrorFromBackendApi http error message for when the backend apis returns an error.
func HttpErrorFromBackendApi(w http.ResponseWriter) {
	http.Error(w, "Error from backend api", http.StatusBadGateway)
}

// HttpUnsupportedMethod http error message for when the rest method is invalid.
func HttpUnsupportedMethod(w http.ResponseWriter) {
	http.Error(w, "Method not supported.", http.StatusMethodNotAllowed)
}

// HttpUnknownServerError http error message for when the server has an undefined error or an error the user should not know.
func HttpUnknownServerError(w http.ResponseWriter) {
	http.Error(w, "Server side error, please try again later", http.StatusInternalServerError)
}

// HttpNoContent http error for when no contet is returned.
func HttpNoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}
