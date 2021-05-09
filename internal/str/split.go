package str

import "strings"

func LastPart(str string, separator string) string {
	tmp := strings.Split(str, separator)
	return tmp[len(tmp)-1]
}
