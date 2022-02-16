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
		constants.COUNTRY_API + countryName + constants.COUNTRY_API_CONSTRAINTS))[0]
}

// getCountryBasedOnCode Gets the country based on the country code from the country api.
func getCountryBasedOnCode(countryCode string) model.CountryApi {
	//TODO: add constraint, currently not working..
	return jsonparser.DecodeCountryInfo(client.GetResponseFromWebPage(
		constants.COUNTRY_API_ALPHA_CODE + countryCode))[0]
}

// GetNeighbouringCountries takes a countryApi struct and returns a map containing countryApi instances
// of the neighbouring countries.
func GetNeighbouringCountries(country model.CountryApi) map[string]model.CountryApi {
	neighbouringCountriesAlphaCodes := country.BordersTo
	var countriesFullName = map[string]model.CountryApi{}
	countriesFullName[country.Name["common"].(string)] = country

	for _, borderingCountry := range neighbouringCountriesAlphaCodes {
		obtainedCountry := getCountryBasedOnCode(borderingCountry)
		countriesFullName[obtainedCountry.Name["common"].(string)] = obtainedCountry
	}

	return countriesFullName
}
