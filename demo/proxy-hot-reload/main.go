package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Reloader struct {
	router *mux.Router
}

type Routes map[string]string

var conf = flag.String("conf", "./proxy.conf", "Proxy configuration location.")

func (r *Reloader) Reload() error {
	raw, err := ioutil.ReadFile(*conf)
	if err != nil {
		return err
	}
	var routes Routes
	err = json.Unmarshal(raw, &routes)
	if err != nil {
		return err
	}
	r.router = NewRouter(routes)
	return nil
}

func (rl *Reloader) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/reload" {
		if err := rl.Reload(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			w.Write([]byte("Reloaded successfully"))
		}
		return
	}
	rl.router.ServeHTTP(w, r)
}

// func (f *Flipper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	f.router.ServeHTTP(w, r)
// 	if r.URL.Path == "/flip" {
// 		f.router = NewRouter(handlers2)
// 	} else {
// 		f.router = NewRouter(handlers1)
// 	}
//
// 	f.router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
// 		fmt.Println(route.GetPathTemplate())
// 		return nil
// 	})
// }

// var handlers2 = map[string]http.HandlerFunc{
// 	"flop": func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Content-Type", "text/html")
// 		w.Write([]byte("<pre>go to /flip</pre>"))
// 	},
// }
//
// var handlers1 = map[string]http.HandlerFunc{
// 	"flip": func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Content-Type", "text/html")
// 		w.Write([]byte("<pre>go to /flop</pre>"))
// 	},
// }

func main() {
	flipper := &Reloader{NewRouter(map[string]string{
		"foo": "bar",
	})}
	// router = NewRouter(handlers2)
	http.ListenAndServe(":3750", flipper)
}

func NewRouter(routes map[string]string) *mux.Router {
	r := mux.NewRouter()

	for route, destination := range routes {
		route = strings.Trim(route, "/")
		r.PathPrefix("/" + route).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// temp handler for test
			w.Write([]byte("forward to " + destination))
		})
	}

	return r
}
