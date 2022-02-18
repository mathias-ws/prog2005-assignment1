package model_logic

import (
	"assignment-1/client"
	"assignment-1/constants"
	"assignment-1/jsonparser"
	"assignment-1/model"
	"errors"
)

// GetCountry Gets the country based on the country name from the country api.
func GetCountry(countryName string) (model.CountryApi, error) {
	country := jsonparser.DecodeCountryInfo(client.GetResponseFromWebPage(
		constants.COUNTRY_API + countryName + constants.COUNTRY_API_CONSTRAINTS))

	if len(country) == 0 || country == nil {
		return model.CountryApi{}, errors.New("unable to retrieve country")
	}

	return country[0], nil
}

// getCountryBasedOnCode Gets the country based on the country code from the country api.
func getCountryBasedOnCode(countryCode string) model.CountryApi {
	return jsonparser.DecodeCountryInfo(client.GetResponseFromWebPage(
		constants.COUNTRY_API_ALPHA_CODE + countryCode + constants.COUNTRY_API_CONSTRAINTS))[0]
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
