package gors

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
)

// Route 基础路由，定义一个接口路由，包括method、path和处理器
type Route interface {
	Method() string
	Path() string
	Handler() gin.HandlerFunc
}

type route struct {
	method  string
	path    string
	handler gin.HandlerFunc
}

func (r *route) Method() string { return r.method }

func (r *route) Path() string { return r.path }

func (r *route) Handler() gin.HandlerFunc { return r.handler }

// RichRoute 富路由，在 Route 基础上，为路由增加了中间件功能。
type RichRoute interface {
	Method() string
	Path() string
	Handlers() []gin.HandlerFunc
}

type richRoute struct {
	route       Route
	middlewares []gin.HandlerFunc
}

func (r *richRoute) Method() string { return r.route.Method() }

func (r *richRoute) Path() string { return r.route.Path() }

func (r *richRoute) Handlers() []gin.HandlerFunc {
	return append(slices.Clone(r.middlewares), r.route.Handler())
}

// NewRoute 创建一个路由
func NewRoute(method string, path string, handler gin.HandlerFunc) Route {
	return &route{method: method, path: path, handler: handler}
}

// NewRichRoute 创建一个富路由
func NewRichRoute(route Route, middlewares ...gin.HandlerFunc) RichRoute {
	return &richRoute{route: route, middlewares: middlewares}
}

// NewRichRoutes 创建一个多个富路由
func NewRichRoutes(routes []Route, middlewares ...gin.HandlerFunc) []RichRoute {
	var rs []RichRoute
	for _, route := range routes {
		rs = append(rs, NewRichRoute(route, middlewares...))
	}
	return rs
}
