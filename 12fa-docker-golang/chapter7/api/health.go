package api

import "net/http"
import "app/common"

type Health struct {
}

func (t Health) Register() {
	http.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		response := "OK"
		common.Respond(w, response, nil)
	})
}
