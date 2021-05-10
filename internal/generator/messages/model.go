package messages

import (
	"github.com/apoprotsky/protoc-gen-tpl/internal/generator/fields"
)

type Model struct {
	ProtoFile string
	GoFile    string
	GoPackage string
	Name      string
	Fields    []*fields.Model
	GoMax     int
}
