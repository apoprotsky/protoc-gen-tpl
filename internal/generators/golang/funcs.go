package golang

import (
	"reflect"
	"strings"
	"text/template"

	"github.com/apoprotsky/prototpl/internal/str"
)

var funcs = template.FuncMap{
	"last": func(x int, a interface{}) bool {
		return x == reflect.ValueOf(a).Len()-1
	},
}

func getModuleFromParameter(parameter string) string {
	options := strings.Split(parameter, ",")
	for _, option := range options {
		tmp := strings.Split(option, "=")
		if tmp[0] == "module" {
			return tmp[1]
		}
	}
	return ""
}

func getFileDirectory(go_package string, module string) string {
	path := strings.Split(go_package, ";")[0]
	path = strings.TrimPrefix(path, module)
	path = strings.TrimPrefix(path, "/")
	if path != "" {
		path += "/"
	}
	return path
}

func getPackageName(go_package string) string {
	tmp := strings.Split(go_package, ";")
	if len(tmp) > 1 {
		return tmp[1]
	}
	return str.LastPart(tmp[0], "/")
}
