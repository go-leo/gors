package svc

import (
	"context"
	"fmt"
	"github.com/go-leo/gors"
	"github.com/go-leo/gors/example/api/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"net/http"
)

var _ helloworld.GreeterServer = new(HelloWorldService)

type HelloWorldService struct {
	helloworld.UnimplementedGreeterServer
}

func (svc *HelloWorldService) DELETEUriBindingJSONRender(ctx context.Context, request *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{
		Message: fmt.Sprintf("hi %s, age: %d, salary: %f, token: %s", request.GetName(), request.GetAge(), request.GetSalary(), request.GetToken()),
	}, nil
}

func (svc *HelloWorldService) GETUriBindingIndentedJSONRender(ctx context.Context, request *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{
		Message: fmt.Sprintf("hi %s, age: %d, salary: %f, token: %s", request.GetName(), request.GetAge(), request.GetSalary(), request.GetToken()),
	}, nil
}

func (svc *HelloWorldService) GETUriQueryBindingSecureJSONRender(ctx context.Context, request *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{
		Message: fmt.Sprintf("hi %s, age: %d, salary: %f, token: %s", request.GetName(), request.GetAge(), request.GetSalary(), request.GetToken()),
	}, nil
}

func (svc *HelloWorldService) POSTHeaderFormPostBindingJSONPJSONRender(ctx context.Context, request *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{
		Message: fmt.Sprintf("hi %s, age: %d, salary: %f, token: %s", request.GetName(), request.GetAge(), request.GetSalary(), request.GetToken()),
	}, nil
}

func (svc *HelloWorldService) PATCHHeaderProtoFormBindingPureJSONRender(ctx context.Context, request *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{
		Message: fmt.Sprintf("hi %s, age: %d, salary: %f, token: %s", request.GetName(), request.GetAge(), request.GetSalary(), request.GetToken()),
	}, nil
}

func (svc *HelloWorldService) PUTHeaderJSONBindingAsciiJSONRender(ctx context.Context, request *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{
		Message: fmt.Sprintf("hi %s, age: %d, salary: %f, token: %s", request.GetName(), request.GetAge(), request.GetSalary(), request.GetToken()),
	}, nil
}

func (svc *HelloWorldService) POSTProtoBufBindingProtoBufRender(ctx context.Context, request *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{
		Message: fmt.Sprintf("hi %s, age: %d, salary: %f, token: %s", request.GetName(), request.GetAge(), request.GetSalary(), request.GetToken()),
	}, nil
}

func (svc *HelloWorldService) POSTProtoJSONBindingProtoJSONRender(ctx context.Context, request *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{
		Message: fmt.Sprintf("hi %s, age: %d, salary: %f, token: %s", request.GetName(), request.GetAge(), request.GetSalary(), request.GetToken()),
	}, nil
}

func (svc *HelloWorldService) HeaderMsgPackBindingMsgPackRender(ctx context.Context, request *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{
		Message: fmt.Sprintf("hi %s, age: %d, salary: %f, token: %s", request.GetName(), request.GetAge(), request.GetSalary(), request.GetToken()),
	}, nil
}

func (svc *HelloWorldService) POSTCustomBindingCustomRender(ctx context.Context, request *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{
		Message: fmt.Sprintf("hi %s, age: %d, salary: %f, token: %s", request.GetName(), request.GetAge(), request.GetSalary(), request.GetToken()),
	}, nil
}

func (svc *HelloWorldService) NotDefine(ctx context.Context, request *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{
		Message: fmt.Sprintf("hi %s, age: %d, salary: %f, token: %s", request.GetName(), request.GetAge(), request.GetSalary(), request.GetToken()),
	}, nil
}

func (svc *HelloWorldService) POSTSetHeaderTrailer(ctx context.Context, request *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	if err := grpc.SetHeader(ctx, metadata.New(map[string]string{"SetHeader": "SetHeader"})); err != nil {
		return nil, err
	}
	if err := grpc.SendHeader(ctx, metadata.New(map[string]string{"SendHeader": "SendHeader"})); err != nil {
		return nil, err
	}
	if err := grpc.SetTrailer(ctx, metadata.New(map[string]string{"SetTrailer": "SetTrailer"})); err != nil {
		return nil, err
	}
	return &helloworld.HelloReply{
		Message: fmt.Sprintf("hi %s, age: %d, salary: %f, token: %s", request.GetName(), request.GetAge(), request.GetSalary(), request.GetToken()),
	}, nil
}

func (svc *HelloWorldService) POSTError(ctx context.Context, request *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return nil, gors.Error{
		StatusCode: http.StatusConflict,
		Code:       4000,
		Message:    "StatusConflict",
	}
}

func (svc *HelloWorldService) POSTGRPCStatus(ctx context.Context, request *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return nil, status.New(codes.PermissionDenied, "PermissionDenied").Err()
}
