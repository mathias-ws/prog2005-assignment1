package handlers

import (
	"assignment-1/client"
	"assignment-1/constants"
	"assignment-1/jsonparser"
	"assignment-1/model"
	"net/http"
	"net/url"
)

// urlHandler Gets the search string from the user
func urlHandler(url *url.URL) string {
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

// UnisearchHandler is the handler called from the web server.
func UnisearchHandler(w http.ResponseWriter, r *http.Request) {
	jsonparser.EncodeUni(w, model.Combine(jsonparser.DecodeUniInfo(
		client.GetResponseFromWebPage(urlHandler(r.URL)))))
}
