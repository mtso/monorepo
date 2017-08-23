package main

import (
	"bytes"
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

	// var found []byte

	strings := bytes.Split(buf, []byte{'\n'})

	for i, str := range strings {
		// str := strings[0]
		// i := 1
		for j := byte('A'); j < byte('Z'); j++ {
			decoded := xorWith(str, j)

			if len(alphaPattern.FindAll(decoded, -1)) > 53 {
				spaceCount := count(decoded, ' ')
				// if spaceCount > 0 {
				fmt.Printf("%q %d %q line=%d\n", j, spaceCount, decoded, i)
				// }
			}
		}

		for j := byte('a'); j < byte('z'); j++ {
			decoded := xorWith(str, j)

			if len(alphaPattern.FindAll(decoded, -1)) > 53 {
				spaceCount := count(decoded, ' ')
				// if spaceCount > 0 {
				fmt.Printf("%q %d %q line=%d\n", j, spaceCount, decoded, i)
				// }
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
