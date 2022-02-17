package model_logic

import (
	"assignment-1/client"
	"assignment-1/constants"
	"assignment-1/jsonparser"
	"assignment-1/model"
	"strings"
)

// combineStructs combines the university info with the country info and returns the data struct.
func combineStructs(uni model.UniversityInfo, country model.CountryApi) model.University {
	return model.University{
		Name:      uni.Name,
		Country:   uni.Country,
		Isocode:   uni.Isocode,
		Webpages:  uni.Webpages,
		Languages: country.Languages,
		Map:       country.Maps["openStreetMaps"],
	}
}

// GetUniversitiesBorderingTo gets the
func GetUniversitiesBorderingTo(universityName string, searchCountry string) []model.University {
	var combinedUniversities []model.University
	baseUrlToSearch := strings.Builder{}
	baseUrlToSearch.WriteString(constants.UNIVERSITY_API)
	baseUrlToSearch.WriteString(constants.URL_PARAM_NAME)
	baseUrlToSearch.WriteString(constants.URL_PARAM_EQUALS)
	baseUrlToSearch.WriteString(universityName)
	baseUrlToSearch.WriteString(constants.URL_PARAM_AND)
	baseUrlToSearch.WriteString(constants.URL_PARAM_COUNTRY)
	baseUrlToSearch.WriteString(constants.URL_PARAM_EQUALS)

	countries := GetNeighbouringCountries(GetCountry(searchCountry))

	for _, country := range countries {
		urlToSearch := strings.Builder{}
		urlToSearch.WriteString(baseUrlToSearch.String())
		urlToSearch.WriteString(country.Name["common"].(string))

		universities := jsonparser.DecodeUniInfo(client.GetResponseFromWebPage(
			strings.ReplaceAll(urlToSearch.String(), " ", "%20")))

		for _, obtainedUniversity := range universities {
			combinedUniversities = append(combinedUniversities, combineStructs(obtainedUniversity,
				countries[obtainedUniversity.Country]))
		}
	}

	return combinedUniversities
}

// Combine takes a slice of university info and combines every element with its country info and appends it to
//a new slice. The combined slice is returned. The method caches the result from the countries api to minimize
//the number of requests.
func Combine(universities []model.UniversityInfo) []model.University {
	var combinedUniversityList []model.University
	var countriesCache = map[string]model.CountryApi{}

	for _, obtainedUniversity := range universities {
		var country model.CountryApi
		countryName := obtainedUniversity.Country

		if value, ok := countriesCache[countryName]; ok {
			country = value
		} else {
			countriesCache[countryName] = GetCountry(countryName)
			country = countriesCache[countryName]
		}

		combinedUniversityList = append(combinedUniversityList, combineStructs(
			obtainedUniversity, country))
	}
	return combinedUniversityList
}
