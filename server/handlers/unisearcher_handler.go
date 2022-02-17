package handlers

import (
	"assignment-1/client"
	"assignment-1/jsonparser"
	"assignment-1/model_logic"
	"assignment-1/server/url"
	"net/http"
)

// UnisearchHandler is the handler for the unisearch endpoint that checks for supported methods.
func UnisearchHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGetRequestUniSearch(w, r)
	default:
		// Returns method not supported for unsupported rest methods.
		http.Error(w, "Method not supported.", http.StatusMethodNotAllowed)
	}
}

// handleGetRequestUniSearch handles the get requests for the endpoint.
func handleGetRequestUniSearch(w http.ResponseWriter, r *http.Request) {
	combinedUniversities, err := model_logic.Combine(jsonparser.DecodeUniInfo(
		client.GetResponseFromWebPage(url.GenerateUniversitySearchString(r.URL))))

	// Enters the if when no results in the university api is found.
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	jsonparser.Encode(w, combinedUniversities)
}
