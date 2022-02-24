package web_server

import (
	"assignment-1/constants"
	"assignment-1/model_logic"
	"assignment-1/web_server/handlers"
	"log"
	"net/http"
	"os"
)

// setHandlers sets all the web handlers that the server has.
func setHandlers() {
	http.HandleFunc(constants.UNISEARCH_LOCATION, handlers.UnisearchHandler)
	http.HandleFunc(constants.NEIGHBOUR_UNIS_LOCATION, handlers.NeighbourUnisHandler)
	http.HandleFunc(constants.DIAG_LOCATION, handlers.DiagHandler)
	http.HandleFunc(constants.DEFAULT_LOCATION, handlers.DefaultHandler)
}

// StartWebServer starts the webserver for the api.
func StartWebServer() {
	port := os.Getenv("PORT")

	// Checks if the port env variable is set, if not it will set the default port.
	if port == "" {
		log.Println("Port variable not set, using default: " + constants.PORT)
		port = constants.PORT
	}

	setHandlers()

	model_logic.SetStartTime()

	log.Println("Webserver started on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
