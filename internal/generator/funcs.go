package generator

import (
	"strings"

	"github.com/apoprotsky/protoc-gen-tpl/internal/generator/fields"
	"github.com/apoprotsky/protoc-gen-tpl/internal/generator/messages"
	"github.com/apoprotsky/protoc-gen-tpl/internal/generator/tags"
	"github.com/apoprotsky/protoc-gen-tpl/internal/generator/types"
	"github.com/apoprotsky/protoc-gen-tpl/internal/str"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

func getPrefixFromParameter(parameter string) string {
	options := strings.Split(parameter, ",")
	for _, option := range options {
		tmp := strings.Split(option, "=")
		if tmp[0] == "prefix" {
			return tmp[1]
		}
	}
	return ""
}

func getProtoFileByName(
	request *pluginpb.CodeGeneratorRequest,
	name string,
) *descriptorpb.FileDescriptorProto {
	protoFiles := request.GetProtoFile()
	for _, protoFile := range protoFiles {
		if protoFile.GetName() == name {
			return protoFile
		}
	}
	return nil
}

func genMessagesFromRequest(request *pluginpb.CodeGeneratorRequest) []*messages.Model {
	messages := []*messages.Model{}
	prefix := getPrefixFromParameter(request.GetParameter())
	for _, filename := range request.GetFileToGenerate() {
		protoFile := getProtoFileByName(request, filename)
		filename = str.LastPart(filename, "/")
		// go
		goDir := getGoFileDirectory(protoFile, prefix)
		goFile := goDir + strings.Replace(filename, ".proto", ".go", -1)
		goPackage := getGoPackageName(protoFile)
		// typescript
		typescriptDir := getTypescriptFileDirectory(protoFile, prefix)
		typescriptFile := typescriptDir + strings.Replace(filename, ".proto", ".ts", -1)
		typescriptPackage := getTypescriptPackageName(protoFile)
		// php
		phpDir := getPhpFileDirectory(protoFile, prefix)
		phpPackage := getPhpPackageName(protoFile)
		//
		protoMessages := protoFile.MessageType
		for _, protoMessage := range protoMessages {
			genMessage := genMessageFromProtoMessage(protoMessage)
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

func genMessageFromProtoMessage(protoMessage *descriptorpb.DescriptorProto) *messages.Model {
	genMessage := messages.Model{
		Name:   str.ToUpperCamelCase(protoMessage.GetName()),
		Fields: []*fields.Model{},
		GoMax:  0,
	}
	genFields, goMax := genFieldsFromProtoFields(protoMessage.GetField())
	genMessage.Fields = genFields
	genMessage.GoMax = goMax
	return &genMessage
}

func genFieldsFromProtoFields(protoFields []*descriptorpb.FieldDescriptorProto) (result []*fields.Model, goMax int) {
	genFields := []*fields.Model{}
	goMaxName := 0
	for _, protoField := range protoFields {
		genField := genFieldFromProtoField(protoField)
		genFields = append(genFields, genField)
		if len(genField.GoName) > goMax {
			goMax = len(genField.GoName)
		}
	}
	return genFields, goMaxName
}

func genFieldFromProtoField(protoField *descriptorpb.FieldDescriptorProto) *fields.Model {
	genField := fields.Model{
		GoName:         str.ToUpperCamelCase(protoField.GetName()),
		GoIsArray:      protoField.GetLabel() == descriptorpb.FieldDescriptorProto_LABEL_REPEATED,
		GoIsPointer:    protoField.GetType() == descriptorpb.FieldDescriptorProto_TYPE_MESSAGE,
		GoType:         types.GetGoType(protoField.GetType()),
		GoTags:         []*tags.Model{},
		TypescriptName: str.ToLowerCamelCase(protoField.GetName()),
		TypescriptType: types.GetTypescriptType(protoField.GetType()),
		PhpName:        str.ToLowerCamelCase(protoField.GetName()),
		PhpType:        types.GetPhpType(protoField.GetType()),
	}
	genField.GoTags = append(genField.GoTags, &tags.Model{Name: "json", Value: protoField.GetJsonName()})
	return &genField
}
