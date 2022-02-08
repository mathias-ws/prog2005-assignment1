package jsonparser

import (
	"assignment-1/model"
	"encoding/json"
	"io"
)

func EncodeUni(w io.Writer, list []model.University) {
	json.NewEncoder(w).Encode(list)
}
