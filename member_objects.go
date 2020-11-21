package pkgo

import (
	"regexp"
	"strconv"
	"strings"
)

var colorRe *regexp.Regexp

func init() {
	colorRe = regexp.MustCompile("(i)[\\dabcdef]{6}")
}

// Color holds the color for a member
type Color string

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

// String returns a <prefix>text<suffix> formatted version of the proxy tag
func (p *ProxyTag) String() string {
	return p.Prefix + "text" + p.Suffix
}
