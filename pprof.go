package ginpprof

import (
	"net/http"
	"net/http/pprof"
	"strings"

	"github.com/gin-gonic/gin"
)

// Options provides potential route registration configuration options
type Options struct {
	// RoutePrefix is an optional path prefix. If left unspecified, `/debug/pprof`
	// is used as the default path prefix.
	RoutePrefix string

	// LocalhostOnly is an optional option. If set to true, pprof routes will only
	// be accessible from localhost.
	LocalhostOnly bool
}

// Register the standard HandlerFuncs from the net/http/pprof package with
// the provided gin.Engine. opts is a optional. If a `nil` value is passed,
// the default path prefix is used.
func Register(r *gin.Engine, opts *Options) {
	prefix := routePrefix(opts)
	lhonly := localhostOnly(localhostOnlyOpt(opts))

	r.GET(prefix+"/block", lhonly, pprofHandler(pprof.Index))
	r.GET(prefix+"/heap", lhonly, pprofHandler(pprof.Index))
	r.GET(prefix+"/profile", lhonly, pprofHandler(pprof.Profile))
	r.POST(prefix+"/symbol", lhonly, pprofHandler(pprof.Symbol))
	r.GET(prefix+"/symbol", lhonly, pprofHandler(pprof.Symbol))
	r.GET(prefix+"/trace", lhonly, pprofHandler(pprof.Trace))
}

func pprofHandler(h http.HandlerFunc) gin.HandlerFunc {
	handler := http.HandlerFunc(h)
	return func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	}
}

func routePrefix(opts *Options) string {
	if opts == nil {
		return "/debug/pprof"
	}
	return opts.RoutePrefix
}

func localhostOnlyOpt(opts *Options) bool {
	if opts == nil {
		return false
	}
	return opts.LocalhostOnly
}

// LocalhostOnly only lets in connections from localhost
func localhostOnly(enable bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		if enable && !isLocalhost(c.ClientIP()) {
			c.String(http.StatusForbidden, "Forbidden")
			c.Abort()
			return
		}

		c.Next()
	}
}

// isLocalhost checks that a given addr is lcoalhost
func isLocalhost(addr string) bool {
	// either it starts with localhost
	if strings.HasPrefix(addr, "localhost") {
		return true
	}

	// or it starts with 127.0.0.1
	if strings.HasPrefix(addr, "127.0.0.1") {
		return true
	}

	// or it starts with [::1]
	if strings.HasPrefix(addr, "[::1]") {
		return true
	}

	return false
}
