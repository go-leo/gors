// Code generated by "gors -service ReaderBytes"; DO NOT EDIT.

package demo

import (
	gin "github.com/gin-gonic/gin"
	render "github.com/gin-gonic/gin/render"
	gors "github.com/go-leo/gors"
	io "io"
	http "net/http"
)

func ReaderBytesRoutes(srv ReaderBytes) []gors.Route {
	return []gors.Route{
		gors.NewRoute(
			http.MethodGet,
			"/api/ReaderBytes/Get",
			func(c *gin.Context) {
				var req io.Reader
				var resp []byte
				var err error
				req = c.Request.Body
				ctx := gors.NewContext(c)
				resp, err = srv.GetReaderBytes(ctx, req)
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
			http.MethodPost,
			"/api/ReaderBytes/Post",
			func(c *gin.Context) {
				var req io.Reader
				var resp []byte
				var err error
				req = c.Request.Body
				ctx := gors.NewContext(c)
				resp, err = srv.PostReaderBytes(ctx, req)
				switch e := err.(type) {
				case nil:
					statusCode := gors.HttpStatusCode(c, resp)
					c.Render(statusCode, render.Data{ContentType: "text/go", Data: resp})
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
