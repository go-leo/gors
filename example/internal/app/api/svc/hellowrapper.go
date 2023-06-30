package svc

import (
	"context"
	"fmt"
	"github.com/go-leo/gors"
	"github.com/go-leo/gors/example/api/hellowrapper"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

var _ hellowrapper.GreeterServer = new(HelloWrapperService)

type HelloWrapperService struct {
	hellowrapper.UnimplementedGreeterServer
}

func (svc *HelloWrapperService) POSTProtoBufBindingProtoBufRender(ctx context.Context, request *hellowrapper.HelloRequest) (*hellowrapper.HelloReply, error) {
	return &hellowrapper.HelloReply{
		Message: fmt.Sprintf("hi %s, age: %d, salary: %f, token: %s", request.GetName(), request.GetAge(), request.GetSalary(), request.GetToken()),
	}, nil
}

func (svc *HelloWrapperService) POSTProtoJSONBindingProtoJSONRender(ctx context.Context, request *hellowrapper.HelloRequest) (*hellowrapper.HelloReply, error) {
	return &hellowrapper.HelloReply{
		Message: fmt.Sprintf("hi %s, age: %d, salary: %f, token: %s", request.GetName(), request.GetAge(), request.GetSalary(), request.GetToken()),
	}, nil
}

func (svc *HelloWrapperService) NotDefine(ctx context.Context, request *hellowrapper.HelloRequest) (*hellowrapper.HelloReply, error) {
	return &hellowrapper.HelloReply{
		Message: fmt.Sprintf("hi %s, age: %d, salary: %f, token: %s", request.GetName(), request.GetAge(), request.GetSalary(), request.GetToken()),
	}, nil
}

func (svc *HelloWrapperService) POSTError(ctx context.Context, request *hellowrapper.HelloRequest) (*hellowrapper.HelloReply, error) {
	return nil, gors.Error{
		StatusCode: http.StatusConflict,
		Code:       4000,
		Message:    "StatusConflict",
	}
}

func (svc *HelloWrapperService) POSTGRPCStatus(ctx context.Context, request *hellowrapper.HelloRequest) (*hellowrapper.HelloReply, error) {
	return nil, status.New(codes.PermissionDenied, "PermissionDenied").Err()
}
