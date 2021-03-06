package web_client

import (
	"assignment-1/constants"
	"assignment-1/custom_errors"
	"log"
	"net/http"
)

// GetResponseFromWebPage method that takes an urlHandlingClient and gets a response from the webpage.
func GetResponseFromWebPage(url string) (*http.Response, error) {
	request, errorFromRequest := http.NewRequest(http.MethodGet, url, nil)

	if errorFromRequest != nil {
		log.Println("Error when creating the request:", errorFromRequest.Error())
		return nil, custom_errors.GetUnableToReachBackendApisError()
	}

	// Setting the content type header
	request.Header.Add("content-type", "application/json")

	// Instantiate the webClient
	webClient := &http.Client{}

	// Setting timeout for web web_client
	webClient.Timeout = constants.CLIENT_TIMEOUT

	// Sending the request
	response, errorFromResponse := webClient.Do(request)

	if errorFromResponse != nil {
		log.Println("Error in the response:", errorFromResponse.Error())
		return nil, custom_errors.GetUnableToReachBackendApisError()
	}

	return response, nil
}
