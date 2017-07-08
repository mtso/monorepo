package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	// h := sha256.Sum224([]byte("hello~"))
	// s := fmt.Sprintf("%x", h)

	fmt.Printf("%T\n", hashid([]byte("hello~")))
}

func hashid(in []byte) string {
	h := sha256.Sum224(in)
	return fmt.Sprintf("%x", h)[:24]
}
