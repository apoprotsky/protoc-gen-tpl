package types

import (
	"os"

	"google.golang.org/protobuf/types/descriptorpb"
)

const (
	// GoFloat64 float64 type
	GoFloat64 Type = "float64"
	// GoFloat32 float32 type
	GoFloat32 Type = "float32"
	// GoInt64 int64 type
	GoInt64 Type = "int64"
	// GoUInt64 uint64 type
	GoUInt64 Type = "uint64"
	// GoInt32 int32 type
	GoInt32 Type = "int32"
	// GoUInt32 uint32 type
	GoUInt32 Type = "uint32"
	// GoBool bool type
	GoBool Type = "bool"
	// GoString string type
	GoString Type = "string"
)

var goTypes = map[descriptorpb.FieldDescriptorProto_Type]Type{
	descriptorpb.FieldDescriptorProto_TYPE_DOUBLE:   GoFloat64,
	descriptorpb.FieldDescriptorProto_TYPE_FLOAT:    GoFloat32,
	descriptorpb.FieldDescriptorProto_TYPE_INT64:    GoInt64,
	descriptorpb.FieldDescriptorProto_TYPE_SINT64:   GoInt64,
	descriptorpb.FieldDescriptorProto_TYPE_SFIXED64: GoInt64,
	descriptorpb.FieldDescriptorProto_TYPE_UINT64:   GoUInt64,
	descriptorpb.FieldDescriptorProto_TYPE_FIXED64:  GoUInt64,
	descriptorpb.FieldDescriptorProto_TYPE_INT32:    GoInt32,
	descriptorpb.FieldDescriptorProto_TYPE_SINT32:   GoInt32,
	descriptorpb.FieldDescriptorProto_TYPE_SFIXED32: GoInt32,
	descriptorpb.FieldDescriptorProto_TYPE_UINT32:   GoUInt32,
	descriptorpb.FieldDescriptorProto_TYPE_FIXED32:  GoUInt32,
	descriptorpb.FieldDescriptorProto_TYPE_BOOL:     GoBool,
	descriptorpb.FieldDescriptorProto_TYPE_STRING:   GoString,
	descriptorpb.FieldDescriptorProto_TYPE_BYTES:    GoString,
}

// GetGoType returns golang type by protobuf type
func GetGoType(fieldType descriptorpb.FieldDescriptorProto_Type) Type {
	value, ok := goTypes[fieldType]
	if !ok {
		println("go: unknown or unsupported field type: " + fieldType.String())
		os.Exit(1)
	}
	return value
}
