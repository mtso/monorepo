package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"golang.org/x/sys/unix"
)

const SIZE = 4096
const MESSAGE = "hello~ hello~ hello~"

func main() {
	writedata := make([]byte, SIZE)

	filename := os.Args[1]
	fmt.Println(filename)

	ioutil.WriteFile(filename, writedata, 777)

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 700) //os.O_RDWR | os.O_APPEND, 0777)
	if err != nil {
		panic(err)
	}

	mtype := unix.PROT_READ | unix.PROT_WRITE // | unix.PROT_GROWSUP

	data, err := unix.Mmap(int(file.Fd()), 0, SIZE, mtype, unix.MAP_SHARED)
	if err != nil {
		panic(err)
	}

	fmt.Println(len(data))

	msg := MESSAGE

	for i := 0; i < len(msg); i++ {
		data[i] = msg[i]
	}
}
