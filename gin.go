package gors

import "github.com/gin-gonic/gin"

func UriParams(c *gin.Context) map[string][]string {
	m := make(map[string][]string)
	for _, v := range c.Params {
		m[v.Key] = []string{v.Value}
	}
	return m
}

func QueryParams(c *gin.Context) map[string][]string {
	return c.Request.URL.Query()
}

func HeaderParams(c *gin.Context) map[string][]string {
	return c.Request.Header
}
