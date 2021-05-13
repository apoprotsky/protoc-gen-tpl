package generator

import (
	"strings"

	"github.com/apoprotsky/protoc-gen-tpl/internal/str"
	"google.golang.org/protobuf/types/descriptorpb"
)

func getPhpFileDirectory(protoFile *descriptorpb.FileDescriptorProto, prefix string) string {
	return strings.ReplaceAll(protoFile.GetPackage(), ".", "/") + "/"
}

func getPhpPackageName(protoFile *descriptorpb.FileDescriptorProto) string {
	parts := strings.Split(protoFile.GetPackage(), ".")
	parts = str.TitleArray(parts)
	return strings.Join(parts, "\\")
}
