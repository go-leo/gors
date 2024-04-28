# gors

gors（Golang Restful service）可以生成Go语言的Restful服务。
在proto上，通过google.api.annotations语法，可以生成Gin路由。


# 快速开始
## 1. 安装
如果通过grpc生成Restful服务，需要安装protoc-gen-go-gors：
```
go install github.com/go-leo/gors/cmd/protoc-gen-gors@latest
```

## 2. 定义Restful服务
或者通过grpc定义接口：
```proto
syntax = "proto3";

package tests.example.message.v1;

import "google/api/annotations.proto";

option go_package = "github.com/go-leo/gors/example/api/tests/example/v1;example";

service Messaging {
  rpc GetMessage(GetMessageRequest) returns (Message) {
    option (google.api.http) = {
      get : "/v1/messages/{message_id}"
    };
  }

  rpc CreateMessage(Message) returns (Message) {
    option (google.api.http) = {
      post : "/v1/messages/{message_id}"
      body : "*"
    };
  }
}

message GetMessageRequest {
  string message_id = 1;
  uint64 user_id = 2;
}

message Message {
  string message_id = 1;
  uint64 user_id = 2;
  string content = 3;
  optional string maybe = 4;
}
```

## 3. 生成gin路由代码
通过protocolbuf生成，运行一下命令，生产文件service_gors.go
```
	protoc \
		--proto_path=. \
		--proto_path=example/api \
		--go_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_out=. \
		--go-grpc_opt=paths=source_relative \
		--gors_out=. \
		--gors_opt=paths=source_relative \
		--gors_opt=naming=proto \
		--openapi_out=. \
		--openapi_opt=paths=source_relative \
		--openapi_opt=output_mode=source_relative \
		--openapi_opt=fq_schema_naming=true \
		--openapi_opt=depth=5 \
		--openapi_opt=naming=proto \
		example/api/tests/example/message.proto
```
生产文件message_gors.pb.go
```go
package example

import (
	context "context"
	gin "github.com/gin-gonic/gin"
	gors "github.com/go-leo/gors"
	binding "github.com/go-leo/gors/pkg/binding"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
)

func MessagingServiceRoutes(svc MessagingService, opts ...gors.Option) []gors.Route {
	options := gors.NewOptions(opts...)
	wrapper := &_MessagingServiceWrapper{svc: svc, options: options}
	_ = wrapper
	return []gors.Route{
		gors.NewRoute("GET", "/v1/messages/:message_id", _Messaging_GetMessage_GORS_Handler(wrapper, options, _Messaging_GetMessage_GORS_Handler_GET_71b8052a59ef2e1e6bb26f276891271b_Binding())),
		gors.NewRoute("POST", "/v1/messages/:message_id", _Messaging_CreateMessage_GORS_Handler(wrapper, options, _Messaging_CreateMessage_GORS_Handler_POST_71b8052a59ef2e1e6bb26f276891271b_Binding())),
	}
}

func MessagingServerRoutes(srv MessagingServer, opts ...gors.Option) []gors.Route {
	options := gors.NewOptions(opts...)
	wrapper := &_MessagingServerWrapper{srv: srv, options: options}
	_ = wrapper
	return []gors.Route{
		gors.NewRoute("GET", "/v1/messages/:message_id", _Messaging_GetMessage_GORS_Handler(wrapper, options, _Messaging_GetMessage_GORS_Handler_GET_71b8052a59ef2e1e6bb26f276891271b_Binding())),
		gors.NewRoute("POST", "/v1/messages/:message_id", _Messaging_CreateMessage_GORS_Handler(wrapper, options, _Messaging_CreateMessage_GORS_Handler_POST_71b8052a59ef2e1e6bb26f276891271b_Binding())),
	}
}

func MessagingClientRoutes(cli MessagingClient, opts ...gors.Option) []gors.Route {
	options := gors.NewOptions(opts...)
	wrapper := &_MessagingClientWrapper{cli: cli, options: options}
	_ = wrapper
	return []gors.Route{
		gors.NewRoute("GET", "/v1/messages/:message_id", _Messaging_GetMessage_GORS_Handler(wrapper, options, _Messaging_GetMessage_GORS_Handler_GET_71b8052a59ef2e1e6bb26f276891271b_Binding())),
		gors.NewRoute("POST", "/v1/messages/:message_id", _Messaging_CreateMessage_GORS_Handler(wrapper, options, _Messaging_CreateMessage_GORS_Handler_POST_71b8052a59ef2e1e6bb26f276891271b_Binding())),
	}
}

// MessagingService is the service API for Messaging service.
type MessagingService interface {
	GetMessage(context.Context, *GetMessageRequest) (*Message, error)
	CreateMessage(context.Context, *Message) (*Message, error)
}

var _ MessagingService = (*_MessagingServiceWrapper)(nil)

type _MessagingServiceWrapper struct {
	svc     MessagingService
	options *gors.Options
}

func (wrapper *_MessagingServiceWrapper) GetMessage(ctx context.Context, request *GetMessageRequest) (*Message, error) {
	return wrapper.svc.GetMessage(ctx, request)
}

func (wrapper *_MessagingServiceWrapper) CreateMessage(ctx context.Context, request *Message) (*Message, error) {
	return wrapper.svc.CreateMessage(ctx, request)
}

var _ MessagingService = (*_MessagingServerWrapper)(nil)

// _MessagingServerWrapper implement MessagingService and wrap gRPC MessagingServer
type _MessagingServerWrapper struct {
	srv     MessagingServer
	options *gors.Options
}

func (wrapper *_MessagingServerWrapper) GetMessage(ctx context.Context, request *GetMessageRequest) (*Message, error) {
	rpcMethodName := "/tests.example.message.v1.Messaging/GetMessage"
	stream := gors.NewServerTransportStream(rpcMethodName)
	ctx = grpc.NewContextWithServerTransportStream(ctx, stream)
	resp, err := wrapper.srv.GetMessage(ctx, request)
	gors.AddGRPCMetadata(ctx, stream.Header(), stream.Trailer(), wrapper.options.OutgoingHeaderMatcher)
	return resp, err
}

func (wrapper *_MessagingServerWrapper) CreateMessage(ctx context.Context, request *Message) (*Message, error) {
	rpcMethodName := "/tests.example.message.v1.Messaging/CreateMessage"
	stream := gors.NewServerTransportStream(rpcMethodName)
	ctx = grpc.NewContextWithServerTransportStream(ctx, stream)
	resp, err := wrapper.srv.CreateMessage(ctx, request)
	gors.AddGRPCMetadata(ctx, stream.Header(), stream.Trailer(), wrapper.options.OutgoingHeaderMatcher)
	return resp, err
}

var _ MessagingService = (*_MessagingClientWrapper)(nil)

// _MessagingClientWrapper implement MessagingService and wrap gRPC MessagingClient
type _MessagingClientWrapper struct {
	cli     MessagingClient
	options *gors.Options
}

func (wrapper *_MessagingClientWrapper) GetMessage(ctx context.Context, request *GetMessageRequest) (*Message, error) {
	var headerMD, trailerMD metadata.MD
	resp, err := wrapper.cli.GetMessage(ctx, request, grpc.Header(&headerMD), grpc.Trailer(&trailerMD))
	gors.AddGRPCMetadata(ctx, headerMD, trailerMD, wrapper.options.OutgoingHeaderMatcher)
	return resp, err
}

func (wrapper *_MessagingClientWrapper) CreateMessage(ctx context.Context, request *Message) (*Message, error) {
	var headerMD, trailerMD metadata.MD
	resp, err := wrapper.cli.CreateMessage(ctx, request, grpc.Header(&headerMD), grpc.Trailer(&trailerMD))
	gors.AddGRPCMetadata(ctx, headerMD, trailerMD, wrapper.options.OutgoingHeaderMatcher)
	return resp, err
}

func _Messaging_GetMessage_GORS_Handler(svc MessagingService, options *gors.Options, binding *binding.HttpRuleBinding) func(c *gin.Context) {
	return func(c *gin.Context) {
		var rpcMethodName = "/tests.example.message.v1.Messaging/GetMessage"
		var ctx = gors.NewContext(c, rpcMethodName)
		var req *GetMessageRequest
		var resp *Message
		var err error
		req = new(GetMessageRequest)
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

func _Messaging_CreateMessage_GORS_Handler(svc MessagingService, options *gors.Options, binding *binding.HttpRuleBinding) func(c *gin.Context) {
	return func(c *gin.Context) {
		var rpcMethodName = "/tests.example.message.v1.Messaging/CreateMessage"
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

func _Messaging_GetMessage_GORS_Handler_GET_71b8052a59ef2e1e6bb26f276891271b_Binding() *binding.HttpRuleBinding {
	return &binding.HttpRuleBinding{
		Path: []*binding.PathRule{
			{Name: "message_id", Type: "string"},
		},
		Query: []*binding.QueryRule{
			{Name: "user_id", Type: "string"},
		},
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
```

## 4. 实现服务
通过grpc定义接口：
```go
package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-leo/gors"
	"github.com/go-leo/gors/example/api/tests/example"
	"google.golang.org/protobuf/proto"
	"net"
	"net/http"
)

func main() {
	engine := gin.New()
	engine = gors.AppendRoutes(engine, example.MessagingServiceRoutes(NewMessagingService())...)
	srv := http.Server{Handler: engine}
	listen, err := net.Listen("tcp", ":8088")
	if err != nil {
		panic(err)
	}
	err = srv.Serve(listen)
	if err != nil {
		panic(err)
	}
}

type MessagingService struct {
}

func (m MessagingService) GetMessage(ctx context.Context, request *example.GetMessageRequest) (*example.Message, error) {
	return &example.Message{
		MessageId: request.GetMessageId(),
		UserId:    request.GetUserId(),
		Content:   "this is content",
		Maybe:     proto.String("this is maybe"),
	}, nil
}

func (m MessagingService) CreateMessage(ctx context.Context, message *example.Message) (*example.Message, error) {
	return message, nil
}

func NewMessagingService() example.MessagingService {
	return &MessagingService{}
}
```

## 5. 启动服务
```
go run main.go
```

## 6. 发生请求
```
curl --location --request GET 'http://localhost:8088/v1/messages/1234?user_id=1234' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)' \
--header 'Accept: */*' \
--header 'Host: localhost:8088' \
--header 'Connection: keep-alive'

{"messageId":"1234","userId":"1234","content":"this is content","maybe":"this is maybe"}%    
```

```
curl --location --request POST 'http://localhost:8088/v1/messages/1111' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)' \
--header 'Content-Type: application/json' \
--header 'Accept: */*' \
--header 'Host: localhost:8088' \
--header 'Connection: keep-alive' \
--data-raw '{
    "userId": "1234",
    "content": "this is content",
    "maybe": "this is maybe"
}'

{"messageId":"1111","userId":"1234","content":"this is content","maybe":"this is maybe"}%   
```

# 例子
[tests](example%2Fapi%2Ftests)