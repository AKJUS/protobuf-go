// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cmd/protoc-gen-go/testdata/extensions/extra/extra.proto

package extra

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

type ExtraMessage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Data          []byte                 `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExtraMessage) Reset() {
	*x = ExtraMessage{}
	mi := &file_cmd_protoc_gen_go_testdata_extensions_extra_extra_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExtraMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtraMessage) ProtoMessage() {}

func (x *ExtraMessage) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_protoc_gen_go_testdata_extensions_extra_extra_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtraMessage.ProtoReflect.Descriptor instead.
func (*ExtraMessage) Descriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_go_testdata_extensions_extra_extra_proto_rawDescGZIP(), []int{0}
}

func (x *ExtraMessage) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_cmd_protoc_gen_go_testdata_extensions_extra_extra_proto protoreflect.FileDescriptor

const file_cmd_protoc_gen_go_testdata_extensions_extra_extra_proto_rawDesc = "" +
	"\n" +
	"7cmd/protoc-gen-go/testdata/extensions/extra/extra.proto\x12\x1egoproto.protoc.extension.extra\"\"\n" +
	"\fExtraMessage\x12\x12\n" +
	"\x04data\x18\x01 \x01(\fR\x04dataBHZFgoogle.golang.org/protobuf/cmd/protoc-gen-go/testdata/extensions/extra"

var (
	file_cmd_protoc_gen_go_testdata_extensions_extra_extra_proto_rawDescOnce sync.Once
	file_cmd_protoc_gen_go_testdata_extensions_extra_extra_proto_rawDescData []byte
)

func file_cmd_protoc_gen_go_testdata_extensions_extra_extra_proto_rawDescGZIP() []byte {
	file_cmd_protoc_gen_go_testdata_extensions_extra_extra_proto_rawDescOnce.Do(func() {
		file_cmd_protoc_gen_go_testdata_extensions_extra_extra_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_cmd_protoc_gen_go_testdata_extensions_extra_extra_proto_rawDesc), len(file_cmd_protoc_gen_go_testdata_extensions_extra_extra_proto_rawDesc)))
	})
	return file_cmd_protoc_gen_go_testdata_extensions_extra_extra_proto_rawDescData
}

var file_cmd_protoc_gen_go_testdata_extensions_extra_extra_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_cmd_protoc_gen_go_testdata_extensions_extra_extra_proto_goTypes = []any{
	(*ExtraMessage)(nil), // 0: goproto.protoc.extension.extra.ExtraMessage
}
var file_cmd_protoc_gen_go_testdata_extensions_extra_extra_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cmd_protoc_gen_go_testdata_extensions_extra_extra_proto_init() }
func file_cmd_protoc_gen_go_testdata_extensions_extra_extra_proto_init() {
	if File_cmd_protoc_gen_go_testdata_extensions_extra_extra_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_cmd_protoc_gen_go_testdata_extensions_extra_extra_proto_rawDesc), len(file_cmd_protoc_gen_go_testdata_extensions_extra_extra_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cmd_protoc_gen_go_testdata_extensions_extra_extra_proto_goTypes,
		DependencyIndexes: file_cmd_protoc_gen_go_testdata_extensions_extra_extra_proto_depIdxs,
		MessageInfos:      file_cmd_protoc_gen_go_testdata_extensions_extra_extra_proto_msgTypes,
	}.Build()
	File_cmd_protoc_gen_go_testdata_extensions_extra_extra_proto = out.File
	file_cmd_protoc_gen_go_testdata_extensions_extra_extra_proto_goTypes = nil
	file_cmd_protoc_gen_go_testdata_extensions_extra_extra_proto_depIdxs = nil
}
