package str

import "strings"

// TitleArray converts array of strings to Title form
func TitleArray(strs []string) []string {
	for index, str := range strs {
		strs[index] = strings.Title(str)
	}
	return strs
}
