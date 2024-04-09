// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: example/api/protoservice/protoservice.proto

package protoservice

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
	ProtoService_Method_FullMethodName = "/protoservice.ProtoService/Method"
)

// ProtoServiceClient is the client API for ProtoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProtoServiceClient interface {
	// @GORS @POST @Path(/Method) @ProtoJSONBinding @ProtoJSONRender
	Method(ctx context.Context, in *HelloRequest1, opts ...grpc.CallOption) (*HelloReply1, error)
}

type protoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProtoServiceClient(cc grpc.ClientConnInterface) ProtoServiceClient {
	return &protoServiceClient{cc}
}

func (c *protoServiceClient) Method(ctx context.Context, in *HelloRequest1, opts ...grpc.CallOption) (*HelloReply1, error) {
	out := new(HelloReply1)
	err := c.cc.Invoke(ctx, ProtoService_Method_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProtoServiceServer is the server API for ProtoService service.
// All implementations must embed UnimplementedProtoServiceServer
// for forward compatibility
type ProtoServiceServer interface {
	// @GORS @POST @Path(/Method) @ProtoJSONBinding @ProtoJSONRender
	Method(context.Context, *HelloRequest1) (*HelloReply1, error)
	mustEmbedUnimplementedProtoServiceServer()
}

// UnimplementedProtoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProtoServiceServer struct {
}

func (UnimplementedProtoServiceServer) Method(context.Context, *HelloRequest1) (*HelloReply1, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Method not implemented")
}
func (UnimplementedProtoServiceServer) mustEmbedUnimplementedProtoServiceServer() {}

// UnsafeProtoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProtoServiceServer will
// result in compilation errors.
type UnsafeProtoServiceServer interface {
	mustEmbedUnimplementedProtoServiceServer()
}

func RegisterProtoServiceServer(s grpc.ServiceRegistrar, srv ProtoServiceServer) {
	s.RegisterService(&ProtoService_ServiceDesc, srv)
}

func _ProtoService_Method_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProtoServiceServer).Method(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProtoService_Method_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProtoServiceServer).Method(ctx, req.(*HelloRequest1))
	}
	return interceptor(ctx, in, info, handler)
}

// ProtoService_ServiceDesc is the grpc.ServiceDesc for ProtoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProtoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protoservice.ProtoService",
	HandlerType: (*ProtoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Method",
			Handler:    _ProtoService_Method_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "example/api/protoservice/protoservice.proto",
}

const (
	ProtoService2_Method_FullMethodName = "/protoservice.ProtoService2/Method"
)

// ProtoService2Client is the client API for ProtoService2 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProtoService2Client interface {
	// @GORS @POST @Path(/Method) @ProtoJSONBinding @ProtoJSONRender
	Method(ctx context.Context, in *HelloRequest1, opts ...grpc.CallOption) (*HelloReply1, error)
}

type protoService2Client struct {
	cc grpc.ClientConnInterface
}

func NewProtoService2Client(cc grpc.ClientConnInterface) ProtoService2Client {
	return &protoService2Client{cc}
}

func (c *protoService2Client) Method(ctx context.Context, in *HelloRequest1, opts ...grpc.CallOption) (*HelloReply1, error) {
	out := new(HelloReply1)
	err := c.cc.Invoke(ctx, ProtoService2_Method_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProtoService2Server is the server API for ProtoService2 service.
// All implementations must embed UnimplementedProtoService2Server
// for forward compatibility
type ProtoService2Server interface {
	// @GORS @POST @Path(/Method) @ProtoJSONBinding @ProtoJSONRender
	Method(context.Context, *HelloRequest1) (*HelloReply1, error)
	mustEmbedUnimplementedProtoService2Server()
}

// UnimplementedProtoService2Server must be embedded to have forward compatible implementations.
type UnimplementedProtoService2Server struct {
}

func (UnimplementedProtoService2Server) Method(context.Context, *HelloRequest1) (*HelloReply1, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Method not implemented")
}
func (UnimplementedProtoService2Server) mustEmbedUnimplementedProtoService2Server() {}

// UnsafeProtoService2Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProtoService2Server will
// result in compilation errors.
type UnsafeProtoService2Server interface {
	mustEmbedUnimplementedProtoService2Server()
}

func RegisterProtoService2Server(s grpc.ServiceRegistrar, srv ProtoService2Server) {
	s.RegisterService(&ProtoService2_ServiceDesc, srv)
}

func _ProtoService2_Method_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProtoService2Server).Method(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProtoService2_Method_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProtoService2Server).Method(ctx, req.(*HelloRequest1))
	}
	return interceptor(ctx, in, info, handler)
}

// ProtoService2_ServiceDesc is the grpc.ServiceDesc for ProtoService2 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProtoService2_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protoservice.ProtoService2",
	HandlerType: (*ProtoService2Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Method",
			Handler:    _ProtoService2_Method_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "example/api/protoservice/protoservice.proto",
}
