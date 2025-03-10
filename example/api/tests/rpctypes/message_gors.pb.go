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
// source: example/api/tests/rpctypes/message.proto

package rpctypes

import (
	context "context"
	gin "github.com/gin-gonic/gin"
	gors "github.com/go-leo/gors"
	binding "github.com/go-leo/gors/pkg/binding"
	status "google.golang.org/genproto/googleapis/rpc/status"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func StatusServiceRoutes(svc StatusService, opts ...gors.Option) []gors.Route {
	options := gors.NewOptions(opts...)
	wrapper := &_StatusServiceWrapper{svc: svc, options: options}
	_ = wrapper
	return []gors.Route{
		gors.NewRoute("GET", "/v1/status", _Status_GetStatus_GORS_Handler(wrapper, options, _Status_GetStatus_GORS_Handler_GET_b2f15b089d05a01f62c53b6e8eccb1e9_Binding())),
	}
}

func StatusServerRoutes(srv StatusServer, opts ...gors.Option) []gors.Route {
	options := gors.NewOptions(opts...)
	wrapper := &_StatusServerWrapper{srv: srv, options: options}
	_ = wrapper
	return []gors.Route{
		gors.NewRoute("GET", "/v1/status", _Status_GetStatus_GORS_Handler(wrapper, options, _Status_GetStatus_GORS_Handler_GET_b2f15b089d05a01f62c53b6e8eccb1e9_Binding())),
	}
}

func StatusClientRoutes(cli StatusClient, opts ...gors.Option) []gors.Route {
	options := gors.NewOptions(opts...)
	wrapper := &_StatusClientWrapper{cli: cli, options: options}
	_ = wrapper
	return []gors.Route{
		gors.NewRoute("GET", "/v1/status", _Status_GetStatus_GORS_Handler(wrapper, options, _Status_GetStatus_GORS_Handler_GET_b2f15b089d05a01f62c53b6e8eccb1e9_Binding())),
	}
}

// StatusService is the service API for Status service.
type StatusService interface {
	GetStatus(context.Context, *emptypb.Empty) (*status.Status, error)
}

var _ StatusService = (*_StatusServiceWrapper)(nil)

type _StatusServiceWrapper struct {
	svc     StatusService
	options *gors.Options
}

func (wrapper *_StatusServiceWrapper) GetStatus(ctx context.Context, request *emptypb.Empty) (*status.Status, error) {
	return wrapper.svc.GetStatus(ctx, request)
}

var _ StatusService = (*_StatusServerWrapper)(nil)

// _StatusServerWrapper implement StatusService and wrap gRPC StatusServer
type _StatusServerWrapper struct {
	srv     StatusServer
	options *gors.Options
}

func (wrapper *_StatusServerWrapper) GetStatus(ctx context.Context, request *emptypb.Empty) (*status.Status, error) {
	rpcMethodName := "/tests.rpctypes.message.v1.Status/GetStatus"
	stream := gors.NewServerTransportStream(rpcMethodName)
	ctx = grpc.NewContextWithServerTransportStream(ctx, stream)
	resp, err := wrapper.srv.GetStatus(ctx, request)
	gors.AddGRPCMetadata(ctx, stream.Header(), stream.Trailer(), wrapper.options.OutgoingHeaderMatcher)
	return resp, err
}

var _ StatusService = (*_StatusClientWrapper)(nil)

// _StatusClientWrapper implement StatusService and wrap gRPC StatusClient
type _StatusClientWrapper struct {
	cli     StatusClient
	options *gors.Options
}

func (wrapper *_StatusClientWrapper) GetStatus(ctx context.Context, request *emptypb.Empty) (*status.Status, error) {
	var headerMD, trailerMD metadata.MD
	resp, err := wrapper.cli.GetStatus(ctx, request, grpc.Header(&headerMD), grpc.Trailer(&trailerMD))
	gors.AddGRPCMetadata(ctx, headerMD, trailerMD, wrapper.options.OutgoingHeaderMatcher)
	return resp, err
}

func _Status_GetStatus_GORS_Handler(svc StatusService, options *gors.Options, binding *binding.HttpRuleBinding) func(c *gin.Context) {
	return func(c *gin.Context) {
		var rpcMethodName = "/tests.rpctypes.message.v1.Status/GetStatus"
		var ctx = gors.NewContext(c, rpcMethodName)
		var req *emptypb.Empty
		var resp *status.Status
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
		resp, err = svc.GetStatus(ctx, req)
		if err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler, options.ResponseWrapper)
			return
		}
		gors.ResponseRender(ctx, gors.StatusCode(ctx), resp, "", gors.ProtoJSONRender(options.ProtoJSONMarshalOptions), options.ResponseWrapper)
	}
}

func _Status_GetStatus_GORS_Handler_GET_b2f15b089d05a01f62c53b6e8eccb1e9_Binding() *binding.HttpRuleBinding {
	return &binding.HttpRuleBinding{}
}
