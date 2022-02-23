package handlers

import (
	"assignment-1/custom_errors"
	"assignment-1/json_parser"
	"assignment-1/model_logic"
	"assignment-1/web_client"
	"assignment-1/web_server/url"
	"net/http"
)

// UnisearchHandler is the handler for the unisearch endpoint that checks for supported methods.
func UnisearchHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGetRequestUniSearch(w, r)
	default:
		// Returns method not supported for unsupported rest methods.
		custom_errors.HttpUnsupportedMethod(w)
		return
	}
}

// handleGetRequestUniSearch handles the get requests for the endpoint.
func handleGetRequestUniSearch(w http.ResponseWriter, r *http.Request) {
	urlToSearch := url.GenerateUniversitySearchString(r.URL)

	// Checks for an empty search string
	if urlToSearch == "" {
		custom_errors.HttpSearchParameters(w)
		return
	}

	response, err := web_client.GetResponseFromWebPage(urlToSearch)

	// Checks if there has been an error when fetching the api.
	if err != nil && err.Error() == custom_errors.GetUnableToReachBackendApisError().Error() {
		custom_errors.HttpErrorFromBackendApi(w)
		return
	}

	combinedUniversities, err := model_logic.Combine(json_parser.DecodeUniInfo(response))

	if err != nil {
		if err.Error() == custom_errors.GetUnableToReachBackendApisError().Error() {
			custom_errors.HttpErrorFromBackendApi(w)
			return
		} else {
			// Enters the if when no results in the university api is found.
			custom_errors.HttpNoContent(w)
			return
		}
	}

	err = json_parser.Encode(w, combinedUniversities)

	// Checks for errors in the encoding process.
	if err != nil {
		custom_errors.HttpUnknownServerError(w)
		return
	}
}
