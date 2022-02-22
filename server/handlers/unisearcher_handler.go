package handlers

import (
	"assignment-1/client"
	"assignment-1/customErrors"
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
		customErrors.HttpUnsupportedMethod(w)
		return
	}
}

// handleGetRequestUniSearch handles the get requests for the endpoint.
func handleGetRequestUniSearch(w http.ResponseWriter, r *http.Request) {
	urlToSearch := url.GenerateUniversitySearchString(r.URL)

	if urlToSearch == "" {
		customErrors.HttpSearchParameters(w)
		return
	}

	response, err := client.GetResponseFromWebPage(urlToSearch)

	if err != nil {
		if err.Error() == customErrors.GetUnableToReachBackendApisError().Error() {
			customErrors.HttpErrorFromBackendApi(w)
			return
		}
	}

	combinedUniversities, err := model_logic.Combine(jsonparser.DecodeUniInfo(response))

	if err != nil {
		// Enters the if when no results in the university api is found.
		customErrors.HttpNoContent(w)
		return
	}

	err = jsonparser.Encode(w, combinedUniversities)

	if err != nil {
		customErrors.HttpUnknownServerError(w)
		return
	}
}
