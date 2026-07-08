package main

import "strings"

var badWords = map[string]struct{}{
	"kerfuffle": {},
	"sharbert":  {},
	"fornax":    {},
}

func filterProfaneWords(chirp string) string {

	words := strings.Split(chirp, " ")
	for i := 0; i < len(words); i++ {
		word := strings.ToLower(words[i])
		if _, ok := badWords[word]; ok {
			words[i] = "****"
		}
	}
	cleaned := strings.Join(words, " ")
	return cleaned
}
