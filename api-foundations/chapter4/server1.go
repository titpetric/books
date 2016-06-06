package main

import (
	"fmt"
	"log"
	"net/http"
)

func requestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Request: %s\n", r.URL.Path)
	fmt.Fprintf(w, "Hello world!")
}

func main() {
	fmt.Printf("Starting server on port :80\n")
	http.HandleFunc("/", requestHandler)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
