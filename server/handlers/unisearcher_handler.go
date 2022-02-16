package handlers

import (
	"assignment-1/client"
	"assignment-1/jsonparser"
	"assignment-1/model_logic"
	"assignment-1/server/url"
	"net/http"
)

// UnisearchHandler is the handler called from the web server.
func UnisearchHandler(w http.ResponseWriter, r *http.Request) {
	jsonparser.Encode(w, model_logic.Combine(jsonparser.DecodeUniInfo(
		client.GetResponseFromWebPage(url.GenerateUniversitySearchString(r.URL)))))
}
