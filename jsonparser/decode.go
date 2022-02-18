package jsonparser

import (
	"assignment-1/model"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

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
// The checking for list or object is inspired from this stack overflow comment:
// https://stackoverflow.com/a/61837617
func DecodeCountryInfo(httpResponse *http.Response) []model.CountryApi {
	var p []byte
	var list []model.CountryApi
	var countryObject model.CountryApi

	// If nothing is found in the country api nothing is returned.
	if httpResponse.StatusCode == http.StatusBadRequest || httpResponse.StatusCode == http.StatusNotFound {
		return nil
	}

	// Turns the response body into a byte array
	p, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return nil
	}

	// Checks if the response is a list or an object and handles them accordingly.
	if p[0] == '[' {
		if err := json.Unmarshal(p, &list); err != nil {
			log.Println(err)
		}
	} else if p[0] == '{' {
		if err := json.Unmarshal(p, &countryObject); err != nil {
			log.Println(err)
		}
		list = append(list, countryObject)
	}

	return list
}
