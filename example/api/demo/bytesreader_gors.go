// Code generated by "gors -service BytesReader"; DO NOT EDIT.

package demo

import (
	gin "github.com/gin-gonic/gin"
	gors "github.com/go-leo/gors"
	io "io"
	http "net/http"
)

func BytesReaderRoutes(srv BytesReader, opts ...gors.Option) []gors.Route {
	options := gors.NewOptions(opts...)
	return []gors.Route{
		gors.NewRoute(http.MethodGet, "/api/BytesReader/Get", _BytesReader_GetBytesReader_Handler(srv, options)),
		gors.NewRoute(http.MethodPatch, "/api/BytesReader/Patch", _BytesReader_PatchBytesReader_Handler(srv, options)),
	}
}

func _BytesReader_GetBytesReader_Handler(srv BytesReader, options *gors.Options) func(c *gin.Context) {
	return func(c *gin.Context) {
		var rpcMethodName = "/demo.BytesReader/GetBytesReader"
		var ctx = gors.NewContext(c, rpcMethodName)
		var req []byte
		var resp io.Reader
		var err error
		if err = gors.RequestBind(
			ctx, &req, options.Tag,
			gors.BytesBinding,
		); err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler, options.ResponseWrapper)
			return
		}
		resp, err = srv.GetBytesReader(ctx, req)
		if err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler, options.ResponseWrapper)
			return
		}
		gors.ResponseRender(ctx, gors.StatusCode(ctx), resp, "video/mpeg4", gors.ReaderRender, options.ResponseWrapper)
	}
}

func _BytesReader_PatchBytesReader_Handler(srv BytesReader, options *gors.Options) func(c *gin.Context) {
	return func(c *gin.Context) {
		var rpcMethodName = "/demo.BytesReader/PatchBytesReader"
		var ctx = gors.NewContext(c, rpcMethodName)
		var req []byte
		var resp io.Reader
		var err error
		if err = gors.RequestBind(
			ctx, &req, options.Tag,
			gors.BytesBinding,
		); err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler, options.ResponseWrapper)
			return
		}
		resp, err = srv.PatchBytesReader(ctx, req)
		if err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler, options.ResponseWrapper)
			return
		}
		gors.ResponseRender(ctx, gors.StatusCode(ctx), resp, "video/mpeg4", gors.ReaderRender, options.ResponseWrapper)
	}
}
