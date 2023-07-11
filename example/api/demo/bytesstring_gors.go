// Code generated by "gors -service BytesString"; DO NOT EDIT.

package demo

import (
	gin "github.com/gin-gonic/gin"
	gors "github.com/go-leo/gors"
	http "net/http"
)

func BytesStringRoutes(srv BytesString, opts ...gors.Option) []gors.Route {
	options := gors.New(opts...)
	return []gors.Route{
		gors.NewRoute(http.MethodGet, "/api/BytesString/Get", _BytesString_GetBytesString_Handler(srv, options)),
		gors.NewRoute(http.MethodPut, "/api/BytesString/Put", _BytesString_PutBytesString_Handler(srv, options)),
	}
}
func _BytesString_GetBytesString_Handler(srv BytesString, options *gors.Options) func(c *gin.Context) {
	return func(c *gin.Context) {
		var rpcMethodName = "/demo.BytesString/GetBytesString"
		var ctx = gors.NewContext(c, rpcMethodName)
		var req []byte
		var resp string
		var err error
		if err = gors.RequestBind(
			ctx, &req, options.Tag,
			gors.BytesBinding,
		); err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler, options.ResponseWrapper)
			return
		}
		resp, err = srv.GetBytesString(ctx, req)
		if err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler, options.ResponseWrapper)
			return
		}
		gors.ResponseRender(ctx, gors.StatusCode(ctx), resp, "text/html", gors.HTMLRender, options.ResponseWrapper)
	}
}

func _BytesString_PutBytesString_Handler(srv BytesString, options *gors.Options) func(c *gin.Context) {
	return func(c *gin.Context) {
		var rpcMethodName = "/demo.BytesString/PutBytesString"
		var ctx = gors.NewContext(c, rpcMethodName)
		var req []byte
		var resp string
		var err error
		if err = gors.RequestBind(
			ctx, &req, options.Tag,
			gors.BytesBinding,
		); err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler, options.ResponseWrapper)
			return
		}
		resp, err = srv.PutBytesString(ctx, req)
		if err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler, options.ResponseWrapper)
			return
		}
		gors.ResponseRender(ctx, gors.StatusCode(ctx), resp, "", gors.RedirectRender, options.ResponseWrapper)
	}
}
