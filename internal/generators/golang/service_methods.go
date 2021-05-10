package golang

import (
	"github.com/apoprotsky/protoc-gen-tpl/internal/generator/messages"
	"github.com/apoprotsky/protoc-gen-tpl/internal/template"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
)

// GenerateFiles generates go files content
func (svc *Service) GenerateFiles(request *plugin.CodeGeneratorRequest, messages []*messages.Model) []*plugin.CodeGeneratorResponse_File {
	files := map[string]*plugin.CodeGeneratorResponse_File{}
	messagesByFiles := map[string]*model{}

	for _, message := range messages {
		if _, ok := files[message.GoFile]; !ok {
			name := message.GoFile
			files[message.GoFile] = &plugin.CodeGeneratorResponse_File{
				Name: &name,
			}
			messagesByFiles[message.GoFile] = newModel(message.GoPackage, message.GoMax)
		}
		messagesByFiles[message.GoFile].addMessage(message)
	}

	templateService := template.GetService()

	result := []*plugin.CodeGeneratorResponse_File{}
	for _, file := range files {
		content := templateService.ExecuteTemplate(template.DefaultGoTemplate, messagesByFiles[file.GetName()])
		file.Content = &content
		result = append(result, file)
	}
	return result
}
