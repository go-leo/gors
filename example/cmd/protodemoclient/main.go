package main

import (
	"github.com/go-leo/gors/example/api/protodemo"
	"github.com/go-leo/gors/example/internal/app/api/svc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/go-leo/gors"
)

func main() {
	go func() {
		server := grpc.NewServer()
		protodemo.RegisterProtoDemoServer(server, new(svc.ProtoDemoServer))
		listen, err := net.Listen("tcp", ":9090")
		if err != nil {
			panic(err)
		}
		err = server.Serve(listen)
		if err != nil {
			panic(err)
		}
	}()

	time.Sleep(time.Second)
	dial, err := grpc.Dial(":9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	client := protodemo.NewProtoDemoClient(dial)

	engine := gin.New()
	engine = gors.AppendRoutes(engine, protodemo.ProtoDemoClientRoutes(client)...)
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
