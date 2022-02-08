package handlers

import (
	"assignment-1/client"
	"assignment-1/constants"
	"assignment-1/jsonparser"
	"assignment-1/model"
	"net/http"
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

// getCountry Gets the country based on the country name from the country api.
func getCountry(countryName string) model.CountryApi {
	return jsonparser.DecodeCountryInfo(client.GetResponseFromWebPage(
		constants.COUNTRY_API + countryName))[0]
}

// combine takes a slice of university info and combines every element with its country info and appends it to
//a new slice. The combined slice is returned.
func combine(list []model.UniversityInfo) []model.University {
	var combinedUniversityList []model.University
	for _, obtainedUniversity := range list {
		combinedUniversityList = append(combinedUniversityList, combineStructs(
			obtainedUniversity, getCountry(obtainedUniversity.Country)))
	}
	return combinedUniversityList
}

// UnisearchHandler is the handler called from the web server.
func UnisearchHandler(w http.ResponseWriter, r *http.Request) {
	jsonparser.EncodeUni(w, combine(jsonparser.DecodeUniInfo(
		client.GetResponseFromWebPage("http://universities.hipolabs.com/search?name=Molde"))))
}
