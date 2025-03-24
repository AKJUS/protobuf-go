// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cmd/protoc-gen-go/testdata/protoeditions/legacy_enum.proto

package protoeditions

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/gofeaturespb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

// EnumTypeWithLegacyUnmarshalJSON comment.
type EnumTypeWithLegacyUnmarshalJSON int32

const (
	// EnumTypeWithLegacyUnmarshalJSON_ONE comment.
	EnumTypeWithLegacyUnmarshalJSON_FIRST EnumTypeWithLegacyUnmarshalJSON = 1
	// EnumTypeWithLegacyUnmarshalJSON_TWO comment.
	EnumTypeWithLegacyUnmarshalJSON_SECOND EnumTypeWithLegacyUnmarshalJSON = 2
)

// Enum value maps for EnumTypeWithLegacyUnmarshalJSON.
var (
	EnumTypeWithLegacyUnmarshalJSON_name = map[int32]string{
		1: "FIRST",
		2: "SECOND",
	}
	EnumTypeWithLegacyUnmarshalJSON_value = map[string]int32{
		"FIRST":  1,
		"SECOND": 2,
	}
)

func (x EnumTypeWithLegacyUnmarshalJSON) Enum() *EnumTypeWithLegacyUnmarshalJSON {
	p := new(EnumTypeWithLegacyUnmarshalJSON)
	*p = x
	return p
}

func (x EnumTypeWithLegacyUnmarshalJSON) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EnumTypeWithLegacyUnmarshalJSON) Descriptor() protoreflect.EnumDescriptor {
	return file_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto_enumTypes[0].Descriptor()
}

func (EnumTypeWithLegacyUnmarshalJSON) Type() protoreflect.EnumType {
	return &file_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto_enumTypes[0]
}

func (x EnumTypeWithLegacyUnmarshalJSON) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *EnumTypeWithLegacyUnmarshalJSON) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = EnumTypeWithLegacyUnmarshalJSON(num)
	return nil
}

// Deprecated: Use EnumTypeWithLegacyUnmarshalJSON.Descriptor instead.
func (EnumTypeWithLegacyUnmarshalJSON) EnumDescriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto_rawDescGZIP(), []int{0}
}

type EnumWithoutUnmarshalJSON int32

const (
	EnumWithoutUnmarshalJSON_WITHOUT_UNMARSHAL_JSON_FOO EnumWithoutUnmarshalJSON = 0
	EnumWithoutUnmarshalJSON_WITHOUT_UNMARSHAL_JSON_BAR EnumWithoutUnmarshalJSON = 1
	EnumWithoutUnmarshalJSON_WITHOUT_UNMARSHAL_JSON_BAZ EnumWithoutUnmarshalJSON = 2
)

// Enum value maps for EnumWithoutUnmarshalJSON.
var (
	EnumWithoutUnmarshalJSON_name = map[int32]string{
		0: "WITHOUT_UNMARSHAL_JSON_FOO",
		1: "WITHOUT_UNMARSHAL_JSON_BAR",
		2: "WITHOUT_UNMARSHAL_JSON_BAZ",
	}
	EnumWithoutUnmarshalJSON_value = map[string]int32{
		"WITHOUT_UNMARSHAL_JSON_FOO": 0,
		"WITHOUT_UNMARSHAL_JSON_BAR": 1,
		"WITHOUT_UNMARSHAL_JSON_BAZ": 2,
	}
)

func (x EnumWithoutUnmarshalJSON) Enum() *EnumWithoutUnmarshalJSON {
	p := new(EnumWithoutUnmarshalJSON)
	*p = x
	return p
}

func (x EnumWithoutUnmarshalJSON) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EnumWithoutUnmarshalJSON) Descriptor() protoreflect.EnumDescriptor {
	return file_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto_enumTypes[1].Descriptor()
}

func (EnumWithoutUnmarshalJSON) Type() protoreflect.EnumType {
	return &file_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto_enumTypes[1]
}

func (x EnumWithoutUnmarshalJSON) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EnumWithoutUnmarshalJSON.Descriptor instead.
func (EnumWithoutUnmarshalJSON) EnumDescriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto_rawDescGZIP(), []int{1}
}

// NestedEnumType1A comment.
type ContainerForNestedEnum_NestedEnum int32

const (
	// NestedEnum_VALUE comment.
	ContainerForNestedEnum_VALUE ContainerForNestedEnum_NestedEnum = 0
)

// Enum value maps for ContainerForNestedEnum_NestedEnum.
var (
	ContainerForNestedEnum_NestedEnum_name = map[int32]string{
		0: "VALUE",
	}
	ContainerForNestedEnum_NestedEnum_value = map[string]int32{
		"VALUE": 0,
	}
)

func (x ContainerForNestedEnum_NestedEnum) Enum() *ContainerForNestedEnum_NestedEnum {
	p := new(ContainerForNestedEnum_NestedEnum)
	*p = x
	return p
}

func (x ContainerForNestedEnum_NestedEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ContainerForNestedEnum_NestedEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto_enumTypes[2].Descriptor()
}

func (ContainerForNestedEnum_NestedEnum) Type() protoreflect.EnumType {
	return &file_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto_enumTypes[2]
}

func (x ContainerForNestedEnum_NestedEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *ContainerForNestedEnum_NestedEnum) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = ContainerForNestedEnum_NestedEnum(num)
	return nil
}

// Deprecated: Use ContainerForNestedEnum_NestedEnum.Descriptor instead.
func (ContainerForNestedEnum_NestedEnum) EnumDescriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto_rawDescGZIP(), []int{0, 0}
}

type ContainerForNestedEnum struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ContainerForNestedEnum) Reset() {
	*x = ContainerForNestedEnum{}
	mi := &file_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ContainerForNestedEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContainerForNestedEnum) ProtoMessage() {}

func (x *ContainerForNestedEnum) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContainerForNestedEnum.ProtoReflect.Descriptor instead.
func (*ContainerForNestedEnum) Descriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto_rawDescGZIP(), []int{0}
}

var File_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto protoreflect.FileDescriptor

const file_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto_rawDesc = "" +
	"\n" +
	":cmd/protoc-gen-go/testdata/protoeditions/legacy_enum.proto\x12\x1cgoproto.protoc.protoeditions\x1a!google/protobuf/go_features.proto\"1\n" +
	"\x16ContainerForNestedEnum\"\x17\n" +
	"\n" +
	"NestedEnum\x12\t\n" +
	"\x05VALUE\x10\x00*>\n" +
	"\x1fEnumTypeWithLegacyUnmarshalJSON\x12\t\n" +
	"\x05FIRST\x10\x01\x12\n" +
	"\n" +
	"\x06SECOND\x10\x02\x1a\x04:\x02\x10\x02*\x83\x01\n" +
	"\x18EnumWithoutUnmarshalJSON\x12\x1e\n" +
	"\x1aWITHOUT_UNMARSHAL_JSON_FOO\x10\x00\x12\x1e\n" +
	"\x1aWITHOUT_UNMARSHAL_JSON_BAR\x10\x01\x12\x1e\n" +
	"\x1aWITHOUT_UNMARSHAL_JSON_BAZ\x10\x02\x1a\a:\x05\xd2>\x02\b\x00BMZCgoogle.golang.org/protobuf/cmd/protoc-gen-go/testdata/protoeditions\x92\x03\x05\xd2>\x02\b\x01b\beditionsp\xe8\a"

var (
	file_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto_rawDescOnce sync.Once
	file_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto_rawDescData []byte
)

func file_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto_rawDescGZIP() []byte {
	file_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto_rawDescOnce.Do(func() {
		file_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto_rawDesc), len(file_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto_rawDesc)))
	})
	return file_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto_rawDescData
}

var file_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto_goTypes = []any{
	(EnumTypeWithLegacyUnmarshalJSON)(0),   // 0: goproto.protoc.protoeditions.EnumTypeWithLegacyUnmarshalJSON
	(EnumWithoutUnmarshalJSON)(0),          // 1: goproto.protoc.protoeditions.EnumWithoutUnmarshalJSON
	(ContainerForNestedEnum_NestedEnum)(0), // 2: goproto.protoc.protoeditions.ContainerForNestedEnum.NestedEnum
	(*ContainerForNestedEnum)(nil),         // 3: goproto.protoc.protoeditions.ContainerForNestedEnum
}
var file_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto_init() }
func file_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto_init() {
	if File_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto_rawDesc), len(file_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto_rawDesc)),
			NumEnums:      3,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto_goTypes,
		DependencyIndexes: file_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto_depIdxs,
		EnumInfos:         file_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto_enumTypes,
		MessageInfos:      file_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto_msgTypes,
	}.Build()
	File_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto = out.File
	file_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto_goTypes = nil
	file_cmd_protoc_gen_go_testdata_protoeditions_legacy_enum_proto_depIdxs = nil
}
