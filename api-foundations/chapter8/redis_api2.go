package main

import "foundations/api"
import "foundations/bootstrap"
import "fmt"
import "log"
import "encoding/json"
import "net/http"

var apiService api.Registry

func respondWithError(w http.ResponseWriter, err error) {
	response := map[string]string{}
	response["error"] = fmt.Sprintf("%s", err)
	response_json, _ := json.MarshalIndent(response, "", "\t")
	fmt.Fprintf(w, string(response_json[:]))
}
func respondWith(w http.ResponseWriter, response interface{}) {
	response_json, _ := json.MarshalIndent(response, "", "\t")
	fmt.Fprintf(w, string(response_json[:]))
}
func respond(w http.ResponseWriter, response interface{}, err error) {
	if err != nil {
		respondWithError(w, err)
		return
	}
	respondWith(w, response)
}

func main() {
	redisPool := bootstrap.RedigoPool()
	defer redisPool.Close()

	apiService = api.Registry{Name: "api"}
	http.HandleFunc("/getAll", func(w http.ResponseWriter, r *http.Request) {
		response, err := apiService.GetAll()
		respond(w, response, err)
	})
	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		aResponse, aErr := make(chan interface{}, 1), make(chan error, 1)
		key := r.FormValue("key")
		go func() {
			response, err := apiService.Get(key)
			aResponse <- response
			aErr <- err
		}()
		response := <-aResponse
		err := <-aErr
		respond(w, response, err)
	})
	http.HandleFunc("/set", func(w http.ResponseWriter, r *http.Request) {
		key := r.FormValue("key")
		value := r.FormValue("value")
		response, err := apiService.Set(key, value)
		respond(w, response, err)
	})

	fmt.Printf("Starting server on port :80\n")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
