// Repeating-Key XOR
//
// Run with:
// $ cat opening.txt | go run encrypt.go | echo $(cat)

package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

type Key []byte

// read input
// for each character
//   xor with key char
//   increment key index
func main() {
	key := flag.String("key", "ICE", "Encryption key.")

	in, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	out := Key(*key).Encrypt(in)

	fmt.Print(hex.EncodeToString(out))
}

// Encrypt encodes the input byte array with a repeating-key XOR method.
func (k Key) Encrypt(in []byte) []byte {
	out := make([]byte, len(in))
	for i := 0; i < len(in); i++ {
		out[i] = in[i] ^ k[i%len(k)]
	}
	return out
}
