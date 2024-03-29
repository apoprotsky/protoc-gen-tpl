package fields

import (
	"github.com/apoprotsky/protoc-gen-tpl/internal/generator/tags"
	"github.com/apoprotsky/protoc-gen-tpl/internal/generator/types"
)

// Model of field
type Model struct {
	GoName            string
	GoIsArray         bool
	GoIsPointer       bool
	GoType            types.Type
	GoTags            []*tags.Model
	TypescriptName    string
	TypescriptType    types.Type
	TypescriptIsArray bool
	PhpName           string
	PhpType           types.Type
}
