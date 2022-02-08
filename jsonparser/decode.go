package jsonparser

import (
	"assignment-1/model"
	"encoding/json"
	"log"
	"net/http"
)

//TODO: Verify valid response?

// DecodeUniInfo function that takes a http request and decodes the json body.
func DecodeUniInfo(httpResponse *http.Response) []model.UniversityInfo {
	decoder := json.NewDecoder(httpResponse.Body)
	var list []model.UniversityInfo

	if err := decoder.Decode(&list); err != nil {
		log.Fatal(err)
	}

	return list
}

// DecodeCountryInfo function that takes a http request and decodes the json body.
func DecodeCountryInfo(httpResponse *http.Response) []model.CountryApi {
	decoder := json.NewDecoder(httpResponse.Body)
	var list []model.CountryApi

	if err := decoder.Decode(&list); err != nil {
		log.Fatal(err)
	}

	return list
}
