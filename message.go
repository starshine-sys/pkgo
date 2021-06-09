package pkgo

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Message is a proxied message
type Message struct {
	Timestamp time.Time `json:"timestamp"`

	ID       Snowflake `json:"id"`
	Original Snowflake `json:"original"`
	Sender   Snowflake `json:"sender"`
	Channel  Snowflake `json:"channel"`

	System System `json:"system"`

	Member Member `json:"member"`
}

// Snowflake is a Discord snowflake
type Snowflake uint64

func (s Snowflake) String() string {
	return fmt.Sprintf("%d", uint64(s))
}

// ParseSnowflake parses a snowflake from a string
func ParseSnowflake(str string) (Snowflake, error) {
	s, err := strconv.ParseUint(str, 0, 0)
	if err != nil {
		return 0, err
	}
	return Snowflake(s), nil
}

// UnmarshalJSON ...
func (s *Snowflake) UnmarshalJSON(v []byte) error {
	if string(v) == "null" {
		*s = 0
		return nil
	}

	p, err := strconv.ParseUint(strings.Trim(string(v), `"`), 0, 0)
	if err != nil {
		return err
	}

	*s = Snowflake(p)
	return nil
}

// GetMessage gets a message by Discord snowflake
func (s *Session) GetMessage(id Snowflake) (m *Message, err error) {
	err = s.getEndpoint("/msg/"+id.String(), &m)
	if err != nil {
		switch err.(type) {
		case *ErrStatusNot200:
			return m, ErrMsgNotFound
		default:
			return m, err
		}
	}
	return m, err
}
