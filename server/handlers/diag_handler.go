package handlers

import (
	"assignment-1/jsonparser"
	"assignment-1/model_logic"
	"net/http"
)

// DiagHandler gets the diag info struct and encodes it.
func DiagHandler(w http.ResponseWriter, r *http.Request) {
	jsonparser.Encode(w, model_logic.GetDiagInfo())
}
