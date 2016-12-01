package main

import "gopkg.in/gin-gonic/gin.v1"
import "gopkg.in/gin-contrib/pprof.v1"

func main() {
	router := gin.Default()
	pprof.Register(router, &pprof.Options{
		// default is "debug/pprof"
		RoutePrefix: "debug/pprof",
	})
	router.Run(":8080")
}
