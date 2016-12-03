package pprof

import "testing"

func Test_routePrefix(t *testing.T) {
	type args struct {
		opts *Options
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"default value", args{nil}, "/debug/pprof"},
		{"test user input value", args{&Options{
			RoutePrefix: "test/pprof",
		}}, "test/pprof"},
	}
	for _, tt := range tests {
		if got := routePrefix(tt.args.opts); got != tt.want {
			t.Errorf("%q. routePrefix() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
