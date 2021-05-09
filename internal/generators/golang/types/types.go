package types

import (
	"errors"

	"google.golang.org/protobuf/types/descriptorpb"
)

type Type string

const (
	Float64 Type = "float64"
	Float32 Type = "float32"
	Int64   Type = "int64"
	UInt64  Type = "uint64"
	Int32   Type = "int32"
	UInt32  Type = "uint32"
	Bool    Type = "bool"
	String  Type = "string"
	// Message Type = "message"
	// Enum Type = "enum"
)

func GetType(fieldType descriptorpb.FieldDescriptorProto_Type) Type {
	switch fieldType {
	case descriptorpb.FieldDescriptorProto_TYPE_DOUBLE:
		return Float64
	case descriptorpb.FieldDescriptorProto_TYPE_FLOAT:
		return Float32
	case descriptorpb.FieldDescriptorProto_TYPE_INT64:
		return Int64
	case descriptorpb.FieldDescriptorProto_TYPE_SINT64:
		return Int64
	case descriptorpb.FieldDescriptorProto_TYPE_SFIXED64:
		return Int64
	case descriptorpb.FieldDescriptorProto_TYPE_UINT64:
		return UInt64
	case descriptorpb.FieldDescriptorProto_TYPE_FIXED64:
		return UInt64
	case descriptorpb.FieldDescriptorProto_TYPE_INT32:
		return Int32
	case descriptorpb.FieldDescriptorProto_TYPE_SINT32:
		return Int32
	case descriptorpb.FieldDescriptorProto_TYPE_SFIXED32:
		return Int32
	case descriptorpb.FieldDescriptorProto_TYPE_UINT32:
		return UInt32
	case descriptorpb.FieldDescriptorProto_TYPE_FIXED32:
		return UInt32
	case descriptorpb.FieldDescriptorProto_TYPE_BOOL:
		return Bool
	case descriptorpb.FieldDescriptorProto_TYPE_STRING:
		return String
	case descriptorpb.FieldDescriptorProto_TYPE_BYTES:
		return String
		// case descriptorpb.FieldDescriptorProto_TYPE_MESSAGE:
		// return Message
		// case descriptorpb.FieldDescriptorProto_TYPE_ENUM:
	}
	panic(errors.New("unknown or unsupported field type: " + fieldType.String()))
}
