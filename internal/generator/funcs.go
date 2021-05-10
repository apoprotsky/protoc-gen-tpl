package generator

import (
	"strings"

	"github.com/apoprotsky/protoc-gen-tpl/internal/generator/fields"
	"github.com/apoprotsky/protoc-gen-tpl/internal/generator/messages"
	"github.com/apoprotsky/protoc-gen-tpl/internal/generator/tags"
	"github.com/apoprotsky/protoc-gen-tpl/internal/generator/types"
	"github.com/apoprotsky/protoc-gen-tpl/internal/str"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"google.golang.org/protobuf/types/descriptorpb"
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

func getGoFileDirectory(go_package string, module string) string {
	path := strings.Split(go_package, ";")[0]
	path = strings.TrimPrefix(path, module)
	path = strings.TrimPrefix(path, "/")
	if path != "" {
		path += "/"
	}
	return path
}

func getGoPackageName(go_package string) string {
	tmp := strings.Split(go_package, ";")
	if len(tmp) > 1 {
		return tmp[1]
	}
	return str.LastPart(tmp[0], "/")
}

func genMessagesFromRequest(request *plugin.CodeGeneratorRequest) []*messages.Model {
	messages := []*messages.Model{}
	prefix := getPrefixFromParameter(request.GetParameter())
	for index, filename := range request.FileToGenerate {
		protoFile := request.ProtoFile[index]
		filename = str.LastPart(filename, "/")
		// go
		go_package := protoFile.GetOptions().GetGoPackage()
		goDir := getGoFileDirectory(go_package, prefix)
		goFile := goDir + strings.Replace(filename, ".proto", ".go", -1)
		// Get messages information
		protoMessages := protoFile.MessageType
		for _, protoMessage := range protoMessages {
			genMessage := genMessageFromProtoMessage(protoMessage)
			genMessage.ProtoFile = protoFile.GetName()
			genMessage.GoFile = goFile
			genMessage.GoPackage = getGoPackageName(go_package)
			messages = append(messages, genMessage)
		}
	}
	return messages
}

func genMessageFromProtoMessage(protoMessage *descriptorpb.DescriptorProto) *messages.Model {
	genMessage := messages.Model{
		Name:   protoMessage.GetName(),
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
		GoName:      strings.Title(strings.ToLower(protoField.GetName())),
		GoIsArray:   protoField.GetLabel() == descriptorpb.FieldDescriptorProto_LABEL_REPEATED,
		GoIsPointer: protoField.GetType() == descriptorpb.FieldDescriptorProto_TYPE_MESSAGE,
		GoType:      types.GetGoType(protoField.GetType()),
		GoTags:      []*tags.Model{},
	}
	genField.GoTags = append(genField.GoTags, &tags.Model{Name: "json", Value: protoField.GetJsonName()})
	return &genField
}
