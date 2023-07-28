package internal

import (
	"encoding/json"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	jsonContent, err := json.Marshal(payload)
	if err != nil {
		respondWithError(w, 400, err.Error())
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonContent)
	w.WriteHeader(code)
}
