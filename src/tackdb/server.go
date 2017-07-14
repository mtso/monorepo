package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"net"
)

const (
	SCHEME  = "tcp"
	VERSION = "0.0.1"
)

// Options
var (
	port = *(flag.String("port", "3750", "Port to bind to. Defaults to 3750."))
)

// Database
var DB map[string][]byte

func main() {
	fmt.Println("Listening on", port)
	portconn, err := net.Listen(SCHEME, ":"+port)
	if err != nil {
		panic(err)
	}
	defer portconn.Close()

	clientid := 1.0
	for {
		conn, err := portconn.Accept()
		if err != nil {
			panic(err)
		}

		go handleConnection(conn, clientid)
		clientid += 1
	}
}

var startupmsg = []byte("tackdb v" + VERSION + "\n")

func handleConnection(conn net.Conn, id float64) {
	defer conn.Close()

	conn.Write(startupmsg)

	for {
		msg, err := bufio.NewReader(conn).ReadBytes('\n')
		if err == io.EOF {
			fmt.Println("Client closed.", fmt.Sprintf("id=%v", id))
			break
		} else if err != nil {
			panic(err)
		}

		pieces := bytes.Split(msg, []byte(" "))
		fmt.Printf("%s\n", pieces)
		if len(pieces) == 1 && pieces[0][0] == 10 {
			resp := []byte("NOCOMMAND\n")
			fmt.Printf("Message received (%d): %q: %b\n", len(msg), msg, msg)
			fmt.Println("unrecognized command")
			conn.Write(resp)
			continue
		}

		fmt.Printf("Message received (%d): %b\n", len(msg), msg)
		resp := []byte("Echo: ")
		resp = append(resp, msg...)
		conn.Write([]byte(resp))
	}
}
