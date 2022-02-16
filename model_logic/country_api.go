package model_logic

import (
	"assignment-1/client"
	"assignment-1/constants"
	"assignment-1/jsonparser"
	"assignment-1/model"
)

// GetCountry Gets the country based on the country name from the country api.
func GetCountry(countryName string) model.CountryApi {
	return jsonparser.DecodeCountryInfo(client.GetResponseFromWebPage(
		constants.COUNTRY_API + countryName))[0]
}

// getCountryBasedOnCode Gets the country based on the country code from the country api.
func getCountryBasedOnCode(countryCode string) model.CountryApi {
	return jsonparser.DecodeCountryInfo(client.GetResponseFromWebPage(
		constants.COUNTRY_API_ALPHA_CODE + countryCode))[0]
}

// GetNeighbouringCountries takes a countryApi struct and returns a map containing countryApi instances
// of the neighbouring countries.
func GetNeighbouringCountries(country model.CountryApi) map[string]model.CountryApi {
	neighbouringCountriesAlphaCodes := country.BordersTo
	var neighbouringCountriesFullName = map[string]model.CountryApi{}

	for _, borderingCountry := range neighbouringCountriesAlphaCodes {
		obtainedCountry := getCountryBasedOnCode(borderingCountry)
		neighbouringCountriesFullName[obtainedCountry.Name["common"].(string)] = obtainedCountry
	}

	return neighbouringCountriesFullName
}
