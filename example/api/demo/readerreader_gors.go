// Code generated by "gors -service ReaderReader"; DO NOT EDIT.

package demo

import (
	gin "github.com/gin-gonic/gin"
	render "github.com/gin-gonic/gin/render"
	gors "github.com/go-leo/gors"
	io "io"
	http "net/http"
)

func ReaderReaderRoutes(srv ReaderReader) []gors.Route {
	return []gors.Route{
		gors.NewRoute(
			http.MethodGet,
			"/api/ReaderReader/Get",
			func(c *gin.Context) {
				var req io.Reader
				var resp io.Reader
				var err error
				req = c.Request.Body
				ctx := gors.NewContext(c)
				resp, err = srv.GetReaderReader(ctx, req)
				if err != nil {
					if httpErr, ok := err.(*gors.HttpError); ok {
						c.String(httpErr.StatusCode(), httpErr.Error())
						_ = c.Error(err).SetType(gin.ErrorTypePublic)
						return
					}
					c.String(http.StatusInternalServerError, err.Error())
					_ = c.Error(err).SetType(gin.ErrorTypePrivate)
					return
				}
				statusCode := gors.GetCodeFromContext(ctx)
				c.Render(statusCode, render.Reader{ContentType: "", ContentLength: -1, Reader: resp})
			},
		),
		gors.NewRoute(
			http.MethodHead,
			"/api/ReaderReader/head",
			func(c *gin.Context) {
				var req io.Reader
				var resp io.Reader
				var err error
				req = c.Request.Body
				ctx := gors.NewContext(c)
				resp, err = srv.HeadReaderReader(ctx, req)
				if err != nil {
					if httpErr, ok := err.(*gors.HttpError); ok {
						c.String(httpErr.StatusCode(), httpErr.Error())
						_ = c.Error(err).SetType(gin.ErrorTypePublic)
						return
					}
					c.String(http.StatusInternalServerError, err.Error())
					_ = c.Error(err).SetType(gin.ErrorTypePrivate)
					return
				}
				statusCode := gors.GetCodeFromContext(ctx)
				c.Render(statusCode, render.Reader{ContentType: "", ContentLength: -1, Reader: resp})
			},
		),
	}
}
