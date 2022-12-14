// Code generated by "gors -service StringReader"; DO NOT EDIT.

package demo

import (
	gin "github.com/gin-gonic/gin"
	render "github.com/gin-gonic/gin/render"
	gors "github.com/go-leo/gors"
	io "io"
	http "net/http"
)

func StringReaderRoutes(srv StringReader) []gors.Route {
	return []gors.Route{
		gors.NewRoute(
			http.MethodGet,
			"/api/StringReader/Get",
			func(c *gin.Context) {
				var req string
				var resp io.Reader
				var err error
				var body []byte
				body, err = io.ReadAll(c.Request.Body)
				if err != nil {
					c.String(http.StatusBadRequest, err.Error())
					_ = c.Error(err).SetType(gin.ErrorTypeBind)
					return
				}
				req = string(body)
				ctx := gors.NewContext(c)
				resp, err = srv.GetStringRender(ctx, req)
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
				c.Render(statusCode, render.Reader{ContentType: "video/mpeg4", ContentLength: -1, Reader: resp})
			},
		),
		gors.NewRoute(
			http.MethodOptions,
			"/api/StringReader/Options",
			func(c *gin.Context) {
				var req string
				var resp io.Reader
				var err error
				var body []byte
				body, err = io.ReadAll(c.Request.Body)
				if err != nil {
					c.String(http.StatusBadRequest, err.Error())
					_ = c.Error(err).SetType(gin.ErrorTypeBind)
					return
				}
				req = string(body)
				ctx := gors.NewContext(c)
				resp, err = srv.OptionsStringReader(ctx, req)
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
				c.Render(statusCode, render.Reader{ContentType: "video/mpeg4", ContentLength: -1, Reader: resp})
			},
		),
	}
}
