package generator

import (
	"github.com/apoprotsky/protoc-gen-tpl/internal/generators/golang"
	"github.com/apoprotsky/protoc-gen-tpl/internal/generators/typescript"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
)

// GenerateCode creates responce for protoc
func (svc *Service) GenerateCode(request *plugin.CodeGeneratorRequest) proto.Message {
	messages := genMessagesFromRequest(request)

	files := []*pluginpb.CodeGeneratorResponse_File{}

	golangService := golang.GetService()
	files = append(files, golangService.GenerateFiles(request, messages)...)

	typescriptService := typescript.GetService()
	files = append(files, typescriptService.GenerateFiles(request, messages)...)

	return &plugin.CodeGeneratorResponse{File: files}
}
