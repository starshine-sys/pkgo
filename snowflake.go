package pkgo

import (
	"fmt"
	"strconv"
	"strings"
)

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
