// Single-file, multi-goroutine TCP client/server exchange.
package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func startServer() {
	conn, err := net.Listen("tcp", ":3750")
	if err != nil {
		panic(err)
	}

	for {
		client, err := conn.Accept()
		if err != nil {
			panic(err)
		}

		msg, err := bufio.NewReader(client).ReadString('\n')
		if err != nil {
			panic(err)
		}
		fmt.Println(msg)
		client.Write([]byte("echo: " + msg))
	}
}

func main() {
	go startServer()

	t := time.Tick(500)
	select {
	case <-t:
		conn, err := net.Dial("tcp", "127.0.0.1:3750")
		if err != nil {
			panic(err)
		}

		conn.Write([]byte("hello~\n"))

		msg, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			panic(err)
		}

		fmt.Println(msg)
	}
}
