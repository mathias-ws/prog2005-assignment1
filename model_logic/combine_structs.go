package model_logic

import (
	"assignment-1/constants"
	"assignment-1/custom_errors"
	"assignment-1/json_parser"
	"assignment-1/model"
	"assignment-1/web_client"
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

// generateBaseUrlForCountrySearch generates the bas url that the GetUniversitiesBorderingTo function uses.
func generateBaseUrlForCountrySearch(universityName string) string {
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

	// Goes through all the countries and searches for universities matching the conditions within the given country.
	for _, country := range countries {
		urlToSearch := strings.Builder{}
		urlToSearch.WriteString(generateBaseUrlForCountrySearch(universityName))
		urlToSearch.WriteString(country.Name["common"].(string))

		response, err := web_client.GetResponseFromWebPage(
			strings.ReplaceAll(urlToSearch.String(), " ", "%20"))

		if err != nil {
			return nil, err
		}

		universities := json_parser.DecodeUniInfo(response)

		// Goes through the all the universities in a given country, combines it with the country info and adds it to the list.
		for _, obtainedUniversity := range universities {
			combinedUniversities = append(combinedUniversities, combineStructs(obtainedUniversity,
				countries[obtainedUniversity.Country]))

			// Ends the loop if the limit is reached.
			if len(combinedUniversities) > limit-1 && limit != 0 {
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
		countryName := obtainedUniversity.Country

		// If already present in cache.
		if value, ok := countriesCache[countryName]; ok {
			country = value
		} else {
			// Gets a new country struct and adds it to the cache.
			if returnedCountry, err := GetCountry(countryName); err != nil {
				log.Println(err)
				return []model.University{}, err
			} else {
				countriesCache[countryName] = returnedCountry
			}

			country = countriesCache[countryName]
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
