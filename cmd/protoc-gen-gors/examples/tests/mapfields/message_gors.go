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
// source: cmd/protoc-gen-gors/examples/tests/mapfields/message.proto

package message

import (
	context "context"
	gin "github.com/gin-gonic/gin"
	gors "github.com/go-leo/gors"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
)

func MessagingServiceRoutes(svc MessagingService, opts ...gors.Option) []gors.Route {
	options := gors.NewOptions(opts...)
	wrapper := &_MessagingServiceWrapper{svc: svc, options: options}
	return []gors.Route{
		gors.NewRoute("PATCH", "/v1/messages/:message_id", _Messaging_UpdateMessage_GORS_Handler(wrapper, options)),
	}
}

func MessagingServerRoutes(srv MessagingServer, opts ...gors.Option) []gors.Route {
	options := gors.NewOptions(opts...)
	wrapper := &_MessagingServerWrapper{srv: srv, options: options}
	return []gors.Route{
		gors.NewRoute("PATCH", "/v1/messages/:message_id", _Messaging_UpdateMessage_GORS_Handler(wrapper, options)),
	}
}

func MessagingClientRoutes(cli MessagingClient, opts ...gors.Option) []gors.Route {
	options := gors.NewOptions(opts...)
	wrapper := &_MessagingClientWrapper{cli: cli, options: options}
	return []gors.Route{
		gors.NewRoute("PATCH", "/v1/messages/:message_id", _Messaging_UpdateMessage_GORS_Handler(wrapper, options)),
	}
}

// MessagingService is the service API for Messaging service.
type MessagingService interface {
	UpdateMessage(context.Context, *Message) (*Message, error)
}

var _ MessagingService = (*_MessagingServiceWrapper)(nil)

type _MessagingServiceWrapper struct {
	svc     MessagingService
	options *gors.Options
}

func (wrapper *_MessagingServiceWrapper) UpdateMessage(ctx context.Context, request *Message) (*Message, error) {
	return wrapper.svc.UpdateMessage(ctx, request)
}

var _ MessagingService = (*_MessagingServerWrapper)(nil)

// _MessagingServerWrapper implement MessagingService and wrap gRPC MessagingServer
type _MessagingServerWrapper struct {
	srv     MessagingServer
	options *gors.Options
}

func (wrapper *_MessagingServerWrapper) UpdateMessage(ctx context.Context, request *Message) (*Message, error) {
	rpcMethodName := "/tests.mapfields.message.v1.Messaging/UpdateMessage"
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

func (wrapper *_MessagingClientWrapper) UpdateMessage(ctx context.Context, request *Message) (*Message, error) {
	var headerMD, trailerMD metadata.MD
	resp, err := wrapper.cli.UpdateMessage(ctx, request, grpc.Header(&headerMD), grpc.Trailer(&trailerMD))
	gors.AddGRPCMetadata(ctx, headerMD, trailerMD, wrapper.options.OutgoingHeaderMatcher)
	return resp, err
}

func _Messaging_UpdateMessage_GORS_Handler(svc MessagingService, options *gors.Options) func(c *gin.Context) {
	return func(c *gin.Context) {
		var rpcMethodName = "/tests.mapfields.message.v1.Messaging/UpdateMessage"
		var ctx = gors.NewContext(c, rpcMethodName)
		var req *Message
		var resp *Message
		var err error
		req = new(Message)
		if ctx, err = gors.NewGRPCContext(ctx, options.IncomingHeaderMatcher, options.MetadataAnnotators); err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler, options.ResponseWrapper)
			return
		}
		resp, err = svc.UpdateMessage(ctx, req)
		if err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler, options.ResponseWrapper)
			return
		}
	}
}
