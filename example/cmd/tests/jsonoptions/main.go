package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-leo/gors"
	"github.com/go-leo/gors/example/api/tests/jsonoptions"
	"net"
	"net/http"
)

func main() {
	engine := gin.New()
	engine = gors.AppendRoutes(engine, jsonoptions.MessagingServiceRoutes(NewMessagingService())...)
	srv := http.Server{Handler: engine}
	listen, err := net.Listen("tcp", ":8092")
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

func (m MessagingService) CreateMessage(ctx context.Context, message *jsonoptions.Message) (*jsonoptions.Message, error) {
	return &jsonoptions.Message{
		MessageId: message.GetMessageId(),
		BodyText:  message.GetBodyText(),
		NotUsed:   message.GetNotUsed(),
	}, nil
}

func (m MessagingService) UpdateMessage(ctx context.Context, message2 *jsonoptions.Message2) (*jsonoptions.Message2, error) {
	return &jsonoptions.Message2{
		MessageId: message2.GetMessageId(),
		BodyText:  message2.GetBodyText(),
		NotUsed:   message2.GetNotUsed(),
	}, nil
}

func NewMessagingService() jsonoptions.MessagingService {
	return &MessagingService{}
}
