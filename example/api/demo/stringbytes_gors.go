// Code generated by "gors -service StringBytes"; DO NOT EDIT.

package demo

import (
	gin "github.com/gin-gonic/gin"
	render "github.com/gin-gonic/gin/render"
	gors "github.com/go-leo/gors"
	io "io"
	http "net/http"
)

func StringBytesRoutes(srv StringBytes) []gors.Route {
	return []gors.Route{
		gors.NewRoute(
			http.MethodGet,
			"/api/StringBytes/Get",
			func(c *gin.Context) {
				var req string
				var resp []byte
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
				resp, err = srv.GetStringBytes(ctx, req)
				switch e := err.(type) {
				case nil:
					statusCode := gors.HttpStatusCode(c, resp)
					c.Render(statusCode, render.Data{ContentType: "", Data: resp})
					return
				case *gors.HttpError:
					c.String(e.StatusCode(), e.Error())
					_ = c.Error(e).SetType(gin.ErrorTypePublic)
					return
				default:
					c.String(http.StatusInternalServerError, err.Error())
					_ = c.Error(e).SetType(gin.ErrorTypePrivate)
					return
				}
			},
		),
		gors.NewRoute(
			http.MethodOptions,
			"/api/StringBytes/Options",
			func(c *gin.Context) {
				var req string
				var resp []byte
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
				resp, err = srv.OptionsStringBytes(ctx, req)
				switch e := err.(type) {
				case nil:
					statusCode := gors.HttpStatusCode(c, resp)
					c.Render(statusCode, render.Data{ContentType: "", Data: resp})
					return
				case *gors.HttpError:
					c.String(e.StatusCode(), e.Error())
					_ = c.Error(e).SetType(gin.ErrorTypePublic)
					return
				default:
					c.String(http.StatusInternalServerError, err.Error())
					_ = c.Error(e).SetType(gin.ErrorTypePrivate)
					return
				}
			},
		),
	}
}
