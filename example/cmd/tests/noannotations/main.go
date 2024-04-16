package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-leo/gors"
	"github.com/go-leo/gors/example/api/tests/noannotations"
	"net"
	"net/http"
)

func main() {
	engine := gin.New()
	engine = gors.AppendRoutes(engine, noannotations.Messaging1ServiceRoutes(NewMessaging1Service())...)
	engine = gors.AppendRoutes(engine, noannotations.Messaging2ServiceRoutes(NewMessaging2Service())...)
	srv := http.Server{Handler: engine}
	listen, err := net.Listen("tcp", ":8094")
	if err != nil {
		panic(err)
	}
	err = srv.Serve(listen)
	if err != nil {
		panic(err)
	}
}

type Messaging1Service struct {
}

func (m Messaging1Service) UpdateMessage(ctx context.Context, message *noannotations.Message) (*noannotations.Message, error) {
	return message, nil
}

func NewMessaging1Service() noannotations.Messaging1Service {
	return &Messaging1Service{}
}

type Messaging2Service struct {
}

func (m Messaging2Service) UpdateMessage(ctx context.Context, message *noannotations.Message) (*noannotations.Message, error) {
	return message, nil
}

func NewMessaging2Service() noannotations.Messaging2Service {
	return &Messaging2Service{}
}
