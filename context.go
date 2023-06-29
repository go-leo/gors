package gors

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type key struct{}

// NewContext 向context注入gin.Context
func NewContext(c *gin.Context, rpcMethodName string) context.Context {
	ctx := c.Request.Context()
	ctx = context.WithValue(ctx, rpcMethodKey{}, rpcMethodName)
	ctx = context.WithValue(ctx, key{}, c)
	c.Request = c.Request.WithContext(ctx)
	return c.Request.Context()
}

// FromContext 从context获取gin.Contenxt
func FromContext(ctx context.Context) *gin.Context {
	v, _ := ctx.Value(key{}).(*gin.Context)
	return v
}

type rpcMethodKey struct{}

// RPCMethod returns the method string for the server context. The returned
// string is in the format of "/package.Service/method".
func RPCMethod(ctx context.Context) (string, bool) {
	method, ok := ctx.Value(rpcMethodKey{}).(string)
	return method, ok
}

// SetStatusCode 向context设置status code
func SetStatusCode(ctx context.Context, code int) {
	FromContext(ctx).Set("HTTP_STATUS_CODE", code)
}

// StatusCode 从context获取status code
func StatusCode(ctx context.Context) int {
	code, exists := FromContext(ctx).Get("HTTP_STATUS_CODE")
	if !exists {
		return http.StatusOK
	}
	return code.(int)
}

// SetHeader 向context设置header
func SetHeader(ctx context.Context, header http.Header) {
	FromContext(ctx).Set("HTTP_HEADER", header)
}

// Header 从context获取header
func Header(ctx context.Context) http.Header {
	c := FromContext(ctx)
	header, exists := c.Get("HTTP_HEADER")
	if !exists {
		header = http.Header{}
		c.Set("HTTP_HEADER", header)
	}
	return header.(http.Header)
}

// SetTrailer 向context设置trailer
func SetTrailer(ctx context.Context, trailer http.Header) {
	FromContext(ctx).Set("HTTP_TRAILER", trailer)
}

// Trailer 从context获取trailer
func Trailer(ctx context.Context) http.Header {
	c := FromContext(ctx)
	trailer, exists := c.Get("HTTP_TRAILER")
	if !exists {
		trailer = http.Header{}
		c.Set("HTTP_TRAILER", trailer)
	}
	return trailer.(http.Header)
}
