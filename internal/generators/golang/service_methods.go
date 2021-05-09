package golang

import (
	"bytes"
	"strings"
	"text/template"

	"github.com/apoprotsky/prototpl/internal/generators/golang/types"
	"github.com/apoprotsky/prototpl/internal/str"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

func (svc *Service) GenerateCode(request *plugin.CodeGeneratorRequest) proto.Message {

	module := getModuleFromParameter(request.GetParameter())

	var responce plugin.CodeGeneratorResponse
	responce.File = []*plugin.CodeGeneratorResponse_File{}

	for index, filename := range request.FileToGenerate {
		print(filename, " > ")
		filename = str.LastPart(filename, "/")
		protoFile := request.ProtoFile[index]
		go_package := protoFile.GetOptions().GetGoPackage()

		ourDirName := getFileDirectory(go_package, module)
		outFileName := ourDirName + strings.Replace(filename, ".proto", ".go", -1)
		println(ourDirName, "; ", outFileName)

		var outFile plugin.CodeGeneratorResponse_File
		outFile.Name = &outFileName

		goFile := types.File{
			Package: getPackageName(go_package),
			Structs: []types.Struct{},
		}

		protoMessages := protoFile.MessageType
		for _, protoMessage := range protoMessages {

			goStruct := types.Struct{
				Name:               protoMessage.GetName(),
				Fields:             []types.Field{},
				MaxFieldNameLength: 0,
			}

			protoFields := protoMessage.GetField()
			for _, protoField := range protoFields {

				goField := types.Field{
					Name:      strings.Title(strings.ToLower(protoField.GetName())),
					Type:      types.GetType(protoField.GetType()),
					IsArray:   protoField.GetLabel() == descriptorpb.FieldDescriptorProto_LABEL_REPEATED,
					IsPointer: protoField.GetType() == descriptorpb.FieldDescriptorProto_TYPE_MESSAGE,
					Tags:      []types.Tag{},
				}

				goField.Tags = append(goField.Tags, types.Tag{Name: "json", Value: protoField.GetJsonName()})

				goStruct.Fields = append(goStruct.Fields, goField)
				if len(goField.Name) > goStruct.MaxFieldNameLength {
					goStruct.MaxFieldNameLength = len(goField.Name)
				}

			}

			goFile.Structs = append(goFile.Structs, goStruct)
		}

		tmpl := template.Must(template.New("go").Funcs(funcs).Parse(defaultTemplate))
		var buffer bytes.Buffer
		tmpl.Execute(&buffer, &goFile)
		content := buffer.String()

		outFile.Content = &content
		responce.File = append(responce.File, &outFile)
	}

	return &responce

}
