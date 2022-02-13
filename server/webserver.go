package server

import (
	"assignment-1/constants"
	"assignment-1/server/handlers"
	"log"
	"net/http"
	"os"
)

func StartWebServer() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Println("Port variable not set, using default: " + constants.PORT)
		port = constants.PORT
	}

	http.HandleFunc(constants.UNISEARCH_LOCATION, handlers.UnisearchHandler)
	http.HandleFunc(constants.NEIGHBOUR_UNIS_LOCATION, handlers.NeighbourUnisHandler)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
