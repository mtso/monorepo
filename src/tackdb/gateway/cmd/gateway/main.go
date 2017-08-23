package main

import (
	"log"
	"net/http"
	"os"

	flag "github.com/ogier/pflag"

	"gitlab.com/tackdb/gateway"
)

var (
	port       = flag.StringP("port", "p", "443", "Default port binds to 80")
	configdir  = flag.StringP("config-directory", "d", os.Getenv("HOME"), "Home directory.")
	configname = flag.StringP("config-name", "n", "gateway.conf", "Name of initial config file.")
)

func main() {
	flag.Parse()

	gateway.ConfigDir = *configdir
	router := &gateway.ReloadingRouter{}
	if err := router.Load(*configname); err != nil {
		log.Fatal(err)
	}

	log.Fatal(http.ListenAndServeTLS(":"+*port, "server.crt", "server.key", router))
}
