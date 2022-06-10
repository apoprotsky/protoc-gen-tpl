package typescript

import "github.com/apoprotsky/protoc-gen-tpl/internal/generator/messages"

type model struct {
	TypescriptImports []string
	Messages          []*messages.Model
}

func newModel(message *messages.Model) *model {
	return &model{
		TypescriptImports: message.TypescriptImports,
	}
}
