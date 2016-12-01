package main

import "github.com/gin-gonic/gin"
import "github.com/gin-contrib/pprof"

func main() {
	router := gin.Default()
	pprof.Register(router, &pprof.Options{
		// default is "debug/pprof"
		RoutePrefix: "debug/pprof",
	})
	router.Run(":8080")
}
