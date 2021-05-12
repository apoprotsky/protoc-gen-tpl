package str

import "strings"

func ArrayTitle(strs []string) []string {
	for index, str := range strs {
		strs[index] = strings.Title(str)
	}
	return strs
}
