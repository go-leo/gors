// Code generated by "gors -service ReaderString"; DO NOT EDIT.
package demo

import (
	gin "github.com/gin-gonic/gin"
	render "github.com/gin-gonic/gin/render"
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
				c.Render(statusCode, render.Data{ContentType: "text/plain; charset=utf-8", Data: []byte(resp)})
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
				c.Render(statusCode, render.Data{ContentType: "text/go", Data: []byte(resp)})
			},
		),
	}
}
