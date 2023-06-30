package svc

import (
	"context"
	"fmt"
	"github.com/go-leo/gors"
	"github.com/go-leo/gors/example/api/protodemo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"net/http"
)

var _ protodemo.ProtoDemoServer = new(ProtoDemoServer)

type ProtoDemoServer struct {
	protodemo.UnsafeProtoDemoServer
}

func (svc *ProtoDemoServer) DELETEUriBindingJSONRender(ctx context.Context, request *protodemo.HelloRequest) (*protodemo.HelloReply, error) {
	return &protodemo.HelloReply{
		Message: fmt.Sprintf("hi %s, age: %d, salary: %f, token: %s", request.GetName(), request.GetAge(), request.GetSalary(), request.GetToken()),
	}, nil
}

func (svc *ProtoDemoServer) GETUriBindingIndentedJSONRender(ctx context.Context, request *protodemo.HelloRequest) (*protodemo.HelloReply, error) {
	return &protodemo.HelloReply{
		Message: fmt.Sprintf("hi %s, age: %d, salary: %f, token: %s", request.GetName(), request.GetAge(), request.GetSalary(), request.GetToken()),
	}, nil
}

func (svc *ProtoDemoServer) GETUriQueryBindingSecureJSONRender(ctx context.Context, request *protodemo.HelloRequest) (*protodemo.HelloReply, error) {
	return &protodemo.HelloReply{
		Message: fmt.Sprintf("hi %s, age: %d, salary: %f, token: %s", request.GetName(), request.GetAge(), request.GetSalary(), request.GetToken()),
	}, nil
}

func (svc *ProtoDemoServer) POSTHeaderFormPostBindingJSONPJSONRender(ctx context.Context, request *protodemo.HelloRequest) (*protodemo.HelloReply, error) {
	return &protodemo.HelloReply{
		Message: fmt.Sprintf("hi %s, age: %d, salary: %f, token: %s", request.GetName(), request.GetAge(), request.GetSalary(), request.GetToken()),
	}, nil
}

func (svc *ProtoDemoServer) PATCHHeaderProtoFormBindingPureJSONRender(ctx context.Context, request *protodemo.HelloRequest) (*protodemo.HelloReply, error) {
	return &protodemo.HelloReply{
		Message: fmt.Sprintf("hi %s, age: %d, salary: %f, token: %s", request.GetName(), request.GetAge(), request.GetSalary(), request.GetToken()),
	}, nil
}

func (svc *ProtoDemoServer) PUTHeaderJSONBindingAsciiJSONRender(ctx context.Context, request *protodemo.HelloRequest) (*protodemo.HelloReply, error) {
	return &protodemo.HelloReply{
		Message: fmt.Sprintf("hi %s, age: %d, salary: %f, token: %s", request.GetName(), request.GetAge(), request.GetSalary(), request.GetToken()),
	}, nil
}

func (svc *ProtoDemoServer) POSTProtoBufBindingProtoBufRender(ctx context.Context, request *protodemo.HelloRequest) (*protodemo.HelloReply, error) {
	return &protodemo.HelloReply{
		Message: fmt.Sprintf("hi %s, age: %d, salary: %f, token: %s", request.GetName(), request.GetAge(), request.GetSalary(), request.GetToken()),
	}, nil
}

func (svc *ProtoDemoServer) POSTProtoJSONBindingProtoJSONRender(ctx context.Context, request *protodemo.HelloRequest) (*protodemo.HelloReply, error) {
	return &protodemo.HelloReply{
		Message: fmt.Sprintf("hi %s, age: %d, salary: %f, token: %s", request.GetName(), request.GetAge(), request.GetSalary(), request.GetToken()),
	}, nil
}

func (svc *ProtoDemoServer) HeaderMsgPackBindingMsgPackRender(ctx context.Context, request *protodemo.HelloRequest) (*protodemo.HelloReply, error) {
	return &protodemo.HelloReply{
		Message: fmt.Sprintf("hi %s, age: %d, salary: %f, token: %s", request.GetName(), request.GetAge(), request.GetSalary(), request.GetToken()),
	}, nil
}

func (svc *ProtoDemoServer) POSTCustomBindingCustomRender(ctx context.Context, request *protodemo.HelloRequest) (*protodemo.HelloReply, error) {
	return &protodemo.HelloReply{
		Message: fmt.Sprintf("hi %s, age: %d, salary: %f, token: %s", request.GetName(), request.GetAge(), request.GetSalary(), request.GetToken()),
	}, nil
}

func (svc *ProtoDemoServer) NotDefine(ctx context.Context, request *protodemo.HelloRequest) (*protodemo.HelloReply, error) {
	return &protodemo.HelloReply{
		Message: fmt.Sprintf("hi %s, age: %d, salary: %f, token: %s", request.GetName(), request.GetAge(), request.GetSalary(), request.GetToken()),
	}, nil
}

func (svc *ProtoDemoServer) POSTSetHeaderTrailer(ctx context.Context, request *protodemo.HelloRequest) (*protodemo.HelloReply, error) {
	if err := grpc.SetHeader(ctx, metadata.New(map[string]string{"SetHeader": "SetHeader"})); err != nil {
		return nil, err
	}
	if err := grpc.SendHeader(ctx, metadata.New(map[string]string{"SendHeader": "SendHeader"})); err != nil {
		return nil, err
	}
	if err := grpc.SetTrailer(ctx, metadata.New(map[string]string{"SetTrailer": "SetTrailer"})); err != nil {
		return nil, err
	}
	return &protodemo.HelloReply{
		Message: fmt.Sprintf("hi %s, age: %d, salary: %f, token: %s", request.GetName(), request.GetAge(), request.GetSalary(), request.GetToken()),
	}, nil
}

func (svc *ProtoDemoServer) POSTError(ctx context.Context, request *protodemo.HelloRequest) (*protodemo.HelloReply, error) {
	return nil, gors.Error{
		StatusCode: http.StatusConflict,
		Code:       4000,
		Message:    "StatusConflict",
	}
}

func (svc *ProtoDemoServer) POSTGRPCStatus(ctx context.Context, request *protodemo.HelloRequest) (*protodemo.HelloReply, error) {
	return nil, status.New(codes.PermissionDenied, "PermissionDenied").Err()
}
