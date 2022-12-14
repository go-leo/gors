// Code generated by "gors -service StringString"; DO NOT EDIT.

package demo

import (
	gin "github.com/gin-gonic/gin"
	render "github.com/gin-gonic/gin/render"
	gors "github.com/go-leo/gors"
	io "io"
	http "net/http"
)

func StringStringRoutes(srv StringString) []gors.Route {
	return []gors.Route{
		gors.NewRoute(
			http.MethodGet,
			"/api/StringString/Get",
			func(c *gin.Context) {
				var req string
				var resp string
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
				resp, err = srv.GetStringString(ctx, req)
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
		gors.NewRoute(
			http.MethodPatch,
			"/api/StringString/Patch",
			func(c *gin.Context) {
				var req string
				var resp string
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
				resp, err = srv.PatchStringString(ctx, req)
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
