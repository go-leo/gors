package gors

import "github.com/gin-gonic/gin"

type IRoute interface {
	Method() string
	AbsolutePath() string
	Handler() gin.HandlerFunc
}

type Route struct {
	HTTPMethod  string
	Path        string
	HandlerFunc gin.HandlerFunc
}

func (r *Route) Method() string {
	return r.HTTPMethod
}

func (r *Route) AbsolutePath() string {
	return r.Path
}

func (r *Route) Handler() gin.HandlerFunc {
	return r.Handler()
}
