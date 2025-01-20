package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

const vowels string = "aeiouy"   //these 6 letters account for 39.8% of the characters found in english sentences
const badChars string = "%|{}[]" //if it contains any of these it cannot be the solution
const letters string = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM1234567890"
const targetPercent float64 = .398

func main() {
	hexStrings := readInFileByLines()

	hexResult := checkEachLine(hexStrings)

	fmt.Printf("%s", hexResult)
}

func readInFileByLines() []string {
	file, err := os.ReadFile("challenge4-data.txt")
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(string(file), "\n")
}

func checkEachLine(arr []string) []byte {
	var resultBytes []byte = make([]byte, 30) //change this from being hardcoded
	var prevPercent float64 = 1
	for i := 0; i < len(arr); i++ {
		bytes, percent := scoreEachLetter(decodeHex([]byte(arr[i])))

		if percent < prevPercent {
			prevPercent = percent
			copy(resultBytes, bytes)
		}
	}
	return resultBytes
}

func scoreEachLetter(encodedValue []byte) ([]byte, float64) {
	var resultPercent float64 = 0
	var tempBytes []byte = make([]byte, len(encodedValue))
	var resultBytes []byte = make([]byte, len(encodedValue))

	for i := 0; i < len(letters); i++ {
		var vowelCounter int = 0
		var badCounter int = 0

		for j := 0; j < len(encodedValue); j++ {
			var charResult = encodedValue[j] ^ letters[i]

			if strings.Contains(badChars, string(charResult)) {
				break
			} else if strings.Contains(vowels, string(charResult)) {
				vowelCounter++
			}

			tempBytes[j] = charResult
		}
		if badCounter == 0 {
			var currentPercent float64 = float64(vowelCounter) / float64(len(tempBytes))

			//the closer to zero the result of the subtraction is the closer to the target percent, therefore the smaller value is closer
			if math.Abs(currentPercent-targetPercent) < math.Abs(resultPercent-targetPercent) {
				resultPercent = currentPercent

				copy(resultBytes, tempBytes)
			}
		}
	}
	return resultBytes, math.Abs(resultPercent - targetPercent)
}

func decodeHex(src []byte) []byte {
	dst := make([]byte, hex.DecodedLen(len(src)))
	hex.Decode(dst, src)
	return dst
}
