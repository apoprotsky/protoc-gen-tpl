package types

import (
	"os"

	"google.golang.org/protobuf/types/descriptorpb"
)

const (
	// TypescriptBigint bigint type
	TypescriptBigint Type = "bigint"
	// TypescriptNumber number type
	TypescriptNumber Type = "number"
	// TypescriptBoolean boolean type
	TypescriptBoolean Type = "boolean"
	// TypescriptString string type
	TypescriptString Type = "string"
	// TypescriptEnum enum type
	TypescriptEnum Type = "enum"
	// TypescriptMessage message type
	TypescriptMessage Type = "message"
)

var typescriptTypes = map[descriptorpb.FieldDescriptorProto_Type]Type{
	descriptorpb.FieldDescriptorProto_TYPE_DOUBLE:   TypescriptBigint,
	descriptorpb.FieldDescriptorProto_TYPE_FLOAT:    TypescriptNumber,
	descriptorpb.FieldDescriptorProto_TYPE_INT64:    TypescriptBigint,
	descriptorpb.FieldDescriptorProto_TYPE_SINT64:   TypescriptBigint,
	descriptorpb.FieldDescriptorProto_TYPE_SFIXED64: TypescriptBigint,
	descriptorpb.FieldDescriptorProto_TYPE_UINT64:   TypescriptBigint,
	descriptorpb.FieldDescriptorProto_TYPE_FIXED64:  TypescriptBigint,
	descriptorpb.FieldDescriptorProto_TYPE_INT32:    TypescriptNumber,
	descriptorpb.FieldDescriptorProto_TYPE_SINT32:   TypescriptNumber,
	descriptorpb.FieldDescriptorProto_TYPE_SFIXED32: TypescriptNumber,
	descriptorpb.FieldDescriptorProto_TYPE_UINT32:   TypescriptNumber,
	descriptorpb.FieldDescriptorProto_TYPE_FIXED32:  TypescriptNumber,
	descriptorpb.FieldDescriptorProto_TYPE_BOOL:     TypescriptBoolean,
	descriptorpb.FieldDescriptorProto_TYPE_STRING:   TypescriptString,
	descriptorpb.FieldDescriptorProto_TYPE_BYTES:    TypescriptString,
	descriptorpb.FieldDescriptorProto_TYPE_ENUM:     TypescriptEnum,
	descriptorpb.FieldDescriptorProto_TYPE_MESSAGE:  TypescriptMessage,
}

// GetTypescriptType returns typescript type by protobuf type
func GetTypescriptType(fieldType descriptorpb.FieldDescriptorProto_Type) Type {
	value, ok := typescriptTypes[fieldType]
	if !ok {
		println("typescript: unknown or unsupported field type: " + fieldType.String())
		os.Exit(1)
	}
	return value
}
