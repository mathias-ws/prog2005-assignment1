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

	http.HandleFunc(constants.UNISEARCH_LOCATION, handlers.Unisearch_handler)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
