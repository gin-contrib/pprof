package main

import (
	"net/http"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	pprof.Register(router)
	adminGroup := router.Group("/admin", func(c *gin.Context) {
		if c.Request.Header.Get("Authorization") != "foobar" {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		c.Next()
	})
	pprof.RouteRegister(adminGroup, "pprof")
	router.Run(":8080")
}
