package model_logic

import (
	"assignment-1/constants"
	"assignment-1/custom_errors"
	"assignment-1/json_parser"
	"assignment-1/model"
	"assignment-1/web_client"
	"assignment-1/web_client/urlHandlingClient"
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

// GetCountryByAlphaTwoCode Gets the country based on the country alpha two code from the country api.
func GetCountryByAlphaTwoCode(countryAlphaTwoCode string) (model.CountryApi, error) {
	response, err := web_client.GetResponseFromWebPage(
		constants.COUNTRY_API_ALPHA_CODE + countryAlphaTwoCode + constants.COUNTRY_API_CONSTRAINTS_AND)

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
			if strings.ToLower(country[i].Name["common"].(string)) == strings.ToLower(countryAlphaTwoCode) {
				return country[i], nil
			}
		}
	}

	return country[0], nil
}

// GetNeighbouringCountries takes a countryApi struct and returns a map containing countryApi instances
// of the neighbouring countries.
// Gets the country based on the country code from the country api. It is sending in a list of country codes to return
// a list
func GetNeighbouringCountries(country model.CountryApi) ([]model.CountryApi, error) {
	countryCodes := country.BordersTo

	response, err := web_client.GetResponseFromWebPage(urlHandlingClient.GenerateUrlCountries(countryCodes))

	// Checks for errors in the fetching from the api.
	if err != nil {
		return nil, err
	}

	// Adding the countries to a list and adding the original country.
	borderingCountries := json_parser.DecodeCountryInfo(response)
	borderingCountries = append(borderingCountries, country)

	return borderingCountries, nil
}
