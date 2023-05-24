package gors

import "github.com/gin-gonic/gin"

type Binding interface {
	Bind(c *gin.Context) error
}
