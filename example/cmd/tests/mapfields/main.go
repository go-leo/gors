package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-leo/gors"
	"github.com/go-leo/gors/example/api/tests/mapfields"
	"net"
	"net/http"
)

func main() {
	engine := gin.New()
	engine = gors.AppendRoutes(engine, mapfields.MessagingServiceRoutes(NewMessagingService())...)
	srv := http.Server{Handler: engine}
	listen, err := net.Listen("tcp", ":8093")
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

func (m MessagingService) UpdateMessage(ctx context.Context, message *mapfields.Message) (*mapfields.Message, error) {
	return &mapfields.Message{
		MessageId:      message.GetMessageId(),
		AnotherMessage: message.GetAnotherMessage(),
		SubMessage:     message.GetSubMessage(),
		StringList:     message.GetStringList(),
		SubMessageList: message.GetSubMessageList(),
		ObjectList:     message.GetObjectList(),
		StringsMap:     message.GetStringsMap(),
		SubMessagesMap: message.GetSubMessagesMap(),
		ObjectsMap:     message.GetObjectsMap(),
	}, nil
}

func NewMessagingService() mapfields.MessagingService {
	return &MessagingService{}
}
