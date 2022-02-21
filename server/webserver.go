package server

import (
	"assignment-1/constants"
	"assignment-1/model_logic"
	"assignment-1/server/handlers"
	"log"
	"net/http"
	"os"
)

func setHandlers() {
	http.HandleFunc(constants.UNISEARCH_LOCATION, handlers.UnisearchHandler)
	http.HandleFunc(constants.NEIGHBOUR_UNIS_LOCATION, handlers.NeighbourUnisHandler)
	http.HandleFunc(constants.DIAG_LOCATION, handlers.DiagHandler)
}

func StartWebServer() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Println("Port variable not set, using default: " + constants.PORT)
		port = constants.PORT
	}

	setHandlers()

	model_logic.SetStartTime()

	log.Println("Webserver started on port:", port)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
