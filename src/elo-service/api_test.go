package main

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func NewAssert(t *testing.T) func(interface{}, interface{}, ...string) {
	return func(a, b interface{}, msg ...string) {
		if !reflect.DeepEqual(a, b) {
			t.Errorf("%s %v != %v", msg, a, b)
		}
	}
}

func Test_NewLeague(t *testing.T) {
	assert := NewAssert(t)

	ts := httptest.NewServer(NewRouter())
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	assert(err, nil, "Should make GET request to root")
	assert(resp.StatusCode, 200, "Should reach server root")
}
