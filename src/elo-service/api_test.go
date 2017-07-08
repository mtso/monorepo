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

var persistedLeagueid string

func TestMain(m *testing.M) {
	app := newApp()
	defer app.Db.Close()
	ts := httptest.NewServer(app.Handler)
	defer ts.Close()

	buf := bytes.NewBuffer([]byte(`{"title":"Persisted League"}`))
	resp, err := http.Post(ts.URL+"/api/new", "application/json", buf)
	if err != nil {
		panic(err)
	}

	body, err := ParseResponseBody(resp)
	if err != nil {
		panic(err)
	}

	league, _ := body["league"]
	leagueid, _ := league.(map[string]interface{})["id"]
	persistedLeagueid = leagueid.(string)

	buf = bytes.NewBuffer([]byte(`{"winner":"foo","loser":"bar"}`))
	_, err = http.Post(ts.URL+"/api/"+persistedLeagueid, "application/json", buf)
	if err != nil {
		panic(err)
	}

	m.Run()
}

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

func Test_AddGame(t *testing.T) {
	assert := NewAssert(t)
	must := NewMust(t)

	app := newApp()
	defer app.Db.Close()
	ts := httptest.NewServer(app.Handler)
	defer ts.Close()

	buf := bytes.NewBuffer([]byte(`{"title":"New World Cup"}`))
	resp, err := http.Post(ts.URL+"/api/new", "application/json", buf)
	body, err := ParseResponseBody(resp)
	must(err, nil, "Should JSON-encode response")
	league, ok := body["league"]
	must(ok, true, "Should contain league")

	leagueid, ok := league.(map[string]interface{})["id"]

	buf = bytes.NewBuffer([]byte(`{"winner":"foo","loser":"bar"}`))
	resp, err = http.Post(ts.URL+"/api/"+leagueid.(string), "application/json", buf)
	assert(resp.StatusCode, 200, "Should be status OK")

	body, err = ParseResponseBody(resp)
	must(err, nil, "Should JSON-encode response")

	game, ok := body["game"]
	must(ok, true, "Should have game field")

	winner, ok := game.(map[string]interface{})["winner"]
	must(ok, true, "Should have winner field")

	username, ok := winner.(map[string]interface{})["username"]
	assert(username, "foo", "Should be the POSTed name")
}

func Test_GetGames(t *testing.T) {
	assert := NewAssert(t)
	must := NewMust(t)

	app := newApp()
	defer app.Db.Close()
	ts := httptest.NewServer(app.Handler)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/api/" + persistedLeagueid + "/games")
	must(err, nil, "Should be successful GET request")

	body, err := ParseResponseBody(resp)
	must(err, nil, "Should decode JSON response")

	games, ok := body["games"]
	must(ok, true, "Should contain games field")

	assert(reflect.ValueOf(games), reflect.Slice, "Should be an array")
}

func Test_GetPlayers(t *testing.T) {
	assert := NewAssert(t)
	must := NewMust(t)

	app := newApp()
	defer app.Db.Close()
	ts := httptest.NewServer(app.Handler)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/api/" + persistedLeagueid + "/players")
	must(err, nil, "Should be successful GET request")

	body, err := ParseResponseBody(resp)
	must(err, nil, "Should decode JSON response")

	players, ok := body["players"]
	must(ok, true, "Should contain players field")

	assert(reflect.ValueOf(players), reflect.Slice, "Should be an array")
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
