package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/pkg/errors"

	"github.com/titpetric/factory/resputil"
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

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	login := JWT{}.new()

	mux := chi.NewRouter()
	mux.Use(cors.Handler)
	mux.Use(middleware.Logger)
	mux.Use(login.Verifier())

	// Protected API endpoints
	mux.Group(func(mux chi.Router) {
		// Error out on invalid/empty JWT here
		mux.Use(login.Authenticator())
		{
			mux.Get("/time", requestTime)
			mux.Route("/say", func(mux chi.Router) {
				mux.Get("/{name}", requestSay)
				mux.Get("/", requestSay)
			})
		}
	})

	// Public API endpoints
	mux.Group(func(mux chi.Router) {
		// Print info about claim
		mux.Get("/info", func(w http.ResponseWriter, r *http.Request) {
			owner := login.Decode(r)
			resputil.JSON(w, owner, errors.New("Not logged in"))
		})
	})

	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		fmt.Println("ListenAndServe:", err)
	}
}
