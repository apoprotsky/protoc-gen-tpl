package generator

import (
	"github.com/apoprotsky/protoc-gen-tpl/internal/generators/golang"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
)

func (svc *Service) GenerateCode(request *plugin.CodeGeneratorRequest) proto.Message {
	messages := genMessagesFromRequest(request)

	files := []*pluginpb.CodeGeneratorResponse_File{}

	golangService := golang.GetService()
	files = append(files, golangService.GenerateFiles(request, messages)...)

	return &plugin.CodeGeneratorResponse{File: files}
}
