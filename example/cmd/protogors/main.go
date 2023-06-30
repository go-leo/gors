package main

import (
	"github.com/go-leo/gors/example/api/protodemo"
	"github.com/go-leo/gors/example/internal/app/api/svc"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/go-leo/gors"
)

func main() {
	engine := gin.New()
	engine = gors.AppendRoutes(engine, protodemo.ProtoDemoServerRoutes(new(svc.ProtoDemoServer))...)
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
