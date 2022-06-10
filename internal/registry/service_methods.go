package registry

import (
	"strings"

	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

// ParseRequest gather information from request
func (svc *Service) ParseRequest(request *pluginpb.CodeGeneratorRequest) {
	protoFiles := request.GetProtoFile()
	for _, protoFile := range protoFiles {
		svc.filesByName[protoFile.GetName()] = protoFile
		svc.filesByPackage[protoFile.GetPackage()] = protoFile
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

func (svc *Service) GetFileByPackage(name string) *descriptorpb.FileDescriptorProto {
	if name[0] == '.' {
		name = name[1:]
	}
	parts := strings.Split(name, ".")
	name = strings.Join(parts[:len(parts)-1], ".")
	return svc.filesByPackage[name]
}

// GetExtensionByNumber returns extension descriptor
func (svc *Service) GetExtensionByNumber(number int32) *descriptorpb.FieldDescriptorProto {
	return svc.extensionsByNumber[number]
}
