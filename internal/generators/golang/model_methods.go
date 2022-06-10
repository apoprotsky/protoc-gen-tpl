package golang

import (
	"sort"

	"github.com/apoprotsky/protoc-gen-tpl/internal/generator/messages"
	"github.com/apoprotsky/protoc-gen-tpl/internal/slice"
)

func (m *model) addMessage(msg *messages.Model) {
	m.Messages = append(m.Messages, msg)
	m.GoImports = append(m.GoImports, msg.GoImports...)
}

func (m *model) optimizeImports() {
	m.GoImports = slice.Unique(m.GoImports)
	sort.Strings(m.GoImports)
}
