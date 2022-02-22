package url

import (
	"assignment-1/constants"
	"assignment-1/custom_errors"
	"assignment-1/utilities"
	"net/url"
	"strconv"
	"strings"
)

// GetNameAndCountry Gets the search string from the user containing uni name and country.
func GetNameAndCountry(url *url.URL) (string, string, error) {
	searchParameters := url.Query()

	// Checking that the search terms are valid strings and exists.
	if searchParameters.Has(constants.URL_PARAM_NAME) && searchParameters.Has(constants.URL_PARAM_COUNTRY) &&
		utilities.CheckIfStringIsNotEmpty(searchParameters[constants.URL_PARAM_NAME][0]) &&
		utilities.CheckIfStringIsNotEmpty(searchParameters[constants.URL_PARAM_COUNTRY][0]) {
		return searchParameters[constants.URL_PARAM_NAME][0], searchParameters[constants.URL_PARAM_COUNTRY][0], nil
	} else {
		return "", "", custom_errors.GetParameterError()
	}
}

// GenerateUniversitySearchString generates a search string for the university api based on the user inputted url.
func GenerateUniversitySearchString(url *url.URL) string {
	searchStrings := url.Query()
	urlToSearch := strings.Builder{}
	urlToSearch.WriteString(constants.UNIVERSITY_API)

	if searchStrings.Has(constants.URL_PARAM_NAME) {
		if !utilities.CheckIfStringIsNotEmpty(searchStrings[constants.URL_PARAM_NAME][0]) {
			return ""
		}

		urlToSearch.WriteString(constants.URL_PARAM_NAME_CONTAINS)
		urlToSearch.WriteString(constants.URL_PARAM_EQUALS)
		urlToSearch.WriteString(strings.ReplaceAll(searchStrings[constants.URL_PARAM_NAME][0], " ", "%20"))
	}

	if searchStrings.Has(constants.URL_PARAM_COUNTRY) {
		if !utilities.CheckIfStringIsNotEmpty(searchStrings[constants.URL_PARAM_COUNTRY][0]) {
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

// GetLimit returns the limit specified by the user.
func GetLimit(url *url.URL) (int, error) {
	obtainedQuery := url.Query()

	if obtainedQuery.Has(constants.URL_PARAM_LIMIT) {
		limit, err := strconv.Atoi(obtainedQuery[constants.URL_PARAM_LIMIT][0])

		if !(limit > 0) || err != nil {
			return 0, custom_errors.GetInvalidLimitError()
		}

		return limit, nil
	}

	return 0, nil
}