package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	hexStringOne := []byte("1c0111001f010100061a024b53535009181c")
	hexStringTwo := []byte("686974207468652062756c6c277320657965")

	hexResultOne := decodeHex(hexStringOne)
	//fmt.Println(hexResultOne)

	hexResultTwo := decodeHex(hexStringTwo)
	//fmt.Println(hexResultTwo)

	result := xor(hexResultOne, hexResultTwo)
	fmt.Printf("%s", encodeHex(result))
}

func decodeHex(src []byte) []byte {
	dst := make([]byte, hex.DecodedLen(len(src)))
	hex.Decode(dst, src)
	return dst
}

func encodeHex(src []byte) []byte {
	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)
	return dst
}

func xor(first, second []byte) []byte {
	resultBytes := make([]byte, len(first))

	for i := 0; i < len(first); i++ {
		resultBytes[i] = first[i] ^ second[i]
	}
	return resultBytes
}
