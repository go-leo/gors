package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-leo/gors"
	"github.com/go-leo/gors/example/api/tests/example"
	"google.golang.org/protobuf/proto"
	"net"
	"net/http"
)

func main() {
	engine := gin.New()
	engine = gors.AppendRoutes(engine, example.MessagingServiceRoutes(NewMessagingService())...)
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

func (m MessagingService) GetMessages(ctx context.Context, request *example.GetMessageRequest) (*example.Message, error) {
	return &example.Message{
		MessageId: request.GetMessageId(),
		UserId:    request.GetUserId(),
		Content:   "this is content",
		Maybe:     proto.String("this is maybe"),
		Tags:      request.GetTags(),
	}, nil
}

func (m MessagingService) GetMessage(ctx context.Context, request *example.GetMessageRequest) (*example.Message, error) {
	return &example.Message{
		MessageId: request.GetMessageId(),
		UserId:    request.GetUserId(),
		Content:   "this is content",
		Maybe:     proto.String("this is maybe"),
		Tags:      request.GetTags(),
	}, nil
}

func (m MessagingService) CreateMessage(ctx context.Context, message *example.Message) (*example.Message, error) {
	return message, nil
}

func (m MessagingService) UpdateMessage(ctx context.Context, message *example.Message) (*example.Message, error) {
	return &example.Message{
		MessageId: message.GetMessageId(),
		UserId:    0,
		Content:   message.GetContent(),
		Maybe:     nil,
		Tags:      nil,
	}, nil
}

func NewMessagingService() example.MessagingService {
	return &MessagingService{}
}
