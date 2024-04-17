package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-leo/gors"
	"github.com/go-leo/gors/example/api/tests/rpctypes"
	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"net"
	"net/http"
)

func main() {
	engine := gin.New()
	engine = gors.AppendRoutes(engine, rpctypes.StatusServiceRoutes(NewMessagingService())...)
	srv := http.Server{Handler: engine}
	listen, err := net.Listen("tcp", ":8098")
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

func (m MessagingService) GetStatus(ctx context.Context, empty *emptypb.Empty) (*status.Status, error) {
	return &status.Status{
		Code:    int32(code.Code_OK),
		Message: code.Code_OK.String(),
		Details: nil,
	}, nil
}

func NewMessagingService() rpctypes.StatusService {
	return &MessagingService{}
}
