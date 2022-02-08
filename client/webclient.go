package client

import (
	"fmt"
	"net/http"
)

//TODO: Send http error når ingenting matcher søket

// GetResponseFromWebPage method that takes an url and gets a json response from the webpage.
func GetResponseFromWebPage(url string) *http.Response {
	request, errorFromRequest := http.NewRequest(http.MethodGet, url, nil)

	if errorFromRequest != nil {
		fmt.Errorf("Error when creating the request:", errorFromRequest.Error())
	}

	// Setting the content type header
	request.Header.Add("content-type", "application/json")

	// Instantiate the webClient
	webClient := &http.Client{}

	// Sending the request
	response, errorFromResponse := webClient.Do(request)

	if errorFromResponse != nil {
		fmt.Errorf("Error in the response:", errorFromResponse.Error())
	}

	return response
}
