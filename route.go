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

func HandleRoutes(iRoutes gin.IRoutes, routes []Route) gin.IRoutes {
	for _, route := range routes {
		iRoutes = iRoutes.Handle(route.Method(), route.Path(), route.Handler())
	}
	return iRoutes
}

type RichRoute interface {
	Method() string
	Path() string
	Handlers() []gin.HandlerFunc
}

func NewRichRoute(route Route, middlewares ...gin.HandlerFunc) RichRoute {
	return &richRoute{route: route, middlewares: middlewares}
}

func HandleRichRoutes(iRoutes gin.IRoutes, routes []RichRoute) gin.IRoutes {
	for _, route := range routes {
		iRoutes = iRoutes.Handle(route.Method(), route.Path(), route.Handlers()...)
	}
	return iRoutes
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

type richRoute struct {
	route       Route
	middlewares []gin.HandlerFunc
}

func (r *richRoute) Method() string {
	return r.route.Method()
}

func (r *richRoute) Path() string {
	return r.route.Path()
}

func (r *richRoute) Handlers() []gin.HandlerFunc {
	return append(r.middlewares, r.route.Handler())
}
