package main

import (
	"github.com/go-leo/gors/example/api/protodemo"
	"github.com/go-leo/gors/example/internal/app/api/svc"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/go-leo/gors"
)

func main() {
	engine := gin.New()
	engine = gors.AppendRoutes(
		engine,
		protodemo.ProtoDemoServerRoutes(
			new(svc.ProtoDemoServer),
			gors.ResponseWrapper(func(resp any) any {
				switch t := resp.(type) {
				case *gors.Status:
					return &protodemo.CommonReply{
						Code: t.Code,
						Msg:  t.Message,
					}
				default:
					a, _ := anypb.New(resp.(proto.Message))
					return &protodemo.CommonReply{
						Code: 200,
						Msg:  "ok",
						Data: a,
					}
				}
			}))...)
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
