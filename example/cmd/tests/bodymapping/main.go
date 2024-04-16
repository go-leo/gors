package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-leo/gors"
	"github.com/go-leo/gors/example/api/tests/bodymapping"
	"google.golang.org/protobuf/types/known/emptypb"
	"net"
	"net/http"
)

func main() {
	engine := gin.New()
	engine = gors.AppendRoutes(engine, bodymapping.MessagingServiceRoutes(NewMessagingService())...)
	srv := http.Server{Handler: engine}
	listen, err := net.Listen("tcp", ":8090")
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

func (m MessagingService) UpdateMessage(ctx context.Context, message *bodymapping.Message) (*bodymapping.Message, error) {
	return &bodymapping.Message{
		MessageId: message.GetMessageId(),
		Text:      message.GetText(),
	}, nil
}

func (m MessagingService) DoAny(ctx context.Context, empty *emptypb.Empty) (*emptypb.Empty, error) {
	return new(emptypb.Empty), nil
}

func NewMessagingService() bodymapping.MessagingService {
	return &MessagingService{}
}
