package main

import (
	"encoding/hex"
	"fmt"
)

const unencryptedString string = "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
const key string = "ICE"

func main() {
	fmt.Printf("%s", encodeHex(encrypt([]byte(unencryptedString))))
}

func encrypt(unencrypted []byte) []byte {
	var resultBytes []byte

	for i := 0; i < len(unencrypted); i++ {
		resultBytes = append(resultBytes, unencrypted[i]^key[i%3])
	}

	return resultBytes
}

func encodeHex(src []byte) []byte {
	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)
	return dst
}
