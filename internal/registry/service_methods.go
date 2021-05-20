package registry

import (
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

// ParseRequest gather information from request
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

// GetFileByName returns file descriptor
func (svc *Service) GetFileByName(name string) *descriptorpb.FileDescriptorProto {
	return svc.filesByName[name]
}

// GetExtensionByNumber returns extension descriptor
func (svc *Service) GetExtensionByNumber(number int32) *descriptorpb.FieldDescriptorProto {
	return svc.extensionsByNumber[number]
}
