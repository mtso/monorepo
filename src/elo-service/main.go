package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello~"))
	})
	http.ListenAndServe(":3750", nil)
}

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	return r
}
