package messages

import (
	"github.com/apoprotsky/protoc-gen-tpl/internal/generator/fields"
)

// Model of message
type Model struct {
	ProtoFile         string
	GoFile            string
	GoPackage         string
	GoMaxName         int
	GoMaxType         int
	GoImports         []string
	TypescriptFile    string
	TypescriptPackage string
	TypescriptImports []string
	PhpFile           string
	PhpPackage        string
	Name              string
	Fields            []*fields.Model
}
