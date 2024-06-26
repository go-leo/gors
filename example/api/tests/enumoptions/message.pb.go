// Copyright 2020 Google LLC.
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
// source: example/api/tests/enumoptions/message.proto

package enumoptions

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

type Kind int32

const (
	Kind_UNKNOWN_KIND Kind = 0
	Kind_KIND_1       Kind = 1
	Kind_KIND_2       Kind = 2
)

// Enum value maps for Kind.
var (
	Kind_name = map[int32]string{
		0: "UNKNOWN_KIND",
		1: "KIND_1",
		2: "KIND_2",
	}
	Kind_value = map[string]int32{
		"UNKNOWN_KIND": 0,
		"KIND_1":       1,
		"KIND_2":       2,
	}
)

func (x Kind) Enum() *Kind {
	p := new(Kind)
	*p = x
	return p
}

func (x Kind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Kind) Descriptor() protoreflect.EnumDescriptor {
	return file_example_api_tests_enumoptions_message_proto_enumTypes[0].Descriptor()
}

func (Kind) Type() protoreflect.EnumType {
	return &file_example_api_tests_enumoptions_message_proto_enumTypes[0]
}

func (x Kind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Kind.Descriptor instead.
func (Kind) EnumDescriptor() ([]byte, []int) {
	return file_example_api_tests_enumoptions_message_proto_rawDescGZIP(), []int{0}
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind      Kind   `protobuf:"varint,1,opt,name=kind,proto3,enum=tests.enumoptions.message.v1.Kind" json:"kind,omitempty"`
	MessageId string `protobuf:"bytes,2,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	BodyText  string `protobuf:"bytes,3,opt,name=body_text,json=bodyText,proto3" json:"body_text,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_api_tests_enumoptions_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_example_api_tests_enumoptions_message_proto_msgTypes[0]
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
	return file_example_api_tests_enumoptions_message_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetKind() Kind {
	if x != nil {
		return x.Kind
	}
	return Kind_UNKNOWN_KIND
}

func (x *Message) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

func (x *Message) GetBodyText() string {
	if x != nil {
		return x.BodyText
	}
	return ""
}

var File_example_api_tests_enumoptions_message_proto protoreflect.FileDescriptor

var file_example_api_tests_enumoptions_message_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x65,
	0x73, 0x74, 0x73, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x74,
	0x65, 0x73, 0x74, 0x73, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7d, 0x0a, 0x07, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x36, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x22, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x4b, 0x69, 0x6e, 0x64, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x62,
	0x6f, 0x64, 0x79, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x62, 0x6f, 0x64, 0x79, 0x54, 0x65, 0x78, 0x74, 0x2a, 0x30, 0x0a, 0x04, 0x4b, 0x69, 0x6e, 0x64,
	0x12, 0x10, 0x0a, 0x0c, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x4b, 0x49, 0x4e, 0x44,
	0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x31, 0x10, 0x01, 0x12, 0x0a,
	0x0a, 0x06, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x32, 0x10, 0x02, 0x32, 0x99, 0x01, 0x0a, 0x09, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x12, 0x8b, 0x01, 0x0a, 0x0d, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x25, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x73, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x1a, 0x25, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x26,
	0x3a, 0x09, 0x62, 0x6f, 0x64, 0x79, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x22, 0x19, 0x2f, 0x76, 0x31,
	0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x7b, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x45, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x6c, 0x65, 0x6f, 0x2f, 0x67, 0x6f, 0x72, 0x73,
	0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x65, 0x73,
	0x74, 0x73, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76,
	0x31, 0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_example_api_tests_enumoptions_message_proto_rawDescOnce sync.Once
	file_example_api_tests_enumoptions_message_proto_rawDescData = file_example_api_tests_enumoptions_message_proto_rawDesc
)

func file_example_api_tests_enumoptions_message_proto_rawDescGZIP() []byte {
	file_example_api_tests_enumoptions_message_proto_rawDescOnce.Do(func() {
		file_example_api_tests_enumoptions_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_example_api_tests_enumoptions_message_proto_rawDescData)
	})
	return file_example_api_tests_enumoptions_message_proto_rawDescData
}

var file_example_api_tests_enumoptions_message_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_example_api_tests_enumoptions_message_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_example_api_tests_enumoptions_message_proto_goTypes = []interface{}{
	(Kind)(0),       // 0: tests.enumoptions.message.v1.Kind
	(*Message)(nil), // 1: tests.enumoptions.message.v1.Message
}
var file_example_api_tests_enumoptions_message_proto_depIdxs = []int32{
	0, // 0: tests.enumoptions.message.v1.Message.kind:type_name -> tests.enumoptions.message.v1.Kind
	1, // 1: tests.enumoptions.message.v1.Messaging.CreateMessage:input_type -> tests.enumoptions.message.v1.Message
	1, // 2: tests.enumoptions.message.v1.Messaging.CreateMessage:output_type -> tests.enumoptions.message.v1.Message
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_example_api_tests_enumoptions_message_proto_init() }
func file_example_api_tests_enumoptions_message_proto_init() {
	if File_example_api_tests_enumoptions_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_example_api_tests_enumoptions_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_example_api_tests_enumoptions_message_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_example_api_tests_enumoptions_message_proto_goTypes,
		DependencyIndexes: file_example_api_tests_enumoptions_message_proto_depIdxs,
		EnumInfos:         file_example_api_tests_enumoptions_message_proto_enumTypes,
		MessageInfos:      file_example_api_tests_enumoptions_message_proto_msgTypes,
	}.Build()
	File_example_api_tests_enumoptions_message_proto = out.File
	file_example_api_tests_enumoptions_message_proto_rawDesc = nil
	file_example_api_tests_enumoptions_message_proto_goTypes = nil
	file_example_api_tests_enumoptions_message_proto_depIdxs = nil
}
