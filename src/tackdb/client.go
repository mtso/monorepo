package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:3750")
	if err != nil {
		panic(err)
	}

	serverOutput := bufio.NewReader(conn)
	stdin := bufio.NewReader(os.Stdin)

	msg, err := serverOutput.ReadString('\n')
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", msg)

	for {
		fmt.Print("Text to send: ")
		text, err := stdin.ReadString('\n')
		if err != nil {
			panic(err)
		}

		fmt.Fprintf(conn, text+"\n")

		msg, err := serverOutput.ReadString('\n')
		fmt.Println(msg)
		if err == io.EOF {
			fmt.Println("server disconnected.")
			break
		} else if err != nil {
			panic(err)
		}

		fmt.Print("Server response: ", msg)
	}
}
