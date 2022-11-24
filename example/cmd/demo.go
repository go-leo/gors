package main

import (
	"net"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/go-leo/gors"
	"github.com/go-leo/gors/example/api/demo"
	"github.com/go-leo/gors/example/internal/app/api/svc"
)

func main() {
	var routers []gors.Router
	routers = append(routers, demo.BindingRenderRouters(new(svc.BindingRenderService))...)
	routers = append(routers, demo.BytesBytesRouters(new(svc.BytesBytesService))...)
	routers = append(routers, demo.BytesReaderRouters(new(svc.BytesReaderService))...)
	routers = append(routers, demo.BytesStringRouters(new(svc.BytesStringService))...)
	routers = append(routers, demo.ReaderBytesRouters(new(svc.ReaderBytesService))...)
	routers = append(routers, demo.ReaderReaderRouters(new(svc.ReaderReaderService))...)
	routers = append(routers, demo.ReaderStringRouters(new(svc.ReaderStringService))...)
	routers = append(routers, demo.StringBytesRouters(new(svc.StringBytesService))...)
	routers = append(routers, demo.StringReaderRouters(new(svc.StringReaderService))...)
	routers = append(routers, demo.StringStringRouters(new(svc.StringStringService))...)

	engine := gin.New()
	for _, router := range routers {
		engine.Handle(router.HTTPMethod, router.Path, router.HandlerFunc)
	}

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
