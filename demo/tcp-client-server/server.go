package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	scheme := "tcp"
	port := "3750"

	fmt.Println("Listening on", port)
	ln, err := net.Listen(scheme, ":"+port)
	check(err)

	for {
		conn, err := ln.Accept()
		check(err)

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	msg, err := bufio.NewReader(conn).ReadString('\n')
	check(err)

	fmt.Print("Message received:", string(msg))
	resp := "Echo: " + string(msg)
	conn.Write([]byte(resp + "\n"))
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
