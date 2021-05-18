package typescript

import (
	"github.com/apoprotsky/protoc-gen-tpl/internal/generator/messages"
	"github.com/apoprotsky/protoc-gen-tpl/internal/template"
	"google.golang.org/protobuf/types/pluginpb"
)

// GenerateFiles generates typescript files content
func (svc *Service) GenerateFiles(
	request *pluginpb.CodeGeneratorRequest,
	messages []*messages.Model,
) []*pluginpb.CodeGeneratorResponse_File {
	files := map[string]*pluginpb.CodeGeneratorResponse_File{}
	messagesByFiles := map[string]*model{}

	for _, message := range messages {
		if _, ok := files[message.TypescriptFile]; !ok {
			name := message.TypescriptFile
			files[message.TypescriptFile] = &pluginpb.CodeGeneratorResponse_File{
				Name: &name,
			}
			messagesByFiles[message.TypescriptFile] = newModel()
		}
		messagesByFiles[message.TypescriptFile].addMessage(message)
	}

	templateService := template.GetService()

	result := []*pluginpb.CodeGeneratorResponse_File{}
	for _, file := range files {
		content := templateService.ExecuteTemplate(
			template.DefaultTypescriptTemplate,
			messagesByFiles[file.GetName()],
		)
		file.Content = &content
		result = append(result, file)
	}
	return result
}
