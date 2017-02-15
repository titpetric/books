package api

import "net/http"
import "app/common"

// Health check API
type Health struct {
}

// Register Health check API endpoints
func (t Health) Register() {
	http.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		response := "OK"
		common.Respond(w, response, nil)
	})
}
