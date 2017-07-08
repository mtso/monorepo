package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func Test_NewLeague(t *testing.T) {
	assert := NewAssert(t)
	must := NewMust(t)

	app := newApp()
	defer app.Db.Close()
	ts := httptest.NewServer(app.Handler)
	defer ts.Close()

	buf := bytes.NewBuffer([]byte(`{"title":"Debauchery Tea Party"}`))

	resp, err := http.Post(ts.URL+"/api/new", "application/json", buf)
	assert(resp.StatusCode, 200, "Should get OK status")

	body, err := ParseResponseBody(resp)
	must(err, nil, "Should JSON-encode response")

	rawleague, ok := body["league"]
	must(ok, true, "Should return league object")

	league := rawleague.(map[string]interface{})

	assert(fmt.Sprintf("%T", league["id"]), "string", "Should contain a string ID")
	assert(league["title"], "Debauchery Tea Party", "Should return league title")
}

// NewAssert makes an assert func(i{}, i{}, ...string).
// Test continues on fail.
func NewAssert(t *testing.T) func(interface{}, interface{}, ...string) {
	return func(a, b interface{}, msg ...string) {
		if !reflect.DeepEqual(a, b) {
			t.Errorf("%s %v != %v", msg, a, b)
		}
	}
}

// NewMust makes a must func(i{}, i{}, ...string).
// Fails test on not equal.
func NewMust(t *testing.T) func(interface{}, interface{}, ...string) {
	return func(a, b interface{}, msg ...string) {
		if !reflect.DeepEqual(a, b) {
			t.Fatalf("%s %v != %v", msg, a, b)
		}
	}
}

// Parses a JSON response body into map[string]interface{}.
func ParseResponseBody(resp *http.Response) (map[string]interface{}, error) {
	decoder := json.NewDecoder(resp.Body)
	var raw interface{}

	err := decoder.Decode(&raw)
	if err != nil {
		return nil, err
	}

	js := raw.(map[string]interface{})
	return js, nil
}
