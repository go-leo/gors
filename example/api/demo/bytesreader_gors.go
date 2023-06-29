// Code generated by "gors -service BytesReader"; DO NOT EDIT.

package demo

import (
	gin "github.com/gin-gonic/gin"
	gors "github.com/go-leo/gors"
	io "io"
	http "net/http"
)

func BytesReaderRoutes(srv BytesReader, opts ...gors.Option) []gors.Route {
	options := gors.New(opts...)
	_ = options
	return []gors.Route{
		gors.NewRoute(
			http.MethodGet,
			"/api/BytesReader/Get",
			func(c *gin.Context) {
				var rpcMethodName = "/demo.BytesReader/GetBytesReader"
				var ctx = gors.NewContext(c, rpcMethodName)
				var req []byte
				var resp io.Reader
				var err error
				var body []byte
				body, err = io.ReadAll(c.Request.Body)
				if err != nil {
					gors.ErrorRender(ctx, err, options.ErrorHandler)
					return
				}
				req = body
				resp, err = srv.GetBytesReader(ctx, req)
				if err != nil {
					gors.ErrorRender(ctx, err, options.ErrorHandler)
					return
				}
				gors.ResponseRender(ctx, gors.StatusCode(ctx), resp, "video/mpeg4", gors.ReaderRender, options.ResponseWrapper)
			},
		),
		gors.NewRoute(
			http.MethodPatch,
			"/api/BytesReader/Patch",
			func(c *gin.Context) {
				var rpcMethodName = "/demo.BytesReader/PatchBytesReader"
				var ctx = gors.NewContext(c, rpcMethodName)
				var req []byte
				var resp io.Reader
				var err error
				var body []byte
				body, err = io.ReadAll(c.Request.Body)
				if err != nil {
					gors.ErrorRender(ctx, err, options.ErrorHandler)
					return
				}
				req = body
				resp, err = srv.PatchBytesReader(ctx, req)
				if err != nil {
					gors.ErrorRender(ctx, err, options.ErrorHandler)
					return
				}
				gors.ResponseRender(ctx, gors.StatusCode(ctx), resp, "video/mpeg4", gors.ReaderRender, options.ResponseWrapper)
			},
		),
	}
}
