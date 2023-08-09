// Code generated by "gors -service StringString"; DO NOT EDIT.

package demo

import (
	gin "github.com/gin-gonic/gin"
	gors "github.com/go-leo/gors"
	http "net/http"
)

func StringStringRoutes(srv StringString, opts ...gors.Option) []gors.Route {
	options := gors.NewOptions(opts...)
	return []gors.Route{
		gors.NewRoute(http.MethodGet, "/api/StringString/Get", _StringString_GetStringString_Handler(srv, options)),
		gors.NewRoute(http.MethodPatch, "/api/StringString/Patch", _StringString_PatchStringString_Handler(srv, options)),
	}
}

func _StringString_GetStringString_Handler(srv StringString, options *gors.Options) func(c *gin.Context) {
	return func(c *gin.Context) {
		var rpcMethodName = "/demo.StringString/GetStringString"
		var ctx = gors.NewContext(c, rpcMethodName)
		var req string
		var resp string
		var err error
		if err = gors.RequestBind(
			ctx, &req, options.Tag(),
			gors.StringBinding,
		); err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler(), options.ResponseWrapper())
			return
		}
		resp, err = srv.GetStringString(ctx, req)
		if err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler(), options.ResponseWrapper())
			return
		}
		gors.ResponseRender(ctx, gors.StatusCode(ctx), resp, "text/go", gors.StringRender, options.ResponseWrapper())
	}
}

func _StringString_PatchStringString_Handler(srv StringString, options *gors.Options) func(c *gin.Context) {
	return func(c *gin.Context) {
		var rpcMethodName = "/demo.StringString/PatchStringString"
		var ctx = gors.NewContext(c, rpcMethodName)
		var req string
		var resp string
		var err error
		if err = gors.RequestBind(
			ctx, &req, options.Tag(),
			gors.StringBinding,
		); err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler(), options.ResponseWrapper())
			return
		}
		resp, err = srv.PatchStringString(ctx, req)
		if err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler(), options.ResponseWrapper())
			return
		}
		gors.ResponseRender(ctx, gors.StatusCode(ctx), resp, "application/go", gors.StringRender, options.ResponseWrapper())
	}
}
