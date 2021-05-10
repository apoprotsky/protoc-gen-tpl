package template

import (
	"reflect"
	"text/template"
)

var funcs = template.FuncMap{
	"last": func(x int, a interface{}) bool {
		return x == reflect.ValueOf(a).Len()-1
	},
}
