package model_logic

import (
	"assignment-1/client"
	"assignment-1/constants"
	"assignment-1/jsonparser"
	"assignment-1/model"
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

func GetUniversitiesBorderingTo(universityName string, searchCountry string) []model.University {
	var combinedUniversities []model.University
	url := constants.UNIVERSITY_API + "name=" + universityName + "&country="
	countries := GetNeighbouringCountries(GetCountry(searchCountry))

	for _, country := range countries {
		urlToSearch := url + country.Name["common"].(string)
		universities := jsonparser.DecodeUniInfo(client.GetResponseFromWebPage(urlToSearch))

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
func Combine(list []model.UniversityInfo) []model.University {
	var combinedUniversityList []model.University
	var countriesCache = map[string]model.CountryApi{}
	for _, obtainedUniversity := range list {
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
