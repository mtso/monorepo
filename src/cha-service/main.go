package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/mtso/cha"
	"go.uber.org/zap"
)

var port = flag.String("port", "3750", "Port to listen on.")

func main() {
	handleFunc := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		path := r.URL.Path
		hash := cha.Hash([]byte(path))
		w.Write(hash)

		logger, err := zap.NewProduction()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		defer logger.Sync()

		delta := time.Now().Sub(start)

		logger.Info("",
			zap.String("method", r.Method),
			zap.String("path", path),
			zap.Duration("response", delta),
			zap.ByteString("hash", hash),
		)
	}

	/*
		timeResponse := func(h http.HandlerFunc) http.HandlerFunc {
			return func(w http.ResponseWriter, r *http.Request) {
				start := time.Now()
				h(w, r)
				delta := time.Now().Sub(start)
				log.Printf("time=%s", delta)
			}
		}
	*/

	s := &http.Server{
		Addr:    ":" + *port,
		Handler: http.HandlerFunc(handleFunc),
	}

	log.Fatal(s.ListenAndServe())
}
