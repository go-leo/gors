package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-leo/gors"
	"github.com/go-leo/gors/example/api/tests/enumoptions"
	"net"
	"net/http"
)

func main() {
	engine := gin.New()
	engine = gors.AppendRoutes(engine, enumoptions.MessagingServiceRoutes(NewMessagingService())...)
	srv := http.Server{Handler: engine}
	listen, err := net.Listen("tcp", ":8091")
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

func (m MessagingService) CreateMessage(ctx context.Context, message *enumoptions.Message) (*enumoptions.Message, error) {
	return &enumoptions.Message{
		Kind:      message.GetKind(),
		MessageId: message.GetMessageId(),
		BodyText:  message.GetBodyText(),
	}, nil
}

func NewMessagingService() enumoptions.MessagingService {
	return &MessagingService{}
}
