package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:3750")
	if err != nil {
		panic(err)
	}

	resp := bufio.NewReader(conn)
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Text to send: ")
		text, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		fmt.Fprintf(conn, text+"\n")

		msg, err := resp.ReadString('\n')
		if err != nil {
			panic(err)
		}

		fmt.Print("Server response: ", msg)
	}
}
