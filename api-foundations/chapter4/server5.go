package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func requestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Request: %s\n", r.URL.Path)
	fmt.Fprintf(w, "Hello world!")
}

func requestHelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Request: %s\n", r.URL.Path)

	vars := mux.Vars(r)
	name, ok := vars["name"]
	if ok && name != "" {
		fmt.Fprintf(w, "Hello %s!", name)
	} else {
		fmt.Fprintf(w, "Hello ... you.")
	}
}

func main() {
	fmt.Printf("Starting server on port :80\n")

	m := mux.NewRouter()

	hey := m.PathPrefix("/hey").Subrouter()

	hey.HandleFunc("/{name}/", requestHelloHandler)
	hey.HandleFunc("/{name}", requestHelloHandler)
	hey.HandleFunc("/", requestHelloHandler)
	// hey.HandleFunc("/", requestHelloHandler);

	http.Handle("/", m)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
