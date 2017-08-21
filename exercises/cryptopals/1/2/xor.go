// Fixed XOR
//
// Write a function that takes two equal-length buffers
// and produces their XOR combination.
// input1: 1c0111001f010100061a024b53535009181c
// input2: 686974207468652062756c6c277320657965
// expected: 746865206b696420646f6e277420706c6179
package main

import (
	"fmt"
)

func main() {
	input := "1c0111001f010100061a024b53535009181c"
	input2 := "686974207468652062756c6c277320657965"
	fmt.Println(DecodeHex(input).XorWith(DecodeHex(input2)).EncodeHex())
}

var htoi = map[byte]byte{
	'0': 0,
	'1': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'a': 10,
	'b': 11,
	'c': 12,
	'd': 13,
	'e': 14,
	'f': 15,
}

var itoh = map[byte]byte{
	0:  '0',
	1:  '1',
	2:  '2',
	3:  '3',
	4:  '4',
	5:  '5',
	6:  '6',
	7:  '7',
	8:  '8',
	9:  '9',
	10: 'a',
	11: 'b',
	12: 'c',
	13: 'd',
	14: 'e',
	15: 'f',
}

type data []byte

func DecodeHex(encoded string) data {
	var buf data
	if isEven(len(encoded)) {
		buf = make([]byte, len(encoded)/2)
	} else {
		buf = make([]byte, len(encoded)/2+1)
	}

	for i := 0; i < len(encoded); i++ {
		b := i / 2
		buf[b] = buf[b] | htoi[encoded[i]]

		if isEven(i) {
			buf[b] = buf[b] << 4
		}
	}

	return buf
}

func (d data) EncodeHex() string {
	buf := make([]byte, 0)

	for i := 0; i < len(d); i++ {
		buf = append(buf, itoh[d[i]>>4])
		buf = append(buf, itoh[d[i]&15])
	}

	return string(buf)
}

func (d data) XorWith(e data) data {
	f := make(data, len(d))
	for i := 0; i < len(d); i++ {
		f[i] = d[i] ^ e[i]
	}
	return f
}

func isEven(i int) bool { return i%2 == 0 }
