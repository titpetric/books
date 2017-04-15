package golf

import (
	"encoding/json"
	"fmt"
)

// Workaround for json encoding 64 bit floats.  When using
// a normal float it's marshalled in scientific notation instead
// of decimal notation
type jsonFloat struct {
	val float64
}

func newJsonFloat(val float64) *jsonFloat {
	return &jsonFloat{val: val}
}
func (jf *jsonFloat) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%0f", jf.val)), nil
}

func generateMsgJson(msg *Message) (string, error) {
	obj := make(map[string]interface{}, 0)

	obj["version"] = msg.version
	obj["host"] = msg.Hostname
	obj["level"] = msg.Level

	obj["short_message"] = msg.ShortMessage
	if len(msg.FullMessage) > 0 {
		obj["full_message"] = msg.FullMessage
	}

	ts := float64(msg.Timestamp.UnixNano()) * float64(0.000000001)
	obj["timestamp"] = newJsonFloat(ts)

	// First add all the logger level attrs if it exists
	if msg.logger != nil {
		for attrName, attrVal := range msg.logger.attrs {
			obj[fmt.Sprintf("_%v", attrName)] = attrVal
		}
	}

	// Next add all the message level attrs. Those override
	// logger level attrs
	for attrName, attrVal := range msg.Attrs {
		obj[fmt.Sprintf("_%v", attrName)] = attrVal
	}

	data, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
