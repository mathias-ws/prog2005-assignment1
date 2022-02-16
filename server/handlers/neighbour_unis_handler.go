package handlers

import (
	"assignment-1/jsonparser"
	"assignment-1/model_logic"
	"assignment-1/server/url"
	"net/http"
)

// NeighbourUnisHandler handler for the neighbouruni endpoint.
func NeighbourUnisHandler(w http.ResponseWriter, r *http.Request) {
	jsonparser.Encode(w, model_logic.GetUniversitiesBorderingTo(url.GetNameAndCountry(r.URL)))
}
