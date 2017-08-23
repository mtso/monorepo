package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

var alphaPattern = regexp.MustCompile("[a-zA-Z ]")

func main() {
	buf, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	strings := bytes.Split(buf, []byte{'\n'})

	for i, str := range strings {
		str, err := hex.DecodeString(string(str))
		if err != nil {
			fmt.Println(err)
			continue
		}

		for j := byte(0); j < byte(255); j++ {
			decoded := xorWith(str, j)

			numSpace := count(decoded, ' ')

			if numSpace > 4 {
				fmt.Printf("%q %d %q line=%d\n", j, numSpace, decoded, i)
			}
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

func count(buf []byte, search byte) (count int) {
	for i := 0; i < len(buf); i++ {
		if buf[i] == search {
			count++
		}
	}
	return
}
