package json_parser

import (
	"assignment-1/model"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// DecodeUniInfo function that takes a http request and decodes the json body into a slice UniversityInfo structs.
func DecodeUniInfo(httpResponse *http.Response) []model.UniversityInfo {
	decoder := json.NewDecoder(httpResponse.Body)
	var list []model.UniversityInfo

	// Checks for errors in the decoding process.
	if err := decoder.Decode(&list); err != nil {
		log.Fatal(err)
	}

	return list
}

// DecodeCountryInfo function that takes a http request and decodes the json body into a slice of CountryApi structs.
// The checking for list or object is inspired from this stack overflow comment:
// https://stackoverflow.com/a/61837617
func DecodeCountryInfo(httpResponse *http.Response) []model.CountryApi {
	var responseBytes []byte
	var list []model.CountryApi
	var countryObject model.CountryApi

	// If nothing is found in the country api nothing is returned.
	if httpResponse.StatusCode == http.StatusBadRequest || httpResponse.StatusCode == http.StatusNotFound {
		return nil
	}

	// Turns the response body into a byte array
	responseBytes, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		log.Println(err)
		return nil
	}

	// Checks if the response is a list or an object and handles them accordingly.
	// This is needed because the country api might return an object when searching for alpha codes with constraints.
	if responseBytes[0] == '[' {
		// When the byte array starts with a [ it means that the country api has returned a list.
		if err := json.Unmarshal(responseBytes, &list); err != nil {
			log.Println(err)
		}
	} else if responseBytes[0] == '{' {
		// When the byte array starts with a { it means that the country api has returned an object.
		if err := json.Unmarshal(responseBytes, &countryObject); err != nil {
			log.Println(err)
		}
		list = append(list, countryObject)
	}

	return list
}
