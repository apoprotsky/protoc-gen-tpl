package golang

import (
	"github.com/apoprotsky/protoc-gen-tpl/internal/generator/messages"
	"github.com/apoprotsky/protoc-gen-tpl/internal/template"
	"google.golang.org/protobuf/types/pluginpb"
)

// GenerateFiles generates go files content
func (svc *Service) GenerateFiles(
	request *pluginpb.CodeGeneratorRequest,
	messages []*messages.Model,
) []*pluginpb.CodeGeneratorResponse_File {
	files := map[string]*pluginpb.CodeGeneratorResponse_File{}
	messagesByFiles := map[string]*model{}

	for _, message := range messages {
		if _, ok := files[message.GoFile]; !ok {
			name := message.GoFile
			files[message.GoFile] = &pluginpb.CodeGeneratorResponse_File{
				Name: &name,
			}
			messagesByFiles[message.GoFile] = newModel(message)
		}
		messagesByFiles[message.GoFile].addMessage(message)
	}

	templateService := template.GetService()

	result := []*pluginpb.CodeGeneratorResponse_File{}
	for _, file := range files {
		name := file.GetName()
		messagesByFiles[name].optimizeImports()
		content := templateService.ExecuteTemplate(template.DefaultGoTemplate, messagesByFiles[name])
		file.Content = &content
		result = append(result, file)
	}
	return result
}
