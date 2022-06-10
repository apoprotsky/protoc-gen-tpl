package registry

import (
	gs "github.com/apoprotsky/go-services"
	"google.golang.org/protobuf/types/descriptorpb"
)

// Service struct
type Service struct {
	filesByName        map[string]*descriptorpb.FileDescriptorProto
	filesByPackage     map[string]*descriptorpb.FileDescriptorProto
	extensionsByNumber map[int32]*descriptorpb.FieldDescriptorProto
}

// GoService initializes service
func (svc *Service) GoService() {
	svc.filesByName = map[string]*descriptorpb.FileDescriptorProto{}
	svc.filesByPackage = map[string]*descriptorpb.FileDescriptorProto{}
	svc.extensionsByNumber = map[int32]*descriptorpb.FieldDescriptorProto{}
}

// GetService returns instance of service
func GetService() *Service {
	return gs.Get((*Service)(nil)).(*Service)
}
