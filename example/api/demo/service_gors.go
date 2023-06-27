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
				if err = gors.ShouldBind(
					c, req, "",
					gors.UriBinding,
				); err != nil {
					gors.HandleBadRequest(c, err)
					return
				}
				ctx := gors.NewContext(c)
				resp, err = srv.Method(ctx, req)
				gors.MustRender(c, resp, err, "", gors.JSONRender)
			},
		),
	}
}
