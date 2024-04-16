package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-leo/gors"
	"github.com/go-leo/gors/example/api/tests/additional_bindings"
	"net"
	"net/http"
)

func main() {
	engine := gin.New()
	engine = gors.AppendRoutes(engine, additional_bindings.MessagingServiceRoutes(NewMessagingService())...)
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

func (m MessagingService) UpdateMessage(ctx context.Context, message *additional_bindings.Message) (*additional_bindings.Message, error) {
	return &additional_bindings.Message{
		MessageId: message.GetMessageId(),
		Text:      message.GetText(),
	}, nil
}

func NewMessagingService() additional_bindings.MessagingService {
	return &MessagingService{}
}
