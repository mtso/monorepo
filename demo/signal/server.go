package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var message []byte

func main() {
	if msg, err := ioutil.ReadFile("./message.txt"); err != nil {
		log.Println(err)
	} else {
		message = msg
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(message)
	})

	signals := make(chan os.Signal)
	signal.Notify(signals, syscall.SIGIO)

	go func() {
		for {
			select {
			case <-signals:
				if msg, err := ioutil.ReadFile("./message.txt"); err != nil {
					log.Println(err)
				} else {
					message = msg
					log.Printf("Reloaded message to: %q", msg)
				}
			}
		}
	}()

	pid := []byte(fmt.Sprintf("%d", syscall.Getpid()))
	ioutil.WriteFile("./pid", pid, os.ModePerm)
	http.ListenAndServe(":3750", nil)
}
