package model_logic

import (
	"assignment-1/client"
	"assignment-1/constants"
	"assignment-1/model"
	"fmt"
	"net/http"
	"time"
)

// Variable containing the start time of the server.
var startTime time.Time

// SetStartTime sets the startTime variable to the current time.
func SetStartTime() {
	startTime = time.Now()
}

// GetUptime returns the current time since the startTime variable was set (the uptime).
func getUptime() string {
	return fmt.Sprintf("%ds", int(time.Since(startTime).Seconds()))
}

// getStatusCode gets the status code from a webpage specified by an url.
func getStatusCode(url string) (int, error) {
	statusCode, err := client.GetResponseFromWebPage(url)

	if err != nil {
		return 0, err
	}

	return statusCode.StatusCode, err
}

// GetDiagInfo gets the diagnosis information and turns it into a struct.
func GetDiagInfo() model.Diagnostics {
	countryApiStatusCode, errCountry := getStatusCode(constants.COUNTRY_API_ROOT_URL)
	universityApiStatusCode, errUni := getStatusCode(constants.UNIVERSITY_API_ROOT_URL)

	if errCountry != nil {
		countryApiStatusCode = http.StatusGatewayTimeout
	}
	if errUni != nil {
		universityApiStatusCode = http.StatusGatewayTimeout
	}

	return model.Diagnostics{
		CountryApiStatus:    countryApiStatusCode,
		UniversityApiStatus: universityApiStatusCode,
		Uptime:              getUptime(),
		Version:             constants.PROGRAM_VERSION,
	}
}
