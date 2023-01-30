package gors

import "github.com/gin-gonic/gin"

func AppendRoutes[R gin.IRoutes](iRoutes R, routes ...Route) R {
	for _, route := range routes {
		_ = iRoutes.Handle(route.Method(), route.Path(), route.Handler())
	}
	return iRoutes
}

func AppendRichRoutes[R gin.IRoutes](iRoutes R, routes ...RichRoute) R {
	for _, route := range routes {
		_ = iRoutes.Handle(route.Method(), route.Path(), route.Handlers()...)
	}
	return iRoutes
}
