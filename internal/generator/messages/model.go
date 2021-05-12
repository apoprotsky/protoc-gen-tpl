package messages

import (
	"github.com/apoprotsky/protoc-gen-tpl/internal/generator/fields"
)

// Model of message
type Model struct {
	ProtoFile         string
	GoFile            string
	GoPackage         string
	TypescriptFile    string
	TypescriptPackage string
	Name              string
	Fields            []*fields.Model
	GoMax             int
}
