package generator

import (
	"sort"
	"strings"

	"github.com/apoprotsky/protoc-gen-tpl/internal/generator/fields"
	"github.com/apoprotsky/protoc-gen-tpl/internal/generator/fields/options"
	"github.com/apoprotsky/protoc-gen-tpl/internal/generator/messages"
	"github.com/apoprotsky/protoc-gen-tpl/internal/generator/tags"
	"github.com/apoprotsky/protoc-gen-tpl/internal/generator/types"
	"github.com/apoprotsky/protoc-gen-tpl/internal/registry"
	"github.com/apoprotsky/protoc-gen-tpl/internal/slice"
	"github.com/apoprotsky/protoc-gen-tpl/internal/str"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/types/descriptorpb"
)

func genMessageFromProtoMessage(protoMessage *descriptorpb.DescriptorProto, goPackage string) *messages.Model {
	genMessage := messages.Model{
		Name:      str.ToUpperCamelCase(protoMessage.GetName()),
		Fields:    []*fields.Model{},
		GoMaxName: 0,
		GoMaxType: 0,
	}
	genFields, goMaxName, goMaxType, goImports, typescriptImports := genFieldsFromProtoFields(protoMessage.GetField(), goPackage)
	genMessage.Fields = genFields
	genMessage.GoMaxName = goMaxName
	genMessage.GoMaxType = goMaxType
	genMessage.GoImports = goImports
	genMessage.TypescriptImports = typescriptImports
	return &genMessage
}

func genFieldsFromProtoFields(protoFields []*descriptorpb.FieldDescriptorProto, goPackage string) (
	result []*fields.Model, goMaxName int, goMaxType int, goImports []string, typescriptImports []string) {
	genFields := []*fields.Model{}
	goMaxNameLength := 0
	goMaxTypeLength := 0
	var goMessageImports []string
	var typescriptMessageImports []string
	for _, protoField := range protoFields {
		genField, genGoImports, genTypescriptImports := genFieldFromProtoField(protoField, goPackage)
		genFields = append(genFields, genField)
		if len(genField.GoName) > goMaxNameLength {
			goMaxNameLength = len(genField.GoName)
		}
		if len(genField.GoType) > goMaxTypeLength {
			goMaxTypeLength = len(genField.GoType)
		}
		goMessageImports = append(goMessageImports, genGoImports...)
		typescriptMessageImports = append(typescriptMessageImports, genTypescriptImports...)
	}
	goMessageImports = slice.Unique(goMessageImports)
	typescriptMessageImports = slice.Unique(typescriptMessageImports)
	sort.Strings(goMessageImports)
	return genFields, goMaxNameLength, goMaxTypeLength, goMessageImports, typescriptMessageImports
}

func genFieldFromProtoField(protoField *descriptorpb.FieldDescriptorProto, goPackage string) (
	field *fields.Model, goImports []string, typescriptImports []string) {
	var genGoImports []string
	var genTypescriptImports []string
	goFieldType := types.GetGoType(protoField.GetType())
	typescriptFieldType := types.GetTypescriptType(protoField.GetType())
	phpFieldType := types.GetPhpType(protoField.GetType())
	if goFieldType == "message" {
		typeName := protoField.GetTypeName()
		typeName = typeName[1:]
		typeName = strings.TrimPrefix(typeName, goPackage)
		goFieldType = types.GoTypeToAliased(typeName)
		goAlias := types.GetAliasFromGoType(typeName)
		typescriptFieldType = types.Type(str.LastPart(typeName, "."))
		if len(goAlias) > 0 {
			registryService := registry.GetService()
			generatorService := GetService()
			file := registryService.GetFileByPackage(protoField.GetTypeName())
			genGoImports = append(
				genGoImports,
				goAlias+" \""+file.GetOptions().GetGoPackage()+"\"",
			)
			genTypescriptImport := generatorService.getTypescriptFileDirectory(file) + file.GetName()
			genTypescriptImport = strings.Replace(genTypescriptImport, ".proto", "", -1)
			genTypescriptImports = append(
				genTypescriptImports,
				"import { "+string(typescriptFieldType)+" } from '"+genTypescriptImport+"'",
			)
		}
		phpFieldType = types.Type(typeName)
	}
	genField := fields.Model{
		GoName:            str.ToUpperCamelCase(protoField.GetName()),
		GoIsArray:         protoField.GetLabel() == descriptorpb.FieldDescriptorProto_LABEL_REPEATED,
		GoIsPointer:       protoField.GetType() == descriptorpb.FieldDescriptorProto_TYPE_MESSAGE,
		GoType:            goFieldType,
		GoTags:            []*tags.Model{},
		TypescriptName:    str.ToLowerCamelCase(protoField.GetName()),
		TypescriptType:    typescriptFieldType,
		TypescriptIsArray: protoField.GetLabel() == descriptorpb.FieldDescriptorProto_LABEL_REPEATED,
		PhpName:           str.ToLowerCamelCase(protoField.GetName()),
		PhpType:           phpFieldType,
	}
	if genField.GoIsPointer {
		genField.GoType = "*" + genField.GoType
	}
	if genField.GoIsArray {
		genField.GoType = "[]" + genField.GoType
	}
	genField.GoTags = append(genField.GoTags, &tags.Model{Name: "json", Value: protoField.GetJsonName()})
	protoOptions := protoField.GetOptions()
	if protoOptions != nil {
		optionsStrings := strings.Split(prototext.MarshalOptions{Multiline: true}.Format(protoOptions), "\n")
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
	return &genField, genGoImports, genTypescriptImports
}
