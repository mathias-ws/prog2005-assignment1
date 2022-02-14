package jsonparser

import (
	"assignment-1/model"
	"encoding/json"
	"io"
	"log"
)

func EncodeUni(w io.Writer, list []model.University) {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")

	if err := encoder.Encode(list); err != nil {
		log.Fatal("Unable to encode data: ", err)
	}
}
