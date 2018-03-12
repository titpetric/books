package main

import (
	"fmt"
	"net/http"
	"time"
)

func requestTime(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The current time is: %s\n", time.Now())
}

func requestSay(w http.ResponseWriter, r *http.Request) {
	val := r.FormValue("name")
	if val != "" {
		fmt.Fprintf(w, "Hello %s!\n", val)
	} else {
		fmt.Fprintf(w, "Hello ... you.\n")
	}
}

func main() {
	fmt.Println("Starting server on port :3000")

	http.HandleFunc("/time", requestTime)
	http.HandleFunc("/say", requestSay)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("ListenAndServe:", err)
	}
}
