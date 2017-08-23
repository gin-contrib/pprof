package pprof

import "testing"

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
