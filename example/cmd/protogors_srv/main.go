package main

import (
	"github.com/go-leo/gors/example/api/helloworld"
	"github.com/go-leo/gors/example/api/hellowrapper"
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
	engine = gors.AppendRoutes(engine, helloworld.GreeterServerRoutes(new(svc.HelloWorldService))...)

	engine = gors.AppendRoutes(engine,
		hellowrapper.GreeterServerRoutes(
			new(svc.HelloWrapperService),
			gors.ResponseWrapper(func(resp any) any {
				switch t := resp.(type) {
				case *gors.Status:
					return &hellowrapper.CommonReply{
						Code: t.Code,
						Msg:  t.Message,
					}
				default:
					a, _ := anypb.New(resp.(proto.Message))
					return &hellowrapper.CommonReply{
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
