package typescript

import "github.com/apoprotsky/protoc-gen-tpl/internal/generator/messages"

func (m *model) addMessage(msg *messages.Model) {
	m.Messages = append(m.Messages, msg)
}
