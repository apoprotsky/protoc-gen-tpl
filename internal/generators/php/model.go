package php

import "github.com/apoprotsky/protoc-gen-tpl/internal/generator/messages"

type model struct {
	PhpPackage string
	Message    *messages.Model
}

func newModel() *model {
	return &model{}
}
