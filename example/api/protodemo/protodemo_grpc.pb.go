// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: example/api/protodemo/protodemo.proto

package protodemo

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
	ProtoDemo_DELETEUriBindingJSONRender_FullMethodName                = "/protodemo.ProtoDemo/DELETEUriBindingJSONRender"
	ProtoDemo_GETUriBindingIndentedJSONRender_FullMethodName           = "/protodemo.ProtoDemo/GETUriBindingIndentedJSONRender"
	ProtoDemo_GETUriQueryBindingSecureJSONRender_FullMethodName        = "/protodemo.ProtoDemo/GETUriQueryBindingSecureJSONRender"
	ProtoDemo_PATCHHeaderProtoFormBindingPureJSONRender_FullMethodName = "/protodemo.ProtoDemo/PATCHHeaderProtoFormBindingPureJSONRender"
	ProtoDemo_PUTHeaderJSONBindingAsciiJSONRender_FullMethodName       = "/protodemo.ProtoDemo/PUTHeaderJSONBindingAsciiJSONRender"
	ProtoDemo_POSTProtoBufBindingProtoBufRender_FullMethodName         = "/protodemo.ProtoDemo/POSTProtoBufBindingProtoBufRender"
	ProtoDemo_POSTProtoJSONBindingProtoJSONRender_FullMethodName       = "/protodemo.ProtoDemo/POSTProtoJSONBindingProtoJSONRender"
	ProtoDemo_POSTCustomBindingCustomRender_FullMethodName             = "/protodemo.ProtoDemo/POSTCustomBindingCustomRender"
	ProtoDemo_NotDefine_FullMethodName                                 = "/protodemo.ProtoDemo/NotDefine"
	ProtoDemo_POSTSetHeaderTrailer_FullMethodName                      = "/protodemo.ProtoDemo/POSTSetHeaderTrailer"
	ProtoDemo_POSTError_FullMethodName                                 = "/protodemo.ProtoDemo/POSTError"
	ProtoDemo_POSTGRPCStatus_FullMethodName                            = "/protodemo.ProtoDemo/POSTGRPCStatus"
)

// ProtoDemoClient is the client API for ProtoDemo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProtoDemoClient interface {
	// @GORS @DELETE @Path(/UriBinding/JSONRender/:name) @UriBinding @JSONRender
	DELETEUriBindingJSONRender(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	// @GORS @GET @Path(/UriBinding/IndentedJSONRender/:name) @UriBinding @IndentedJSONRender
	GETUriBindingIndentedJSONRender(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	// @GORS @GET @Path(/UriQueryBinding/SecureJSONRender/:name) @UriBinding @QueryBinding @SecureJSONRender
	GETUriQueryBindingSecureJSONRender(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	// @GORS @PATCH @Path(/HeaderProtoFormBinding/PureJSONRender) @HeaderBinding @FormBinding @PureJSONRender
	PATCHHeaderProtoFormBindingPureJSONRender(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	// @GORS @PUT @Path(/HeaderJSONBinding/AsciiJSONRender) @HeaderBinding @JSONBinding @AsciiJSONRender
	PUTHeaderJSONBindingAsciiJSONRender(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	// @GORS @POST @Path(/ProtoBufBinding/ProtoBufRender) @ProtoBufBinding @ProtoBufRender
	POSTProtoBufBindingProtoBufRender(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	// @GORS @POST @Path(/ProtoJSONBinding/ProtoJSONRender) @ProtoJSONBinding @ProtoJSONRender
	POSTProtoJSONBindingProtoJSONRender(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	// @GORS @POST @Path(/CustomBinding/CustomRender) @CustomBinding @CustomRender
	POSTCustomBindingCustomRender(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	NotDefine(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	// @GORS @POST @Path(/POSTSetHeaderTrailer) @ProtoJSONBinding @ProtoJSONRender
	POSTSetHeaderTrailer(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	// @GORS @POST @Path(/Error) @ProtoJSONBinding @ProtoJSONRender
	POSTError(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	// @GORS @POST @Path(/GRPCStatus) @ProtoJSONBinding @ProtoJSONRender
	POSTGRPCStatus(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type protoDemoClient struct {
	cc grpc.ClientConnInterface
}

func NewProtoDemoClient(cc grpc.ClientConnInterface) ProtoDemoClient {
	return &protoDemoClient{cc}
}

func (c *protoDemoClient) DELETEUriBindingJSONRender(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, ProtoDemo_DELETEUriBindingJSONRender_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *protoDemoClient) GETUriBindingIndentedJSONRender(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, ProtoDemo_GETUriBindingIndentedJSONRender_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *protoDemoClient) GETUriQueryBindingSecureJSONRender(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, ProtoDemo_GETUriQueryBindingSecureJSONRender_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *protoDemoClient) PATCHHeaderProtoFormBindingPureJSONRender(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, ProtoDemo_PATCHHeaderProtoFormBindingPureJSONRender_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *protoDemoClient) PUTHeaderJSONBindingAsciiJSONRender(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, ProtoDemo_PUTHeaderJSONBindingAsciiJSONRender_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *protoDemoClient) POSTProtoBufBindingProtoBufRender(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, ProtoDemo_POSTProtoBufBindingProtoBufRender_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *protoDemoClient) POSTProtoJSONBindingProtoJSONRender(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, ProtoDemo_POSTProtoJSONBindingProtoJSONRender_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *protoDemoClient) POSTCustomBindingCustomRender(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, ProtoDemo_POSTCustomBindingCustomRender_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *protoDemoClient) NotDefine(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, ProtoDemo_NotDefine_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *protoDemoClient) POSTSetHeaderTrailer(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, ProtoDemo_POSTSetHeaderTrailer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *protoDemoClient) POSTError(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, ProtoDemo_POSTError_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *protoDemoClient) POSTGRPCStatus(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, ProtoDemo_POSTGRPCStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProtoDemoServer is the server API for ProtoDemo service.
// All implementations must embed UnimplementedProtoDemoServer
// for forward compatibility
type ProtoDemoServer interface {
	// @GORS @DELETE @Path(/UriBinding/JSONRender/:name) @UriBinding @JSONRender
	DELETEUriBindingJSONRender(context.Context, *HelloRequest) (*HelloReply, error)
	// @GORS @GET @Path(/UriBinding/IndentedJSONRender/:name) @UriBinding @IndentedJSONRender
	GETUriBindingIndentedJSONRender(context.Context, *HelloRequest) (*HelloReply, error)
	// @GORS @GET @Path(/UriQueryBinding/SecureJSONRender/:name) @UriBinding @QueryBinding @SecureJSONRender
	GETUriQueryBindingSecureJSONRender(context.Context, *HelloRequest) (*HelloReply, error)
	// @GORS @PATCH @Path(/HeaderProtoFormBinding/PureJSONRender) @HeaderBinding @FormBinding @PureJSONRender
	PATCHHeaderProtoFormBindingPureJSONRender(context.Context, *HelloRequest) (*HelloReply, error)
	// @GORS @PUT @Path(/HeaderJSONBinding/AsciiJSONRender) @HeaderBinding @JSONBinding @AsciiJSONRender
	PUTHeaderJSONBindingAsciiJSONRender(context.Context, *HelloRequest) (*HelloReply, error)
	// @GORS @POST @Path(/ProtoBufBinding/ProtoBufRender) @ProtoBufBinding @ProtoBufRender
	POSTProtoBufBindingProtoBufRender(context.Context, *HelloRequest) (*HelloReply, error)
	// @GORS @POST @Path(/ProtoJSONBinding/ProtoJSONRender) @ProtoJSONBinding @ProtoJSONRender
	POSTProtoJSONBindingProtoJSONRender(context.Context, *HelloRequest) (*HelloReply, error)
	// @GORS @POST @Path(/CustomBinding/CustomRender) @CustomBinding @CustomRender
	POSTCustomBindingCustomRender(context.Context, *HelloRequest) (*HelloReply, error)
	NotDefine(context.Context, *HelloRequest) (*HelloReply, error)
	// @GORS @POST @Path(/POSTSetHeaderTrailer) @ProtoJSONBinding @ProtoJSONRender
	POSTSetHeaderTrailer(context.Context, *HelloRequest) (*HelloReply, error)
	// @GORS @POST @Path(/Error) @ProtoJSONBinding @ProtoJSONRender
	POSTError(context.Context, *HelloRequest) (*HelloReply, error)
	// @GORS @POST @Path(/GRPCStatus) @ProtoJSONBinding @ProtoJSONRender
	POSTGRPCStatus(context.Context, *HelloRequest) (*HelloReply, error)
	mustEmbedUnimplementedProtoDemoServer()
}

// UnimplementedProtoDemoServer must be embedded to have forward compatible implementations.
type UnimplementedProtoDemoServer struct {
}

func (UnimplementedProtoDemoServer) DELETEUriBindingJSONRender(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DELETEUriBindingJSONRender not implemented")
}
func (UnimplementedProtoDemoServer) GETUriBindingIndentedJSONRender(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GETUriBindingIndentedJSONRender not implemented")
}
func (UnimplementedProtoDemoServer) GETUriQueryBindingSecureJSONRender(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GETUriQueryBindingSecureJSONRender not implemented")
}
func (UnimplementedProtoDemoServer) PATCHHeaderProtoFormBindingPureJSONRender(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PATCHHeaderProtoFormBindingPureJSONRender not implemented")
}
func (UnimplementedProtoDemoServer) PUTHeaderJSONBindingAsciiJSONRender(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PUTHeaderJSONBindingAsciiJSONRender not implemented")
}
func (UnimplementedProtoDemoServer) POSTProtoBufBindingProtoBufRender(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method POSTProtoBufBindingProtoBufRender not implemented")
}
func (UnimplementedProtoDemoServer) POSTProtoJSONBindingProtoJSONRender(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method POSTProtoJSONBindingProtoJSONRender not implemented")
}
func (UnimplementedProtoDemoServer) POSTCustomBindingCustomRender(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method POSTCustomBindingCustomRender not implemented")
}
func (UnimplementedProtoDemoServer) NotDefine(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotDefine not implemented")
}
func (UnimplementedProtoDemoServer) POSTSetHeaderTrailer(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method POSTSetHeaderTrailer not implemented")
}
func (UnimplementedProtoDemoServer) POSTError(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method POSTError not implemented")
}
func (UnimplementedProtoDemoServer) POSTGRPCStatus(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method POSTGRPCStatus not implemented")
}
func (UnimplementedProtoDemoServer) mustEmbedUnimplementedProtoDemoServer() {}

// UnsafeProtoDemoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProtoDemoServer will
// result in compilation errors.
type UnsafeProtoDemoServer interface {
	mustEmbedUnimplementedProtoDemoServer()
}

func RegisterProtoDemoServer(s grpc.ServiceRegistrar, srv ProtoDemoServer) {
	s.RegisterService(&ProtoDemo_ServiceDesc, srv)
}

func _ProtoDemo_DELETEUriBindingJSONRender_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProtoDemoServer).DELETEUriBindingJSONRender(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProtoDemo_DELETEUriBindingJSONRender_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProtoDemoServer).DELETEUriBindingJSONRender(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProtoDemo_GETUriBindingIndentedJSONRender_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProtoDemoServer).GETUriBindingIndentedJSONRender(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProtoDemo_GETUriBindingIndentedJSONRender_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProtoDemoServer).GETUriBindingIndentedJSONRender(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProtoDemo_GETUriQueryBindingSecureJSONRender_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProtoDemoServer).GETUriQueryBindingSecureJSONRender(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProtoDemo_GETUriQueryBindingSecureJSONRender_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProtoDemoServer).GETUriQueryBindingSecureJSONRender(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProtoDemo_PATCHHeaderProtoFormBindingPureJSONRender_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProtoDemoServer).PATCHHeaderProtoFormBindingPureJSONRender(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProtoDemo_PATCHHeaderProtoFormBindingPureJSONRender_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProtoDemoServer).PATCHHeaderProtoFormBindingPureJSONRender(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProtoDemo_PUTHeaderJSONBindingAsciiJSONRender_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProtoDemoServer).PUTHeaderJSONBindingAsciiJSONRender(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProtoDemo_PUTHeaderJSONBindingAsciiJSONRender_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProtoDemoServer).PUTHeaderJSONBindingAsciiJSONRender(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProtoDemo_POSTProtoBufBindingProtoBufRender_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProtoDemoServer).POSTProtoBufBindingProtoBufRender(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProtoDemo_POSTProtoBufBindingProtoBufRender_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProtoDemoServer).POSTProtoBufBindingProtoBufRender(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProtoDemo_POSTProtoJSONBindingProtoJSONRender_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProtoDemoServer).POSTProtoJSONBindingProtoJSONRender(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProtoDemo_POSTProtoJSONBindingProtoJSONRender_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProtoDemoServer).POSTProtoJSONBindingProtoJSONRender(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProtoDemo_POSTCustomBindingCustomRender_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProtoDemoServer).POSTCustomBindingCustomRender(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProtoDemo_POSTCustomBindingCustomRender_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProtoDemoServer).POSTCustomBindingCustomRender(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProtoDemo_NotDefine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProtoDemoServer).NotDefine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProtoDemo_NotDefine_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProtoDemoServer).NotDefine(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProtoDemo_POSTSetHeaderTrailer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProtoDemoServer).POSTSetHeaderTrailer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProtoDemo_POSTSetHeaderTrailer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProtoDemoServer).POSTSetHeaderTrailer(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProtoDemo_POSTError_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProtoDemoServer).POSTError(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProtoDemo_POSTError_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProtoDemoServer).POSTError(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProtoDemo_POSTGRPCStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProtoDemoServer).POSTGRPCStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProtoDemo_POSTGRPCStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProtoDemoServer).POSTGRPCStatus(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProtoDemo_ServiceDesc is the grpc.ServiceDesc for ProtoDemo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProtoDemo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protodemo.ProtoDemo",
	HandlerType: (*ProtoDemoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DELETEUriBindingJSONRender",
			Handler:    _ProtoDemo_DELETEUriBindingJSONRender_Handler,
		},
		{
			MethodName: "GETUriBindingIndentedJSONRender",
			Handler:    _ProtoDemo_GETUriBindingIndentedJSONRender_Handler,
		},
		{
			MethodName: "GETUriQueryBindingSecureJSONRender",
			Handler:    _ProtoDemo_GETUriQueryBindingSecureJSONRender_Handler,
		},
		{
			MethodName: "PATCHHeaderProtoFormBindingPureJSONRender",
			Handler:    _ProtoDemo_PATCHHeaderProtoFormBindingPureJSONRender_Handler,
		},
		{
			MethodName: "PUTHeaderJSONBindingAsciiJSONRender",
			Handler:    _ProtoDemo_PUTHeaderJSONBindingAsciiJSONRender_Handler,
		},
		{
			MethodName: "POSTProtoBufBindingProtoBufRender",
			Handler:    _ProtoDemo_POSTProtoBufBindingProtoBufRender_Handler,
		},
		{
			MethodName: "POSTProtoJSONBindingProtoJSONRender",
			Handler:    _ProtoDemo_POSTProtoJSONBindingProtoJSONRender_Handler,
		},
		{
			MethodName: "POSTCustomBindingCustomRender",
			Handler:    _ProtoDemo_POSTCustomBindingCustomRender_Handler,
		},
		{
			MethodName: "NotDefine",
			Handler:    _ProtoDemo_NotDefine_Handler,
		},
		{
			MethodName: "POSTSetHeaderTrailer",
			Handler:    _ProtoDemo_POSTSetHeaderTrailer_Handler,
		},
		{
			MethodName: "POSTError",
			Handler:    _ProtoDemo_POSTError_Handler,
		},
		{
			MethodName: "POSTGRPCStatus",
			Handler:    _ProtoDemo_POSTGRPCStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "example/api/protodemo/protodemo.proto",
}
