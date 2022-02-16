package url

import (
	"assignment-1/constants"
	"net/url"
)

// GetNameAndCountry Gets the search string from the user containing uni name and country.
func GetNameAndCountry(url *url.URL) (string, string) {
	searchStrings := url.Query()
	return searchStrings["name"][0], searchStrings["country"][0]
}

// GenerateUniversitySearchString generates a search string for the university api based on the user inputted url.
func GenerateUniversitySearchString(url *url.URL) string {
	searchStrings := url.Query()

	urlToSearch := constants.UNIVERSITY_API

	if searchStrings.Has("name") {
		urlToSearch += "name=" + searchStrings["name"][0]
	}
	if searchStrings.Has("country") {
		urlToSearch += "&country=" + searchStrings["country"][0]
	}

	return urlToSearch
}
