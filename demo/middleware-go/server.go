package main

import (
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("hello~"))
}

type logWriter struct {
	// Must embed http.ResponseWriter (to inherit
	// the other methods: Header() and Write()).
	http.ResponseWriter
	statusCode int
}

func (rw *logWriter) WriteHeader(code int) {
	rw.statusCode = code
	// Use original interface name to access embedded type.
	rw.ResponseWriter.WriteHeader(code)
}

func logStatus(handler http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Initialize logWriter.
		lw := &logWriter{w, http.StatusOK}

		// Pass logWriter into handler.
		handler(lw, r)

		log.Println(lw.statusCode, r.URL.Path)
	})
}

func main() {

	http.HandleFunc("/", logStatus(handler))

	http.ListenAndServe(":3750", nil)
}
