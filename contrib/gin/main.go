package main

import (
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.New()
	engine.GET("/:name/:id", func(c *gin.Context) {
		person := make(map[string]interface{})
		for _, key := range []string{"name", "id"} {
			person[key] = c.Param(key)
		}
		c.JSON(http.StatusOK, gin.H{"name": person["name"], "uuid": person["id"]})
	})
	srv := http.Server{
		Handler: engine,
	}
	listen, err := net.Listen("tcp", ":8088")
	if err != nil {
		panic(err)
	}
	err = srv.Serve(listen)
	if err != nil {
		panic(err)
	}
}
