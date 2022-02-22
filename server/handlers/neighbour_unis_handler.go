package handlers

import (
	"assignment-1/customErrors"
	"assignment-1/jsonparser"
	"assignment-1/model_logic"
	"assignment-1/server/url"
	"net/http"
)

// NeighbourUnisHandler handler for the neighbouruni endpoint that checks for supported methods.
func NeighbourUnisHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGetRequestNeighbourUnis(w, r)
	default:
		// Returns method not supported for unsupported rest methods.
		customErrors.HttpUnsupportedMethod(w)
		return
	}
}

// handleGetRequestNeighbourUnis handles the get requests for the endpoint.
func handleGetRequestNeighbourUnis(w http.ResponseWriter, r *http.Request) {
	uniName, country, err := url.GetNameAndCountry(r.URL)

	// When invalid parameters.
	if err != nil {
		customErrors.HttpSearchParameters(w)
		return
	}

	limit, err := url.GetLimit(r.URL)

	// When the limit is invalid
	if err != nil {
		customErrors.HttpSearchParameters(w)
		return
	}

	valuesToEncode, err := model_logic.GetUniversitiesBorderingTo(uniName, country, limit)

	if err != nil {
		if err.Error() == customErrors.GetUnableToReachBackendApisError().Error() {
			// When the error is related to the backend apis.
			customErrors.HttpErrorFromBackendApi(w)
			return
		} else {
			// Enters when the country does not exist in the country api.
			customErrors.HttpNoContent(w)
			return
		}
	}

	err = jsonparser.Encode(w, valuesToEncode)

	if err != nil {
		customErrors.HttpUnknownServerError(w)
		return
	}
}
