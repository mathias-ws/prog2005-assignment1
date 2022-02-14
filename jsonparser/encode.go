package jsonparser

import (
	"encoding/json"
	"io"
	"log"
)

func Encode(w io.Writer, valueToEncode interface{}) {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")

	if err := encoder.Encode(valueToEncode); err != nil {
		log.Fatal("Unable to encode data: ", err)
	}
}
