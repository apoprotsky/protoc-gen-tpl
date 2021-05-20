package registry

import (
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

func (svc *Service) ParseRequest(request *pluginpb.CodeGeneratorRequest) {
	protoFiles := request.GetProtoFile()
	for _, protoFile := range protoFiles {
		svc.filesByName[protoFile.GetName()] = protoFile
		protoExtensions := protoFile.GetExtension()
		for _, protoExtension := range protoExtensions {
			svc.extensionsByNumber[protoExtension.GetNumber()] = protoExtension
		}
	}
}

func (svc *Service) GetFileByName(name string) *descriptorpb.FileDescriptorProto {
	return svc.filesByName[name]
}

func (svc *Service) GetExtensionByNumber(number int32) *descriptorpb.FieldDescriptorProto {
	return svc.extensionsByNumber[number]
}
