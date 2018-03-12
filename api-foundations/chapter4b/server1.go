package main

import (
	"fmt"
	"net/http"
	"time"
)

func requestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The current time is: %s\n", time.Now())
}

func main() {
	fmt.Println("Starting server on port :3000")
	http.HandleFunc("/", requestHandler)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("ListenAndServe:", err)
	}
}
