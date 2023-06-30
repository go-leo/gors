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
	engine := gin.New()
	engine = gors.AppendRoutes(engine, demo.BytesBytesRoutes(new(svc.BytesBytesService))...)
	engine = gors.AppendRoutes(engine, demo.BytesReaderRoutes(new(svc.BytesReaderService))...)
	engine = gors.AppendRoutes(engine, demo.BytesStringRoutes(new(svc.BytesStringService))...)
	engine = gors.AppendRoutes(engine, demo.ReaderBytesRoutes(new(svc.ReaderBytesService))...)
	engine = gors.AppendRoutes(engine, demo.ReaderReaderRoutes(new(svc.ReaderReaderService))...)
	engine = gors.AppendRoutes(engine, demo.ReaderStringRoutes(new(svc.ReaderStringService))...)
	engine = gors.AppendRoutes(engine, demo.StringBytesRoutes(new(svc.StringBytesService))...)
	engine = gors.AppendRoutes(engine, demo.StringReaderRoutes(new(svc.StringReaderService))...)
	engine = gors.AppendRoutes(engine, demo.StringStringRoutes(new(svc.StringStringService))...)
	engine = gors.AppendRoutes(engine, demo.CustomBinderRenderRoutes(new(svc.Custom))...)
	engine = gors.AppendRoutes(engine, demo.ObjObjRoutes(new(svc.ObjObjService))...)
	engine = gors.AppendRoutes(engine, demo.ServiceRoutes(new(svc.Service))...)

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
