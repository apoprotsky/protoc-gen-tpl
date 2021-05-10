package golang

import "github.com/apoprotsky/protoc-gen-tpl/internal/generator/messages"

type model struct {
	GoPackage string
	GoMax     int
	Messages  []*messages.Model
}

func newModel(goPackage string, goMax int) *model {
	return &model{
		GoPackage: goPackage,
		GoMax:     goMax,
		Messages:  []*messages.Model{},
	}
}
