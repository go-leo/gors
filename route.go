package gors

import "github.com/gin-gonic/gin"

type Router struct {
	HTTPMethods  []string
	Path         string
	HandlerFuncs []gin.HandlerFunc
}
