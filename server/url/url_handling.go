package url

import (
	"assignment-1/constants"
	"assignment-1/utils"
	"errors"
	"net/url"
	"strings"
)

// GetNameAndCountry Gets the search string from the user containing uni name and country.
func GetNameAndCountry(url *url.URL) (string, string, error) {
	searchParameters := url.Query()

	// Checking that the search terms are valid strings and exists.
	if searchParameters.Has(constants.URL_PARAM_NAME) && searchParameters.Has(constants.URL_PARAM_COUNTRY) &&
		utils.CheckIfStringIsNotEmpty(searchParameters[constants.URL_PARAM_NAME][0]) &&
		utils.CheckIfStringIsNotEmpty(searchParameters[constants.URL_PARAM_COUNTRY][0]) {
		return searchParameters[constants.URL_PARAM_NAME][0], searchParameters[constants.URL_PARAM_COUNTRY][0], nil
	} else {
		return "", "", errors.New("missing parameters or wrong parameters")
	}
}

// GenerateUniversitySearchString generates a search string for the university api based on the user inputted url.
func GenerateUniversitySearchString(url *url.URL) string {
	searchStrings := url.Query()
	urlToSearch := strings.Builder{}
	urlToSearch.WriteString(constants.UNIVERSITY_API)

	if searchStrings.Has(constants.URL_PARAM_NAME) {
		if !utils.CheckIfStringIsNotEmpty(searchStrings[constants.URL_PARAM_NAME][0]) {
			return ""
		}

		urlToSearch.WriteString(constants.URL_PARAM_NAME)
		urlToSearch.WriteString(constants.URL_PARAM_EQUALS)
		urlToSearch.WriteString(strings.ReplaceAll(searchStrings[constants.URL_PARAM_NAME][0], " ", "%20"))
	}

	if searchStrings.Has(constants.URL_PARAM_COUNTRY) {
		if !utils.CheckIfStringIsNotEmpty(searchStrings[constants.URL_PARAM_COUNTRY][0]) {
			return ""
		}

		// Adds an and if the name parameter is used as well.
		if strings.Contains(urlToSearch.String(), constants.URL_PARAM_NAME) {
			urlToSearch.WriteString(constants.URL_PARAM_AND)
		}

		urlToSearch.WriteString(constants.URL_PARAM_COUNTRY)
		urlToSearch.WriteString(constants.URL_PARAM_EQUALS)
		urlToSearch.WriteString(strings.ReplaceAll(searchStrings[constants.URL_PARAM_COUNTRY][0], " ", "%20"))
	}

	if !strings.Contains(urlToSearch.String(), constants.URL_PARAM_NAME) &&
		!strings.Contains(urlToSearch.String(), constants.URL_PARAM_COUNTRY) {
		return ""
	}

	return urlToSearch.String()
}
