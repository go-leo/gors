package gors

import "github.com/gin-gonic/gin"

type Render interface {
	Render(c *gin.Context, statusCode int) error
}
