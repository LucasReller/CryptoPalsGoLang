package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {
	hexString := []byte("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	hexResult := decodeHex(hexString)
	base64Result := encodeBase64(hexResult)
	fmt.Printf("%s", base64Result)
}

func decodeHex(src []byte) []byte {
	dst := make([]byte, hex.DecodedLen(len(src)))
	hex.Decode(dst, src)
	return dst
}

func encodeBase64(src []byte) []byte {
	dst := make([]byte, base64.StdEncoding.EncodedLen(len(src)))
	base64.StdEncoding.Encode(dst, src)
	return dst
}
