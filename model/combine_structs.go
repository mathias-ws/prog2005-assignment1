package model

import (
	"assignment-1/client"
	"assignment-1/constants"
	"assignment-1/jsonparser"
)

// combineStructs combines the university info with the country info and returns the data struct.
func combineStructs(uni UniversityInfo, country CountryApi) University {
	return University{
		Name:      uni.Name,
		Country:   uni.Country,
		Isocode:   uni.Isocode,
		Webpages:  uni.Webpages,
		Languages: country.Languages,
		Map:       country.Maps["openStreetMaps"],
	}
}

// getCountry Gets the country based on the country name from the country api.
func getCountry(countryName string) CountryApi {
	return jsonparser.DecodeCountryInfo(client.GetResponseFromWebPage(
		constants.COUNTRY_API + countryName))[0]
}

// Combine takes a slice of university info and combines every element with its country info and appends it to
//a new slice. The combined slice is returned. The method caches the result from the countries api to minimize
//the number of requests.
func Combine(list []UniversityInfo) []University {
	var combinedUniversityList []University
	var countriesCache = map[string]CountryApi{}
	for _, obtainedUniversity := range list {
		var country CountryApi
		countryName := obtainedUniversity.Country

		if value, ok := countriesCache[countryName]; ok {
			country = value
		} else {
			countriesCache[countryName] = getCountry(countryName)
			country = countriesCache[countryName]
		}

		combinedUniversityList = append(combinedUniversityList, combineStructs(
			obtainedUniversity, country))
	}
	return combinedUniversityList
}
