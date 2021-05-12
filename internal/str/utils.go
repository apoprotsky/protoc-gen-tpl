package str

import "strings"

func arrayTitle(strs []string) []string {
	for index, str := range strs {
		strs[index] = strings.Title(str)
	}
	return strs
}
