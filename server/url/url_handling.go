package url

import (
	"assignment-1/constants"
	"errors"
	"net/url"
	"strings"
)

// GetNameAndCountry Gets the search string from the user containing uni name and country.
func GetNameAndCountry(url *url.URL) (string, string, error) {
	searchParameters := url.Query()
	if searchParameters.Has(constants.URL_PARAM_NAME) && searchParameters.Has(constants.URL_PARAM_COUNTRY) {
		return searchParameters[constants.URL_PARAM_NAME][0], searchParameters[constants.URL_PARAM_COUNTRY][0], nil
	} else {
		return "", "", errors.New("missing parameters or wrong parameters")
	}
}

// GenerateUniversitySearchString generates a search string for the university api based on the user inputted url.
func GenerateUniversitySearchString(url *url.URL) string {
	searchStrings := url.Query()
	urlTest := strings.Builder{}
	urlTest.WriteString(constants.UNIVERSITY_API)

	if searchStrings.Has(constants.URL_PARAM_NAME) {
		urlTest.WriteString(constants.URL_PARAM_NAME)
		urlTest.WriteString(constants.URL_PARAM_EQUALS)
		urlTest.WriteString(strings.ReplaceAll(searchStrings[constants.URL_PARAM_NAME][0], " ", "%20"))
	}

	if searchStrings.Has(constants.URL_PARAM_COUNTRY) {
		// Adds an and if the name parameter is used as well.
		if strings.Contains(urlTest.String(), constants.URL_PARAM_NAME) {
			urlTest.WriteString(constants.URL_PARAM_AND)
		}

		urlTest.WriteString(constants.URL_PARAM_COUNTRY)
		urlTest.WriteString(constants.URL_PARAM_EQUALS)
		urlTest.WriteString(strings.ReplaceAll(searchStrings[constants.URL_PARAM_COUNTRY][0], " ", "%20"))
	}

	return urlTest.String()
}
