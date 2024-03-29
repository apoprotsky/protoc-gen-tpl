package generator

import (
	"strings"

	"github.com/apoprotsky/protoc-gen-tpl/internal/str"
	"google.golang.org/protobuf/types/descriptorpb"
)

func (svc *Service) getGoFileDirectory(protoFile *descriptorpb.FileDescriptorProto, prefix string) string {
	path := strings.Split(protoFile.GetOptions().GetGoPackage(), ";")[0]
	path = strings.TrimPrefix(path, prefix)
	path = strings.TrimPrefix(path, "/")
	if path != "" {
		path += "/"
	}
	return path
}

func (svc *Service) getGoPackageName(protoFile *descriptorpb.FileDescriptorProto) string {
	tmp := strings.Split(protoFile.GetOptions().GetGoPackage(), ";")
	if len(tmp) > 1 {
		return tmp[1]
	}
	return str.LastPart(tmp[0], "/")
}

func (svc *Service) getGoFullPackagePrefix(protoFile *descriptorpb.FileDescriptorProto) string {
	return protoFile.GetPackage() + "."
}
