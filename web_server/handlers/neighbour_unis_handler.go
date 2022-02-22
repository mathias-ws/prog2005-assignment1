package handlers

import (
	"assignment-1/custom_errors"
	"assignment-1/json_parser"
	"assignment-1/model_logic"
	"assignment-1/web_server/url"
	"net/http"
)

// NeighbourUnisHandler handler for the neighbouruni endpoint that checks for supported methods.
func NeighbourUnisHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGetRequestNeighbourUnis(w, r)
	default:
		// Returns method not supported for unsupported rest methods.
		custom_errors.HttpUnsupportedMethod(w)
		return
	}
}

// handleGetRequestNeighbourUnis handles the get requests for the endpoint.
func handleGetRequestNeighbourUnis(w http.ResponseWriter, r *http.Request) {
	uniName, country, err := url.GetNameAndCountry(r.URL)

	// When invalid parameters.
	if err != nil {
		custom_errors.HttpSearchParameters(w)
		return
	}

	limit, err := url.GetLimit(r.URL)

	// When the limit is invalid
	if err != nil {
		custom_errors.HttpSearchParameters(w)
		return
	}

	valuesToEncode, err := model_logic.GetUniversitiesBorderingTo(uniName, country, limit)

	if err != nil {
		if err.Error() == custom_errors.GetUnableToReachBackendApisError().Error() {
			// When the error is related to the backend apis.
			custom_errors.HttpErrorFromBackendApi(w)
			return
		} else {
			// Enters when the country does not exist in the country api.
			custom_errors.HttpNoContent(w)
			return
		}
	}

	err = json_parser.Encode(w, valuesToEncode)

	// Checks for errors when encoding the data.
	if err != nil {
		custom_errors.HttpUnknownServerError(w)
		return
	}
}
