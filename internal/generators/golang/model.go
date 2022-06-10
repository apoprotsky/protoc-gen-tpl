package golang

import "github.com/apoprotsky/protoc-gen-tpl/internal/generator/messages"

type model struct {
	GoPackage string
	GoMaxName int
	GoMaxType int
	GoImports []string
	Messages  []*messages.Model
}

func newModel(message *messages.Model) *model {
	return &model{
		GoPackage: message.GoPackage,
		GoMaxName: message.GoMaxName,
		GoMaxType: message.GoMaxType,
		GoImports: message.GoImports,
	}
}
