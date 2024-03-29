package str

import (
	"regexp"
	"strings"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

// ToSnakeCase form
func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

// ToUpperCamelCase form
func ToUpperCamelCase(str string) string {
	words := splitInWords(str)
	words = TitleArray(words)
	return strings.Join(words, "")
}

// ToLowerCamelCase form
func ToLowerCamelCase(str string) string {
	words := splitInWords(str)
	words = TitleArray(words)
	words[0] = strings.ToLower(words[0])
	return strings.Join(words, "")
}
