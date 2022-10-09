package main

import (
	"github.com/gin-gonic/gin"

	"github.com/go-leo/gors/contrib/gin/example/api/add"
)

func main() {
	engine := gin.New()
	engine.GET("/:a/:b", func(c *gin.Context) {
		var a add.SumRequest
		c.ShouldBindUri(&a)
		c.BindUri(&a)
	})
}
