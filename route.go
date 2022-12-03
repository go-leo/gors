package gors

import "github.com/gin-gonic/gin"

type Route interface {
	Method() string
	Path() string
	Handler() gin.HandlerFunc
}

func NewRoute(method string, path string, handler gin.HandlerFunc) Route {
	return &route{method: method, path: path, handler: handler}
}

type route struct {
	method  string
	path    string
	handler gin.HandlerFunc
}

func (r *route) Method() string {
	return r.method
}

func (r *route) Path() string {
	return r.path
}

func (r *route) Handler() gin.HandlerFunc {
	return r.handler
}
