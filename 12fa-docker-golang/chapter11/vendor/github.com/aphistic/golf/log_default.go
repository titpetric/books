package golf

import (
	"errors"
	"fmt"
)

var defaultLogger *Logger

// Set the default Logger. Any calls to log messages not associated with a
// specific Logger (such as calling l.Dbg()) will use the default logger.
func DefaultLogger(l *Logger) {
	defaultLogger = l
}

func genDefaultMsg(attrs map[string]interface{}, level int, msg string, va ...interface{}) *Message {
	newMsg := defaultLogger.NewMessage()
	newMsg.Level = level
	if len(va) > 0 {
		newMsg.ShortMessage = fmt.Sprintf(msg, va...)
	} else {
		newMsg.ShortMessage = msg
	}
	newMsg.Attrs = attrs
	return newMsg
}

func logDefaultMsg(attrs map[string]interface{}, level int, msg string, va ...interface{}) error {
	if defaultLogger == nil {
		return errors.New("A default logger is not set.")
	}

	newMsg := genDefaultMsg(attrs, level, msg, va...)
	return defaultLogger.client.QueueMsg(newMsg)
}

// Log a message 'msg' at LEVEL_DBG level on the default logger
func Dbg(msg string) error {
	return logDefaultMsg(nil, LEVEL_DBG, msg)
}

// Log a message at LEVEL_DBG with 'format' populated with values from 'va' on the default logger
func Dbgf(format string, va ...interface{}) error {
	return logDefaultMsg(nil, LEVEL_DBG, format, va...)
}

// Log a message at LEVEL_DBG with 'format' populated with values from 'va' on the default logger.
// The attributes from 'attrs' will be included with the message, overriding any that may
// be set at the Logger level
func Dbgm(attrs map[string]interface{}, format string, va ...interface{}) error {
	return logDefaultMsg(attrs, LEVEL_DBG, format, va...)
}

// Log a message 'msg' at LEVEL_INFO level on the default logger
func Info(msg string) error {
	return logDefaultMsg(nil, LEVEL_INFO, msg)
}

// Log a message at LEVEL_INFO with 'format' populated with values from 'va' on the default logger
func Infof(format string, va ...interface{}) error {
	return logDefaultMsg(nil, LEVEL_INFO, format, va...)
}

// Log a message at LEVEL_INFO with 'format' populated with values from 'va'  on the default logger.
// The attributes from 'attrs' will be included with the message, overriding any that may
// be set at the Logger level
func Infom(attrs map[string]interface{}, format string, va ...interface{}) error {
	return logDefaultMsg(attrs, LEVEL_INFO, format, va...)
}

// Log a message 'msg' at LEVEL_NOTICE level on the default logger
func Notice(msg string) error {
	return logDefaultMsg(nil, LEVEL_NOTICE, msg)
}

// Log a message at LEVEL_NOTICE with 'format' populated with values from 'va' on the default logger
func Noticef(format string, va ...interface{}) error {
	return logDefaultMsg(nil, LEVEL_NOTICE, format, va...)
}

// Log a message at LEVEL_NOTICE with 'format' populated with values from 'va' on the default logger.
// The attributes from 'attrs' will be included with the message, overriding any that may
// be set at the Logger level
func Noticem(attrs map[string]interface{}, format string, va ...interface{}) error {
	return logDefaultMsg(attrs, LEVEL_NOTICE, format, va...)
}

// Log a message 'msg' at LEVEL_WARN level on the default logger
func Warn(msg string) error {
	return logDefaultMsg(nil, LEVEL_WARN, msg)
}

// Log a message at LEVEL_WARN with 'format' populated with values from 'va' on the default logger
func Warnf(format string, va ...interface{}) error {
	return logDefaultMsg(nil, LEVEL_WARN, format, va...)
}

// Log a message at LEVEL_WARN with 'format' populated with values from 'va' on the default logger.
// The attributes from 'attrs' will be included with the message, overriding any that may
// be set at the Logger level
func Warnm(attrs map[string]interface{}, format string, va ...interface{}) error {
	return logDefaultMsg(attrs, LEVEL_WARN, format, va...)
}

// Log a message 'msg' at LEVEL_ERR level on the default logger on the default logger
func Err(msg string) error {
	return logDefaultMsg(nil, LEVEL_ERR, msg)
}

// Log a message at LEVEL_ERR with 'format' populated with values from 'va' on the default logger
func Errf(format string, va ...interface{}) error {
	return logDefaultMsg(nil, LEVEL_ERR, format, va...)
}

// Log a message at LEVEL_ERR with 'format' populated with values from 'va' on the default logger.
// The attributes from 'attrs' will be included with the message, overriding any that may
// be set at the Logger level
func Errm(attrs map[string]interface{}, format string, va ...interface{}) error {
	return logDefaultMsg(attrs, LEVEL_ERR, format, va...)
}

// Log a message 'msg' at LEVEL_CRIT level on the default logger
func Crit(msg string) error {
	return logDefaultMsg(nil, LEVEL_CRIT, msg)
}

// Log a message at LEVEL_CRIT with 'format' populated with values from 'va' on the default logger
func Critf(format string, va ...interface{}) error {
	return logDefaultMsg(nil, LEVEL_CRIT, format, va...)
}

// Log a message at LEVEL_CRIT with 'format' populated with values from 'va' on the default logger.
// The attributes from 'attrs' will be included with the message, overriding any that may
// be set at the Logger level
func Critm(attrs map[string]interface{}, format string, va ...interface{}) error {
	return logDefaultMsg(attrs, LEVEL_CRIT, format, va...)
}

// Log a message 'msg' at LEVEL_ALERT level on the default logger
func Alert(msg string) error {
	return logDefaultMsg(nil, LEVEL_ALERT, msg)
}

// Log a message at LEVEL_ALERT with 'format' populated with values from 'va' on the default logger
func Alertf(format string, va ...interface{}) error {
	return logDefaultMsg(nil, LEVEL_ALERT, format, va...)
}

// Log a message at LEVEL_ALERT with 'format' populated with values from 'va' on the default logger.
// The attributes from 'attrs' will be included with the message, overriding any that may
// be set at the Logger level
func Alertm(attrs map[string]interface{}, format string, va ...interface{}) error {
	return logDefaultMsg(attrs, LEVEL_ALERT, format, va...)
}

// Log a message 'msg' at LEVEL_EMERG level on the default logger
func Emerg(msg string) error {
	return logDefaultMsg(nil, LEVEL_EMERG, msg)
}

// Log a message at LEVEL_EMERG with 'format' populated with values from 'va' on the default logger
func Emergf(format string, va ...interface{}) error {
	return logDefaultMsg(nil, LEVEL_EMERG, format, va...)
}

// Log a message at LEVEL_EMERG with 'format' populated with values from 'va' on the default logger.
// The attributes from 'attrs' will be included with the message, overriding any that may
// be set at the Logger level
func Emergm(attrs map[string]interface{}, format string, va ...interface{}) error {
	return logDefaultMsg(attrs, LEVEL_EMERG, format, va...)
}
