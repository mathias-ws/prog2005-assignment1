package model_logic

import (
	"assignment-1/constants"
	"assignment-1/custom_errors"
	"assignment-1/json_parser"
	"assignment-1/model"
	"assignment-1/web_client"
	"strings"
)

// GetCountry Gets the country based on the country name from the country api.
func GetCountry(countryName string) (model.CountryApi, error) {
	response, err := web_client.GetResponseFromWebPage(
		constants.COUNTRY_API + countryName + constants.COUNTRY_API_CONSTRAINTS)

	// Checks for errors in the fetching from the api.
	if err != nil {
		return model.CountryApi{}, err
	}

	country := json_parser.DecodeCountryInfo(response)

	// Returns an error if no structs are returned.
	if len(country) == 0 || country == nil {
		return model.CountryApi{}, custom_errors.GetUnableToGetCountryError()
	} else if len(country) > 1 {
		// Checking to find the correct country when the api returns multiple countries.
		for i := range country {
			if strings.ToLower(country[i].Name["common"].(string)) == strings.ToLower(countryName) {
				return country[i], nil
			}
		}
	}

	return country[0], nil
}

// getCountryBasedOnAlphaCode Gets the country based on the country code from the country api.
func getCountryBasedOnAlphaCode(countryCode string) (model.CountryApi, error) {
	response, err := web_client.GetResponseFromWebPage(
		constants.COUNTRY_API_ALPHA_CODE + countryCode + constants.COUNTRY_API_CONSTRAINTS)

	// Checks for errors in the fetching from the api.
	if err != nil {
		return model.CountryApi{}, err
	}

	return json_parser.DecodeCountryInfo(response)[0], nil
}

// GetNeighbouringCountries takes a countryApi struct and returns a map containing countryApi instances
// of the neighbouring countries.
func GetNeighbouringCountries(country model.CountryApi) (map[string]model.CountryApi, error) {
	neighbouringCountriesAlphaCodes := country.BordersTo
	var countriesFullName = map[string]model.CountryApi{}

	// Adds the current country into the map.
	countriesFullName[country.Name["common"].(string)] = country

	// Goes through all the bordering country codes and gets the country api struct, and adds it to the map.
	for _, borderingCountry := range neighbouringCountriesAlphaCodes {
		obtainedCountry, err := getCountryBasedOnAlphaCode(borderingCountry)

		// Checks for errors in the fetching from the api.
		if err != nil {
			return nil, err
		}

		countriesFullName[obtainedCountry.Name["common"].(string)] = obtainedCountry
	}

	return countriesFullName, nil
}
