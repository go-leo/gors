// Code generated by "gors -service ReaderString"; DO NOT EDIT.

package demo

import (
	gin "github.com/gin-gonic/gin"
	gors "github.com/go-leo/gors"
	io "io"
	http "net/http"
)

func ReaderStringRoutes(srv ReaderString) []gors.Route {
	return []gors.Route{
		gors.NewRoute(
			http.MethodGet,
			"/api/ReaderString/Get",
			func(c *gin.Context) {
				var req io.Reader
				var resp string
				var err error
				req = c.Request.Body
				ctx := gors.NewContext(c)
				resp, err = srv.GetReaderString(ctx, req)
				gors.MustRender(c, resp, err, "text/plain; charset=utf-8", gors.TextRender)
			},
		),
		gors.NewRoute(
			http.MethodPost,
			"/api/ReaderString/Post",
			func(c *gin.Context) {
				var req io.Reader
				var resp string
				var err error
				req = c.Request.Body
				ctx := gors.NewContext(c)
				resp, err = srv.PostReaderString(ctx, req)
				gors.MustRender(c, resp, err, "text/go", gors.StringRender)
			},
		),
	}
}
