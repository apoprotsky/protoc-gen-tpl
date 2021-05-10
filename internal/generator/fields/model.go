package fields

import (
	"github.com/apoprotsky/protoc-gen-tpl/internal/generator/tags"
	"github.com/apoprotsky/protoc-gen-tpl/internal/generator/types"
)

type Model struct {
	GoName      string
	GoIsArray   bool
	GoIsPointer bool
	GoType      types.Type
	GoTags      []*tags.Model
}
