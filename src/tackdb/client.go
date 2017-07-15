package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:3750")
	if err != nil {
		panic(err)
	}

	// Wrap input/output pipes
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
		fmt.Println(text, text1, text2)
		if err != nil {
			panic(err)
		}

		// text = strings.Trim(text, "\n")

		fmt.Fprintf(conn, text)

		msg, err := serverOutput.ReadString('\n')
		if err == io.EOF {
			fmt.Println("server disconnected.")
			break
		} else if err != nil {
			panic(err)
		}

		// Log server response
		msg = strings.Trim(msg, "\n")
		fmt.Println(msg)
	}
}
