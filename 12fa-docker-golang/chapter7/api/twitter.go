package api

import "time"
import "net/http"
import "encoding/json"
import "app/services"
import "github.com/garyburd/redigo/redis"

// Twitter post message API
type Twitter struct {
}

// TwitterMessage - Structure of a post
type TwitterMessage struct {
	Message string
	Time    time.Time
}

// Register Twitter API endpoints
func (t *Twitter) Register() {
	http.HandleFunc("/api/twitter/list", func(w http.ResponseWriter, r *http.Request) {
		response, err := t.List()
		services.Respond(w, response, err)
	})
	http.HandleFunc("/api/twitter/add", func(w http.ResponseWriter, r *http.Request) {
		message := t.Message(r.FormValue("message"))
		response := "OK"
		err := t.Store(message)
		services.Respond(w, response, err)
	})
}

// Message creates a new TwitterMessage
func (t *Twitter) Message(message string) *TwitterMessage {
	return &TwitterMessage{message, time.Now()}
}

// List TwitterMessages for pagination
func (t *Twitter) List() ([]*TwitterMessage, error) {
	results := make([]*TwitterMessage, 0)

	conn, err := services.GetRedis()
	if err != nil {
		return nil, err
	}

	messages, err := redis.Strings(conn.Do("LRANGE", "twitter", 0, 4))
	for _, message := range messages {
		messageJSON := &TwitterMessage{}
		err = json.Unmarshal([]byte(message), messageJSON)
		if err != nil {
			continue
		}
		results = append(results, messageJSON)
	}
	return results, nil
}

// Store a new TwitterMessage
func (t *Twitter) Store(message *TwitterMessage) error {
	conn, err := services.GetRedis()
	if err != nil {
		return err
	}
	messageJSON, _ := json.MarshalIndent(message, "", "\t")
	_, err = conn.Do("LPUSH", "twitter", string(messageJSON[:]))
	return err
}
