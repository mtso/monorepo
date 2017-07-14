package server

import "testing"

func TestTrim(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"hello", "hello"},
		{"what\n", "what"},
	}

	for _, tt := range cases {
		got := trim(tt.in)
		if got != tt.want {
			t.Errorf("Expected trim(%q) == %q but got %q", tt.in, tt.want, got)
		}
	}
}
