package main

import (
	"github.com/gin-gonic/gin"

	"github.com/go-leo/gors"
)

func main() {
	engine := gin.New()
	var routes []*gors.Route

	for _, route := range routes {
		if len(route.Methods) <= 0 {
			engine.Any(route.Path, gin.WrapF(route.Handler))
			continue
		}
		for _, method := range route.Methods {
			engine.Handle(method, route.Path, gin.WrapF(route.Handler))
		}
	}
}
