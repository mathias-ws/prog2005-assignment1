package client

import (
	"assignment-1/constants"
	"log"
	"net/http"
)

//TODO: Send http error når ingenting matcher søket

// GetResponseFromWebPage method that takes an url and gets a json response from the webpage.
func GetResponseFromWebPage(url string) *http.Response {
	request, errorFromRequest := http.NewRequest(http.MethodGet, url, nil)

	if errorFromRequest != nil {
		log.Println("Error when creating the request:", errorFromRequest.Error())
	}

	// Setting the content type header
	request.Header.Add("content-type", "application/json")

	// Instantiate the webClient
	webClient := &http.Client{}

	// Setting timeout for web client
	webClient.Timeout = constants.CLIENT_TIMEOUT

	// Sending the request
	response, errorFromResponse := webClient.Do(request)

	if errorFromResponse != nil {
		log.Println("Error in the response:", errorFromResponse.Error())
	}

	return response
}
