package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:3750")
	check(err)

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, err := reader.ReadString('\n')
		check(err)

		fmt.Fprintf(conn, text+"\n")
		msg, err := bufio.NewReader(conn).ReadString('\n')
		check(err)

		fmt.Print("Message from server:", msg)
	}
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
