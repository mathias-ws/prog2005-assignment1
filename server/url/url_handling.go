package url

import (
	"assignment-1/constants"
	"net/url"
	"strings"
)

// GetNameAndCountry Gets the search string from the user containing uni name and country.
func GetNameAndCountry(url *url.URL) (string, string) {
	searchStrings := url.Query()
	return searchStrings[constants.URL_PARAM_NAME][0], searchStrings[constants.URL_PARAM_COUNTRY][0]
}

// GenerateUniversitySearchString generates a search string for the university api based on the user inputted url.
func GenerateUniversitySearchString(url *url.URL) string {
	searchStrings := url.Query()
	urlToSearch := constants.UNIVERSITY_API

	if searchStrings.Has(constants.URL_PARAM_NAME) {
		urlToSearch += constants.URL_PARAM_NAME + constants.URL_PARAM_EQUALS +
			strings.ReplaceAll(searchStrings[constants.URL_PARAM_NAME][0], " ", "%20")
	}
	if searchStrings.Has(constants.URL_PARAM_COUNTRY) {
		// Adds an and if the name parameter is used as well.
		if strings.Contains(urlToSearch, constants.URL_PARAM_NAME) {
			urlToSearch += constants.URL_PARAM_AND
		}

		urlToSearch += constants.URL_PARAM_COUNTRY + constants.URL_PARAM_EQUALS +
			strings.ReplaceAll(searchStrings[constants.URL_PARAM_COUNTRY][0], " ", "%20")

	}

	return urlToSearch
}
