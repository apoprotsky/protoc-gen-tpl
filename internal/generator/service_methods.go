package generator

import (
	"strings"

	"github.com/apoprotsky/protoc-gen-tpl/internal/generators/golang"
	"github.com/apoprotsky/protoc-gen-tpl/internal/generators/php"
	"github.com/apoprotsky/protoc-gen-tpl/internal/generators/typescript"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
)

// GenerateCode creates response for protoc
func (svc *Service) GenerateCode(request *plugin.CodeGeneratorRequest) proto.Message {
	svc.getLanguagesFromParameter(request.GetParameter())
	messages := genMessagesFromRequest(request)

	files := []*pluginpb.CodeGeneratorResponse_File{}

	if svc.golang {
		golangService := golang.GetService()
		files = append(files, golangService.GenerateFiles(request, messages)...)
	}

	if svc.typescript {
		typescriptService := typescript.GetService()
		files = append(files, typescriptService.GenerateFiles(request, messages)...)
	}

	if svc.php {
		phpService := php.GetService()
		files = append(files, phpService.GenerateFiles(request, messages)...)
	}

	return &plugin.CodeGeneratorResponse{File: files}
}

func (svc *Service) getLanguagesFromParameter(parameter string) {
	options := strings.Split(parameter, ",")
	for _, option := range options {
		tmp := strings.Split(option, "=")
		if tmp[0] != "lang" {
			continue
		}
		langs := strings.Split(tmp[1], ";")
		for _, lang := range langs {
			lang = strings.ToLower(lang)
			switch lang {
			case "go":
				svc.golang = true
			case "golang":
				svc.golang = true
			case "ts":
				svc.typescript = true
			case "typescript":
				svc.typescript = true
			case "php":
				svc.php = true
			}
		}
	}
}
