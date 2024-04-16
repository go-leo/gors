package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-leo/gors"
	"github.com/go-leo/gors/example/api/tests/allofwrap"
	"net"
	"net/http"
)

func main() {
	engine := gin.New()
	engine = gors.AppendRoutes(engine, allofwrap.MessagingServiceRoutes(NewMessagingService())...)
	srv := http.Server{Handler: engine}
	listen, err := net.Listen("tcp", ":8089")
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

func (m MessagingService) UpdateMessage(ctx context.Context, message *allofwrap.Message) (*allofwrap.Message, error) {
	return &allofwrap.Message{
		Sub:       message.GetSub(),
		SubInput:  nil,
		SubOutput: message.GetSubInput(),
		SubDesc:   message.GetSubDesc(),
		Subs:      []*allofwrap.Message_Sub{message.GetSubInput()},
		MessageId: message.GetMessageId(),
	}, nil
}

func NewMessagingService() allofwrap.MessagingService {
	return &MessagingService{}
}
