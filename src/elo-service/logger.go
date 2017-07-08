package main

import (
	"log"
	"net/http"
)

type middleware struct {
	handler http.HandlerFunc
}

func (m middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.handler(w, r)
}

type logwriter struct {
	http.ResponseWriter
	status int
}

func (lw *logwriter) WriteHeader(code int) {
	lw.status = code
	lw.ResponseWriter.WriteHeader(code)
}

func LogHandler(handler http.Handler) http.Handler {
	return &middleware{
		handler: func(w http.ResponseWriter, r *http.Request) {
			lw := &logwriter{w, 200}

			handler.ServeHTTP(lw, r)

			log.Printf("%d %s %s", lw.status, r.Method, r.URL.Path)
		},
	}
}
