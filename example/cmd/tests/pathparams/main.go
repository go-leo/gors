package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-leo/gors"
	"github.com/go-leo/gors/example/api/tests/pathparams"
	"github.com/go-leo/gox/convx"
	"github.com/go-leo/gox/mathx/randx"
	"google.golang.org/protobuf/proto"
	"net"
	"net/http"
	"strings"
)

func main() {
	engine := gin.New()
	engine = gors.AppendRoutes(engine, pathparams.MessagingServiceRoutes(NewMessagingService())...)
	srv := http.Server{Handler: engine}
	listen, err := net.Listen("tcp", ":8096")
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

func (m MessagingService) GetMessage(ctx context.Context, request *pathparams.GetMessageRequest) (*pathparams.Message, error) {
	return &pathparams.Message{
		MessageId: request.GetMessageId(),
		UserId:    request.GetUserId(),
		Content:   randx.HexString(16),
		Maybe:     proto.String(randx.WordString(16)),
	}, nil
}

func (m MessagingService) GetUserMessage(ctx context.Context, request *pathparams.GetMessageRequest) (*pathparams.Message, error) {
	return &pathparams.Message{
		MessageId: request.GetMessageId(),
		UserId:    request.GetUserId(),
		Content:   randx.HexString(16),
		Maybe:     proto.String(randx.WordString(16)),
	}, nil
}

func (m MessagingService) CreateMessage(ctx context.Context, message *pathparams.Message) (*pathparams.Message, error) {
	return &pathparams.Message{
		MessageId: message.GetMessageId(),
		UserId:    message.GetUserId(),
		Content:   message.GetContent(),
		Maybe:     proto.String(message.GetMaybe()),
	}, nil
}

func (m MessagingService) UpdateMessage(ctx context.Context, request *pathparams.UpdateMessageRequest) (*pathparams.Message, error) {
	split := strings.Split(request.GetName(), "/")
	return &pathparams.Message{
		MessageId: split[1],
		UserId:    convx.ToUint64(split[3]),
		Content:   request.GetContent(),
		Maybe:     proto.String(request.GetMaybe()),
	}, nil
}

func NewMessagingService() pathparams.MessagingService {
	return &MessagingService{}
}
