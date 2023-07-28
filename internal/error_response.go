package internal

import (
	"encoding/json"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, error interface{}) {
	jsonContent, err := json.Marshal(error)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(code)
	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonContent)
}
