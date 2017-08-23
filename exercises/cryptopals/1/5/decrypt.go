package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	in, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	decoded := make([]byte, len(in)/2)

	_, err = hex.Decode(decoded, in)
	if err != nil {
		panic(err)
	}

	out := decrypt(decoded, []byte("ICE"))

	fmt.Printf("%s", out)
}

func decrypt(in []byte, key []byte) []byte {
	out := make([]byte, len(in))
	for i := 0; i < len(in); i++ {
		out[i] = in[i] ^ key[i%len(key)]
	}
	return out
}
