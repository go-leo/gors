// Copyright 2021 Google LLC.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: cmd/protoc-gen-gors/examples/tests/noannotations/message.proto

package message

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Label string `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_protoc_gen_gors_examples_tests_noannotations_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_protoc_gen_gors_examples_tests_noannotations_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_gors_examples_tests_noannotations_message_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Message) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

var File_cmd_protoc_gen_gors_examples_tests_noannotations_message_proto protoreflect.FileDescriptor

var file_cmd_protoc_gen_gors_examples_tests_noannotations_message_proto_rawDesc = []byte{
	0x0a, 0x3e, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x67, 0x6f, 0x72, 0x73, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x74,
	0x65, 0x73, 0x74, 0x73, 0x2f, 0x6e, 0x6f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x6e, 0x6f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2f,
	0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x32,
	0x96, 0x01, 0x0a, 0x0a, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x31, 0x12, 0x87,
	0x01, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x27, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x6e, 0x6f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x27, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x73, 0x2e, 0x6e, 0x6f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x3a, 0x01, 0x2a, 0x32, 0x19, 0x2f,
	0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x7b, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x32, 0x71, 0x0a, 0x0a, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x69, 0x6e, 0x67, 0x32, 0x12, 0x63, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x27, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e,
	0x6e, 0x6f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x1a, 0x27, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x6e, 0x6f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x42, 0x63, 0x5a, 0x61, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69,
	0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f,
	0x6e, 0x6f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cmd_protoc_gen_gors_examples_tests_noannotations_message_proto_rawDescOnce sync.Once
	file_cmd_protoc_gen_gors_examples_tests_noannotations_message_proto_rawDescData = file_cmd_protoc_gen_gors_examples_tests_noannotations_message_proto_rawDesc
)

func file_cmd_protoc_gen_gors_examples_tests_noannotations_message_proto_rawDescGZIP() []byte {
	file_cmd_protoc_gen_gors_examples_tests_noannotations_message_proto_rawDescOnce.Do(func() {
		file_cmd_protoc_gen_gors_examples_tests_noannotations_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmd_protoc_gen_gors_examples_tests_noannotations_message_proto_rawDescData)
	})
	return file_cmd_protoc_gen_gors_examples_tests_noannotations_message_proto_rawDescData
}

var file_cmd_protoc_gen_gors_examples_tests_noannotations_message_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_cmd_protoc_gen_gors_examples_tests_noannotations_message_proto_goTypes = []interface{}{
	(*Message)(nil), // 0: tests.noannotations.message.v1.Message
}
var file_cmd_protoc_gen_gors_examples_tests_noannotations_message_proto_depIdxs = []int32{
	0, // 0: tests.noannotations.message.v1.Messaging1.UpdateMessage:input_type -> tests.noannotations.message.v1.Message
	0, // 1: tests.noannotations.message.v1.Messaging2.UpdateMessage:input_type -> tests.noannotations.message.v1.Message
	0, // 2: tests.noannotations.message.v1.Messaging1.UpdateMessage:output_type -> tests.noannotations.message.v1.Message
	0, // 3: tests.noannotations.message.v1.Messaging2.UpdateMessage:output_type -> tests.noannotations.message.v1.Message
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cmd_protoc_gen_gors_examples_tests_noannotations_message_proto_init() }
func file_cmd_protoc_gen_gors_examples_tests_noannotations_message_proto_init() {
	if File_cmd_protoc_gen_gors_examples_tests_noannotations_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cmd_protoc_gen_gors_examples_tests_noannotations_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cmd_protoc_gen_gors_examples_tests_noannotations_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_cmd_protoc_gen_gors_examples_tests_noannotations_message_proto_goTypes,
		DependencyIndexes: file_cmd_protoc_gen_gors_examples_tests_noannotations_message_proto_depIdxs,
		MessageInfos:      file_cmd_protoc_gen_gors_examples_tests_noannotations_message_proto_msgTypes,
	}.Build()
	File_cmd_protoc_gen_gors_examples_tests_noannotations_message_proto = out.File
	file_cmd_protoc_gen_gors_examples_tests_noannotations_message_proto_rawDesc = nil
	file_cmd_protoc_gen_gors_examples_tests_noannotations_message_proto_goTypes = nil
	file_cmd_protoc_gen_gors_examples_tests_noannotations_message_proto_depIdxs = nil
}
