package main

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	pprof.Register(router)
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
