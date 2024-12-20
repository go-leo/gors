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

// Code generated by protoc-gen-gors. DO NOT EDIT.
// versions:
// - protoc-gen-gors v1.3.0
// - protoc             v4.24.3
// source: example/api/tests/protobuftypes/message.proto

package protobuftypes

import (
	context "context"
	gin "github.com/gin-gonic/gin"
	gors "github.com/go-leo/gors"
	binding "github.com/go-leo/gors/pkg/binding"
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	structpb "google.golang.org/protobuf/types/known/structpb"
)

func MessagingServiceRoutes(svc MessagingService, opts ...gors.Option) []gors.Route {
	options := gors.NewOptions(opts...)
	wrapper := &_MessagingServiceWrapper{svc: svc, options: options}
	_ = wrapper
	return []gors.Route{
		gors.NewRoute("POST", "/v1/messages/:message_id", _Messaging_CreateMessage_GORS_Handler(wrapper, options, _Messaging_CreateMessage_GORS_Handler_POST_71b8052a59ef2e1e6bb26f276891271b_Binding())),
		gors.NewRoute("POST", "/v1/messages:csv", _Messaging_CreateMessagesFromCSV_GORS_Handler(wrapper, options, _Messaging_CreateMessagesFromCSV_GORS_Handler_POST_bb9003b5beb9432c1541343e9e470f29_Binding())),
		gors.NewRoute("GET", "/v1/messages", _Messaging_ListMessages_GORS_Handler(wrapper, options, _Messaging_ListMessages_GORS_Handler_GET_46cfd5872c02a89afb8c3f6fac56cbf9_Binding())),
		gors.NewRoute("GET", "/v1/messages:csv", _Messaging_ListMessagesCSV_GORS_Handler(wrapper, options, _Messaging_ListMessagesCSV_GORS_Handler_GET_bb9003b5beb9432c1541343e9e470f29_Binding())),
		gors.NewRoute("GET", "/v1/messages/:message_id", _Messaging_GetMessage_GORS_Handler(wrapper, options, _Messaging_GetMessage_GORS_Handler_GET_71b8052a59ef2e1e6bb26f276891271b_Binding())),
		gors.NewRoute("PATCH", "/v1/messages/:message_id", _Messaging_UpdateMessage_GORS_Handler(wrapper, options, _Messaging_UpdateMessage_GORS_Handler_PATCH_71b8052a59ef2e1e6bb26f276891271b_Binding())),
	}
}

func MessagingServerRoutes(srv MessagingServer, opts ...gors.Option) []gors.Route {
	options := gors.NewOptions(opts...)
	wrapper := &_MessagingServerWrapper{srv: srv, options: options}
	_ = wrapper
	return []gors.Route{
		gors.NewRoute("POST", "/v1/messages/:message_id", _Messaging_CreateMessage_GORS_Handler(wrapper, options, _Messaging_CreateMessage_GORS_Handler_POST_71b8052a59ef2e1e6bb26f276891271b_Binding())),
		gors.NewRoute("POST", "/v1/messages:csv", _Messaging_CreateMessagesFromCSV_GORS_Handler(wrapper, options, _Messaging_CreateMessagesFromCSV_GORS_Handler_POST_bb9003b5beb9432c1541343e9e470f29_Binding())),
		gors.NewRoute("GET", "/v1/messages", _Messaging_ListMessages_GORS_Handler(wrapper, options, _Messaging_ListMessages_GORS_Handler_GET_46cfd5872c02a89afb8c3f6fac56cbf9_Binding())),
		gors.NewRoute("GET", "/v1/messages:csv", _Messaging_ListMessagesCSV_GORS_Handler(wrapper, options, _Messaging_ListMessagesCSV_GORS_Handler_GET_bb9003b5beb9432c1541343e9e470f29_Binding())),
		gors.NewRoute("GET", "/v1/messages/:message_id", _Messaging_GetMessage_GORS_Handler(wrapper, options, _Messaging_GetMessage_GORS_Handler_GET_71b8052a59ef2e1e6bb26f276891271b_Binding())),
		gors.NewRoute("PATCH", "/v1/messages/:message_id", _Messaging_UpdateMessage_GORS_Handler(wrapper, options, _Messaging_UpdateMessage_GORS_Handler_PATCH_71b8052a59ef2e1e6bb26f276891271b_Binding())),
	}
}

func MessagingClientRoutes(cli MessagingClient, opts ...gors.Option) []gors.Route {
	options := gors.NewOptions(opts...)
	wrapper := &_MessagingClientWrapper{cli: cli, options: options}
	_ = wrapper
	return []gors.Route{
		gors.NewRoute("POST", "/v1/messages/:message_id", _Messaging_CreateMessage_GORS_Handler(wrapper, options, _Messaging_CreateMessage_GORS_Handler_POST_71b8052a59ef2e1e6bb26f276891271b_Binding())),
		gors.NewRoute("POST", "/v1/messages:csv", _Messaging_CreateMessagesFromCSV_GORS_Handler(wrapper, options, _Messaging_CreateMessagesFromCSV_GORS_Handler_POST_bb9003b5beb9432c1541343e9e470f29_Binding())),
		gors.NewRoute("GET", "/v1/messages", _Messaging_ListMessages_GORS_Handler(wrapper, options, _Messaging_ListMessages_GORS_Handler_GET_46cfd5872c02a89afb8c3f6fac56cbf9_Binding())),
		gors.NewRoute("GET", "/v1/messages:csv", _Messaging_ListMessagesCSV_GORS_Handler(wrapper, options, _Messaging_ListMessagesCSV_GORS_Handler_GET_bb9003b5beb9432c1541343e9e470f29_Binding())),
		gors.NewRoute("GET", "/v1/messages/:message_id", _Messaging_GetMessage_GORS_Handler(wrapper, options, _Messaging_GetMessage_GORS_Handler_GET_71b8052a59ef2e1e6bb26f276891271b_Binding())),
		gors.NewRoute("PATCH", "/v1/messages/:message_id", _Messaging_UpdateMessage_GORS_Handler(wrapper, options, _Messaging_UpdateMessage_GORS_Handler_PATCH_71b8052a59ef2e1e6bb26f276891271b_Binding())),
	}
}

// MessagingService is the service API for Messaging service.
type MessagingService interface {
	CreateMessage(context.Context, *Message) (*Message, error)
	CreateMessagesFromCSV(context.Context, *httpbody.HttpBody) (*httpbody.HttpBody, error)
	ListMessages(context.Context, *emptypb.Empty) (*structpb.Value, error)
	// OpenAPI does not allow requestBody in GET operations.
	// But it should not convert it to query params either.
	ListMessagesCSV(context.Context, *httpbody.HttpBody) (*httpbody.HttpBody, error)
	GetMessage(context.Context, *Message) (*Message, error)
	UpdateMessage(context.Context, *Message) (*structpb.Struct, error)
}

var _ MessagingService = (*_MessagingServiceWrapper)(nil)

type _MessagingServiceWrapper struct {
	svc     MessagingService
	options *gors.Options
}

func (wrapper *_MessagingServiceWrapper) CreateMessage(ctx context.Context, request *Message) (*Message, error) {
	return wrapper.svc.CreateMessage(ctx, request)
}

func (wrapper *_MessagingServiceWrapper) CreateMessagesFromCSV(ctx context.Context, request *httpbody.HttpBody) (*httpbody.HttpBody, error) {
	return wrapper.svc.CreateMessagesFromCSV(ctx, request)
}

func (wrapper *_MessagingServiceWrapper) ListMessages(ctx context.Context, request *emptypb.Empty) (*structpb.Value, error) {
	return wrapper.svc.ListMessages(ctx, request)
}

func (wrapper *_MessagingServiceWrapper) ListMessagesCSV(ctx context.Context, request *httpbody.HttpBody) (*httpbody.HttpBody, error) {
	return wrapper.svc.ListMessagesCSV(ctx, request)
}

func (wrapper *_MessagingServiceWrapper) GetMessage(ctx context.Context, request *Message) (*Message, error) {
	return wrapper.svc.GetMessage(ctx, request)
}

func (wrapper *_MessagingServiceWrapper) UpdateMessage(ctx context.Context, request *Message) (*structpb.Struct, error) {
	return wrapper.svc.UpdateMessage(ctx, request)
}

var _ MessagingService = (*_MessagingServerWrapper)(nil)

// _MessagingServerWrapper implement MessagingService and wrap gRPC MessagingServer
type _MessagingServerWrapper struct {
	srv     MessagingServer
	options *gors.Options
}

func (wrapper *_MessagingServerWrapper) CreateMessage(ctx context.Context, request *Message) (*Message, error) {
	rpcMethodName := "/tests.protobuftypes.message.v1.Messaging/CreateMessage"
	stream := gors.NewServerTransportStream(rpcMethodName)
	ctx = grpc.NewContextWithServerTransportStream(ctx, stream)
	resp, err := wrapper.srv.CreateMessage(ctx, request)
	gors.AddGRPCMetadata(ctx, stream.Header(), stream.Trailer(), wrapper.options.OutgoingHeaderMatcher)
	return resp, err
}

func (wrapper *_MessagingServerWrapper) CreateMessagesFromCSV(ctx context.Context, request *httpbody.HttpBody) (*httpbody.HttpBody, error) {
	rpcMethodName := "/tests.protobuftypes.message.v1.Messaging/CreateMessagesFromCSV"
	stream := gors.NewServerTransportStream(rpcMethodName)
	ctx = grpc.NewContextWithServerTransportStream(ctx, stream)
	resp, err := wrapper.srv.CreateMessagesFromCSV(ctx, request)
	gors.AddGRPCMetadata(ctx, stream.Header(), stream.Trailer(), wrapper.options.OutgoingHeaderMatcher)
	return resp, err
}

func (wrapper *_MessagingServerWrapper) ListMessages(ctx context.Context, request *emptypb.Empty) (*structpb.Value, error) {
	rpcMethodName := "/tests.protobuftypes.message.v1.Messaging/ListMessages"
	stream := gors.NewServerTransportStream(rpcMethodName)
	ctx = grpc.NewContextWithServerTransportStream(ctx, stream)
	resp, err := wrapper.srv.ListMessages(ctx, request)
	gors.AddGRPCMetadata(ctx, stream.Header(), stream.Trailer(), wrapper.options.OutgoingHeaderMatcher)
	return resp, err
}

func (wrapper *_MessagingServerWrapper) ListMessagesCSV(ctx context.Context, request *httpbody.HttpBody) (*httpbody.HttpBody, error) {
	rpcMethodName := "/tests.protobuftypes.message.v1.Messaging/ListMessagesCSV"
	stream := gors.NewServerTransportStream(rpcMethodName)
	ctx = grpc.NewContextWithServerTransportStream(ctx, stream)
	resp, err := wrapper.srv.ListMessagesCSV(ctx, request)
	gors.AddGRPCMetadata(ctx, stream.Header(), stream.Trailer(), wrapper.options.OutgoingHeaderMatcher)
	return resp, err
}

func (wrapper *_MessagingServerWrapper) GetMessage(ctx context.Context, request *Message) (*Message, error) {
	rpcMethodName := "/tests.protobuftypes.message.v1.Messaging/GetMessage"
	stream := gors.NewServerTransportStream(rpcMethodName)
	ctx = grpc.NewContextWithServerTransportStream(ctx, stream)
	resp, err := wrapper.srv.GetMessage(ctx, request)
	gors.AddGRPCMetadata(ctx, stream.Header(), stream.Trailer(), wrapper.options.OutgoingHeaderMatcher)
	return resp, err
}

func (wrapper *_MessagingServerWrapper) UpdateMessage(ctx context.Context, request *Message) (*structpb.Struct, error) {
	rpcMethodName := "/tests.protobuftypes.message.v1.Messaging/UpdateMessage"
	stream := gors.NewServerTransportStream(rpcMethodName)
	ctx = grpc.NewContextWithServerTransportStream(ctx, stream)
	resp, err := wrapper.srv.UpdateMessage(ctx, request)
	gors.AddGRPCMetadata(ctx, stream.Header(), stream.Trailer(), wrapper.options.OutgoingHeaderMatcher)
	return resp, err
}

var _ MessagingService = (*_MessagingClientWrapper)(nil)

// _MessagingClientWrapper implement MessagingService and wrap gRPC MessagingClient
type _MessagingClientWrapper struct {
	cli     MessagingClient
	options *gors.Options
}

func (wrapper *_MessagingClientWrapper) CreateMessage(ctx context.Context, request *Message) (*Message, error) {
	var headerMD, trailerMD metadata.MD
	resp, err := wrapper.cli.CreateMessage(ctx, request, grpc.Header(&headerMD), grpc.Trailer(&trailerMD))
	gors.AddGRPCMetadata(ctx, headerMD, trailerMD, wrapper.options.OutgoingHeaderMatcher)
	return resp, err
}

func (wrapper *_MessagingClientWrapper) CreateMessagesFromCSV(ctx context.Context, request *httpbody.HttpBody) (*httpbody.HttpBody, error) {
	var headerMD, trailerMD metadata.MD
	resp, err := wrapper.cli.CreateMessagesFromCSV(ctx, request, grpc.Header(&headerMD), grpc.Trailer(&trailerMD))
	gors.AddGRPCMetadata(ctx, headerMD, trailerMD, wrapper.options.OutgoingHeaderMatcher)
	return resp, err
}

func (wrapper *_MessagingClientWrapper) ListMessages(ctx context.Context, request *emptypb.Empty) (*structpb.Value, error) {
	var headerMD, trailerMD metadata.MD
	resp, err := wrapper.cli.ListMessages(ctx, request, grpc.Header(&headerMD), grpc.Trailer(&trailerMD))
	gors.AddGRPCMetadata(ctx, headerMD, trailerMD, wrapper.options.OutgoingHeaderMatcher)
	return resp, err
}

func (wrapper *_MessagingClientWrapper) ListMessagesCSV(ctx context.Context, request *httpbody.HttpBody) (*httpbody.HttpBody, error) {
	var headerMD, trailerMD metadata.MD
	resp, err := wrapper.cli.ListMessagesCSV(ctx, request, grpc.Header(&headerMD), grpc.Trailer(&trailerMD))
	gors.AddGRPCMetadata(ctx, headerMD, trailerMD, wrapper.options.OutgoingHeaderMatcher)
	return resp, err
}

func (wrapper *_MessagingClientWrapper) GetMessage(ctx context.Context, request *Message) (*Message, error) {
	var headerMD, trailerMD metadata.MD
	resp, err := wrapper.cli.GetMessage(ctx, request, grpc.Header(&headerMD), grpc.Trailer(&trailerMD))
	gors.AddGRPCMetadata(ctx, headerMD, trailerMD, wrapper.options.OutgoingHeaderMatcher)
	return resp, err
}

func (wrapper *_MessagingClientWrapper) UpdateMessage(ctx context.Context, request *Message) (*structpb.Struct, error) {
	var headerMD, trailerMD metadata.MD
	resp, err := wrapper.cli.UpdateMessage(ctx, request, grpc.Header(&headerMD), grpc.Trailer(&trailerMD))
	gors.AddGRPCMetadata(ctx, headerMD, trailerMD, wrapper.options.OutgoingHeaderMatcher)
	return resp, err
}

func _Messaging_CreateMessage_GORS_Handler(svc MessagingService, options *gors.Options, binding *binding.HttpRuleBinding) func(c *gin.Context) {
	return func(c *gin.Context) {
		var rpcMethodName = "/tests.protobuftypes.message.v1.Messaging/CreateMessage"
		var ctx = gors.NewContext(c, rpcMethodName)
		var req *Message
		var resp *Message
		var err error
		req = new(Message)
		if err = gors.RequestBind(
			ctx, req, options.Tag,
			gors.HttpRuleBinding(binding),
		); err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler, options.ResponseWrapper)
			return
		}
		if ctx, err = gors.NewGRPCContext(ctx, options.IncomingHeaderMatcher, options.MetadataAnnotators); err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler, options.ResponseWrapper)
			return
		}
		resp, err = svc.CreateMessage(ctx, req)
		if err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler, options.ResponseWrapper)
			return
		}
		gors.ResponseRender(ctx, gors.StatusCode(ctx), resp, "", gors.ProtoJSONRender(options.ProtoJSONMarshalOptions), options.ResponseWrapper)
	}
}

func _Messaging_CreateMessagesFromCSV_GORS_Handler(svc MessagingService, options *gors.Options, binding *binding.HttpRuleBinding) func(c *gin.Context) {
	return func(c *gin.Context) {
		var rpcMethodName = "/tests.protobuftypes.message.v1.Messaging/CreateMessagesFromCSV"
		var ctx = gors.NewContext(c, rpcMethodName)
		var req *httpbody.HttpBody
		var resp *httpbody.HttpBody
		var err error
		req = new(httpbody.HttpBody)
		if err = gors.RequestBind(
			ctx, req, options.Tag,
			gors.HttpRuleBinding(binding),
		); err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler, options.ResponseWrapper)
			return
		}
		if ctx, err = gors.NewGRPCContext(ctx, options.IncomingHeaderMatcher, options.MetadataAnnotators); err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler, options.ResponseWrapper)
			return
		}
		resp, err = svc.CreateMessagesFromCSV(ctx, req)
		if err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler, options.ResponseWrapper)
			return
		}
		gors.ResponseRender(ctx, gors.StatusCode(ctx), resp, "", gors.ProtoJSONRender(options.ProtoJSONMarshalOptions), options.ResponseWrapper)
	}
}

func _Messaging_ListMessages_GORS_Handler(svc MessagingService, options *gors.Options, binding *binding.HttpRuleBinding) func(c *gin.Context) {
	return func(c *gin.Context) {
		var rpcMethodName = "/tests.protobuftypes.message.v1.Messaging/ListMessages"
		var ctx = gors.NewContext(c, rpcMethodName)
		var req *emptypb.Empty
		var resp *structpb.Value
		var err error
		req = new(emptypb.Empty)
		if err = gors.RequestBind(
			ctx, req, options.Tag,
			gors.HttpRuleBinding(binding),
		); err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler, options.ResponseWrapper)
			return
		}
		if ctx, err = gors.NewGRPCContext(ctx, options.IncomingHeaderMatcher, options.MetadataAnnotators); err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler, options.ResponseWrapper)
			return
		}
		resp, err = svc.ListMessages(ctx, req)
		if err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler, options.ResponseWrapper)
			return
		}
		gors.ResponseRender(ctx, gors.StatusCode(ctx), resp, "", gors.ProtoJSONRender(options.ProtoJSONMarshalOptions), options.ResponseWrapper)
	}
}

func _Messaging_ListMessagesCSV_GORS_Handler(svc MessagingService, options *gors.Options, binding *binding.HttpRuleBinding) func(c *gin.Context) {
	return func(c *gin.Context) {
		var rpcMethodName = "/tests.protobuftypes.message.v1.Messaging/ListMessagesCSV"
		var ctx = gors.NewContext(c, rpcMethodName)
		var req *httpbody.HttpBody
		var resp *httpbody.HttpBody
		var err error
		req = new(httpbody.HttpBody)
		if err = gors.RequestBind(
			ctx, req, options.Tag,
			gors.HttpRuleBinding(binding),
		); err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler, options.ResponseWrapper)
			return
		}
		if ctx, err = gors.NewGRPCContext(ctx, options.IncomingHeaderMatcher, options.MetadataAnnotators); err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler, options.ResponseWrapper)
			return
		}
		resp, err = svc.ListMessagesCSV(ctx, req)
		if err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler, options.ResponseWrapper)
			return
		}
		gors.ResponseRender(ctx, gors.StatusCode(ctx), resp, "", gors.ProtoJSONRender(options.ProtoJSONMarshalOptions), options.ResponseWrapper)
	}
}

func _Messaging_GetMessage_GORS_Handler(svc MessagingService, options *gors.Options, binding *binding.HttpRuleBinding) func(c *gin.Context) {
	return func(c *gin.Context) {
		var rpcMethodName = "/tests.protobuftypes.message.v1.Messaging/GetMessage"
		var ctx = gors.NewContext(c, rpcMethodName)
		var req *Message
		var resp *Message
		var err error
		req = new(Message)
		if err = gors.RequestBind(
			ctx, req, options.Tag,
			gors.HttpRuleBinding(binding),
		); err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler, options.ResponseWrapper)
			return
		}
		if ctx, err = gors.NewGRPCContext(ctx, options.IncomingHeaderMatcher, options.MetadataAnnotators); err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler, options.ResponseWrapper)
			return
		}
		resp, err = svc.GetMessage(ctx, req)
		if err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler, options.ResponseWrapper)
			return
		}
		gors.ResponseRender(ctx, gors.StatusCode(ctx), resp, "", gors.ProtoJSONRender(options.ProtoJSONMarshalOptions), options.ResponseWrapper)
	}
}

func _Messaging_UpdateMessage_GORS_Handler(svc MessagingService, options *gors.Options, binding *binding.HttpRuleBinding) func(c *gin.Context) {
	return func(c *gin.Context) {
		var rpcMethodName = "/tests.protobuftypes.message.v1.Messaging/UpdateMessage"
		var ctx = gors.NewContext(c, rpcMethodName)
		var req *Message
		var resp *structpb.Struct
		var err error
		req = new(Message)
		if err = gors.RequestBind(
			ctx, req, options.Tag,
			gors.HttpRuleBinding(binding),
		); err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler, options.ResponseWrapper)
			return
		}
		if ctx, err = gors.NewGRPCContext(ctx, options.IncomingHeaderMatcher, options.MetadataAnnotators); err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler, options.ResponseWrapper)
			return
		}
		resp, err = svc.UpdateMessage(ctx, req)
		if err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler, options.ResponseWrapper)
			return
		}
		gors.ResponseRender(ctx, gors.StatusCode(ctx), resp, "", gors.ProtoJSONRender(options.ProtoJSONMarshalOptions), options.ResponseWrapper)
	}
}

func _Messaging_CreateMessage_GORS_Handler_POST_71b8052a59ef2e1e6bb26f276891271b_Binding() *binding.HttpRuleBinding {
	return &binding.HttpRuleBinding{
		Path: []*binding.PathRule{
			{Name: "message_id", Type: "string"},
		},
		Body: &binding.BodyRule{
			Name: "*",
			Type: "object",
		},
	}
}
func _Messaging_CreateMessagesFromCSV_GORS_Handler_POST_bb9003b5beb9432c1541343e9e470f29_Binding() *binding.HttpRuleBinding {
	return &binding.HttpRuleBinding{
		Body: &binding.BodyRule{
			Name: "*",
			Type: "string",
		},
	}
}
func _Messaging_ListMessages_GORS_Handler_GET_46cfd5872c02a89afb8c3f6fac56cbf9_Binding() *binding.HttpRuleBinding {
	return &binding.HttpRuleBinding{}
}
func _Messaging_ListMessagesCSV_GORS_Handler_GET_bb9003b5beb9432c1541343e9e470f29_Binding() *binding.HttpRuleBinding {
	return &binding.HttpRuleBinding{}
}
func _Messaging_GetMessage_GORS_Handler_GET_71b8052a59ef2e1e6bb26f276891271b_Binding() *binding.HttpRuleBinding {
	return &binding.HttpRuleBinding{
		Path: []*binding.PathRule{
			{Name: "message_id", Type: "string"},
		},
		Query: []*binding.QueryRule{
			// typ:string
			{Name: "string_type", Type: "string", ItemType: ""},
			// typ:integer
			{Name: "recursive_type.parent_id", Type: "integer", ItemType: ""},
			// typ:integer
			{Name: "recursive_type.child.child_id", Type: "integer", ItemType: ""},
			// typ:integer
			{Name: "recursive_type.child.parent.parent_id", Type: "integer", ItemType: ""},
			// typ:integer
			{Name: "recursive_type.child.parent.child.child_id", Type: "integer", ItemType: ""},
			// typ:string
			{Name: "embedded_type.message_id", Type: "string", ItemType: ""},
			// typ:string
			{Name: "sub_type.message_id", Type: "string", ItemType: ""},
			// typ:string
			{Name: "sub_type.sub_sub_message.message_id", Type: "string", ItemType: ""},
			// typ:array
			{Name: "sub_type.sub_sub_message.integers", Type: "array", ItemType: "integer"},
			// typ:array
			{Name: "repeated_type", Type: "array", ItemType: "string"},
			// typ:
			{Name: "value_type", Type: "", ItemType: ""},
			// typ:array
			{Name: "repeated_value_type", Type: "array", ItemType: ""},
			// typ:boolean
			{Name: "bool_value_type", Type: "boolean", ItemType: ""},
			// typ:string
			{Name: "bytes_value_type", Type: "string", ItemType: ""},
			// typ:integer
			{Name: "int32_value_type", Type: "integer", ItemType: ""},
			// typ:integer
			{Name: "uint32_value_type", Type: "integer", ItemType: ""},
			// typ:string
			{Name: "string_value_type", Type: "string", ItemType: ""},
			// typ:string
			{Name: "int64_value_type", Type: "string", ItemType: ""},
			// typ:string
			{Name: "uint64_value_type", Type: "string", ItemType: ""},
			// typ:number
			{Name: "float_value_type", Type: "number", ItemType: ""},
			// typ:number
			{Name: "double_value_type", Type: "number", ItemType: ""},
			// typ:string
			{Name: "timestamp_type", Type: "string", ItemType: ""},
			// typ:string
			{Name: "duration_type", Type: "string", ItemType: ""},
		},
	}
}
func _Messaging_UpdateMessage_GORS_Handler_PATCH_71b8052a59ef2e1e6bb26f276891271b_Binding() *binding.HttpRuleBinding {
	return &binding.HttpRuleBinding{
		Path: []*binding.PathRule{
			{Name: "message_id", Type: "string"},
		},
		Query: []*binding.QueryRule{
			// typ:string
			{Name: "string_type", Type: "string", ItemType: ""},
			// typ:integer
			{Name: "recursive_type.parent_id", Type: "integer", ItemType: ""},
			// typ:integer
			{Name: "recursive_type.child.child_id", Type: "integer", ItemType: ""},
			// typ:integer
			{Name: "recursive_type.child.parent.parent_id", Type: "integer", ItemType: ""},
			// typ:integer
			{Name: "recursive_type.child.parent.child.child_id", Type: "integer", ItemType: ""},
			// typ:string
			{Name: "embedded_type.message_id", Type: "string", ItemType: ""},
			// typ:string
			{Name: "sub_type.message_id", Type: "string", ItemType: ""},
			// typ:string
			{Name: "sub_type.sub_sub_message.message_id", Type: "string", ItemType: ""},
			// typ:array
			{Name: "sub_type.sub_sub_message.integers", Type: "array", ItemType: "integer"},
			// typ:array
			{Name: "repeated_type", Type: "array", ItemType: "string"},
			// typ:
			{Name: "value_type", Type: "", ItemType: ""},
			// typ:array
			{Name: "repeated_value_type", Type: "array", ItemType: ""},
			// typ:boolean
			{Name: "bool_value_type", Type: "boolean", ItemType: ""},
			// typ:string
			{Name: "bytes_value_type", Type: "string", ItemType: ""},
			// typ:integer
			{Name: "int32_value_type", Type: "integer", ItemType: ""},
			// typ:integer
			{Name: "uint32_value_type", Type: "integer", ItemType: ""},
			// typ:string
			{Name: "string_value_type", Type: "string", ItemType: ""},
			// typ:string
			{Name: "int64_value_type", Type: "string", ItemType: ""},
			// typ:string
			{Name: "uint64_value_type", Type: "string", ItemType: ""},
			// typ:number
			{Name: "float_value_type", Type: "number", ItemType: ""},
			// typ:number
			{Name: "double_value_type", Type: "number", ItemType: ""},
			// typ:string
			{Name: "timestamp_type", Type: "string", ItemType: ""},
			// typ:string
			{Name: "duration_type", Type: "string", ItemType: ""},
		},
		Body: &binding.BodyRule{
			Name: "body",
			Type: "object",
		},
	}
}
