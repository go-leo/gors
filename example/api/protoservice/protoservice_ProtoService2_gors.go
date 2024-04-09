// Code generated by protoc-gen-go-gors. DO NOT EDIT.

package protoservice

import (
	context "context"
	gin "github.com/gin-gonic/gin"
	gors "github.com/go-leo/gors"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
	http "net/http"
)

func ProtoService2ClientRoutes(cli ProtoService2Client, opts ...gors.Option) []gors.Route {
	options := gors.NewOptions(opts...)
	if len(options.Tag) == 0 && !options.DisableDefaultTag {
		options.Tag = "json"
	}
	wrapper := &_ProtoService2ClientWrapper{cli: cli, options: options}
	return []gors.Route{
		gors.NewRoute(http.MethodPost, "/v1/Method", _ProtoService2_Method_GORS_Handler(wrapper, options)),
	}
}

func ProtoService2ServerRoutes(srv ProtoService2Server, opts ...gors.Option) []gors.Route {
	options := gors.NewOptions(opts...)
	if len(options.Tag) == 0 && !options.DisableDefaultTag {
		options.Tag = "json"
	}
	wrapper := &_ProtoService2ServerWrapper{srv: srv, options: options}
	return []gors.Route{
		gors.NewRoute(http.MethodPost, "/v1/Method", _ProtoService2_Method_GORS_Handler(wrapper, options)),
	}
}

func _ProtoService2_Method_GORS_Handler(wrapper ProtoService2Server, options *gors.Options) func(c *gin.Context) {
	return func(c *gin.Context) {
		var rpcMethodName = "/protoservice.ProtoService2/Method"
		var ctx = gors.NewContext(c, rpcMethodName)
		var req *HelloRequest1
		var resp *HelloReply1
		var err error
		req = new(HelloRequest1)
		if err = gors.RequestBind(
			ctx, req, options.Tag,
			gors.ProtoJSONBinding,
		); err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler, options.ResponseWrapper)
			return
		}
		if ctx, err = gors.NewGRPCContext(ctx, options.IncomingHeaderMatcher, options.MetadataAnnotators); err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler, options.ResponseWrapper)
			return
		}
		resp, err = wrapper.Method(ctx, req)
		if err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler, options.ResponseWrapper)
			return
		}
		gors.ResponseRender(ctx, gors.StatusCode(ctx), resp, "application/json", gors.ProtoJSONRender(options.ProtoJSONMarshalOptions), options.ResponseWrapper)
	}
}

type _ProtoService2ClientWrapper struct {
	UnimplementedProtoService2Server
	cli     ProtoService2Client
	options *gors.Options
}

func (wrapper *_ProtoService2ClientWrapper) Method(ctx context.Context, request *HelloRequest1) (*HelloReply1, error) {
	var headerMD, trailerMD metadata.MD
	resp, err := wrapper.cli.Method(ctx, request, grpc.Header(&headerMD), grpc.Trailer(&trailerMD))
	gors.AddGRPCMetadata(ctx, headerMD, trailerMD, wrapper.options.OutgoingHeaderMatcher)
	return resp, err
}

type _ProtoService2ServerWrapper struct {
	UnimplementedProtoService2Server
	srv     ProtoService2Server
	options *gors.Options
}

func (wrapper *_ProtoService2ServerWrapper) Method(ctx context.Context, request *HelloRequest1) (*HelloReply1, error) {
	rpcMethodName := "/protoservice.ProtoService2/Method"
	stream := gors.NewServerTransportStream(rpcMethodName)
	ctx = grpc.NewContextWithServerTransportStream(ctx, stream)
	resp, err := wrapper.srv.Method(ctx, request)
	gors.AddGRPCMetadata(ctx, stream.Header(), stream.Trailer(), wrapper.options.OutgoingHeaderMatcher)
	return resp, err
}
