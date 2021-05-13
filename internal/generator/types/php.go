package types

import (
	"os"

	"google.golang.org/protobuf/types/descriptorpb"
)

const (
	// PhpInt int type
	PhpInt Type = "int"
	// PhpDouble double type
	PhpDouble Type = "double"
	// PhpBoolean boolean type
	PhpBoolean Type = "boolean"
	// PhpString string type
	PhpString Type = "string"
)

var PhpTypes = map[descriptorpb.FieldDescriptorProto_Type]Type{
	descriptorpb.FieldDescriptorProto_TYPE_DOUBLE:   PhpDouble,
	descriptorpb.FieldDescriptorProto_TYPE_FLOAT:    PhpDouble,
	descriptorpb.FieldDescriptorProto_TYPE_INT64:    PhpInt,
	descriptorpb.FieldDescriptorProto_TYPE_SINT64:   PhpInt,
	descriptorpb.FieldDescriptorProto_TYPE_SFIXED64: PhpInt,
	descriptorpb.FieldDescriptorProto_TYPE_UINT64:   PhpInt,
	descriptorpb.FieldDescriptorProto_TYPE_FIXED64:  PhpInt,
	descriptorpb.FieldDescriptorProto_TYPE_INT32:    PhpInt,
	descriptorpb.FieldDescriptorProto_TYPE_SINT32:   PhpInt,
	descriptorpb.FieldDescriptorProto_TYPE_SFIXED32: PhpInt,
	descriptorpb.FieldDescriptorProto_TYPE_UINT32:   PhpInt,
	descriptorpb.FieldDescriptorProto_TYPE_FIXED32:  PhpInt,
	descriptorpb.FieldDescriptorProto_TYPE_BOOL:     PhpBoolean,
	descriptorpb.FieldDescriptorProto_TYPE_STRING:   PhpString,
	descriptorpb.FieldDescriptorProto_TYPE_BYTES:    PhpString,
}

// GetPhpType returns PHP type by protobuf type
func GetPhpType(fieldType descriptorpb.FieldDescriptorProto_Type) Type {
	value, ok := PhpTypes[fieldType]
	if !ok {
		println("php: unknown or unsupported field type: " + fieldType.String())
		os.Exit(1)
	}
	return value
}
