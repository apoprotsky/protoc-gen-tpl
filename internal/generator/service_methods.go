package generator

import (
	"strings"

	"github.com/apoprotsky/protoc-gen-tpl/internal/generator/messages"
	"github.com/apoprotsky/protoc-gen-tpl/internal/generators/golang"
	"github.com/apoprotsky/protoc-gen-tpl/internal/generators/php"
	"github.com/apoprotsky/protoc-gen-tpl/internal/generators/typescript"
	"github.com/apoprotsky/protoc-gen-tpl/internal/str"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
)

// GenerateCode creates response for protoc
func (svc *Service) GenerateCode(request *pluginpb.CodeGeneratorRequest) proto.Message {
	svc.registryService.ParseRequest(request)
	svc.getLanguages(request)
	messages := svc.genMessages(request)

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

	return &pluginpb.CodeGeneratorResponse{File: files}
}

func (svc *Service) getLanguages(request *pluginpb.CodeGeneratorRequest) {
	options := strings.Split(request.GetParameter(), ",")
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

func (svc *Service) setRequest(request *pluginpb.CodeGeneratorRequest) {
	svc.request = request
	options := strings.Split(request.GetParameter(), ",")
	for _, option := range options {
		tmp := strings.Split(option, "=")
		if tmp[0] == "prefix" {
			svc.prefix = tmp[1]
		}
	}
}

func (svc *Service) genMessages(request *pluginpb.CodeGeneratorRequest) []*messages.Model {
	messages := []*messages.Model{}
	svc.setRequest(request)
	for _, filename := range request.GetFileToGenerate() {
		protoFile := svc.registryService.GetFileByName(filename)
		if protoFile == nil {
			panic(filename + " not registred")
		}
		filename = str.LastPart(filename, "/")
		// go
		goDir := svc.getGoFileDirectory(protoFile, svc.prefix)
		goFile := goDir + strings.Replace(filename, ".proto", ".go", -1)
		goPackage := svc.getGoPackageName(protoFile)
		goFillPackagePrefix := svc.getGoFullPackagePrefix(protoFile)
		// typescript
		typescriptDir := svc.getTypescriptFileDirectory(protoFile)
		typescriptFile := typescriptDir + strings.Replace(filename, ".proto", ".ts", -1)
		typescriptPackage := svc.getTypescriptPackageName(protoFile)
		// php
		phpDir := svc.getPhpFileDirectory(protoFile, svc.prefix)
		phpPackage := svc.getPhpPackageName(protoFile)
		//
		protoMessages := protoFile.MessageType
		for _, protoMessage := range protoMessages {
			genMessage := genMessageFromProtoMessage(protoMessage, goFillPackagePrefix)
			genMessage.ProtoFile = protoFile.GetName()
			genMessage.GoFile = goFile
			genMessage.GoPackage = goPackage
			genMessage.TypescriptFile = typescriptFile
			genMessage.TypescriptPackage = typescriptPackage
			genMessage.PhpFile = phpDir + genMessage.Name + ".php"
			genMessage.PhpPackage = phpPackage
			messages = append(messages, genMessage)
		}
	}
	return messages
}
