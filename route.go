package gors

import "github.com/gin-gonic/gin"

type Router struct {
	HTTPMethod  string
	Path        string
	HandlerFunc gin.HandlerFunc
}
