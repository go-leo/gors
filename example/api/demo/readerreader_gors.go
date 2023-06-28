// Code generated by "gors -service ReaderReader"; DO NOT EDIT.

package demo

import (
	gin "github.com/gin-gonic/gin"
	gors "github.com/go-leo/gors"
	io "io"
	http "net/http"
)

func ReaderReaderRoutes(srv ReaderReader, opts ...gors.Option) []gors.Route {
	options := gors.New(opts...)
	_ = options
	return []gors.Route{
		gors.NewRoute(
			http.MethodGet,
			"/api/ReaderReader/Get",
			func(c *gin.Context) {
				var ctx = gors.NewContext(c)
				var req io.Reader
				var resp io.Reader
				var err error
				req = c.Request.Body
				resp, err = srv.GetReaderReader(ctx, req)
				if err != nil {
					gors.ErrorRender(ctx, err, options.ErrorHandler)
					return
				}
				gors.ResponseRender(ctx, gors.StatusCode(ctx), resp, "", gors.ReaderRender, options.ResponseWrapper)
			},
		),
		gors.NewRoute(
			http.MethodHead,
			"/api/ReaderReader/head",
			func(c *gin.Context) {
				var ctx = gors.NewContext(c)
				var req io.Reader
				var resp io.Reader
				var err error
				req = c.Request.Body
				resp, err = srv.HeadReaderReader(ctx, req)
				if err != nil {
					gors.ErrorRender(ctx, err, options.ErrorHandler)
					return
				}
				gors.ResponseRender(ctx, gors.StatusCode(ctx), resp, "", gors.ReaderRender, options.ResponseWrapper)
			},
		),
	}
}
