package model_logic

import (
	"assignment-1/constants"
	"assignment-1/model"
	"assignment-1/web_client"
	"fmt"
	"net/http"
	"time"
)

// Variable containing the start time of the web_server.
var startTime time.Time

// SetStartTime sets the startTime variable to the current time.
func SetStartTime() {
	startTime = time.Now()
}

// GetUptime returns the current time since the startTime variable was set (the uptime).
func getUptime() string {
	return fmt.Sprintf("%ds", int(time.Since(startTime).Seconds()))
}

// getStatusCode gets the status code from a webpage specified by an urlHandlingClient.
func getStatusCode(url string) (int, error) {
	statusCode, err := web_client.GetResponseFromWebPage(url)

	// Checks for errors when fetching the api.
	if err != nil {
		return 0, err
	}

	return statusCode.StatusCode, err
}

// GetDiagInfo gets the diagnosis information and turns it into a struct.
func GetDiagInfo() model.Diagnostics {
	// Gets the status codes.
	countryApiStatusCode, errCountry := getStatusCode(constants.COUNTRY_API_ROOT_URL)
	universityApiStatusCode, errUni := getStatusCode(constants.UNIVERSITY_API_ROOT_URL)

	// If the apis are unreachable set a proper error code.
	if errCountry != nil {
		countryApiStatusCode = http.StatusBadGateway
	}
	if errUni != nil {
		universityApiStatusCode = http.StatusBadGateway
	}

	return model.Diagnostics{
		CountryApiStatus:    countryApiStatusCode,
		UniversityApiStatus: universityApiStatusCode,
		Uptime:              getUptime(),
		Version:             constants.PROGRAM_VERSION,
	}
}
