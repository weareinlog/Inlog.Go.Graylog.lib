package graylog

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// MessageLevel struct message level
type MessageLevel struct {
	Level  Level         `json:"level"`
	Params []interface{} `json:"params"`
}

// Create message level
func (MessageLevel) Create(level Level, v ...interface{}) MessageLevel {
	return MessageLevel{
		Level:  level,
		Params: v,
	}
}

// ToJSON convert level to json
func (m MessageLevel) ToJSON() string {
	str, _ := json.Marshal(m)
	return string(str)
}

func jsonToMessageLevel(b []byte) (MessageLevel, error) {
	msg := MessageLevel{}
	err := json.Unmarshal(b, &msg)
	return msg, err
}

// Message represents the contents of the GELF message.  It is gzipped
// before sending.
type Message struct {
	Version  string                 `json:"version"`
	Host     string                 `json:"host"`
	Short    string                 `json:"short_message"`
	Full     string                 `json:"full_message,omitempty"`
	TimeUnix float64                `json:"timestamp"`
	Level    int32                  `json:"level,omitempty"`
	Facility string                 `json:"facility,omitempty"`
	Extra    map[string]interface{} `json:"-"`
	RawExtra json.RawMessage        `json:"-"`
}

// Level severity levels
type Level int

const (
	LOG_EMERG   Level = 0
	LOG_ALERT   Level = 1
	LOG_CRIT    Level = 2
	LOG_ERR     Level = 3
	LOG_WARNING Level = 4
	LOG_NOTICE  Level = 5
	LOG_INFO    Level = 6
	LOG_DEBUG   Level = 7
)

var levelDescription map[Level]interface{} = map[Level]interface{}{
	LOG_EMERG:   "EMERG",
	LOG_ALERT:   "ALERT",
	LOG_CRIT:    "CRIT",
	LOG_ERR:     "ERR",
	LOG_WARNING: "WARNING",
	LOG_NOTICE:  "NOTICE",
	LOG_INFO:    "INFO",
	LOG_DEBUG:   "DEBUG",
}

func (m *Message) MarshalJSONBuf(buf *bytes.Buffer) error {
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}
	// write up until the final }
	if _, err = buf.Write(b[:len(b)-1]); err != nil {
		return err
	}
	if len(m.Extra) > 0 {
		eb, err := json.Marshal(m.Extra)
		if err != nil {
			return err
		}
		// merge serialized message + serialized extra map
		if err = buf.WriteByte(','); err != nil {
			return err
		}
		// write serialized extra bytes, without enclosing quotes
		if _, err = buf.Write(eb[1 : len(eb)-1]); err != nil {
			return err
		}
	}

	if len(m.RawExtra) > 0 {
		if err := buf.WriteByte(','); err != nil {
			return err
		}

		// write serialized extra bytes, without enclosing quotes
		if _, err = buf.Write(m.RawExtra[1 : len(m.RawExtra)-1]); err != nil {
			return err
		}
	}

	// write final closing quotes
	return buf.WriteByte('}')
}

func (m *Message) UnmarshalJSON(data []byte) error {
	i := make(map[string]interface{}, 16)
	if err := json.Unmarshal(data, &i); err != nil {
		return err
	}
	for k, v := range i {
		if k[0] == '_' {
			if m.Extra == nil {
				m.Extra = make(map[string]interface{}, 1)
			}
			m.Extra[k] = v
			continue
		}

		ok := true
		switch k {
		case "version":
			m.Version, ok = v.(string)
		case "host":
			m.Host, ok = v.(string)
		case "short_message":
			m.Short, ok = v.(string)
		case "full_message":
			m.Full, ok = v.(string)
		case "timestamp":
			m.TimeUnix, ok = v.(float64)
		case "level":
			var level float64
			level, ok = v.(float64)
			m.Level = int32(level)
		case "facility":
			m.Facility, ok = v.(string)
		}

		if !ok {
			return fmt.Errorf("invalid type for field %s", k)
		}
	}
	return nil
}

func (m *Message) toBytes(buf *bytes.Buffer) (messageBytes []byte, err error) {
	if err = m.MarshalJSONBuf(buf); err != nil {
		return nil, err
	}
	messageBytes = buf.Bytes()
	return messageBytes, nil
}

func constructMessage(p []byte, hostname string, facility string, file string, line int, extra map[string]interface{}) (m *Message) {
	// remove trailing and leading whitespace
	p = bytes.TrimSpace(p)

	// If there are newlines in the message, use the first line
	// for the short message and set the full message to the
	// original input.  If the input has no newlines, stick the
	// whole thing in Short.

	if extra == nil {
		extra = make(map[string]interface{})
	}

	extra["_file"] = file
	extra["_line"] = line

	level := int32(LOG_INFO)
	extra["StringLeval"] = levelDescription[LOG_INFO]

	message := p
	index := strings.Index(string(p), "{")

	if index >= 0 {
		mensagem := string(p)[index:len(string(p))]

		msgLevel, err := jsonToMessageLevel([]byte(mensagem))
		if err == nil {
			level = int32(msgLevel.Level)
			extra["StringLeval"] = levelDescription[msgLevel.Level]

		}

		message, err = json.Marshal(msgLevel.Params)
	}

	full := message
	short := message
	if i := bytes.IndexRune(message, '\n'); i > 0 {
		short = message[:i]
		full = message
	}

	m = &Message{
		Version:  "1.1",
		Host:     hostname,
		Short:    string(short),
		Full:     string(full),
		TimeUnix: float64(time.Now().UnixNano()) / float64(time.Second),
		Level:    level, // info
		Facility: facility,
		Extra:    extra,
	}

	return m
}
