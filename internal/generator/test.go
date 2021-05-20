package generator

import (
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

func proto_test(request *pluginpb.CodeGeneratorRequest) {

	fds := &descriptorpb.FileDescriptorSet{File: request.GetProtoFile()}
	files, _ := protodesc.NewFiles(fds)
	files.RangeFiles(func(file protoreflect.FileDescriptor) bool {
		if file.Name() != "Users" {
			return true
		}
		messages := file.Messages()
		for i := 0; i < messages.Len(); i++ {
			message := messages.Get(i)
			if message.Name() != "User" {
				continue
			}
			fields := message.Fields()
			for i := 0; i < fields.Len(); i++ {
				field := fields.Get(i)
				if field.Name() != "email" {
					continue
				}
				options := field.Options().(*descriptorpb.FieldOptions)
				print(prototext.MarshalOptions{Multiline: true, Indent: ""}.Format(options))
			}
		}
		return true
	})
	// for _, file := range files. {
	// 	if file.GetName() != "examples/proto/users.proto" {
	// 		continue
	// 	}
	// 	println("File", file.GetName())
	// }

	// pg, _ := protogen.Options{}.New(request)
	// for _, file := range pg.Files {
	// 	println("##", file.GoPackageName)
	// 	println("###", "Extensions")
	// 	for _, extension := range file.Extensions {
	// 		println("####", extension.GoName)
	// 	}
	// 	println("##", "Messages")
	// 	for _, message := range file.Messages {
	// 		println("###", message.GoIdent.GoName)
	// 		println("{")
	// 		for _, field := range message.Fields {
	// 			println(field.GoName)
	// 		}
	// 		println("}")
	// 	}
	// }
}
