package golf

import (
	"fmt"
)

func (l *Logger) genMsg(attrs map[string]interface{}, level int, msg string, va ...interface{}) *Message {
	newMsg := l.NewMessage()
	newMsg.Level = level
	if len(va) > 0 {
		newMsg.ShortMessage = fmt.Sprintf(msg, va...)
	} else {
		newMsg.ShortMessage = msg
	}
	newMsg.Attrs = attrs
	return newMsg
}

func (l *Logger) logMsg(attrs map[string]interface{}, level int, msg string, va ...interface{}) error {
	newMsg := l.genMsg(attrs, level, msg, va...)
	return l.client.QueueMsg(newMsg)
}

// Dbg logs message 'msg' at LEVEL_DBG level
func (l *Logger) Dbg(msg string) error {
	return l.logMsg(nil, LEVEL_DBG, msg)
}

// Log a message at LEVEL_DBG with 'format' populated with values from 'va'
func (l *Logger) Dbgf(format string, va ...interface{}) error {
	return l.logMsg(nil, LEVEL_DBG, format, va...)
}

// Log a message at LEVEL_DBG with 'format' populated with values from 'va'. The
// attributes from 'attrs' will be included with the message, overriding any that may
// be set at the Logger level
func (l *Logger) Dbgm(attrs map[string]interface{}, format string, va ...interface{}) error {
	return l.logMsg(attrs, LEVEL_DBG, format, va...)
}

// Log a message 'msg' at LEVEL_INFO level
func (l *Logger) Info(msg string) error {
	return l.logMsg(nil, LEVEL_INFO, msg)
}

// Log a message at LEVEL_INFO with 'format' populated with values from 'va'
func (l *Logger) Infof(format string, va ...interface{}) error {
	return l.logMsg(nil, LEVEL_INFO, format, va...)
}

// Log a message at LEVEL_INFO with 'format' populated with values from 'va'. The
// attributes from 'attrs' will be included with the message, overriding any that may
// be set at the Logger level
func (l *Logger) Infom(attrs map[string]interface{}, format string, va ...interface{}) error {
	return l.logMsg(attrs, LEVEL_INFO, format, va...)
}

// Log a message 'msg' at LEVEL_NOTICE level
func (l *Logger) Notice(msg string) error {
	return l.logMsg(nil, LEVEL_NOTICE, msg)
}

// Log a message at LEVEL_NOTICE with 'format' populated with values from 'va'
func (l *Logger) Noticef(format string, va ...interface{}) error {
	return l.logMsg(nil, LEVEL_NOTICE, format, va...)
}

// Log a message at LEVEL_NOTICE with 'format' populated with values from 'va'. The
// attributes from 'attrs' will be included with the message, overriding any that may
// be set at the Logger level
func (l *Logger) Noticem(attrs map[string]interface{}, format string, va ...interface{}) error {
	return l.logMsg(attrs, LEVEL_NOTICE, format, va...)
}

// Log a message 'msg' at LEVEL_WARN level
func (l *Logger) Warn(msg string) error {
	return l.logMsg(nil, LEVEL_WARN, msg)
}

// Log a message at LEVEL_WARN with 'format' populated with values from 'va'
func (l *Logger) Warnf(format string, va ...interface{}) error {
	return l.logMsg(nil, LEVEL_WARN, format, va...)
}

// Log a message at LEVEL_WARN with 'format' populated with values from 'va'. The
// attributes from 'attrs' will be included with the message, overriding any that may
// be set at the Logger level
func (l *Logger) Warnm(attrs map[string]interface{}, format string, va ...interface{}) error {
	return l.logMsg(attrs, LEVEL_WARN, format, va...)
}

// Log a message 'msg' at LEVEL_ERR level
func (l *Logger) Err(msg string) error {
	return l.logMsg(nil, LEVEL_ERR, msg)
}

// Log a message at LEVEL_ERR with 'format' populated with values from 'va'
func (l *Logger) Errf(format string, va ...interface{}) error {
	return l.logMsg(nil, LEVEL_ERR, format, va...)
}

// Log a message at LEVEL_ERR with 'format' populated with values from 'va'. The
// attributes from 'attrs' will be included with the message, overriding any that may
// be set at the Logger level
func (l *Logger) Errm(attrs map[string]interface{}, format string, va ...interface{}) error {
	return l.logMsg(attrs, LEVEL_ERR, format, va...)
}

// Log a message 'msg' at LEVEL_CRIT level
func (l *Logger) Crit(msg string) error {
	return l.logMsg(nil, LEVEL_CRIT, msg)
}

// Log a message at LEVEL_CRIT with 'format' populated with values from 'va'
func (l *Logger) Critf(format string, va ...interface{}) error {
	return l.logMsg(nil, LEVEL_CRIT, format, va...)
}

// Log a message at LEVEL_CRIT with 'format' populated with values from 'va'. The
// attributes from 'attrs' will be included with the message, overriding any that may
// be set at the Logger level
func (l *Logger) Critm(attrs map[string]interface{}, format string, va ...interface{}) error {
	return l.logMsg(attrs, LEVEL_CRIT, format, va...)
}

// Log a message 'msg' at LEVEL_ALERT level
func (l *Logger) Alert(msg string) error {
	return l.logMsg(nil, LEVEL_ALERT, msg)
}

// Log a message at LEVEL_ALERT with 'format' populated with values from 'va'
func (l *Logger) Alertf(format string, va ...interface{}) error {
	return l.logMsg(nil, LEVEL_ALERT, format, va...)
}

// Log a message at LEVEL_ALERT with 'format' populated with values from 'va'. The
// attributes from 'attrs' will be included with the message, overriding any that may
// be set at the Logger level
func (l *Logger) Alertm(attrs map[string]interface{}, format string, va ...interface{}) error {
	return l.logMsg(attrs, LEVEL_ALERT, format, va...)
}

// Log a message 'msg' at LEVEL_EMERG level
func (l *Logger) Emerg(msg string) error {
	return l.logMsg(nil, LEVEL_EMERG, msg)
}

// Log a message at LEVEL_EMERG with 'format' populated with values from 'va'
func (l *Logger) Emergf(format string, va ...interface{}) error {
	return l.logMsg(nil, LEVEL_EMERG, format, va...)
}

// Log a message at LEVEL_EMERG with 'format' populated with values from 'va'. The
// attributes from 'attrs' will be included with the message, overriding any that may
// be set at the Logger level
func (l *Logger) Emergm(attrs map[string]interface{}, format string, va ...interface{}) error {
	return l.logMsg(attrs, LEVEL_EMERG, format, va...)
}
