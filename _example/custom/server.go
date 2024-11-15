package main

import (
	"net/http"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	debugGroup := router.Group("/debug", func(c *gin.Context) {
		if c.Request.Header.Get("Authorization") != "foobar" {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		c.Next()
	})

	pprof.RouteRegister(debugGroup, "pprof")
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
