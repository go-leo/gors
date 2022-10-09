package add

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/go-leo/gors"
)

func NewRouters(srv AddService) []gors.Router {
	return []gors.Router{
		{
			HTTPMethods: []string{http.MethodPost},
			Path:        "/post/sum",
			HandlerFuncs: []gin.HandlerFunc{func(c *gin.Context) {
				var req SumRequest
				if err := c.ShouldBindHeader(&req); err != nil {
					_ = c.Error(err).SetType(gin.ErrorTypeBind)
					c.String(http.StatusBadRequest, err.Error())
					c.Abort()
				}

				resp, err := srv.Sum(c.Request.Context(), &req)
				if err != nil {
					if httpErr, ok := err.(*gors.HttpError); ok {
						c.String(httpErr.StatusCode(), httpErr.Err().Error())
					} else {
						c.String(500, err.Error())
					}
					return
				}
				c.JSON(200, resp)
			}},
		},
	}
}
