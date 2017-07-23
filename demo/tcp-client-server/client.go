package main

import (
	"bufio"
	"fmt"
	"net"
	// "os"
	// "time"
)

func main() {

	done := make(chan int)

	for i := 0; i < 40; i += 1 {
		go func(num int) {
			conn, err := net.Dial("tcp", "127.0.0.1:3750")
			check(err)

			fmt.Fprintf(conn, fmt.Sprintf("ping %d\n", num))
			msg, err := bufio.NewReader(conn).ReadString('\n')
			check(err)

			fmt.Println(msg)
			done <- 0
		}(i)
	}

	for i := 0; i < 40; i += 1 {
		<-done
	}
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
