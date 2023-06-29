package svc

import (
	"context"
	"fmt"
	"github.com/go-leo/gors/example/api/helloworld"
)

var _ helloworld.GreeterServer = new(HelloWorldService)

type HelloWorldService struct {
	helloworld.UnimplementedGreeterServer
}

func (svc *HelloWorldService) DELETEUriBindingJSONRender(ctx context.Context, request *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{
		Message: fmt.Sprintf("hi %s, age: %d, salary: %f", request.GetName(), request.GetAge(), request.GetSalary()),
	}, nil
}

func (svc *HelloWorldService) GETUriBindingIndentedJSONRender(ctx context.Context, request *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{
		Message: fmt.Sprintf("hi %s, age: %d, salary: %f", request.GetName(), request.GetAge(), request.GetSalary()),
	}, nil
}

func (svc *HelloWorldService) GETUriQueryBindingSecureJSONRender(ctx context.Context, request *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{
		Message: fmt.Sprintf("hi %s, age: %d, salary: %f", request.GetName(), request.GetAge(), request.GetSalary()),
	}, nil
}

func (svc *HelloWorldService) POSTHeaderFormPostBindingJSONPJSONRender(ctx context.Context, request *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{
		Message: fmt.Sprintf("hi %s, age: %d, salary: %f", request.GetName(), request.GetAge(), request.GetSalary()),
	}, nil
}

func (svc *HelloWorldService) PATCHHeaderProtoFormBindingPureJSONRender(ctx context.Context, request *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{
		Message: fmt.Sprintf("hi %s, age: %d, salary: %f", request.GetName(), request.GetAge(), request.GetSalary()),
	}, nil
}

func (svc *HelloWorldService) PUTHeaderJSONBindingAsciiJSONRender(ctx context.Context, request *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{
		Message: fmt.Sprintf("hi %s, age: %d, salary: %f", request.GetName(), request.GetAge(), request.GetSalary()),
	}, nil
}

func (svc *HelloWorldService) POSTProtoBufBindingProtoBufRender(ctx context.Context, request *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{
		Message: fmt.Sprintf("hi %s, age: %d, salary: %f", request.GetName(), request.GetAge(), request.GetSalary()),
	}, nil
}

func (svc *HelloWorldService) POSTProtoJSONBindingProtoJSONRender(ctx context.Context, request *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{
		Message: fmt.Sprintf("hi %s, age: %d, salary: %f", request.GetName(), request.GetAge(), request.GetSalary()),
	}, nil
}

func (svc *HelloWorldService) HeaderMsgPackBindingMsgPackRender(ctx context.Context, request *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{
		Message: fmt.Sprintf("hi %s, age: %d, salary: %f", request.GetName(), request.GetAge(), request.GetSalary()),
	}, nil
}

func (svc *HelloWorldService) POSTCustomBindingCustomRender(ctx context.Context, request *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{
		Message: fmt.Sprintf("hi %s, age: %d, salary: %f", request.GetName(), request.GetAge(), request.GetSalary()),
	}, nil
}

func (svc *HelloWorldService) NotDefine(ctx context.Context, request *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{
		Message: fmt.Sprintf("hi %s, age: %d, salary: %f", request.GetName(), request.GetAge(), request.GetSalary()),
	}, nil
}
