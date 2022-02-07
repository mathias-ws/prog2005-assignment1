package server

import (
	"log"
	"net/http"
	"os"
)

func default_handler(w http.ResponseWriter, r *http.Request) {

}

func StartWebServer() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Println("Port variable not set, using default: " + PORT)
		port = PORT
	}

	http.HandleFunc(DEFAULT_LOCATION, default_handler)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
