package model_logic

import (
	"assignment-1/client"
	"assignment-1/constants"
	"assignment-1/jsonparser"
	"assignment-1/model"
	"fmt"
)

// getCountry Gets the country based on the country code from the country api.
func getCountryBasedOnCode(countryCode string) model.CountryApi {
	return jsonparser.DecodeCountryInfo(client.GetResponseFromWebPage(
		constants.COUNTRY_API_ALPHA_CODE + countryCode))[0]
}

// GetNeighbouringCountries takes a country code and returns a slice with the common name of the country.
func GetNeighbouringCountries(country model.CountryApi) []string {
	neighbouringCountriesAlphaCodes := country.BordersTo
	var neighbouringCountriesFullName []string

	for _, borderingCountry := range neighbouringCountriesAlphaCodes {
		neighbouringCountriesFullName = append(neighbouringCountriesFullName, fmt.Sprintf("%v",
			getCountryBasedOnCode(borderingCountry).Name["common"]))
	}

	return neighbouringCountriesFullName
}
