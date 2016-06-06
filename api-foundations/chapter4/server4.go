package main

import (
	"fmt"
	"github.com/bmizerany/pat"
	"log"
	"net/http"
)

func requestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Request: %s\n", r.URL.Path)
	fmt.Fprintf(w, "Hello world!")
}

func requestHelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Request: %s\n", r.URL.Path)
	val := r.URL.Query().Get(":name")
	if val != "" {
		fmt.Fprintf(w, "Hello %s!", val)
	} else {
		fmt.Fprintf(w, "Hello ... you.")
	}
}

func main() {
	fmt.Printf("Starting server on port :80\n")

	m := pat.New()
	m.Get("/hey/:name/", http.HandlerFunc(requestHelloHandler))
	m.Get("/hey/", http.HandlerFunc(requestHelloHandler))
	m.Get("/", http.HandlerFunc(requestHandler))

	http.Handle("/", m)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
