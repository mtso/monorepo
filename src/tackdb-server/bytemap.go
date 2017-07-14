package main

import (
	"fmt"
)

func main() {
	// bytemap := make(map[[]byte][]byte)
	bytemap := make(map[string][]byte)
	key := "foo" // []byte("foo")
	val := []byte("bar")
	bytemap[key] = val

	ret := bytemap[key]
	fmt.Printf("%s\n", ret)
}
