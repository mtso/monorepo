package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	buf, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	first := buf[0]>>4 | buf[len(buf)-1]<<4
	rored := make([]byte, len(buf))
	rored[0] = first

	for i := 1; i < len(buf); i++ {
		rored[i] = buf[i-1]<<4 | buf[i]>>4
	}

	fmt.Printf("%s", rored)
}

func sketch() {
	buf := []byte{127, 65, 66}

	first := buf[0]>>4 | buf[len(buf)-1]<<4

	rored := make([]byte, len(buf))

	rored[0] = first

	for i := 1; i < len(buf); i++ {
		rored[i] = buf[i-1]<<4 | buf[i]>>4
	}

	fmt.Printf("%b %b\n", buf, rored)
	fmt.Printf("%b\n", []byte{255, 1, 2, 1 | 2})
}
