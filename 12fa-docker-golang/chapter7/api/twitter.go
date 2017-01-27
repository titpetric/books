package api

import "time"
import "net/http"
import "encoding/json"
import "app/common"
import "github.com/garyburd/redigo/redis"

type Twitter struct {
}

type TwitterMessage struct {
	Message string
	Time    time.Time
}

func (t Twitter) Register() {
	http.HandleFunc("/api/twitter/list", func(w http.ResponseWriter, r *http.Request) {
		response, err := t.List()
		common.Respond(w, response, err)
	})
	http.HandleFunc("/api/twitter/add", func(w http.ResponseWriter, r *http.Request) {
		message := t.Message(r.FormValue("message"))
		response := "OK"
		err := t.Add(message)
		common.Respond(w, response, err)
	})
}

func (t Twitter) Message(message string) *TwitterMessage {
	return &TwitterMessage{message, time.Now()}
}

func (t Twitter) List() ([]*TwitterMessage, error) {
	results := make([]*TwitterMessage, 0)

	conn, err := common.GetRedis()
	if err != nil {
		return nil, err
	}

	messages, err := redis.Strings(conn.Do("LRANGE", "twitter", 0, 4))
	for _, message := range messages {
		message_json := &TwitterMessage{}
		err = json.Unmarshal([]byte(message), message_json)
		if err != nil {
			continue
		}
		results = append(results, message_json)
	}
	return results, nil
}

func (t Twitter) Add(message *TwitterMessage) error {
	conn, err := common.GetRedis()
	if err != nil {
		return err
	}
	message_json, _ := json.MarshalIndent(message, "", "\t")
	_, err = conn.Do("LPUSH", "twitter", string(message_json[:]))
	return err
}
