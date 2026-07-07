package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	type responseError struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, code, responseError{Error: msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Printf("failed to encode payload %s", err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(code)
	w.Write(jsonData)
}
