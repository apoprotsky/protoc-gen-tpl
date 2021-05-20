package generator

import (
	"strings"

	"github.com/apoprotsky/protoc-gen-tpl/internal/generator/fields"
	"github.com/apoprotsky/protoc-gen-tpl/internal/generator/fields/options"
	"github.com/apoprotsky/protoc-gen-tpl/internal/generator/messages"
	"github.com/apoprotsky/protoc-gen-tpl/internal/generator/tags"
	"github.com/apoprotsky/protoc-gen-tpl/internal/generator/types"
	"github.com/apoprotsky/protoc-gen-tpl/internal/str"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/types/descriptorpb"
)

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
	protoOptions := protoField.GetOptions()
	if protoOptions != nil {
		optionsStrings := strings.Split(prototext.MarshalOptions{Multiline: true, Indent: ""}.Format(protoOptions), "\n")
		for _, optionString := range optionsStrings {
			if optionString == "" {
				continue
			}
			optionModel := options.New(optionString)
			if optionModel != nil {
				genField.GoTags = append(genField.GoTags, &tags.Model{
					Name:  optionModel.Name,
					Value: optionModel.Value,
				})
			}
		}
	}
	return &genField
}
