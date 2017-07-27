package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Flipper struct {
	router *mux.Router
}

func (f *Flipper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f.router.ServeHTTP(w, r)
	if r.URL.Path == "/flip" {
		f.router = NewRouter(handlers2)
	} else {
		f.router = NewRouter(handlers1)
	}

	f.router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		fmt.Println(route.GetPathTemplate())
		return nil
	})
}

var handlers2 = map[string]http.HandlerFunc{
	"flop": func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte("<pre>go to /flip</pre>"))
	},
}

var handlers1 = map[string]http.HandlerFunc{
	"flip": func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte("<pre>go to /flop</pre>"))
	},
}

func main() {
	flipper := &Flipper{NewRouter(handlers2)}
	// router = NewRouter(handlers2)
	http.ListenAndServe(":3750", flipper)
}

func NewRouter(routes map[string]http.HandlerFunc) *mux.Router {
	r := mux.NewRouter()

	for route, handler := range routes {
		route = strings.Trim(route, "/")
		r.PathPrefix("/" + route).HandlerFunc(handler)
	}

	return r
}
