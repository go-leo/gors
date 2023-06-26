package gors

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
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

// SetHTTPStatusCode 向context设置status code
func SetHTTPStatusCode(ctx context.Context, code int) {
	FromContext(ctx).Set("HTTP_STATUS_CODE", code)
}

// HTTPStatusCode 从context获取status code
func HTTPStatusCode(ctx context.Context) int {
	code, exists := FromContext(ctx).Get("HTTP_STATUS_CODE")
	if !exists {
		return http.StatusOK
	}
	return code.(int)
}
