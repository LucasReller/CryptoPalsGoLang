package main

import (
	"crypto/aes"
	"encoding/base64"
	"fmt"
	"log"
	"os"
)

const key string = "YELLOW SUBMARINE"

func main() {
	fmt.Println(string(DecryptAES128Ecb([]byte(readInFileByLines()), []byte(key))))
}

func DecryptAES128Ecb(data, key []byte) []byte {
	cipher, _ := aes.NewCipher(key)
	decryptedResult := make([]byte, len(data))

	for i, j := 0, 16; i < len(data); i, j = i+16, j+16 {
		cipher.Decrypt(decryptedResult[i:j], data[i:j])
	}

	return decryptedResult
}

func readInFileByLines() string {
	file, err := os.ReadFile("challenge7-data.txt")
	if err != nil {
		log.Fatal(err)
	}
	file, err = base64.StdEncoding.DecodeString(string(file))
	if err != nil {
		log.Println(err.Error())
	}
	return string(file)
}
