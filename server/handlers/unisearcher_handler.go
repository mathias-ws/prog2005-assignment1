package handlers

import (
	"assignment-1/client"
	"assignment-1/constants"
	"assignment-1/jsonparser"
	"assignment-1/model"
	"net/http"
)

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

func getCountry(countryName string) model.CountryApi {
	return jsonparser.DecodeCountryInfo(client.GetResponseFromWebPage(
		constants.COUNTRY_API + countryName))[0]
}

func combine(list []model.UniversityInfo) []model.University {
	var combinedUniversityList []model.University
	for _, obtainedUniversity := range list {
		combinedUniversityList = append(combinedUniversityList, combineStructs(
			obtainedUniversity, getCountry(obtainedUniversity.Country)))
	}
	return combinedUniversityList
}

func Unisearch_handler(w http.ResponseWriter, r *http.Request) {
	jsonparser.EncodeUni(w, combine(jsonparser.DecodeUniInfo(
		client.GetResponseFromWebPage("http://universities.hipolabs.com/search?name=Molde"))))
}
