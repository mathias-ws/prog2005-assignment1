package model_logic

import (
	"assignment-1/custom_errors"
	"assignment-1/json_parser"
	"assignment-1/model"
	"assignment-1/web_client"
	"assignment-1/web_server/url"
	"log"
	"strings"
)

// combineStructs combines the university info with the country info and returns the data struct.
func combineStructs(uni model.UniversityInfo, country model.CountryApi) model.University {
	return model.University{
		Name:      uni.Name,
		Country:   uni.Country,
		IsoCode:   uni.IsoCode,
		Webpages:  uni.Webpages,
		Languages: country.Languages,
		Map:       country.Maps["openStreetMaps"],
	}
}

// GetUniversitiesBorderingTo gets the universities based on name and country, including the neighbouring countries.
// It returns a slice of universities.
func GetUniversitiesBorderingTo(universityName string, searchCountry string, limit int) ([]model.University, error) {
	var combinedUniversities []model.University

	// Gets the first country.
	countrySearched, err := GetCountry(searchCountry)

	// Enters the if when the country does not exist in the country api.
	if err != nil {
		return []model.University{}, err
	}

	countries, err := GetNeighbouringCountries(countrySearched)

	// Enters when there has been an issue getting the slice of countries.
	if err != nil {
		return nil, err
	}

	response, err := web_client.GetResponseFromWebPage(
		strings.ReplaceAll(url.GenerateBaseUrlForCountrySearch(universityName), " ", "%20"))

	// Enters when there has been an issue getting the slice of countries.
	if err != nil {
		return nil, err
	}

	universities := json_parser.DecodeUniInfo(response)

	// Goes through the all the universities in a given country, combines it with the country info and adds it to the list.
	for _, obtainedUniversity := range universities {

		// Checks if the university is in the given country and adds it to the list.
		for _, obtainedCountry := range countries {
			if obtainedUniversity.IsoCode == obtainedCountry.Cca2 {
				combinedUniversities = append(combinedUniversities, combineStructs(obtainedUniversity,
					obtainedCountry))
				break
			}
		}

		// Ends the loop if the limit is reached.
		if len(combinedUniversities) > limit-1 && limit != 0 {
			break
		}
	}

	// Checks if any universities were found.
	if len(combinedUniversities) == 0 {
		return nil, custom_errors.GetNoContentFoundError()
	}

	return combinedUniversities, nil
}

// Combine takes a slice of university info and combines every element with its country info and appends it to
// a new slice. The combined slice is returned. The method caches the result from the countries api to minimize
// the number of requests.
func Combine(universities []model.UniversityInfo) ([]model.University, error) {
	var combinedUniversityList []model.University

	// Cache only lives for the duration of the function.
	var countriesCache = map[string]model.CountryApi{}

	// Goes through all the universities and combines them with the country info.
	for _, obtainedUniversity := range universities {
		var country model.CountryApi
		countryCc2Code := obtainedUniversity.IsoCode

		// If already present in cache.
		if value, ok := countriesCache[countryCc2Code]; ok {
			country = value
		} else {
			// Gets a new country struct and adds it to the cache.
			if returnedCountry, err := GetCountryByAlphaTwoCode(countryCc2Code); err != nil {
				log.Println(err)
				return []model.University{}, err
			} else {
				countriesCache[countryCc2Code] = returnedCountry
			}

			country = countriesCache[countryCc2Code]
		}

		// Combines the structs and adds them to the list.
		combinedUniversityList = append(combinedUniversityList, combineStructs(
			obtainedUniversity, country))
	}

	// Checks if no universities are found.
	if len(combinedUniversityList) == 0 {
		return []model.University{}, custom_errors.GetNoContentFoundError()
	}

	return combinedUniversityList, nil
}
