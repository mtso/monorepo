package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
)

const (
	SCHEME = "tcp"
)

var (
	port = *(flag.String("port", "3750", "Port to bind to. Defaults to 3750."))
)

func main() {
	fmt.Println("Listening on", port)
	portconn, err := net.Listen(SCHEME, ":"+port)
	if err != nil {
		panic(err)
	}

	for {
		conn, err := portconn.Accept()
		if err != nil {
			panic(err)
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	for {
		msg, err := bufio.NewReader(conn).ReadBytes('\n')
		if err != nil {
			panic(err)
		}

		fmt.Printf("Message received: %s", string(msg))
		resp := []byte("Echo: ")
		resp = append(resp, msg...)
		conn.Write([]byte(resp))
	}
}
