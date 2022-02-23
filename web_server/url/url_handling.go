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

	// Checks if the name parameter exists.
	if searchStrings.Has(constants.URL_PARAM_NAME) {
		// Checks if it has a valid value.
		if !utilities.CheckIfStringIsNotEmpty(searchStrings[constants.URL_PARAM_NAME][0]) {
			return ""
		}

		// Writes it to the url.
		urlToSearch.WriteString(constants.URL_PARAM_NAME_CONTAINS)
		urlToSearch.WriteString(constants.URL_PARAM_EQUALS)
		urlToSearch.WriteString(strings.ReplaceAll(searchStrings[constants.URL_PARAM_NAME][0], " ", "%20"))
	}

	// Checks if the country parameter exists.
	if searchStrings.Has(constants.URL_PARAM_COUNTRY) {
		// Checks if it has a valid value.
		if !utilities.CheckIfStringIsNotEmpty(searchStrings[constants.URL_PARAM_COUNTRY][0]) {
			return ""
		}

		// Adds an '&' if the name parameter is used as well.
		if strings.Contains(urlToSearch.String(), constants.URL_PARAM_NAME) {
			urlToSearch.WriteString(constants.URL_PARAM_AND)
		}

		// Writes it to the url.
		urlToSearch.WriteString(constants.URL_PARAM_COUNTRY)
		urlToSearch.WriteString(constants.URL_PARAM_EQUALS)
		urlToSearch.WriteString(strings.ReplaceAll(searchStrings[constants.URL_PARAM_COUNTRY][0], " ", "%20"))
	}

	// Checks if no parameters exists.
	if !strings.Contains(urlToSearch.String(), constants.URL_PARAM_NAME) &&
		!strings.Contains(urlToSearch.String(), constants.URL_PARAM_COUNTRY) {
		return ""
	}

	return urlToSearch.String()
}

// GetLimit returns the limit specified by the user.
func GetLimit(url *url.URL) (int, error) {
	obtainedQuery := url.Query()

	// Checks if the limit parameter exists.
	if obtainedQuery.Has(constants.URL_PARAM_LIMIT) {
		limit, err := strconv.Atoi(obtainedQuery[constants.URL_PARAM_LIMIT][0])

		// Checks that the value is valid (bigger than zero).
		if !(limit > 0) || err != nil {
			return 0, custom_errors.GetInvalidLimitError()
		}

		return limit, nil
	}

	return 0, nil
}

// GenerateBaseUrlForCountrySearch generates the bas url that the GetUniversitiesBorderingTo function uses.
func GenerateBaseUrlForCountrySearch(universityName string) string {
	baseUrlToSearch := strings.Builder{}
	baseUrlToSearch.WriteString(constants.UNIVERSITY_API)
	baseUrlToSearch.WriteString(constants.URL_PARAM_NAME_CONTAINS)
	baseUrlToSearch.WriteString(constants.URL_PARAM_EQUALS)
	baseUrlToSearch.WriteString(universityName)
	baseUrlToSearch.WriteString(constants.URL_PARAM_AND)
	baseUrlToSearch.WriteString(constants.URL_PARAM_COUNTRY)
	baseUrlToSearch.WriteString(constants.URL_PARAM_EQUALS)

	return baseUrlToSearch.String()
}

func GenerateUrlCountries(countryCodes []string) string {
	baseUrlToSearch := strings.Builder{}
	baseUrlToSearch.WriteString(constants.COUNTRY_API_ALPHA_CODE)
	baseUrlToSearch.WriteString(strings.Join(countryCodes, ","))
	baseUrlToSearch.WriteString(constants.COUNTRY_API_CONSTRAINTS_AND)

	return baseUrlToSearch.String()
}
