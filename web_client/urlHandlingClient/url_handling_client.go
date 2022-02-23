package urlHandlingClient

import (
	"assignment-1/constants"
	"strings"
)

// GenerateBaseUrlForCountrySearch generates the bas urlHandlingClient that the GetUniversitiesBorderingTo function uses.
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
