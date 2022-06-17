package typescript

import (
	"sort"

	"github.com/apoprotsky/protoc-gen-tpl/internal/generator/messages"
	"github.com/apoprotsky/protoc-gen-tpl/internal/slice"
)

func (m *model) addMessage(msg *messages.Model) {
	m.Messages = append(m.Messages, msg)
	m.TypescriptImports = append(m.TypescriptImports, msg.TypescriptImports...)
}

func (m *model) optimizeImports() {
	m.TypescriptImports = slice.Unique(m.TypescriptImports)
	sort.Strings(m.TypescriptImports)
}
