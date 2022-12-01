package gors

import "github.com/gin-gonic/gin"

type Route struct {
	HTTPMethod  string
	Path        string
	HandlerFunc gin.HandlerFunc
}
