// Code generated by "gors -service ReaderBytes"; DO NOT EDIT.

package demo

import (
	gin "github.com/gin-gonic/gin"
	gors "github.com/go-leo/gors"
	io "io"
	http "net/http"
)

func ReaderBytesRoutes(srv ReaderBytes, opts ...gors.Option) []gors.Route {
	options := gors.NewOptions(opts...)
	return []gors.Route{
		gors.NewRoute(http.MethodGet, "/api/ReaderBytes/Get", _ReaderBytes_GetReaderBytes_Handler(srv, options)),
		gors.NewRoute(http.MethodPost, "/api/ReaderBytes/Post", _ReaderBytes_PostReaderBytes_Handler(srv, options)),
	}
}

func _ReaderBytes_GetReaderBytes_Handler(srv ReaderBytes, options *gors.Options) func(c *gin.Context) {
	return func(c *gin.Context) {
		var rpcMethodName = "/demo.ReaderBytes/GetReaderBytes"
		var ctx = gors.NewContext(c, rpcMethodName)
		var req io.Reader
		var resp []byte
		var err error
		if err = gors.RequestBind(
			ctx, &req, options.Tag,
			gors.ReaderBinding,
		); err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler, options.ResponseWrapper)
			return
		}
		resp, err = srv.GetReaderBytes(ctx, req)
		if err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler, options.ResponseWrapper)
			return
		}
		gors.ResponseRender(ctx, gors.StatusCode(ctx), resp, "", gors.BytesRender, options.ResponseWrapper)
	}
}

func _ReaderBytes_PostReaderBytes_Handler(srv ReaderBytes, options *gors.Options) func(c *gin.Context) {
	return func(c *gin.Context) {
		var rpcMethodName = "/demo.ReaderBytes/PostReaderBytes"
		var ctx = gors.NewContext(c, rpcMethodName)
		var req io.Reader
		var resp []byte
		var err error
		if err = gors.RequestBind(
			ctx, &req, options.Tag,
			gors.ReaderBinding,
		); err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler, options.ResponseWrapper)
			return
		}
		resp, err = srv.PostReaderBytes(ctx, req)
		if err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler, options.ResponseWrapper)
			return
		}
		gors.ResponseRender(ctx, gors.StatusCode(ctx), resp, "text/go", gors.BytesRender, options.ResponseWrapper)
	}
}
