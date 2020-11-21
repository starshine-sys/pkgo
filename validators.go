package pkgo

import "regexp"

var colorRe, idRe, discordIDre *regexp.Regexp

func init() {
	colorRe = regexp.MustCompile("(i)^[\\dabcdef]{6}$")
	idRe = regexp.MustCompile("^[a-z]{5}$")
	discordIDre = regexp.MustCompile("^\\d{16,}$")
}
