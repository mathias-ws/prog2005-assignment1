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
		http.Error(w, "Method not supported.", http.StatusMethodNotAllowed)
	}
}

// handleGetRequestNeighbourUnis handles the get requests for the endpoint.
func handleGetRequestNeighbourUnis(w http.ResponseWriter, r *http.Request) {
	uniName, country, err := url.GetNameAndCountry(r.URL)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	limit, err := url.GetLimit(r.URL)

	// When the limit is invalid
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	valuesToEncode, err := model_logic.GetUniversitiesBorderingTo(uniName, country, limit)

	if err != nil {
		if err.Error() == customErrors.GetUnableToReachBackendApisError().Error() {
			http.Error(w, "Error from backend api", http.StatusBadGateway)
			return
		} else {
			// Enters when the country does not exist in the country api.
			http.Error(w, "No results found for current request.", http.StatusNoContent)
			return
		}
	}

	err = jsonparser.Encode(w, valuesToEncode)

	if err != nil {
		http.Error(w, "Server side error, please try again later", http.StatusInternalServerError)
		return
	}
}
