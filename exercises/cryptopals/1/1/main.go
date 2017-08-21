// Convert hex to base64
// Input string: 49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d
// Output string: SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t
//
// Warning: Always operate on raw bytes, never on encoded strings. Only use hex and base64 for pretty-printing.
//
// Outline:
// decode hex string into byte array
// encode byte array into base64 string
//
package main

import (
	"fmt"
)

var str = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"

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

var itob = map[int]byte{
	0:  'A',
	1:  'B',
	2:  'C',
	3:  'D',
	4:  'E',
	5:  'F',
	6:  'G',
	7:  'H',
	8:  'I',
	9:  'J',
	10: 'K',
	11: 'L',
	12: 'M',
	13: 'N',
	14: 'O',
	15: 'P',
	16: 'Q',
	17: 'R',
	18: 'S',
	19: 'T',
	20: 'U',
	21: 'V',
	22: 'W',
	23: 'X',
	24: 'Y',
	25: 'Z',
	26: 'a',
	27: 'b',
	28: 'c',
	29: 'd',
	30: 'e',
	31: 'f',
	32: 'g',
	33: 'h',
	34: 'i',
	35: 'j',
	36: 'k',
	37: 'l',
	38: 'm',
	39: 'n',
	40: 'o',
	41: 'p',
	42: 'q',
	43: 'r',
	44: 's',
	45: 't',
	46: 'u',
	47: 'v',
	48: 'w',
	49: 'x',
	50: 'y',
	51: 'z',
	52: '0',
	53: '1',
	54: '2',
	55: '3',
	56: '4',
	57: '5',
	58: '6',
	59: '7',
	60: '8',
	61: '9',
	62: '+',
	63: '/',
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

func (d data) ToBase64() string {
	buffer := _buffer64frombytes(len(d))

	i, j := 0, 0
	for ; i < len(d)-2; i += 3 {
		buffer[j] = itob[int(d[i]>>2)]
		buffer[j+1] = itob[int(((d[i]&3)<<4)|d[i+1]>>4)]
		buffer[j+2] = itob[int(((d[i+1]&15)<<2)|(d[i+2]>>6))]
		buffer[j+3] = itob[int(d[i+2]&63)]
		j += 4
	}

	if len(buffer)-j > 0 {
		buffer[j] = itob[int(d[i]>>2)]
	}
	if len(buffer)-j > 1 {
		buffer[j+1] = itob[int(((d[i]&3)<<4)|d[i+1]>>4)]
	}
	if len(buffer)-j > 2 {
		buffer[j+2] = itob[int(((d[i+1]&15)<<2)|(d[i+2]>>6))]
	}

	return string(buffer)
}

func main() {
	buf := DecodeHex(str)

	fmt.Println(buf.ToBase64())
}

func htob(encoded string) []byte {
	var buf []byte
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

func isEven(i int) bool { return i%2 == 0 }

// unpack every 6 bits into a byte array
// then return string constructed from byte array
func bto6(bin []byte) string {
	buffer := _buffer64frombytes(len(bin))

	i, j := 0, 0
	for ; i < len(bin); i += 3 {
		buffer[j] = itob[int(bin[i]>>2)]
		buffer[j+1] = itob[int(((bin[i]&3)<<4)|bin[i+1]>>4)]
		buffer[j+2] = itob[int(((bin[i+1]&15)<<2)|(bin[i+2]>>6))]
		buffer[j+3] = itob[int(bin[i+2]&63)]
		j += 4
	}

	if len(buffer)-j > 0 {
		buffer[j] = itob[int(bin[i]>>2)]
	}
	if len(buffer)-j > 1 {
		buffer[j+1] = itob[int(((bin[i]&3)<<4)|bin[i+1]>>4)]
	}
	if len(buffer)-j > 2 {
		buffer[j+2] = itob[int(((bin[i+1]&15)<<2)|(bin[i+2]>>6))]
	}

	return string(buffer)
}

func _buffer64frombytes(bytes int) []byte {
	buflen := (bytes / 3) * 4

	rem := bytes % 3
	if rem == 1 {
		buflen += 1
	} else if rem == 2 {
		buflen += 2
	}

	return make([]byte, buflen)
}
