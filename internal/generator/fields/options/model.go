package options

import "google.golang.org/protobuf/types/descriptorpb"

// Model of extension
type Model struct {
	Id    int32
	Name  string
	Type  descriptorpb.FieldDescriptorProto_Type
	Value string
}
