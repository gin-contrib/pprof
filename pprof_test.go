package pprof

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func Test_getPrefix(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want string
	}{
		{"default value", nil, "/debug/pprof"},
		{"test user input value", []string{"test/pprof"}, "test/pprof"},
		{"test user input value", []string{"test/pprof", "pprof"}, "test/pprof"},
	}
	for _, tt := range tests {
		if got := getPrefix(tt.args...); got != tt.want {
			t.Errorf("%q. getPrefix() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestRegisterAndRouteRegister(t *testing.T) {
	bearerToken := "Bearer token"
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	Register(r)
	adminGroup := r.Group("/admin", func(c *gin.Context) {
		if c.Request.Header.Get("Authorization") != bearerToken {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		c.Next()
	})
	RouteRegister(adminGroup, "pprof")

	req, _ := http.NewRequest(http.MethodGet, "/debug/pprof/", nil)
	rw := httptest.NewRecorder()
	r.ServeHTTP(rw, req)

	if expected, got := http.StatusOK, rw.Code; expected != got {
		t.Errorf("expected: %d, got: %d", expected, got)
	}

	req, _ = http.NewRequest(http.MethodGet, "/admin/pprof/", nil)
	rw = httptest.NewRecorder()
	r.ServeHTTP(rw, req)

	if expected, got := http.StatusForbidden, rw.Code; expected != got {
		t.Errorf("expected: %d, got: %d", expected, got)
	}

	req, _ = http.NewRequest(http.MethodGet, "/admin/pprof/", nil)
	req.Header.Set("Authorization", bearerToken)
	rw = httptest.NewRecorder()
	r.ServeHTTP(rw, req)

	if expected, got := http.StatusOK, rw.Code; expected != got {
		t.Errorf("expected: %d, got: %d", expected, got)
	}
}
