// Code generated by "gors -service StringBytes"; DO NOT EDIT.

package demo

import (
	gin "github.com/gin-gonic/gin"
	gors "github.com/go-leo/gors"
	http "net/http"
)

func StringBytesRoutes(srv StringBytes, opts ...gors.Option) []gors.Route {
	options := gors.NewOptions(opts...)
	return []gors.Route{
		gors.NewRoute(http.MethodGet, "/api/StringBytes/Get", _StringBytes_GetStringBytes_Handler(srv, options)),
		gors.NewRoute(http.MethodOptions, "/api/StringBytes/Options", _StringBytes_OptionsStringBytes_Handler(srv, options)),
	}
}

func _StringBytes_GetStringBytes_Handler(srv StringBytes, options *gors.Options) func(c *gin.Context) {
	return func(c *gin.Context) {
		var rpcMethodName = "/demo.StringBytes/GetStringBytes"
		var ctx = gors.NewContext(c, rpcMethodName)
		var req string
		var resp []byte
		var err error
		if err = gors.RequestBind(
			ctx, &req, options.Tag,
			gors.StringBinding,
		); err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler, options.ResponseWrapper)
			return
		}
		resp, err = srv.GetStringBytes(ctx, req)
		if err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler, options.ResponseWrapper)
			return
		}
		gors.ResponseRender(ctx, gors.StatusCode(ctx), resp, "", gors.BytesRender, options.ResponseWrapper)
	}
}

func _StringBytes_OptionsStringBytes_Handler(srv StringBytes, options *gors.Options) func(c *gin.Context) {
	return func(c *gin.Context) {
		var rpcMethodName = "/demo.StringBytes/OptionsStringBytes"
		var ctx = gors.NewContext(c, rpcMethodName)
		var req string
		var resp []byte
		var err error
		if err = gors.RequestBind(
			ctx, &req, options.Tag,
			gors.StringBinding,
		); err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler, options.ResponseWrapper)
			return
		}
		resp, err = srv.OptionsStringBytes(ctx, req)
		if err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler, options.ResponseWrapper)
			return
		}
		gors.ResponseRender(ctx, gors.StatusCode(ctx), resp, "", gors.BytesRender, options.ResponseWrapper)
	}
}
