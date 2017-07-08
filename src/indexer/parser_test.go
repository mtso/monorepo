package main

import (
	"testing"
)

func TestGetSummary(t *testing.T) {
	cases := []struct {
		in, want string
		err      error
	}{
		{"# indexer\n\n[summary]::\ntest summary\n", "test summary", nil},
		{"\n[summary]::\ntest summary\n", "test summary", nil},
		{"#\n\ntest summary\n", "", ErrNoSummary},
	}

	for _, c := range cases {
		parsed, err := GetSummary([]byte(c.in))
		if err != c.err {
			t.Errorf("expected %q to return error value %q but got %q", c.in, c.err, err)
			continue
		}

		got := string(parsed)
		if got != c.want {
			t.Errorf("expected %q to match %q in %q", got, c.want, c.in)
		}
	}
}
