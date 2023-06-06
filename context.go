package gors

import (
	"context"
	"github.com/gin-gonic/gin"
)

type key struct{}

// NewContext 向context注入gin.Context
func NewContext(c *gin.Context) context.Context {
	ctx := c.Request.Context()
	ctx = context.WithValue(ctx, key{}, c)
	c.Request = c.Request.WithContext(ctx)
	return c.Request.Context()
}

// FromContext 从context获取gin.Contenxt
func FromContext(ctx context.Context) *gin.Context {
	v, _ := ctx.Value(key{}).(*gin.Context)
	return v
}
