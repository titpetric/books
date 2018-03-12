package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

func requestTime(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The current time is: %s\n", time.Now())
}

func requestSay(w http.ResponseWriter, r *http.Request) {
	val := chi.URLParam(r, "name")
	if val != "" {
		fmt.Fprintf(w, "Hello %s!\n", val)
	} else {
		fmt.Fprintf(w, "Hello ... you.\n")
	}
}

func main() {
	fmt.Println("Starting server on port :3000")

	r := chi.NewRouter()
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)
	r.Use(middleware.Logger)

	r.Get("/time", requestTime)
	r.Route("/say", func(r chi.Router) {
		r.Get("/{name}", requestSay)
		r.Get("/", requestSay)
	})

	err := http.ListenAndServe(":3000", r)
	if err != nil {
		fmt.Println("ListenAndServe:", err)
	}
}
