package svc

import (
	"context"
	"github.com/go-leo/gors/example/api/protoservice"
)

var _ protoservice.ProtoServiceServer = new(ProtoService)

type ProtoService struct {
	protoservice.UnimplementedProtoServiceServer
}

func (p ProtoService) Method(ctx context.Context, request *protoservice.HelloRequest) (*protoservice.HelloReply, error) {
	return &protoservice.HelloReply{Message: "hi " + request.GetName()}, nil
}
