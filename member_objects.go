package pkgo

import (
	"encoding/json"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Color holds the color for a member
type Color string

var colorRe = regexp.MustCompile("(i)^[\\dabcdef]{6}$")

// IsValid returns true if the color is valid for PK
func (c Color) IsValid() bool {
	color := strings.TrimPrefix(string(c), "#")
	return colorRe.MatchString(color)
}

// ToInt converts the color to an int64 value
func (c Color) ToInt() int64 {
	color, _ := strconv.ParseInt(string(c), 16, 0)
	return color
}

// ProxyTag is a single proxy tag for a member, in the format <prefix>text<suffix>
type ProxyTag struct {
	Prefix string `json:"prefix,omitempty"`
	Suffix string `json:"suffix,omitempty"`
}

// MarshalJSON marshals the ProxyTag to json
func (t ProxyTag) MarshalJSON() (b []byte, err error) {
	if t.Prefix == "" && t.Suffix == "" {
		return nil, ErrInvalidProxyTag
	}

	if len(t.String()) > 100 {
		return nil, ErrInvalidProxyTag
	}

	dat := struct {
		Prefix *string `json:"prefix"`
		Suffix *string `json:"suffix"`
	}{}

	if t.Prefix != "" {
		dat.Prefix = &t.Prefix
	}
	if t.Suffix != "" {
		dat.Suffix = &t.Suffix
	}

	return json.Marshal(dat)
}

// String returns a <prefix>text<suffix> formatted version of the proxy tag
func (t ProxyTag) String() string {
	return t.Prefix + "text" + t.Suffix
}

// Birthday is a member's birthday
type Birthday time.Time

// MarshalJSON ...
func (bd Birthday) MarshalJSON() (b []byte, err error) {
	if bd.Time().IsZero() {
		return []byte("null"), nil
	}

	b = []byte(`"`)
	b = time.Time(bd).AppendFormat(b, "2006-01-02")
	return append(b, `"`...), nil
}

// UnmarshalJSON ...
func (bd *Birthday) UnmarshalJSON(v []byte) error {
	if string(v) == "null" {
		*bd = Birthday(time.Time{})
		return nil
	}

	t, err := time.Parse("2006-01-02", strings.Trim(string(v), `"`))
	if err != nil {
		return err
	}

	*bd = Birthday(t)
	return nil
}

// Time returns bd as time.Time
func (bd Birthday) Time() time.Time {
	return time.Time(bd)
}

// ParseBirthday parses a birthday in yyyy-mm-dd or mm-dd format.
func ParseBirthday(in string) (bd Birthday, err error) {
	t, err := time.Parse("2006-01-02", in)
	if err == nil {
		return Birthday(t), nil
	}

	t, err = time.Parse("2006-01-02", "0004-"+in)
	if err == nil {
		return Birthday(t), nil
	}

	return
}

// Privacy is a system or member privacy field.
// Note: an empty Privacy is marshaled as null.
type Privacy string

// MarshalJSON ...
func (p Privacy) MarshalJSON() (b []byte, err error) {
	if p == "" {
		return []byte(`null`), nil
	}

	if p == "public" {
		return []byte(`"public"`), nil
	}

	if p == "private" {
		return []byte(`"private"`), nil
	}

	return nil, ErrPrivacyInvalid
}

// UnmarshalJSON ...
func (p *Privacy) UnmarshalJSON(v []byte) error {
	if string(v) == "null" {
		*p = ""
		return nil
	}

	*p = Privacy(strings.Trim(string(v), `"`))
	return nil
}
