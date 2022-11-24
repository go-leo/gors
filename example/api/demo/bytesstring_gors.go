// Code generated by "gors -service BytesString"; DO NOT EDIT.
package demo

import (
	context "context"
	gin "github.com/gin-gonic/gin"
	render "github.com/gin-gonic/gin/render"
	gors "github.com/go-leo/gors"
	io "io"
	http "net/http"
)

func BytesStringRouters(srv BytesString) []gors.Router {
	return []gors.Router{
		{
			HTTPMethod: http.MethodGet,
			Path:       "/api/BytesString/Get",
			HandlerFunc: func(c *gin.Context) {
				body, err := io.ReadAll(c.Request.Body)
				if err != nil {
					c.String(http.StatusBadRequest, err.Error())
					_ = c.Error(err).SetType(gin.ErrorTypeBind)
					return
				}
				req := body
				var ctx context.Context = c
				ctx = gors.InjectStatusCode(ctx, http.StatusOK)
				ctx = gors.InjectHeader(ctx, c.Writer.Header())
				resp, err := srv.GetBytesString(ctx, req)
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
				statusCode := gors.ExtractStatusCode(ctx)
				c.Render(statusCode, render.Data{ContentType: "text/html; charset=utf-8", Data: []byte(resp)})
			},
		},
		{
			HTTPMethod: http.MethodPut,
			Path:       "/api/BytesString/Put",
			HandlerFunc: func(c *gin.Context) {
				body, err := io.ReadAll(c.Request.Body)
				if err != nil {
					c.String(http.StatusBadRequest, err.Error())
					_ = c.Error(err).SetType(gin.ErrorTypeBind)
					return
				}
				req := body
				var ctx context.Context = c
				ctx = gors.InjectStatusCode(ctx, http.StatusOK)
				ctx = gors.InjectHeader(ctx, c.Writer.Header())
				resp, err := srv.PutBytesString(ctx, req)
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
				statusCode := gors.ExtractStatusCode(ctx)
				c.Redirect(statusCode, resp)
			},
		},
	}
}