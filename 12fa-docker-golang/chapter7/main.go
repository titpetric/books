package main

import (
	"app/api"
	"flag"
	"log"
	"net/http"
)

var apiTwitter api.Twitter

var port = flag.String("port", "8080", "Port for server")

func main() {
	flag.Parse()

	apiTwitter.Register()

	http.Handle("/", http.FileServer(http.Dir("./public_html")))

	log.Printf("Ready and listening on port " + *port + "\n")
	if err := http.ListenAndServe(":"+*port, nil); err != nil {
		panic(err)
	}
}
