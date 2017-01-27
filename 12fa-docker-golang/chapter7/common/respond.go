package common

import "fmt"
import "net/http"
import "encoding/json"

func RespondWithError(w http.ResponseWriter, err interface{}) {
	w.Header().Set("Content-Type", "application/json")
	response := struct {
		Error string `json:"error"`
	}{}
	response.Error = fmt.Sprintf("%s", err)
	response_json, _ := json.MarshalIndent(response, "", "\t")
	fmt.Fprintf(w, string(response_json[:]))
}

func RespondWith(w http.ResponseWriter, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	response_json, err := json.MarshalIndent(response, "", "\t")
	if err != nil {
		RespondWithError(w, err)
		return
	}
	fmt.Fprintf(w, string(response_json[:]))
}

func Respond(w http.ResponseWriter, response interface{}, err interface{}) {
	if err != nil {
		RespondWithError(w, err)
		return
	}
	RespondWith(w, response)
}
