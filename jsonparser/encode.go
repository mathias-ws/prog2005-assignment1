package jsonparser

import (
	"encoding/json"
	"log"
	"net/http"
)

// Encode encodes some data into json and displays it on the website.
func Encode(w http.ResponseWriter, valueToEncode interface{}) error {
	w.Header().Add("Content-Type", "application/json")

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")

	if err := encoder.Encode(valueToEncode); err != nil {
		log.Fatal("Unable to encode data: ", err)
		return err
	}

	return nil
}
