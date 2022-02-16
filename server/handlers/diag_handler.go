package handlers

import (
	"assignment-1/jsonparser"
	"assignment-1/model_logic"
	"net/http"
)

// DiagHandler gets the diag info struct and encodes it.
func DiagHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGetRequestDiag(w)
	default:
		// Returns method not supported for unsupported rest methods.
		http.Error(w, "Method not supported.", http.StatusMethodNotAllowed)
	}

}

// handleGetRequestDiag handles the get requests for the endpoint.
func handleGetRequestDiag(w http.ResponseWriter) {
	jsonparser.Encode(w, model_logic.GetDiagInfo())
}
