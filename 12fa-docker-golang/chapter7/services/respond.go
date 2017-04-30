package services

import "fmt"
import "net/http"
import "encoding/json"

func respondWithError(w http.ResponseWriter, err interface{}) {
	w.Header().Set("Content-Type", "application/json")
	response := struct {
		Error string `json:"error"`
	}{}
	response.Error = fmt.Sprintf("%s", err)
	responseJSON, _ := json.MarshalIndent(response, "", "\t")
	fmt.Fprintf(w, string(responseJSON[:]))
}

func respondWith(w http.ResponseWriter, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	responseJSON, err := json.MarshalIndent(response, "", "\t")
	if err != nil {
		respondWithError(w, err)
		return
	}
	fmt.Fprintf(w, string(responseJSON[:]))
}

// Respond with JSON payload including formatted error response
func Respond(w http.ResponseWriter, response interface{}, err interface{}) {
	if err != nil {
		respondWithError(w, err)
		return
	}
	respondWith(w, response)
}
