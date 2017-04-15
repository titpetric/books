package golf

import (
	"time"
)

// Levels to be used when logging messages.  These match syslog levels as the
// GELF spec specifies.
const (
	LEVEL_EMERG  = iota // Emergency
	LEVEL_ALERT         // Alert
	LEVEL_CRIT          // Critical
	LEVEL_ERR           // Error
	LEVEL_WARN          // Warning
	LEVEL_NOTICE        // Notice
	LEVEL_INFO          // Informational
	LEVEL_DBG           // Debug
)

// A message to be serialized and sent to the GELF server
type Message struct {
	logger *Logger

	version      string                 // GELF version to serialize to
	Level        int                    // Log level for the message (see LEVEL_DBG, etc)
	Hostname     string                 // Hostname of the client
	Timestamp    *time.Time             // Timestamp for the message. Populated automatically if left nil
	ShortMessage string                 // Short log message
	FullMessage  string                 // Full message (optional). Can be used for things like stack traces.
	Attrs        map[string]interface{} // A list of attributes to add to the message
}

// Create a new message associated with a Logger.  When the message is sent, it
// will use the attributes associated with the Logger it was created from in addition
// to its own
func (l *Logger) NewMessage() *Message {
	msg := newMessage()
	msg.logger = l
	msg.Hostname = l.client.hostname
	return msg
}

func newMessage() *Message {
	return newMessageForVersion("1.1")
}
func newMessageForVersion(version string) *Message {
	msg := &Message{
		version: version,

		Attrs: make(map[string]interface{}, 0),
	}
	return msg
}
