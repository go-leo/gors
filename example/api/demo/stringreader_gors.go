// Code generated by "gors -service StringReader"; DO NOT EDIT.

package demo

import (
	gin "github.com/gin-gonic/gin"
	gors "github.com/go-leo/gors"
	io "io"
	http "net/http"
)

func StringReaderRoutes(srv StringReader, opts ...gors.Option) []gors.Route {
	options := gors.New(opts...)
	return []gors.Route{
		gors.NewRoute(http.MethodGet, "/api/StringReader/Get", _StringReader_GetStringRender_Handler(srv, options)),
		gors.NewRoute(http.MethodOptions, "/api/StringReader/Options", _StringReader_OptionsStringReader_Handler(srv, options)),
	}
}
func _StringReader_GetStringRender_Handler(srv StringReader, options *gors.Options) func(c *gin.Context) {
	return func(c *gin.Context) {
		var rpcMethodName = "/demo.StringReader/GetStringRender"
		var ctx = gors.NewContext(c, rpcMethodName)
		var req string
		var resp io.Reader
		var err error
		if err = gors.RequestBind(
			ctx, &req, options.Tag,
			gors.StringBinding,
		); err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler, options.ResponseWrapper)
			return
		}
		resp, err = srv.GetStringRender(ctx, req)
		if err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler, options.ResponseWrapper)
			return
		}
		gors.ResponseRender(ctx, gors.StatusCode(ctx), resp, "video/mpeg4", gors.ReaderRender, options.ResponseWrapper)
	}
}

func _StringReader_OptionsStringReader_Handler(srv StringReader, options *gors.Options) func(c *gin.Context) {
	return func(c *gin.Context) {
		var rpcMethodName = "/demo.StringReader/OptionsStringReader"
		var ctx = gors.NewContext(c, rpcMethodName)
		var req string
		var resp io.Reader
		var err error
		if err = gors.RequestBind(
			ctx, &req, options.Tag,
			gors.StringBinding,
		); err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler, options.ResponseWrapper)
			return
		}
		resp, err = srv.OptionsStringReader(ctx, req)
		if err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler, options.ResponseWrapper)
			return
		}
		gors.ResponseRender(ctx, gors.StatusCode(ctx), resp, "video/mpeg4", gors.ReaderRender, options.ResponseWrapper)
	}
}
