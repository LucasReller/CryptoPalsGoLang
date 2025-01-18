package main

import (
	"encoding/hex"
	"fmt"
)

const string1 string = "this is a test"
const string2 string = "wokka wokka!!!"

func main() {
	fmt.Println(computeEditDistance([]byte(string1), []byte(string2)))
}

func encodeHex(src []byte) []byte {
	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)
	return dst
}

func decodeHex(src []byte) []byte {
	dst := make([]byte, hex.DecodedLen(len(src)))
	hex.Decode(dst, src)
	return dst
}

func computeEditDistance(value1, value2 []byte) int {
	var length int
	var result uint32

	//use the shorter of the two lengths for iterating
	if len(value1) < len(value2) {
		length = len(value1)
	} else {
		length = len(value2)
	}

	for i := 0; i < length; i++ {
		temp := uint32(value1[i] ^ value2[i])
		//Kernighan’s Algorithm
		for temp > 0 {
			result += temp & 1 // check last bit (1 is equivalent to 0001)
			temp = temp >> 1
		}
	}
	return int(result)
}
