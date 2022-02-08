package jsonparser

import (
	"assignment-1/model"
	"encoding/json"
	"log"
	"net/http"
)

//TODO: Verify valid response?

// DecodeUniInfoIn function that takes a http request and decodes the json body.
func DecodeUniInfoIn(httpResponse *http.Response) []model.UniFromApi {
	decoder := json.NewDecoder(httpResponse.Body)
	var list []model.UniFromApi

	if err := decoder.Decode(&list); err != nil {
		log.Fatal(err)
	}

	return list
}
