package generator

import (
	"strings"

	"github.com/apoprotsky/protoc-gen-tpl/internal/str"
	"google.golang.org/protobuf/types/descriptorpb"
)

func getTypescriptFileDirectory(protoFile *descriptorpb.FileDescriptorProto, prefix string) string {
	return strings.ReplaceAll(protoFile.GetPackage(), ".", "/") + "/"
}

func getTypescriptPackageName(protoFile *descriptorpb.FileDescriptorProto) string {
	return str.LastPart(protoFile.GetPackage(), ".")
}
