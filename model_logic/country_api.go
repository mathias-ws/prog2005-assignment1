package model_logic

import (
	"assignment-1/client"
	"assignment-1/constants"
	"assignment-1/customErrors"
	"assignment-1/jsonparser"
	"assignment-1/model"
)

// GetCountry Gets the country based on the country name from the country api.
func GetCountry(countryName string) (model.CountryApi, error) {
	response, err := client.GetResponseFromWebPage(
		constants.COUNTRY_API + countryName + constants.COUNTRY_API_CONSTRAINTS)

	if err != nil {
		return model.CountryApi{}, err
	}

	country := jsonparser.DecodeCountryInfo(response)

	if len(country) == 0 || country == nil {
		return model.CountryApi{}, customErrors.GetUnableToGetCountryError()
	}

	return country[0], nil
}

// getCountryBasedOnCode Gets the country based on the country code from the country api.
func getCountryBasedOnCode(countryCode string) (model.CountryApi, error) {
	response, err := client.GetResponseFromWebPage(
		constants.COUNTRY_API_ALPHA_CODE + countryCode + constants.COUNTRY_API_CONSTRAINTS)

	if err != nil {
		return model.CountryApi{}, err
	}

	return jsonparser.DecodeCountryInfo(response)[0], nil
}

// GetNeighbouringCountries takes a countryApi struct and returns a map containing countryApi instances
// of the neighbouring countries.
func GetNeighbouringCountries(country model.CountryApi) (map[string]model.CountryApi, error) {
	neighbouringCountriesAlphaCodes := country.BordersTo
	var countriesFullName = map[string]model.CountryApi{}
	countriesFullName[country.Name["common"].(string)] = country

	for _, borderingCountry := range neighbouringCountriesAlphaCodes {
		obtainedCountry, err := getCountryBasedOnCode(borderingCountry)

		if err != nil {
			return nil, err
		}

		countriesFullName[obtainedCountry.Name["common"].(string)] = obtainedCountry
	}

	return countriesFullName, nil
}
