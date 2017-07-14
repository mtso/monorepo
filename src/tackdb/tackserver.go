package main

import (
	"flag"
	"log"

	"github.com/mtso/monorepo/src/tackdb/server"
)

var (
	port = *(flag.String("port", "3750", "Port to bind to. Defaults to 3750"))
)

func main() {
	srvr := &server.Server{}
	log.Println("tackdb", "v"+server.VERSION, "listening on", port)
	log.Fatal(srvr.Start(port))
}
