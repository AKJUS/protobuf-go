// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/testprotos/fuzz/fuzz.proto

package fuzz

import (
	test "google.golang.org/protobuf/internal/testprotos/test"
	test3 "google.golang.org/protobuf/internal/testprotos/test3"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

// Fuzz is a container for every message we want to make available to the
// fuzzer.
type Fuzz struct {
	state                   protoimpl.MessageState        `protogen:"open.v1"`
	TestAllTypes            *test.TestAllTypes            `protobuf:"bytes,1,opt,name=test_all_types,json=testAllTypes" json:"test_all_types,omitempty"`
	TestAllExtensions       *test.TestAllExtensions       `protobuf:"bytes,2,opt,name=test_all_extensions,json=testAllExtensions" json:"test_all_extensions,omitempty"`
	TestRequired            *test.TestRequired            `protobuf:"bytes,3,opt,name=test_required,json=testRequired" json:"test_required,omitempty"`
	TestRequiredForeign     *test.TestRequiredForeign     `protobuf:"bytes,4,opt,name=test_required_foreign,json=testRequiredForeign" json:"test_required_foreign,omitempty"`
	TestRequiredGroupFields *test.TestRequiredGroupFields `protobuf:"bytes,5,opt,name=test_required_group_fields,json=testRequiredGroupFields" json:"test_required_group_fields,omitempty"`
	TestPackedTypes         *test.TestPackedTypes         `protobuf:"bytes,6,opt,name=test_packed_types,json=testPackedTypes" json:"test_packed_types,omitempty"`
	TestPackedExtensions    *test.TestPackedExtensions    `protobuf:"bytes,7,opt,name=test_packed_extensions,json=testPackedExtensions" json:"test_packed_extensions,omitempty"`
	TestAllTypes3           *test3.TestAllTypes           `protobuf:"bytes,8,opt,name=test_all_types3,json=testAllTypes3" json:"test_all_types3,omitempty"`
	unknownFields           protoimpl.UnknownFields
	sizeCache               protoimpl.SizeCache
}

func (x *Fuzz) Reset() {
	*x = Fuzz{}
	mi := &file_internal_testprotos_fuzz_fuzz_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Fuzz) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Fuzz) ProtoMessage() {}

func (x *Fuzz) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_fuzz_fuzz_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Fuzz.ProtoReflect.Descriptor instead.
func (*Fuzz) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_fuzz_fuzz_proto_rawDescGZIP(), []int{0}
}

func (x *Fuzz) GetTestAllTypes() *test.TestAllTypes {
	if x != nil {
		return x.TestAllTypes
	}
	return nil
}

func (x *Fuzz) GetTestAllExtensions() *test.TestAllExtensions {
	if x != nil {
		return x.TestAllExtensions
	}
	return nil
}

func (x *Fuzz) GetTestRequired() *test.TestRequired {
	if x != nil {
		return x.TestRequired
	}
	return nil
}

func (x *Fuzz) GetTestRequiredForeign() *test.TestRequiredForeign {
	if x != nil {
		return x.TestRequiredForeign
	}
	return nil
}

func (x *Fuzz) GetTestRequiredGroupFields() *test.TestRequiredGroupFields {
	if x != nil {
		return x.TestRequiredGroupFields
	}
	return nil
}

func (x *Fuzz) GetTestPackedTypes() *test.TestPackedTypes {
	if x != nil {
		return x.TestPackedTypes
	}
	return nil
}

func (x *Fuzz) GetTestPackedExtensions() *test.TestPackedExtensions {
	if x != nil {
		return x.TestPackedExtensions
	}
	return nil
}

func (x *Fuzz) GetTestAllTypes3() *test3.TestAllTypes {
	if x != nil {
		return x.TestAllTypes3
	}
	return nil
}

var File_internal_testprotos_fuzz_fuzz_proto protoreflect.FileDescriptor

const file_internal_testprotos_fuzz_fuzz_proto_rawDesc = "" +
	"\n" +
	"#internal/testprotos/fuzz/fuzz.proto\x12\x12goproto.proto.fuzz\x1a#internal/testprotos/test/test.proto\x1a$internal/testprotos/test3/test.proto\"\xaf\x05\n" +
	"\x04Fuzz\x12F\n" +
	"\x0etest_all_types\x18\x01 \x01(\v2 .goproto.proto.test.TestAllTypesR\ftestAllTypes\x12U\n" +
	"\x13test_all_extensions\x18\x02 \x01(\v2%.goproto.proto.test.TestAllExtensionsR\x11testAllExtensions\x12E\n" +
	"\rtest_required\x18\x03 \x01(\v2 .goproto.proto.test.TestRequiredR\ftestRequired\x12[\n" +
	"\x15test_required_foreign\x18\x04 \x01(\v2'.goproto.proto.test.TestRequiredForeignR\x13testRequiredForeign\x12h\n" +
	"\x1atest_required_group_fields\x18\x05 \x01(\v2+.goproto.proto.test.TestRequiredGroupFieldsR\x17testRequiredGroupFields\x12O\n" +
	"\x11test_packed_types\x18\x06 \x01(\v2#.goproto.proto.test.TestPackedTypesR\x0ftestPackedTypes\x12^\n" +
	"\x16test_packed_extensions\x18\a \x01(\v2(.goproto.proto.test.TestPackedExtensionsR\x14testPackedExtensions\x12I\n" +
	"\x0ftest_all_types3\x18\b \x01(\v2!.goproto.proto.test3.TestAllTypesR\rtestAllTypes3B5Z3google.golang.org/protobuf/internal/testprotos/fuzz"

var (
	file_internal_testprotos_fuzz_fuzz_proto_rawDescOnce sync.Once
	file_internal_testprotos_fuzz_fuzz_proto_rawDescData []byte
)

func file_internal_testprotos_fuzz_fuzz_proto_rawDescGZIP() []byte {
	file_internal_testprotos_fuzz_fuzz_proto_rawDescOnce.Do(func() {
		file_internal_testprotos_fuzz_fuzz_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_internal_testprotos_fuzz_fuzz_proto_rawDesc), len(file_internal_testprotos_fuzz_fuzz_proto_rawDesc)))
	})
	return file_internal_testprotos_fuzz_fuzz_proto_rawDescData
}

var file_internal_testprotos_fuzz_fuzz_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_internal_testprotos_fuzz_fuzz_proto_goTypes = []any{
	(*Fuzz)(nil),                         // 0: goproto.proto.fuzz.Fuzz
	(*test.TestAllTypes)(nil),            // 1: goproto.proto.test.TestAllTypes
	(*test.TestAllExtensions)(nil),       // 2: goproto.proto.test.TestAllExtensions
	(*test.TestRequired)(nil),            // 3: goproto.proto.test.TestRequired
	(*test.TestRequiredForeign)(nil),     // 4: goproto.proto.test.TestRequiredForeign
	(*test.TestRequiredGroupFields)(nil), // 5: goproto.proto.test.TestRequiredGroupFields
	(*test.TestPackedTypes)(nil),         // 6: goproto.proto.test.TestPackedTypes
	(*test.TestPackedExtensions)(nil),    // 7: goproto.proto.test.TestPackedExtensions
	(*test3.TestAllTypes)(nil),           // 8: goproto.proto.test3.TestAllTypes
}
var file_internal_testprotos_fuzz_fuzz_proto_depIdxs = []int32{
	1, // 0: goproto.proto.fuzz.Fuzz.test_all_types:type_name -> goproto.proto.test.TestAllTypes
	2, // 1: goproto.proto.fuzz.Fuzz.test_all_extensions:type_name -> goproto.proto.test.TestAllExtensions
	3, // 2: goproto.proto.fuzz.Fuzz.test_required:type_name -> goproto.proto.test.TestRequired
	4, // 3: goproto.proto.fuzz.Fuzz.test_required_foreign:type_name -> goproto.proto.test.TestRequiredForeign
	5, // 4: goproto.proto.fuzz.Fuzz.test_required_group_fields:type_name -> goproto.proto.test.TestRequiredGroupFields
	6, // 5: goproto.proto.fuzz.Fuzz.test_packed_types:type_name -> goproto.proto.test.TestPackedTypes
	7, // 6: goproto.proto.fuzz.Fuzz.test_packed_extensions:type_name -> goproto.proto.test.TestPackedExtensions
	8, // 7: goproto.proto.fuzz.Fuzz.test_all_types3:type_name -> goproto.proto.test3.TestAllTypes
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_internal_testprotos_fuzz_fuzz_proto_init() }
func file_internal_testprotos_fuzz_fuzz_proto_init() {
	if File_internal_testprotos_fuzz_fuzz_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_internal_testprotos_fuzz_fuzz_proto_rawDesc), len(file_internal_testprotos_fuzz_fuzz_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_testprotos_fuzz_fuzz_proto_goTypes,
		DependencyIndexes: file_internal_testprotos_fuzz_fuzz_proto_depIdxs,
		MessageInfos:      file_internal_testprotos_fuzz_fuzz_proto_msgTypes,
	}.Build()
	File_internal_testprotos_fuzz_fuzz_proto = out.File
	file_internal_testprotos_fuzz_fuzz_proto_goTypes = nil
	file_internal_testprotos_fuzz_fuzz_proto_depIdxs = nil
}
