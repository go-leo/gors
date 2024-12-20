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
// source: example/api/tests/example/message.proto

package example

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
	Messaging_GetMessages_FullMethodName    = "/tests.example.message.v1.Messaging/GetMessages"
	Messaging_GetMessage_FullMethodName     = "/tests.example.message.v1.Messaging/GetMessage"
	Messaging_CreateMessage_FullMethodName  = "/tests.example.message.v1.Messaging/CreateMessage"
	Messaging_UpdateMessage_FullMethodName  = "/tests.example.message.v1.Messaging/UpdateMessage"
	Messaging_StreamRequest_FullMethodName  = "/tests.example.message.v1.Messaging/StreamRequest"
	Messaging_StreamResponse_FullMethodName = "/tests.example.message.v1.Messaging/StreamResponse"
	Messaging_Stream_FullMethodName         = "/tests.example.message.v1.Messaging/Stream"
)

// MessagingClient is the client API for Messaging service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessagingClient interface {
	GetMessages(ctx context.Context, in *GetMessageRequest, opts ...grpc.CallOption) (*Message, error)
	GetMessage(ctx context.Context, in *GetMessageRequest, opts ...grpc.CallOption) (*Message, error)
	CreateMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	UpdateMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	StreamRequest(ctx context.Context, opts ...grpc.CallOption) (Messaging_StreamRequestClient, error)
	StreamResponse(ctx context.Context, in *Message, opts ...grpc.CallOption) (Messaging_StreamResponseClient, error)
	Stream(ctx context.Context, opts ...grpc.CallOption) (Messaging_StreamClient, error)
}

type messagingClient struct {
	cc grpc.ClientConnInterface
}

func NewMessagingClient(cc grpc.ClientConnInterface) MessagingClient {
	return &messagingClient{cc}
}

func (c *messagingClient) GetMessages(ctx context.Context, in *GetMessageRequest, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, Messaging_GetMessages_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingClient) GetMessage(ctx context.Context, in *GetMessageRequest, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, Messaging_GetMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingClient) CreateMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, Messaging_CreateMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingClient) UpdateMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, Messaging_UpdateMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingClient) StreamRequest(ctx context.Context, opts ...grpc.CallOption) (Messaging_StreamRequestClient, error) {
	stream, err := c.cc.NewStream(ctx, &Messaging_ServiceDesc.Streams[0], Messaging_StreamRequest_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &messagingStreamRequestClient{stream}
	return x, nil
}

type Messaging_StreamRequestClient interface {
	Send(*Message) error
	CloseAndRecv() (*Message, error)
	grpc.ClientStream
}

type messagingStreamRequestClient struct {
	grpc.ClientStream
}

func (x *messagingStreamRequestClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *messagingStreamRequestClient) CloseAndRecv() (*Message, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *messagingClient) StreamResponse(ctx context.Context, in *Message, opts ...grpc.CallOption) (Messaging_StreamResponseClient, error) {
	stream, err := c.cc.NewStream(ctx, &Messaging_ServiceDesc.Streams[1], Messaging_StreamResponse_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &messagingStreamResponseClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Messaging_StreamResponseClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type messagingStreamResponseClient struct {
	grpc.ClientStream
}

func (x *messagingStreamResponseClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *messagingClient) Stream(ctx context.Context, opts ...grpc.CallOption) (Messaging_StreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Messaging_ServiceDesc.Streams[2], Messaging_Stream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &messagingStreamClient{stream}
	return x, nil
}

type Messaging_StreamClient interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ClientStream
}

type messagingStreamClient struct {
	grpc.ClientStream
}

func (x *messagingStreamClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *messagingStreamClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MessagingServer is the server API for Messaging service.
// All implementations must embed UnimplementedMessagingServer
// for forward compatibility
type MessagingServer interface {
	GetMessages(context.Context, *GetMessageRequest) (*Message, error)
	GetMessage(context.Context, *GetMessageRequest) (*Message, error)
	CreateMessage(context.Context, *Message) (*Message, error)
	UpdateMessage(context.Context, *Message) (*Message, error)
	StreamRequest(Messaging_StreamRequestServer) error
	StreamResponse(*Message, Messaging_StreamResponseServer) error
	Stream(Messaging_StreamServer) error
	mustEmbedUnimplementedMessagingServer()
}

// UnimplementedMessagingServer must be embedded to have forward compatible implementations.
type UnimplementedMessagingServer struct {
}

func (UnimplementedMessagingServer) GetMessages(context.Context, *GetMessageRequest) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessages not implemented")
}
func (UnimplementedMessagingServer) GetMessage(context.Context, *GetMessageRequest) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessage not implemented")
}
func (UnimplementedMessagingServer) CreateMessage(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMessage not implemented")
}
func (UnimplementedMessagingServer) UpdateMessage(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMessage not implemented")
}
func (UnimplementedMessagingServer) StreamRequest(Messaging_StreamRequestServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamRequest not implemented")
}
func (UnimplementedMessagingServer) StreamResponse(*Message, Messaging_StreamResponseServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamResponse not implemented")
}
func (UnimplementedMessagingServer) Stream(Messaging_StreamServer) error {
	return status.Errorf(codes.Unimplemented, "method Stream not implemented")
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

func _Messaging_GetMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServer).GetMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Messaging_GetMessages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServer).GetMessages(ctx, req.(*GetMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Messaging_GetMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServer).GetMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Messaging_GetMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServer).GetMessage(ctx, req.(*GetMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Messaging_CreateMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServer).CreateMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Messaging_CreateMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServer).CreateMessage(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
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

func _Messaging_StreamRequest_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MessagingServer).StreamRequest(&messagingStreamRequestServer{stream})
}

type Messaging_StreamRequestServer interface {
	SendAndClose(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type messagingStreamRequestServer struct {
	grpc.ServerStream
}

func (x *messagingStreamRequestServer) SendAndClose(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *messagingStreamRequestServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Messaging_StreamResponse_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Message)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MessagingServer).StreamResponse(m, &messagingStreamResponseServer{stream})
}

type Messaging_StreamResponseServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type messagingStreamResponseServer struct {
	grpc.ServerStream
}

func (x *messagingStreamResponseServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func _Messaging_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MessagingServer).Stream(&messagingStreamServer{stream})
}

type Messaging_StreamServer interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type messagingStreamServer struct {
	grpc.ServerStream
}

func (x *messagingStreamServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *messagingStreamServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Messaging_ServiceDesc is the grpc.ServiceDesc for Messaging service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Messaging_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tests.example.message.v1.Messaging",
	HandlerType: (*MessagingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMessages",
			Handler:    _Messaging_GetMessages_Handler,
		},
		{
			MethodName: "GetMessage",
			Handler:    _Messaging_GetMessage_Handler,
		},
		{
			MethodName: "CreateMessage",
			Handler:    _Messaging_CreateMessage_Handler,
		},
		{
			MethodName: "UpdateMessage",
			Handler:    _Messaging_UpdateMessage_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamRequest",
			Handler:       _Messaging_StreamRequest_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "StreamResponse",
			Handler:       _Messaging_StreamResponse_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Stream",
			Handler:       _Messaging_Stream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "example/api/tests/example/message.proto",
}
