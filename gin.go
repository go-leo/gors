package gors

import "github.com/gin-gonic/gin"

func UriParams(c *gin.Context) map[string][]string {
	m := make(map[string][]string)
	for _, v := range c.Params {
		m[v.Key] = []string{v.Value}
	}
	return m
}
