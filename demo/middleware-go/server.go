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

		log.Println("log before")
		// Pass logWriter into handler.
		handler(lw, r)

		log.Println("log after")
		log.Println(lw.statusCode, r.URL.Path)
	})
}

// Attaches a header "foo: bar" to the outgoing response.
func attachHeader() func(http.HandlerFunc) http.HandlerFunc {
	return func(handler http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("foo", "bar")
			log.Println("attach before")
			handler(w, r)
			log.Println("attach after")
		})
	}
}

type middleware func(http.HandlerFunc) http.HandlerFunc

func applyMiddlewares(middlewares ...middleware) middleware {
	if len(middlewares) < 0 {
		return func(h http.HandlerFunc) http.HandlerFunc {
			return h
		}
	}

	return func(h http.HandlerFunc) http.HandlerFunc {
		count := len(middlewares)
		for i := range middlewares {
			// Apply in reverse order (first middleware becomes outer-most layer).
			h = middlewares[count-1-i](h)
		}
		return h
	}
}

type MiddlewareHandlerFunc struct {
	handler http.HandlerFunc
}

func (mh *MiddlewareHandlerFunc) apply(m middleware) *MiddlewareHandlerFunc {
	mh.handler = m(mh.handler)
	return mh
}

func main() {
	h := &MiddlewareHandlerFunc{handler}
	h.apply(logStatus).apply(attachHeader())
	// h := applyMiddlewares(logStatus, attachHeader())(handler)
	http.HandleFunc("/", h.handler)
	http.ListenAndServe(":3750", nil)
}
