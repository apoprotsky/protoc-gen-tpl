package typescript

import (
	"github.com/apoprotsky/protoc-gen-tpl/internal/generator/messages"
	"github.com/apoprotsky/protoc-gen-tpl/internal/template"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
)

// GenerateFiles generates typescript files content
func (svc *Service) GenerateFiles(
	request *plugin.CodeGeneratorRequest,
	messages []*messages.Model,
) []*plugin.CodeGeneratorResponse_File {
	files := map[string]*plugin.CodeGeneratorResponse_File{}
	messagesByFiles := map[string]*model{}

	for _, message := range messages {
		if _, ok := files[message.TypescriptFile]; !ok {
			name := message.TypescriptFile
			files[message.TypescriptFile] = &plugin.CodeGeneratorResponse_File{
				Name: &name,
			}
			messagesByFiles[message.TypescriptFile] = newModel()
		}
		messagesByFiles[message.TypescriptFile].addMessage(message)
	}

	templateService := template.GetService()

	result := []*plugin.CodeGeneratorResponse_File{}
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
