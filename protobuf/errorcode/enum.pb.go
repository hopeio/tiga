// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: lemon/protobuf/errorcode/enum.proto

package errorcode

import (
	_ "github.com/hopeio/lemon/protobuf/utils/enum"
	_ "github.com/hopeio/lemon/protobuf/utils/patch"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ErrCode int32

const (
	SUCCESS            ErrCode = 0
	Canceled           ErrCode = 1
	Unknown            ErrCode = 2
	InvalidArgument    ErrCode = 3
	DeadlineExceeded   ErrCode = 4
	NotFound           ErrCode = 5
	AlreadyExists      ErrCode = 6
	PermissionDenied   ErrCode = 7
	ResourceExhausted  ErrCode = 8
	FailedPrecondition ErrCode = 9
	Aborted            ErrCode = 10
	OutOfRange         ErrCode = 11
	Unimplemented      ErrCode = 12
	Internal           ErrCode = 13
	Unavailable        ErrCode = 14
	DataLoss           ErrCode = 15
	Unauthenticated    ErrCode = 16
	SysError           ErrCode = 10000
	DBError            ErrCode = 21000
	RowExists          ErrCode = 21001
	RedisErr           ErrCode = 22000
	IOError            ErrCode = 30000
	UploadFail         ErrCode = 30001
	UploadCheckFail    ErrCode = 30002
	UploadCheckFormat  ErrCode = 30003
	TimeTooMuch        ErrCode = 30004
	ParamInvalid       ErrCode = 30005
)

// Enum value maps for ErrCode.
var (
	ErrCode_name = map[int32]string{
		0:     "SUCCESS",
		1:     "Canceled",
		2:     "Unknown",
		3:     "InvalidArgument",
		4:     "DeadlineExceeded",
		5:     "NotFound",
		6:     "AlreadyExists",
		7:     "PermissionDenied",
		8:     "ResourceExhausted",
		9:     "FailedPrecondition",
		10:    "Aborted",
		11:    "OutOfRange",
		12:    "Unimplemented",
		13:    "Internal",
		14:    "Unavailable",
		15:    "DataLoss",
		16:    "Unauthenticated",
		10000: "SysError",
		21000: "DBError",
		21001: "RowExists",
		22000: "RedisErr",
		30000: "IOError",
		30001: "UploadFail",
		30002: "UploadCheckFail",
		30003: "UploadCheckFormat",
		30004: "TimeTooMuch",
		30005: "ParamInvalid",
	}
	ErrCode_value = map[string]int32{
		"SUCCESS":            0,
		"Canceled":           1,
		"Unknown":            2,
		"InvalidArgument":    3,
		"DeadlineExceeded":   4,
		"NotFound":           5,
		"AlreadyExists":      6,
		"PermissionDenied":   7,
		"ResourceExhausted":  8,
		"FailedPrecondition": 9,
		"Aborted":            10,
		"OutOfRange":         11,
		"Unimplemented":      12,
		"Internal":           13,
		"Unavailable":        14,
		"DataLoss":           15,
		"Unauthenticated":    16,
		"SysError":           10000,
		"DBError":            21000,
		"RowExists":          21001,
		"RedisErr":           22000,
		"IOError":            30000,
		"UploadFail":         30001,
		"UploadCheckFail":    30002,
		"UploadCheckFormat":  30003,
		"TimeTooMuch":        30004,
		"ParamInvalid":       30005,
	}
)

func (x ErrCode) Enum() *ErrCode {
	p := new(ErrCode)
	*p = x
	return p
}

func (x ErrCode) OrigString() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrCode) Descriptor() protoreflect.EnumDescriptor {
	return file_lemon_protobuf_errorcode_enum_proto_enumTypes[0].Descriptor()
}

func (ErrCode) Type() protoreflect.EnumType {
	return &file_lemon_protobuf_errorcode_enum_proto_enumTypes[0]
}

func (x ErrCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrCode.Descriptor instead.
func (ErrCode) EnumDescriptor() ([]byte, []int) {
	return file_lemon_protobuf_errorcode_enum_proto_rawDescGZIP(), []int{0}
}

var File_lemon_protobuf_errorcode_enum_proto protoreflect.FileDescriptor

var file_lemon_protobuf_errorcode_enum_proto_rawDesc = []byte{
	0x0a, 0x25, 0x70, 0x61, 0x6e, 0x64, 0x6f, 0x72, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x65, 0x6e, 0x75,
	0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x63, 0x6f,
	0x64, 0x65, 0x1a, 0x26, 0x70, 0x61, 0x6e, 0x64, 0x6f, 0x72, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x2f,
	0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x70, 0x61, 0x6e, 0x64,
	0x6f, 0x72, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x75, 0x74, 0x69,
	0x6c, 0x73, 0x2f, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2a, 0x8a, 0x08, 0x0a, 0x07, 0x45, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x17, 0x0a,
	0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x1a, 0x0a, 0x92, 0x9d, 0x20, 0x06,
	0xe6, 0x88, 0x90, 0xe5, 0x8a, 0x9f, 0x12, 0x1e, 0x0a, 0x08, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c,
	0x65, 0x64, 0x10, 0x01, 0x1a, 0x10, 0x92, 0x9d, 0x20, 0x0c, 0xe6, 0x93, 0x8d, 0xe4, 0xbd, 0x9c,
	0xe5, 0x8f, 0x96, 0xe6, 0xb6, 0x88, 0x12, 0x1d, 0x0a, 0x07, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77,
	0x6e, 0x10, 0x02, 0x1a, 0x10, 0x92, 0x9d, 0x20, 0x0c, 0xe6, 0x9c, 0xaa, 0xe7, 0x9f, 0xa5, 0xe9,
	0x94, 0x99, 0xe8, 0xaf, 0xaf, 0x12, 0x28, 0x0a, 0x0f, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x41, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x10, 0x03, 0x1a, 0x13, 0x92, 0x9d, 0x20, 0x0f,
	0xe6, 0x97, 0xa0, 0xe6, 0x95, 0x88, 0xe7, 0x9a, 0x84, 0xe5, 0x8f, 0x82, 0xe6, 0x95, 0xb0, 0x12,
	0x26, 0x0a, 0x10, 0x44, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x45, 0x78, 0x63, 0x65, 0x65,
	0x64, 0x65, 0x64, 0x10, 0x04, 0x1a, 0x10, 0x92, 0x9d, 0x20, 0x0c, 0xe5, 0x93, 0x8d, 0xe5, 0xba,
	0x94, 0xe8, 0xb6, 0x85, 0xe6, 0x97, 0xb6, 0x12, 0x1b, 0x0a, 0x08, 0x4e, 0x6f, 0x74, 0x46, 0x6f,
	0x75, 0x6e, 0x64, 0x10, 0x05, 0x1a, 0x0d, 0x92, 0x9d, 0x20, 0x09, 0xe6, 0x9c, 0xaa, 0xe5, 0x8f,
	0x91, 0xe7, 0x8e, 0xb0, 0x12, 0x23, 0x0a, 0x0d, 0x41, 0x6c, 0x72, 0x65, 0x61, 0x64, 0x79, 0x45,
	0x78, 0x69, 0x73, 0x74, 0x73, 0x10, 0x06, 0x1a, 0x10, 0x92, 0x9d, 0x20, 0x0c, 0xe5, 0xb7, 0xb2,
	0xe7, 0xbb, 0x8f, 0xe5, 0xad, 0x98, 0xe5, 0x9c, 0xa8, 0x12, 0x29, 0x0a, 0x10, 0x50, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x6e, 0x69, 0x65, 0x64, 0x10, 0x07, 0x1a,
	0x13, 0x92, 0x9d, 0x20, 0x0f, 0xe6, 0x93, 0x8d, 0xe4, 0xbd, 0x9c, 0xe6, 0x97, 0xa0, 0xe6, 0x9d,
	0x83, 0xe9, 0x99, 0x90, 0x12, 0x27, 0x0a, 0x11, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x45, 0x78, 0x68, 0x61, 0x75, 0x73, 0x74, 0x65, 0x64, 0x10, 0x08, 0x1a, 0x10, 0x92, 0x9d, 0x20,
	0x0c, 0xe8, 0xb5, 0x84, 0xe6, 0xba, 0x90, 0xe4, 0xb8, 0x8d, 0xe8, 0xb6, 0xb3, 0x12, 0x2b, 0x0a,
	0x12, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x50, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x10, 0x09, 0x1a, 0x13, 0x92, 0x9d, 0x20, 0x0f, 0xe6, 0x93, 0x8d, 0xe4, 0xbd,
	0x9c, 0xe8, 0xa2, 0xab, 0xe6, 0x8b, 0x92, 0xe7, 0xbb, 0x9d, 0x12, 0x1d, 0x0a, 0x07, 0x41, 0x62,
	0x6f, 0x72, 0x74, 0x65, 0x64, 0x10, 0x0a, 0x1a, 0x10, 0x92, 0x9d, 0x20, 0x0c, 0xe6, 0x93, 0x8d,
	0xe4, 0xbd, 0x9c, 0xe7, 0xbb, 0x88, 0xe6, 0xad, 0xa2, 0x12, 0x20, 0x0a, 0x0a, 0x4f, 0x75, 0x74,
	0x4f, 0x66, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x10, 0x0b, 0x1a, 0x10, 0x92, 0x9d, 0x20, 0x0c, 0xe8,
	0xb6, 0x85, 0xe5, 0x87, 0xba, 0xe8, 0x8c, 0x83, 0xe5, 0x9b, 0xb4, 0x12, 0x20, 0x0a, 0x0d, 0x55,
	0x6e, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x65, 0x64, 0x10, 0x0c, 0x1a, 0x0d,
	0x92, 0x9d, 0x20, 0x09, 0xe6, 0x9c, 0xaa, 0xe5, 0xae, 0x9e, 0xe7, 0x8e, 0xb0, 0x12, 0x1e, 0x0a,
	0x08, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x10, 0x0d, 0x1a, 0x10, 0x92, 0x9d, 0x20,
	0x0c, 0xe5, 0x86, 0x85, 0xe9, 0x83, 0xa8, 0xe9, 0x94, 0x99, 0xe8, 0xaf, 0xaf, 0x12, 0x24, 0x0a,
	0x0b, 0x55, 0x6e, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x10, 0x0e, 0x1a, 0x13,
	0x92, 0x9d, 0x20, 0x0f, 0xe6, 0x9c, 0x8d, 0xe5, 0x8a, 0xa1, 0xe4, 0xb8, 0x8d, 0xe5, 0x8f, 0xaf,
	0xe7, 0x94, 0xa8, 0x12, 0x1e, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x4c, 0x6f, 0x73, 0x73, 0x10,
	0x0f, 0x1a, 0x10, 0x92, 0x9d, 0x20, 0x0c, 0xe6, 0x95, 0xb0, 0xe6, 0x8d, 0xae, 0xe4, 0xb8, 0xa2,
	0xe5, 0xa4, 0xb1, 0x12, 0x28, 0x0a, 0x0f, 0x55, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x10, 0x10, 0x1a, 0x13, 0x92, 0x9d, 0x20, 0x0f, 0xe8, 0xba,
	0xab, 0xe4, 0xbb, 0xbd, 0xe6, 0x9c, 0xaa, 0xe9, 0xaa, 0x8c, 0xe8, 0xaf, 0x81, 0x12, 0x1f, 0x0a,
	0x08, 0x53, 0x79, 0x73, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x90, 0x4e, 0x1a, 0x10, 0x92, 0x9d,
	0x20, 0x0c, 0xe7, 0xb3, 0xbb, 0xe7, 0xbb, 0x9f, 0xe9, 0x94, 0x99, 0xe8, 0xaf, 0xaf, 0x12, 0x22,
	0x0a, 0x07, 0x44, 0x42, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x88, 0xa4, 0x01, 0x1a, 0x13, 0x92,
	0x9d, 0x20, 0x0f, 0xe6, 0x95, 0xb0, 0xe6, 0x8d, 0xae, 0xe5, 0xba, 0x93, 0xe9, 0x94, 0x99, 0xe8,
	0xaf, 0xaf, 0x12, 0x24, 0x0a, 0x09, 0x52, 0x6f, 0x77, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x10,
	0x89, 0xa4, 0x01, 0x1a, 0x13, 0x92, 0x9d, 0x20, 0x0f, 0xe8, 0xae, 0xb0, 0xe5, 0xbd, 0x95, 0xe5,
	0xb7, 0xb2, 0xe5, 0xad, 0x98, 0xe5, 0x9c, 0xa8, 0x12, 0x1f, 0x0a, 0x08, 0x52, 0x65, 0x64, 0x69,
	0x73, 0x45, 0x72, 0x72, 0x10, 0xf0, 0xab, 0x01, 0x1a, 0x0f, 0x92, 0x9d, 0x20, 0x0b, 0x52, 0x65,
	0x64, 0x69, 0x73, 0xe9, 0x94, 0x99, 0xe8, 0xaf, 0xaf, 0x12, 0x1b, 0x0a, 0x07, 0x49, 0x4f, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x10, 0xb0, 0xea, 0x01, 0x1a, 0x0c, 0x92, 0x9d, 0x20, 0x08, 0x69, 0x6f,
	0xe9, 0x94, 0x99, 0xe8, 0xaf, 0xaf, 0x12, 0x22, 0x0a, 0x0a, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x46, 0x61, 0x69, 0x6c, 0x10, 0xb1, 0xea, 0x01, 0x1a, 0x10, 0x92, 0x9d, 0x20, 0x0c, 0xe4, 0xb8,
	0x8a, 0xe4, 0xbc, 0xa0, 0xe5, 0xa4, 0xb1, 0xe8, 0xb4, 0xa5, 0x12, 0x2d, 0x0a, 0x0f, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x46, 0x61, 0x69, 0x6c, 0x10, 0xb2, 0xea,
	0x01, 0x1a, 0x16, 0x92, 0x9d, 0x20, 0x12, 0xe6, 0xa3, 0x80, 0xe6, 0x9f, 0xa5, 0xe6, 0x96, 0x87,
	0xe4, 0xbb, 0xb6, 0xe5, 0xa4, 0xb1, 0xe8, 0xb4, 0xa5, 0x12, 0x3b, 0x0a, 0x11, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x10, 0xb3,
	0xea, 0x01, 0x1a, 0x22, 0x92, 0x9d, 0x20, 0x1e, 0xe6, 0x96, 0x87, 0xe4, 0xbb, 0xb6, 0xe6, 0xa0,
	0xbc, 0xe5, 0xbc, 0x8f, 0xe6, 0x88, 0x96, 0xe5, 0xa4, 0xa7, 0xe5, 0xb0, 0x8f, 0xe6, 0x9c, 0x89,
	0xe9, 0x97, 0xae, 0xe9, 0xa2, 0x98, 0x12, 0x29, 0x0a, 0x0b, 0x54, 0x69, 0x6d, 0x65, 0x54, 0x6f,
	0x6f, 0x4d, 0x75, 0x63, 0x68, 0x10, 0xb4, 0xea, 0x01, 0x1a, 0x16, 0x92, 0x9d, 0x20, 0x12, 0xe5,
	0xb0, 0x9d, 0xe8, 0xaf, 0x95, 0xe6, 0xac, 0xa1, 0xe6, 0x95, 0xb0, 0xe8, 0xbf, 0x87, 0xe5, 0xa4,
	0x9a, 0x12, 0x24, 0x0a, 0x0c, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x10, 0xb5, 0xea, 0x01, 0x1a, 0x10, 0x92, 0x9d, 0x20, 0x0c, 0xe5, 0x8f, 0x82, 0xe6, 0x95,
	0xb0, 0xe9, 0x94, 0x99, 0xe8, 0xaf, 0xaf, 0x1a, 0x15, 0xd2, 0xb5, 0x03, 0x0d, 0xf2, 0x01, 0x0a,
	0x4f, 0x72, 0x69, 0x67, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0xe0, 0xa4, 0x1e, 0x00, 0x42, 0x55,
	0x0a, 0x1c, 0x78, 0x79, 0x7a, 0x2e, 0x68, 0x6f, 0x70, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x5a, 0x2c,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x6f, 0x70, 0x65, 0x69,
	0x6f, 0x2f, 0x70, 0x61, 0x6e, 0x64, 0x6f, 0x72, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x63, 0x6f, 0x64, 0x65, 0xd2, 0xb5, 0x03, 0x02,
	0x50, 0x01, 0xd0, 0x3e, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_lemon_protobuf_errorcode_enum_proto_rawDescOnce sync.Once
	file_lemon_protobuf_errorcode_enum_proto_rawDescData = file_lemon_protobuf_errorcode_enum_proto_rawDesc
)

func file_lemon_protobuf_errorcode_enum_proto_rawDescGZIP() []byte {
	file_lemon_protobuf_errorcode_enum_proto_rawDescOnce.Do(func() {
		file_lemon_protobuf_errorcode_enum_proto_rawDescData = protoimpl.X.CompressGZIP(file_lemon_protobuf_errorcode_enum_proto_rawDescData)
	})
	return file_lemon_protobuf_errorcode_enum_proto_rawDescData
}

var file_lemon_protobuf_errorcode_enum_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_lemon_protobuf_errorcode_enum_proto_goTypes = []interface{}{
	(ErrCode)(0), // 0: errorcode.ErrCode
}
var file_lemon_protobuf_errorcode_enum_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_lemon_protobuf_errorcode_enum_proto_init() }
func file_lemon_protobuf_errorcode_enum_proto_init() {
	if File_lemon_protobuf_errorcode_enum_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_lemon_protobuf_errorcode_enum_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_lemon_protobuf_errorcode_enum_proto_goTypes,
		DependencyIndexes: file_lemon_protobuf_errorcode_enum_proto_depIdxs,
		EnumInfos:         file_lemon_protobuf_errorcode_enum_proto_enumTypes,
	}.Build()
	File_lemon_protobuf_errorcode_enum_proto = out.File
	file_lemon_protobuf_errorcode_enum_proto_rawDesc = nil
	file_lemon_protobuf_errorcode_enum_proto_goTypes = nil
	file_lemon_protobuf_errorcode_enum_proto_depIdxs = nil
}
