package golf

// A Logger is a set of attributes associated with a Client. When a message is
// sent with the Logger, the attributes from that Logger will be added to the
// message.
type Logger struct {
	client *Client

	attrs map[string]interface{}
}

func newLogger() *Logger {
	log := &Logger{
		attrs: make(map[string]interface{}, 0),
	}
	return log
}

// Create a new Logger associated with the Client.  Any messages logged with
// this Logger (and any Logger cloned from this) will be sent to Client.
func (c *Client) NewLogger() (*Logger, error) {
	log := newLogger()
	log.client = c

	return log, nil
}

// Create a new Logger with a shallow copy of the original Logger's attributes
func (l *Logger) Clone() *Logger {
	newLogger, _ := l.client.NewLogger()

	for attrKey, attrVal := range l.attrs {
		newLogger.SetAttr(attrKey, attrVal)
	}

	return newLogger
}

// Retrieve the current value of the Logger level attribute named 'name'.
// Returns nil if the attribute was not found
func (l *Logger) Attr(name string) interface{} {
	val, ok := l.attrs[name]
	if !ok {
		return nil
	}
	return val
}

// Set an attribute named 'name' to the value 'val' at the Logger
// level of attributes
func (l *Logger) SetAttr(name string, val interface{}) {
	l.attrs[name] = val
}

// Remove the attribute named 'name' from the Logger level of attributes
func (l *Logger) RemAttr(name string) {
	delete(l.attrs, name)
}
