package resputil

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"

	"github.com/pkg/errors"
)

type stackTracer interface {
	StackTrace() errors.StackTrace
}

type successMessage struct {
	Success struct {
		Message string `json:"message"`
	} `json:"success"`
}

type errorMessage struct {
	Error struct {
		Message string `json:"message"`
		Trace   string `json:"trace,omitempty"`
	} `json:"error"`
}

// Error returns a structured error for API responses
func Error(err error) errorMessage {
	response := errorMessage{}
	// add stack trace to the response if available
	terr, ok := errors.Cause(err).(stackTracer)
	if ok {
		st := terr.StackTrace()
		response.Error.Trace = fmt.Sprintf("%+v", st)
	}
	response.Error.Message = fmt.Sprintf("%s", err)
	return response
}

// Success returns a structured sucess message for API responses
func Success(success ...string) successMessage {
	response := successMessage{}
	response.Success.Message = "OK"
	if len(success) > 0 {
		response.Success.Message = success[0]
	}
	return response
}

// Debug request
func Debug(w http.ResponseWriter, r *http.Request) {
	output, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println("Error dumping request:", err)
		return
	}
	fmt.Println(string(output))
}

// JSON responds with the first non-nil payload, formats error messages
func JSON(w http.ResponseWriter, responses ...interface{}) {
	respond := func(payload interface{}) {
		var result []byte
		var err error
		encode := func(payload interface{}) ([]byte, error) {
			// json, err := json.Marshal(payload)
			return json.MarshalIndent(payload, "", "\t")
		}
		switch value := payload.(type) {
		case errorMessage:
			// main key is "error"
			result, err = encode(value)
		case successMessage:
			// main key is "success"
			result, err = encode(value)
		default:
			// main key is "response"
			result, err = encode(struct {
				Response interface{} `json:"response"`
			}{value})
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(result)
	}

	for _, response := range responses {
		switch value := response.(type) {
		case nil:
			continue
		case func() (interface{}, error):
			result, err := value()
			JSON(w, err, result)
		case func() error:
			err := value()
			if err == nil {
				continue
			}
			respond(Error(err))
		case error:
			respond(Error(value))
		case string:
			if value == "" {
				continue
			}
			respond(value)
		case bool:
			if !value {
				continue
			}
			respond(value)
		case errorMessage:
			respond(value)
		case successMessage:
			respond(value)
		default:
			respond(value)
		}
		// Exit on the first output...
		return
	}
	respond(false)
}
