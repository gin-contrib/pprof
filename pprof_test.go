package ginpprof

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

func Test_isLocalhost(t *testing.T) {
	validTestCases := []string{
		"localhost:12345",
		"localhost",
		"127.0.0.1:12345",
		"127.0.0.1",
		"[::1]:12345",
		"[::1]",
	}

	for _, tc := range validTestCases {
		if !isLocalhost(tc) {
			t.Errorf("should be valid localhost, but is not: %s", tc)
		}
	}

	invalidTestCases := []string{
		"1.2.3.4",
		"1.2.3.4:12345",
		"fake-data",
	}

	for _, tc := range invalidTestCases {
		if isLocalhost(tc) {
			t.Errorf("should not be localhost, but is: %s", tc)
		}
	}
}
