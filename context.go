package gors

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type key struct{}

// NewContext 向context注入gin.Contenxt
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

// SetCodeToContext 向context设置status code
func SetCodeToContext(ctx context.Context, code int) {
	c, _ := ctx.Value(key{}).(*gin.Context)
	c.Set("HTTP_STATUS_CODE", code)
}

// GetCodeFromContext 从context获取status code
func GetCodeFromContext(ctx context.Context) int {
	c, _ := ctx.Value(key{}).(*gin.Context)
	code, exists := c.Get("HTTP_STATUS_CODE")
	if !exists {
		return http.StatusOK
	}
	return code.(int)
}

// InterruptHandle 向context设置中断处理
func InterruptHandle(ctx context.Context) {
	c, _ := ctx.Value(key{}).(*gin.Context)
	c.Set("INTERRUPT_HANDLE", true)
}

// IsInterrupted 从context获取是否中断处理
func IsInterrupted(ctx context.Context) bool {
	c, _ := ctx.Value(key{}).(*gin.Context)
	isInterrupted, exists := c.Get("INTERRUPT_HANDLE")
	if !exists {
		return false
	}
	return isInterrupted.(bool)
}
