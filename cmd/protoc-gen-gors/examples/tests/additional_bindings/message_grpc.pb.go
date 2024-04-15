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

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: cmd/protoc-gen-gors/examples/tests/additional_bindings/message.proto

package message

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Messaging_UpdateMessage_FullMethodName = "/tests.additional_bindings.message.v1.Messaging/UpdateMessage"
)

// MessagingClient is the client API for Messaging service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessagingClient interface {
	UpdateMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
}

type messagingClient struct {
	cc grpc.ClientConnInterface
}

func NewMessagingClient(cc grpc.ClientConnInterface) MessagingClient {
	return &messagingClient{cc}
}

func (c *messagingClient) UpdateMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, Messaging_UpdateMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessagingServer is the server API for Messaging service.
// All implementations must embed UnimplementedMessagingServer
// for forward compatibility
type MessagingServer interface {
	UpdateMessage(context.Context, *Message) (*Message, error)
	mustEmbedUnimplementedMessagingServer()
}

// UnimplementedMessagingServer must be embedded to have forward compatible implementations.
type UnimplementedMessagingServer struct {
}

func (UnimplementedMessagingServer) UpdateMessage(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMessage not implemented")
}
func (UnimplementedMessagingServer) mustEmbedUnimplementedMessagingServer() {}

// UnsafeMessagingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessagingServer will
// result in compilation errors.
type UnsafeMessagingServer interface {
	mustEmbedUnimplementedMessagingServer()
}

func RegisterMessagingServer(s grpc.ServiceRegistrar, srv MessagingServer) {
	s.RegisterService(&Messaging_ServiceDesc, srv)
}

func _Messaging_UpdateMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServer).UpdateMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Messaging_UpdateMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServer).UpdateMessage(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

// Messaging_ServiceDesc is the grpc.ServiceDesc for Messaging service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Messaging_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tests.additional_bindings.message.v1.Messaging",
	HandlerType: (*MessagingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateMessage",
			Handler:    _Messaging_UpdateMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cmd/protoc-gen-gors/examples/tests/additional_bindings/message.proto",
}
