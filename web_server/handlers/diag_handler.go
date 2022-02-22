package handlers

import (
	"assignment-1/custom_errors"
	"assignment-1/json_parser"
	"assignment-1/model_logic"
	"net/http"
)

// DiagHandler checks for the http method and handles the error appropriately.
func DiagHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGetRequestDiag(w)
	default:
		// Returns method not supported for unsupported rest methods.
		custom_errors.HttpUnsupportedMethod(w)
	}

}

// handleGetRequestDiag handles the get requests for the endpoint. Gets the diag info struct and encodes it.
func handleGetRequestDiag(w http.ResponseWriter) {
	err := json_parser.Encode(w, model_logic.GetDiagInfo())

	// Checks for errors in the encoding process.
	if err != nil {
		custom_errors.HttpUnknownServerError(w)
		return
	}
}
