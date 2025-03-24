// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/testprotos/testeditions/testeditions_opaque/test_import.opaque.proto

package testeditions_opaque

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/gofeaturespb"
	reflect "reflect"
	unsafe "unsafe"
)

type ImportEnum int32

const (
	ImportEnum_IMPORT_ZERO ImportEnum = 0
)

// Enum value maps for ImportEnum.
var (
	ImportEnum_name = map[int32]string{
		0: "IMPORT_ZERO",
	}
	ImportEnum_value = map[string]int32{
		"IMPORT_ZERO": 0,
	}
)

func (x ImportEnum) Enum() *ImportEnum {
	p := new(ImportEnum)
	*p = x
	return p
}

func (x ImportEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ImportEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_testprotos_testeditions_testeditions_opaque_test_import_opaque_proto_enumTypes[0].Descriptor()
}

func (ImportEnum) Type() protoreflect.EnumType {
	return &file_internal_testprotos_testeditions_testeditions_opaque_test_import_opaque_proto_enumTypes[0]
}

func (x ImportEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

type ImportMessage struct {
	state         protoimpl.MessageState `protogen:"opaque.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ImportMessage) Reset() {
	*x = ImportMessage{}
	mi := &file_internal_testprotos_testeditions_testeditions_opaque_test_import_opaque_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ImportMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportMessage) ProtoMessage() {}

func (x *ImportMessage) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_testeditions_testeditions_opaque_test_import_opaque_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

type ImportMessage_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

}

func (b0 ImportMessage_builder) Build() *ImportMessage {
	m0 := &ImportMessage{}
	b, x := &b0, m0
	_, _ = b, x
	return m0
}

var File_internal_testprotos_testeditions_testeditions_opaque_test_import_opaque_proto protoreflect.FileDescriptor

const file_internal_testprotos_testeditions_testeditions_opaque_test_import_opaque_proto_rawDesc = "" +
	"\n" +
	"Minternal/testprotos/testeditions/testeditions_opaque/test_import.opaque.proto\x12!opaque.goproto.proto.testeditions\x1a!google/protobuf/go_features.proto\"\x0f\n" +
	"\rImportMessage*\x1d\n" +
	"\n" +
	"ImportEnum\x12\x0f\n" +
	"\vIMPORT_ZERO\x10\x00BYZOgoogle.golang.org/protobuf/internal/testprotos/testeditions/testeditions_opaque\x92\x03\x05\xd2>\x02\x10\x03b\beditionsp\xe8\a"

var file_internal_testprotos_testeditions_testeditions_opaque_test_import_opaque_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_internal_testprotos_testeditions_testeditions_opaque_test_import_opaque_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_internal_testprotos_testeditions_testeditions_opaque_test_import_opaque_proto_goTypes = []any{
	(ImportEnum)(0),       // 0: opaque.goproto.proto.testeditions.ImportEnum
	(*ImportMessage)(nil), // 1: opaque.goproto.proto.testeditions.ImportMessage
}
var file_internal_testprotos_testeditions_testeditions_opaque_test_import_opaque_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_internal_testprotos_testeditions_testeditions_opaque_test_import_opaque_proto_init()
}
func file_internal_testprotos_testeditions_testeditions_opaque_test_import_opaque_proto_init() {
	if File_internal_testprotos_testeditions_testeditions_opaque_test_import_opaque_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_internal_testprotos_testeditions_testeditions_opaque_test_import_opaque_proto_rawDesc), len(file_internal_testprotos_testeditions_testeditions_opaque_test_import_opaque_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_testprotos_testeditions_testeditions_opaque_test_import_opaque_proto_goTypes,
		DependencyIndexes: file_internal_testprotos_testeditions_testeditions_opaque_test_import_opaque_proto_depIdxs,
		EnumInfos:         file_internal_testprotos_testeditions_testeditions_opaque_test_import_opaque_proto_enumTypes,
		MessageInfos:      file_internal_testprotos_testeditions_testeditions_opaque_test_import_opaque_proto_msgTypes,
	}.Build()
	File_internal_testprotos_testeditions_testeditions_opaque_test_import_opaque_proto = out.File
	file_internal_testprotos_testeditions_testeditions_opaque_test_import_opaque_proto_goTypes = nil
	file_internal_testprotos_testeditions_testeditions_opaque_test_import_opaque_proto_depIdxs = nil
}
