# pprof

[![Build Status](https://travis-ci.org/gin-contrib/pprof.svg)](https://travis-ci.org/gin-contrib/pprof)
[![codecov](https://codecov.io/gh/gin-contrib/pprof/branch/master/graph/badge.svg)](https://codecov.io/gh/gin-contrib/pprof)
[![Go Report Card](https://goreportcard.com/badge/github.com/gin-contrib/pprof)](https://goreportcard.com/report/github.com/gin-contrib/pprof)
[![GoDoc](https://godoc.org/github.com/gin-contrib/pprof?status.svg)](https://godoc.org/github.com/gin-contrib/pprof)
[![Join the chat at https://gitter.im/gin-gonic/gin](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/gin-gonic/gin)

gin pprof middleware

> Package pprof serves via its HTTP server runtime profiling data in the format expected by the pprof visualization tool.

## Usage

### Start using it

Download and install it:

```bash
$ go get gopkg.in/gin-contrib/pprof.v1
```

Import it in your code:

```go
import "gopkg.in/gin-contrib/pprof.v1"
```

### Example:

```go
package main

import "gopkg.in/gin-gonic/gin.v1"
import "gopkg.in/gin-contrib/pprof.v1"

func main() {
  router := gin.Default()
  pprof.Register(router, nil)
  router.Run(":8080")
}
```

### change default path prefix:

```go
func main() {
	router := gin.Default()
	pprof.Register(router, &pprof.Options{
		// default is "debug/pprof"
		RoutePrefix: "debug/pprof",
	})
	router.Run(":8080")
}
```

### Use the pprof tool

Then use the pprof tool to look at the heap profile:

```bash
go tool pprof http://localhost:8080/debug/pprof/heap
```

Or to look at a 30-second CPU profile:

```bash
go tool pprof http://localhost:8080/debug/pprof/profile
```

Or to look at the goroutine blocking profile, after calling runtime.SetBlockProfileRate in your program:

```bash
go tool pprof http://localhost:8080/debug/pprof/block
```

Or to collect a 5-second execution trace:

```bash
wget http://localhost:8080/debug/pprof/trace?seconds=5
```
