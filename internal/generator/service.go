package generator

import (
	gs "github.com/apoprotsky/go-services"
	"github.com/apoprotsky/protoc-gen-tpl/internal/registry"
	"google.golang.org/protobuf/types/pluginpb"
)

// Service struct
type Service struct {
	golang          bool
	typescript      bool
	php             bool
	registryService *registry.Service
	request         *pluginpb.CodeGeneratorRequest
	prefix          string
}

// GoService initializes service
func (svc *Service) GoService(registryService *registry.Service) {
	svc.registryService = registryService
}

// GetService returns instance of service
func GetService() *Service {
	return gs.Get((*Service)(nil)).(*Service)
}
