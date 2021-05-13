package str

import (
	"regexp"
	"strings"
)

var matchWordBackward = regexp.MustCompile("(.+?)([A-Z])([^A-Z])")
var matchWordForward = regexp.MustCompile("([^_A-Z])([A-Z])")

func splitInWords(str string) []string {
	tmp := matchWordBackward.ReplaceAllString(str, "${1}_${2}${3}")
	tmp = matchWordForward.ReplaceAllString(tmp, "${1}_${2}")
	words := strings.Split(tmp, "_")
	return words
}
