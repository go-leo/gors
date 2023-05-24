// Code generated by "gors -service Service"; DO NOT EDIT.

package demo

import (
	gin "github.com/gin-gonic/gin"
	gors "github.com/go-leo/gors"
	http "net/http"
)

func ServiceRoutes(srv Service) []gors.Route {
	return []gors.Route{
		gors.NewRoute(
			http.MethodGet,
			"/api/v1/method/:id",
			func(c *gin.Context) {
				var req *MethodReq
				var resp *MethodResp
				var err error
				req = new(MethodReq)
				if err = c.ShouldBindUri(req); err != nil {
					c.String(http.StatusBadRequest, err.Error())
					_ = c.Error(err).SetType(gin.ErrorTypeBind)
					return
				}
				if err = gors.Validate(req); err != nil {
					c.String(http.StatusBadRequest, err.Error())
					_ = c.Error(err).SetType(gin.ErrorTypeBind)
					return
				}
				ctx := gors.NewContext(c)
				resp, err = srv.Method(ctx, req)
				if gors.IsInterrupted(ctx) {
					return
				}
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
				c.JSON(statusCode, resp)
			},
		),
	}
}
