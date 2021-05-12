package str

import "strings"

// LastPart of string using separator
func LastPart(str string, separator string) string {
	tmp := strings.Split(str, separator)
	return tmp[len(tmp)-1]
}
