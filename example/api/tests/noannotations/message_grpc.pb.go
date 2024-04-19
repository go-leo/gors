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

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: example/api/tests/noannotations/message.proto

package noannotations

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
	Messaging1_UpdateMessage_FullMethodName = "/tests.noannotations.message.v1.Messaging1/UpdateMessage"
)

// Messaging1Client is the client API for Messaging1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type Messaging1Client interface {
	UpdateMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
}

type messaging1Client struct {
	cc grpc.ClientConnInterface
}

func NewMessaging1Client(cc grpc.ClientConnInterface) Messaging1Client {
	return &messaging1Client{cc}
}

func (c *messaging1Client) UpdateMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, Messaging1_UpdateMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Messaging1Server is the server API for Messaging1 service.
// All implementations must embed UnimplementedMessaging1Server
// for forward compatibility
type Messaging1Server interface {
	UpdateMessage(context.Context, *Message) (*Message, error)
	mustEmbedUnimplementedMessaging1Server()
}

// UnimplementedMessaging1Server must be embedded to have forward compatible implementations.
type UnimplementedMessaging1Server struct {
}

func (UnimplementedMessaging1Server) UpdateMessage(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMessage not implemented")
}
func (UnimplementedMessaging1Server) mustEmbedUnimplementedMessaging1Server() {}

// UnsafeMessaging1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to Messaging1Server will
// result in compilation errors.
type UnsafeMessaging1Server interface {
	mustEmbedUnimplementedMessaging1Server()
}

func RegisterMessaging1Server(s grpc.ServiceRegistrar, srv Messaging1Server) {
	s.RegisterService(&Messaging1_ServiceDesc, srv)
}

func _Messaging1_UpdateMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Messaging1Server).UpdateMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Messaging1_UpdateMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Messaging1Server).UpdateMessage(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

// Messaging1_ServiceDesc is the grpc.ServiceDesc for Messaging1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Messaging1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tests.noannotations.message.v1.Messaging1",
	HandlerType: (*Messaging1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateMessage",
			Handler:    _Messaging1_UpdateMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "example/api/tests/noannotations/message.proto",
}

const (
	Messaging2_UpdateMessage_FullMethodName = "/tests.noannotations.message.v1.Messaging2/UpdateMessage"
)

// Messaging2Client is the client API for Messaging2 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type Messaging2Client interface {
	UpdateMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
}

type messaging2Client struct {
	cc grpc.ClientConnInterface
}

func NewMessaging2Client(cc grpc.ClientConnInterface) Messaging2Client {
	return &messaging2Client{cc}
}

func (c *messaging2Client) UpdateMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, Messaging2_UpdateMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Messaging2Server is the server API for Messaging2 service.
// All implementations must embed UnimplementedMessaging2Server
// for forward compatibility
type Messaging2Server interface {
	UpdateMessage(context.Context, *Message) (*Message, error)
	mustEmbedUnimplementedMessaging2Server()
}

// UnimplementedMessaging2Server must be embedded to have forward compatible implementations.
type UnimplementedMessaging2Server struct {
}

func (UnimplementedMessaging2Server) UpdateMessage(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMessage not implemented")
}
func (UnimplementedMessaging2Server) mustEmbedUnimplementedMessaging2Server() {}

// UnsafeMessaging2Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to Messaging2Server will
// result in compilation errors.
type UnsafeMessaging2Server interface {
	mustEmbedUnimplementedMessaging2Server()
}

func RegisterMessaging2Server(s grpc.ServiceRegistrar, srv Messaging2Server) {
	s.RegisterService(&Messaging2_ServiceDesc, srv)
}

func _Messaging2_UpdateMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Messaging2Server).UpdateMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Messaging2_UpdateMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Messaging2Server).UpdateMessage(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

// Messaging2_ServiceDesc is the grpc.ServiceDesc for Messaging2 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Messaging2_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tests.noannotations.message.v1.Messaging2",
	HandlerType: (*Messaging2Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateMessage",
			Handler:    _Messaging2_UpdateMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "example/api/tests/noannotations/message.proto",
}