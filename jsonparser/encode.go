package jsonparser

import (
	"assignment-1/model"
	"encoding/json"
	"io"
)

func EncodeUni(w io.Writer, list []model.University) {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")
	encoder.Encode(list)
}
