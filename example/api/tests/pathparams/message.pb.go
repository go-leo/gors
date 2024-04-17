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
// source: example/api/tests/pathparams/message.proto

package pathparams

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

type GetMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageId string `protobuf:"bytes,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	UserId    uint64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetMessageRequest) Reset() {
	*x = GetMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_api_tests_pathparams_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMessageRequest) ProtoMessage() {}

func (x *GetMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_example_api_tests_pathparams_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMessageRequest.ProtoReflect.Descriptor instead.
func (*GetMessageRequest) Descriptor() ([]byte, []int) {
	return file_example_api_tests_pathparams_message_proto_rawDescGZIP(), []int{0}
}

func (x *GetMessageRequest) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

func (x *GetMessageRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type Meta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageId string `protobuf:"bytes,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	UserId    uint64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Content   string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *Meta) Reset() {
	*x = Meta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_api_tests_pathparams_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Meta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Meta) ProtoMessage() {}

func (x *Meta) ProtoReflect() protoreflect.Message {
	mi := &file_example_api_tests_pathparams_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Meta.ProtoReflect.Descriptor instead.
func (*Meta) Descriptor() ([]byte, []int) {
	return file_example_api_tests_pathparams_message_proto_rawDescGZIP(), []int{1}
}

func (x *Meta) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

func (x *Meta) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Meta) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageId string  `protobuf:"bytes,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	UserId    uint64  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Content   string  `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Maybe     *string `protobuf:"bytes,4,opt,name=maybe,proto3,oneof" json:"maybe,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_api_tests_pathparams_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_example_api_tests_pathparams_message_proto_msgTypes[2]
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
	return file_example_api_tests_pathparams_message_proto_rawDescGZIP(), []int{2}
}

func (x *Message) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

func (x *Message) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Message) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Message) GetMaybe() string {
	if x != nil && x.Maybe != nil {
		return *x.Maybe
	}
	return ""
}

type Comment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommentId string `protobuf:"bytes,1,opt,name=comment_id,json=commentId,proto3" json:"comment_id,omitempty"`
	UserId    uint64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Content   string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *Comment) Reset() {
	*x = Comment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_api_tests_pathparams_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Comment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Comment) ProtoMessage() {}

func (x *Comment) ProtoReflect() protoreflect.Message {
	mi := &file_example_api_tests_pathparams_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Comment.ProtoReflect.Descriptor instead.
func (*Comment) Descriptor() ([]byte, []int) {
	return file_example_api_tests_pathparams_message_proto_rawDescGZIP(), []int{3}
}

func (x *Comment) GetCommentId() string {
	if x != nil {
		return x.CommentId
	}
	return ""
}

func (x *Comment) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Comment) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type UpdateMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Content string  `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Maybe   *string `protobuf:"bytes,3,opt,name=maybe,proto3,oneof" json:"maybe,omitempty"`
}

func (x *UpdateMessageRequest) Reset() {
	*x = UpdateMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_api_tests_pathparams_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMessageRequest) ProtoMessage() {}

func (x *UpdateMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_example_api_tests_pathparams_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMessageRequest.ProtoReflect.Descriptor instead.
func (*UpdateMessageRequest) Descriptor() ([]byte, []int) {
	return file_example_api_tests_pathparams_message_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateMessageRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateMessageRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *UpdateMessageRequest) GetMaybe() string {
	if x != nil && x.Maybe != nil {
		return *x.Maybe
	}
	return ""
}

var File_example_api_tests_pathparams_message_proto protoreflect.FileDescriptor

var file_example_api_tests_pathparams_message_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x65,
	0x73, 0x74, 0x73, 0x2f, 0x70, 0x61, 0x74, 0x68, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x2f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x74, 0x65,
	0x73, 0x74, 0x73, 0x2e, 0x70, 0x61, 0x74, 0x68, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4b, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x58, 0x0a, 0x04, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x80,
	0x01, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x05,
	0x6d, 0x61, 0x79, 0x62, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x6d,
	0x61, 0x79, 0x62, 0x65, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6d, 0x61, 0x79, 0x62,
	0x65, 0x22, 0x5b, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x69,
	0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x05, 0x6d, 0x61, 0x79, 0x62, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x6d, 0x61, 0x79, 0x62, 0x65, 0x88, 0x01, 0x01, 0x42,
	0x08, 0x0a, 0x06, 0x5f, 0x6d, 0x61, 0x79, 0x62, 0x65, 0x32, 0xc8, 0x04, 0x0a, 0x09, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x12, 0x85, 0x01, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2e, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x70,
	0x61, 0x74, 0x68, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x70,
	0x61, 0x74, 0x68, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x21, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x2f, 0x7b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x12,
	0x99, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x2e, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x70, 0x61, 0x74, 0x68, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x24, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x70, 0x61, 0x74, 0x68, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x31, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2b,
	0x12, 0x29, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x7b,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x81, 0x01, 0x0a, 0x0d,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x24, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x70, 0x61, 0x74, 0x68, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x1a, 0x24, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x70, 0x61, 0x74, 0x68,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1e, 0x3a, 0x01, 0x2a, 0x22, 0x19, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x2f, 0x7b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x12,
	0x92, 0x01, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x31, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x70, 0x61, 0x74, 0x68, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x70, 0x61, 0x74,
	0x68, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x22, 0x3a, 0x01, 0x2a, 0x32, 0x1d, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65,
	0x3d, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x2a, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x2f, 0x2a, 0x7d, 0x42, 0x43, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x6c, 0x65, 0x6f, 0x2f, 0x67, 0x6f, 0x72, 0x73, 0x2f, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x73,
	0x2f, 0x70, 0x61, 0x74, 0x68, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x70,
	0x61, 0x74, 0x68, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_example_api_tests_pathparams_message_proto_rawDescOnce sync.Once
	file_example_api_tests_pathparams_message_proto_rawDescData = file_example_api_tests_pathparams_message_proto_rawDesc
)

func file_example_api_tests_pathparams_message_proto_rawDescGZIP() []byte {
	file_example_api_tests_pathparams_message_proto_rawDescOnce.Do(func() {
		file_example_api_tests_pathparams_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_example_api_tests_pathparams_message_proto_rawDescData)
	})
	return file_example_api_tests_pathparams_message_proto_rawDescData
}

var file_example_api_tests_pathparams_message_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_example_api_tests_pathparams_message_proto_goTypes = []interface{}{
	(*GetMessageRequest)(nil),    // 0: tests.pathparams.message.v1.GetMessageRequest
	(*Meta)(nil),                 // 1: tests.pathparams.message.v1.Meta
	(*Message)(nil),              // 2: tests.pathparams.message.v1.Message
	(*Comment)(nil),              // 3: tests.pathparams.message.v1.Comment
	(*UpdateMessageRequest)(nil), // 4: tests.pathparams.message.v1.UpdateMessageRequest
}
var file_example_api_tests_pathparams_message_proto_depIdxs = []int32{
	0, // 0: tests.pathparams.message.v1.Messaging.GetMessage:input_type -> tests.pathparams.message.v1.GetMessageRequest
	0, // 1: tests.pathparams.message.v1.Messaging.GetUserMessage:input_type -> tests.pathparams.message.v1.GetMessageRequest
	2, // 2: tests.pathparams.message.v1.Messaging.CreateMessage:input_type -> tests.pathparams.message.v1.Message
	4, // 3: tests.pathparams.message.v1.Messaging.UpdateMessage:input_type -> tests.pathparams.message.v1.UpdateMessageRequest
	2, // 4: tests.pathparams.message.v1.Messaging.GetMessage:output_type -> tests.pathparams.message.v1.Message
	2, // 5: tests.pathparams.message.v1.Messaging.GetUserMessage:output_type -> tests.pathparams.message.v1.Message
	2, // 6: tests.pathparams.message.v1.Messaging.CreateMessage:output_type -> tests.pathparams.message.v1.Message
	2, // 7: tests.pathparams.message.v1.Messaging.UpdateMessage:output_type -> tests.pathparams.message.v1.Message
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_example_api_tests_pathparams_message_proto_init() }
func file_example_api_tests_pathparams_message_proto_init() {
	if File_example_api_tests_pathparams_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_example_api_tests_pathparams_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMessageRequest); i {
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
		file_example_api_tests_pathparams_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Meta); i {
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
		file_example_api_tests_pathparams_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_example_api_tests_pathparams_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Comment); i {
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
		file_example_api_tests_pathparams_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateMessageRequest); i {
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
	file_example_api_tests_pathparams_message_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_example_api_tests_pathparams_message_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_example_api_tests_pathparams_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_example_api_tests_pathparams_message_proto_goTypes,
		DependencyIndexes: file_example_api_tests_pathparams_message_proto_depIdxs,
		MessageInfos:      file_example_api_tests_pathparams_message_proto_msgTypes,
	}.Build()
	File_example_api_tests_pathparams_message_proto = out.File
	file_example_api_tests_pathparams_message_proto_rawDesc = nil
	file_example_api_tests_pathparams_message_proto_goTypes = nil
	file_example_api_tests_pathparams_message_proto_depIdxs = nil
}
