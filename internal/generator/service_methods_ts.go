package generator

import (
	"strings"

	"github.com/apoprotsky/protoc-gen-tpl/internal/str"
	"google.golang.org/protobuf/types/descriptorpb"
)

func (svc *Service) getTypescriptFileDirectory(protoFile *descriptorpb.FileDescriptorProto) string {
	return strings.ReplaceAll(protoFile.GetPackage(), ".", "/") + "/"
}

func (svc *Service) getTypescriptPackageName(protoFile *descriptorpb.FileDescriptorProto) string {
	return str.LastPart(protoFile.GetPackage(), ".")
}
