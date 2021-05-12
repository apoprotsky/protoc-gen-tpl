package typescript

import "github.com/apoprotsky/protoc-gen-tpl/internal/generator/messages"

type model struct {
	Messages []*messages.Model
}

func newModel() *model {
	return &model{
		Messages: []*messages.Model{},
	}
}
