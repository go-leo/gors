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
	var routes []gors.Route
	routes = append(routes, demo.BindingRenderRoutes(new(svc.BindingRenderService))...)
	routes = append(routes, demo.BytesBytesRoutes(new(svc.BytesBytesService))...)
	routes = append(routes, demo.BytesReaderRoutes(new(svc.BytesReaderService))...)
	routes = append(routes, demo.BytesStringRoutes(new(svc.BytesStringService))...)
	routes = append(routes, demo.ReaderBytesRoutes(new(svc.ReaderBytesService))...)
	routes = append(routes, demo.ReaderReaderRoutes(new(svc.ReaderReaderService))...)
	routes = append(routes, demo.ReaderStringRoutes(new(svc.ReaderStringService))...)
	routes = append(routes, demo.StringBytesRoutes(new(svc.StringBytesService))...)
	routes = append(routes, demo.StringReaderRoutes(new(svc.StringReaderService))...)
	routes = append(routes, demo.StringStringRoutes(new(svc.StringStringService))...)

	engine := gin.New()
	for _, route := range routes {
		engine.Handle(route.HTTPMethod, route.Path, route.HandlerFunc)
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
