// Single-byte XOR Cipher
// 1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736
//
// Found: Cooking MC's like a pound of bacon
package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	encrypted := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	buf, err := hex.DecodeString(encrypted)
	if err != nil {
		panic(err)
	}

	for i := byte('A'); i <= 'z'; i++ {
		decoded := xorWith(buf, i)
		spaceCount := countSpaces(decoded)
		if spaceCount > 0 {
			fmt.Printf("char=%q spaces=%d string=%q\n", i, spaceCount, decoded)
		}
	}
}

func xorWith(buf []byte, ch byte) []byte {
	res := make([]byte, len(buf))
	for i := 0; i < len(buf); i++ {
		res[i] = buf[i] ^ ch
	}
	return res
}

func countSpaces(buf []byte) (count int) {
	for i := 0; i < len(buf); i++ {
		if buf[i] == byte(' ') {
			count++
		}
	}
	return
}
