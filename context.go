package gors

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type key struct{}

func NewContext(c *gin.Context) context.Context {
	ctx := c.Request.Context()
	ctx = context.WithValue(ctx, key{}, c)
	c.Request = c.Request.WithContext(ctx)
	return c.Request.Context()
}

func FromContext(ctx context.Context) *gin.Context {
	v, _ := ctx.Value(key{}).(*gin.Context)
	return v
}

func GetCodeFromContext(ctx context.Context) int {
	c, _ := ctx.Value(key{}).(*gin.Context)
	code, exists := c.Get("HTTP_STATUS_CODE")
	if !exists {
		return http.StatusOK
	}
	return code.(int)
}

func SetCodeToContext(ctx context.Context, code int) {
	c, _ := ctx.Value(key{}).(*gin.Context)
	c.Set("HTTP_STATUS_CODE", code)
}
