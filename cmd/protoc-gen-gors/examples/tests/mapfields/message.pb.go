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
// source: cmd/protoc-gen-gors/examples/tests/mapfields/message.proto

package message

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AnotherMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Label string `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *AnotherMessage) Reset() {
	*x = AnotherMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_protoc_gen_gors_examples_tests_mapfields_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnotherMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnotherMessage) ProtoMessage() {}

func (x *AnotherMessage) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_protoc_gen_gors_examples_tests_mapfields_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnotherMessage.ProtoReflect.Descriptor instead.
func (*AnotherMessage) Descriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_gors_examples_tests_mapfields_message_proto_rawDescGZIP(), []int{0}
}

func (x *AnotherMessage) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AnotherMessage) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageId      string                         `protobuf:"bytes,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	AnotherMessage *AnotherMessage                `protobuf:"bytes,2,opt,name=another_message,json=anotherMessage,proto3" json:"another_message,omitempty"`
	SubMessage     *Message_SubMessage            `protobuf:"bytes,3,opt,name=sub_message,json=subMessage,proto3" json:"sub_message,omitempty"`
	StringList     []string                       `protobuf:"bytes,4,rep,name=string_list,json=stringList,proto3" json:"string_list,omitempty"`
	SubMessageList []*Message_SubMessage          `protobuf:"bytes,5,rep,name=sub_message_list,json=subMessageList,proto3" json:"sub_message_list,omitempty"`
	ObjectList     []*structpb.Struct             `protobuf:"bytes,6,rep,name=object_list,json=objectList,proto3" json:"object_list,omitempty"`
	StringsMap     map[string]string              `protobuf:"bytes,7,rep,name=strings_map,json=stringsMap,proto3" json:"strings_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	SubMessagesMap map[string]*Message_SubMessage `protobuf:"bytes,8,rep,name=sub_messages_map,json=subMessagesMap,proto3" json:"sub_messages_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ObjectsMap     map[string]*structpb.Struct    `protobuf:"bytes,9,rep,name=objects_map,json=objectsMap,proto3" json:"objects_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_protoc_gen_gors_examples_tests_mapfields_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_protoc_gen_gors_examples_tests_mapfields_message_proto_msgTypes[1]
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
	return file_cmd_protoc_gen_gors_examples_tests_mapfields_message_proto_rawDescGZIP(), []int{1}
}

func (x *Message) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

func (x *Message) GetAnotherMessage() *AnotherMessage {
	if x != nil {
		return x.AnotherMessage
	}
	return nil
}

func (x *Message) GetSubMessage() *Message_SubMessage {
	if x != nil {
		return x.SubMessage
	}
	return nil
}

func (x *Message) GetStringList() []string {
	if x != nil {
		return x.StringList
	}
	return nil
}

func (x *Message) GetSubMessageList() []*Message_SubMessage {
	if x != nil {
		return x.SubMessageList
	}
	return nil
}

func (x *Message) GetObjectList() []*structpb.Struct {
	if x != nil {
		return x.ObjectList
	}
	return nil
}

func (x *Message) GetStringsMap() map[string]string {
	if x != nil {
		return x.StringsMap
	}
	return nil
}

func (x *Message) GetSubMessagesMap() map[string]*Message_SubMessage {
	if x != nil {
		return x.SubMessagesMap
	}
	return nil
}

func (x *Message) GetObjectsMap() map[string]*structpb.Struct {
	if x != nil {
		return x.ObjectsMap
	}
	return nil
}

type Message_SubMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Label string `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *Message_SubMessage) Reset() {
	*x = Message_SubMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_protoc_gen_gors_examples_tests_mapfields_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message_SubMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message_SubMessage) ProtoMessage() {}

func (x *Message_SubMessage) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_protoc_gen_gors_examples_tests_mapfields_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message_SubMessage.ProtoReflect.Descriptor instead.
func (*Message_SubMessage) Descriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_gors_examples_tests_mapfields_message_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Message_SubMessage) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Message_SubMessage) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

var File_cmd_protoc_gen_gors_examples_tests_mapfields_message_proto protoreflect.FileDescriptor

var file_cmd_protoc_gen_gors_examples_tests_mapfields_message_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x67, 0x6f, 0x72, 0x73, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x74,
	0x65, 0x73, 0x74, 0x73, 0x2f, 0x6d, 0x61, 0x70, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x2f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x74, 0x65,
	0x73, 0x74, 0x73, 0x2e, 0x6d, 0x61, 0x70, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x36, 0x0a, 0x0e, 0x41, 0x6e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x22, 0xd0, 0x07, 0x0a,
	0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x53, 0x0a, 0x0f, 0x61, 0x6e, 0x6f, 0x74, 0x68,
	0x65, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2a, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x6d, 0x61, 0x70, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6e,
	0x6f, 0x74, 0x68, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0e, 0x61, 0x6e,
	0x6f, 0x74, 0x68, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x4f, 0x0a, 0x0b,
	0x73, 0x75, 0x62, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2e, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x6d, 0x61, 0x70, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x75, 0x62, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x0a, 0x73, 0x75, 0x62, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x58,
	0x0a, 0x10, 0x73, 0x75, 0x62, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73,
	0x2e, 0x6d, 0x61, 0x70, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x75,
	0x62, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0e, 0x73, 0x75, 0x62, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x0b, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0a, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x54, 0x0a, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x73, 0x5f, 0x6d, 0x61,
	0x70, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e,
	0x6d, 0x61, 0x70, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x73, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x73, 0x4d, 0x61, 0x70, 0x12, 0x61, 0x0a, 0x10, 0x73, 0x75, 0x62, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x08, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x37, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x6d, 0x61, 0x70, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x75, 0x62, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0e, 0x73, 0x75, 0x62,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x4d, 0x61, 0x70, 0x12, 0x54, 0x0a, 0x0b, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x33, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x6d, 0x61, 0x70, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x4d, 0x61, 0x70,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x4d, 0x61,
	0x70, 0x1a, 0x32, 0x0a, 0x0a, 0x53, 0x75, 0x62, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x1a, 0x3d, 0x0a, 0x0f, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x73,
	0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x1a, 0x71, 0x0a, 0x13, 0x53, 0x75, 0x62, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x44, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x73, 0x2e, 0x6d, 0x61, 0x70, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x53, 0x75, 0x62, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x56, 0x0a, 0x0f, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2d, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32,
	0x8c, 0x01, 0x0a, 0x09, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x12, 0x7f, 0x0a,
	0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x23,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x6d, 0x61, 0x70, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x1a, 0x23, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x6d, 0x61, 0x70, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e,
	0x3a, 0x01, 0x2a, 0x32, 0x19, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x2f, 0x7b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x5f,
	0x5a, 0x5d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x2f, 0x61, 0x70, 0x70, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e,
	0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x74, 0x65, 0x73,
	0x74, 0x73, 0x2f, 0x6d, 0x61, 0x70, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x2f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cmd_protoc_gen_gors_examples_tests_mapfields_message_proto_rawDescOnce sync.Once
	file_cmd_protoc_gen_gors_examples_tests_mapfields_message_proto_rawDescData = file_cmd_protoc_gen_gors_examples_tests_mapfields_message_proto_rawDesc
)

func file_cmd_protoc_gen_gors_examples_tests_mapfields_message_proto_rawDescGZIP() []byte {
	file_cmd_protoc_gen_gors_examples_tests_mapfields_message_proto_rawDescOnce.Do(func() {
		file_cmd_protoc_gen_gors_examples_tests_mapfields_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmd_protoc_gen_gors_examples_tests_mapfields_message_proto_rawDescData)
	})
	return file_cmd_protoc_gen_gors_examples_tests_mapfields_message_proto_rawDescData
}

var file_cmd_protoc_gen_gors_examples_tests_mapfields_message_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_cmd_protoc_gen_gors_examples_tests_mapfields_message_proto_goTypes = []interface{}{
	(*AnotherMessage)(nil),     // 0: tests.mapfields.message.v1.AnotherMessage
	(*Message)(nil),            // 1: tests.mapfields.message.v1.Message
	(*Message_SubMessage)(nil), // 2: tests.mapfields.message.v1.Message.SubMessage
	nil,                        // 3: tests.mapfields.message.v1.Message.StringsMapEntry
	nil,                        // 4: tests.mapfields.message.v1.Message.SubMessagesMapEntry
	nil,                        // 5: tests.mapfields.message.v1.Message.ObjectsMapEntry
	(*structpb.Struct)(nil),    // 6: google.protobuf.Struct
}
var file_cmd_protoc_gen_gors_examples_tests_mapfields_message_proto_depIdxs = []int32{
	0,  // 0: tests.mapfields.message.v1.Message.another_message:type_name -> tests.mapfields.message.v1.AnotherMessage
	2,  // 1: tests.mapfields.message.v1.Message.sub_message:type_name -> tests.mapfields.message.v1.Message.SubMessage
	2,  // 2: tests.mapfields.message.v1.Message.sub_message_list:type_name -> tests.mapfields.message.v1.Message.SubMessage
	6,  // 3: tests.mapfields.message.v1.Message.object_list:type_name -> google.protobuf.Struct
	3,  // 4: tests.mapfields.message.v1.Message.strings_map:type_name -> tests.mapfields.message.v1.Message.StringsMapEntry
	4,  // 5: tests.mapfields.message.v1.Message.sub_messages_map:type_name -> tests.mapfields.message.v1.Message.SubMessagesMapEntry
	5,  // 6: tests.mapfields.message.v1.Message.objects_map:type_name -> tests.mapfields.message.v1.Message.ObjectsMapEntry
	2,  // 7: tests.mapfields.message.v1.Message.SubMessagesMapEntry.value:type_name -> tests.mapfields.message.v1.Message.SubMessage
	6,  // 8: tests.mapfields.message.v1.Message.ObjectsMapEntry.value:type_name -> google.protobuf.Struct
	1,  // 9: tests.mapfields.message.v1.Messaging.UpdateMessage:input_type -> tests.mapfields.message.v1.Message
	1,  // 10: tests.mapfields.message.v1.Messaging.UpdateMessage:output_type -> tests.mapfields.message.v1.Message
	10, // [10:11] is the sub-list for method output_type
	9,  // [9:10] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_cmd_protoc_gen_gors_examples_tests_mapfields_message_proto_init() }
func file_cmd_protoc_gen_gors_examples_tests_mapfields_message_proto_init() {
	if File_cmd_protoc_gen_gors_examples_tests_mapfields_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cmd_protoc_gen_gors_examples_tests_mapfields_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnotherMessage); i {
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
		file_cmd_protoc_gen_gors_examples_tests_mapfields_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_cmd_protoc_gen_gors_examples_tests_mapfields_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message_SubMessage); i {
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
			RawDescriptor: file_cmd_protoc_gen_gors_examples_tests_mapfields_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cmd_protoc_gen_gors_examples_tests_mapfields_message_proto_goTypes,
		DependencyIndexes: file_cmd_protoc_gen_gors_examples_tests_mapfields_message_proto_depIdxs,
		MessageInfos:      file_cmd_protoc_gen_gors_examples_tests_mapfields_message_proto_msgTypes,
	}.Build()
	File_cmd_protoc_gen_gors_examples_tests_mapfields_message_proto = out.File
	file_cmd_protoc_gen_gors_examples_tests_mapfields_message_proto_rawDesc = nil
	file_cmd_protoc_gen_gors_examples_tests_mapfields_message_proto_goTypes = nil
	file_cmd_protoc_gen_gors_examples_tests_mapfields_message_proto_depIdxs = nil
}
