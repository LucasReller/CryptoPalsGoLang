package main

import (
	"encoding/hex"
	"fmt"
	"math"
)

const vowels string = "aeiouy" //these 6 letters account for 39.8% of the characters found in english sentences
const letters string = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"
const targetPercent float64 = 39.8

func main() {
	hexString := []byte("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")

	hexResult := decodeHex(hexString)

	fmt.Printf("%s", scoreEachLetter(hexResult))
}

func decodeHex(src []byte) []byte {
	dst := make([]byte, hex.DecodedLen(len(src)))
	hex.Decode(dst, src)
	return dst
}

func scoreEachLetter(encodedValue []byte) []byte {
	var resultPercent float64 = 0
	var tempBytes []byte = make([]byte, len(encodedValue))
	var resultBytes []byte = make([]byte, len(encodedValue))

	for i := 0; i < len(letters); i++ {
		var vowelCounter int = 0

		for j := 0; j < len(encodedValue); j++ {
			var charResult = encodedValue[j] ^ letters[i]

			for k := 0; k < len(vowels); k++ {
				if charResult == vowels[k] {
					vowelCounter++
				}
			}
			tempBytes[j] = charResult
		}
		var currentPercent float64 = float64(vowelCounter) / float64(len(tempBytes))

		//the closer to zero the result of the subtraction is the closer to the target percent, therefore the smaller value is closer
		if math.Abs(currentPercent-targetPercent) < math.Abs(resultPercent-targetPercent) {
			resultPercent = currentPercent

			//key golang learning form this:
			//make (see above) makes a slice which contains a pointer to the underlying array, as such assignment will assign by pointer not value leading to undesireable behavior
			//however when using copy on a slice it will copy to match the size of the largest slice thus the destination array must be of equal or larger size to copy all values
			copy(resultBytes, tempBytes)
		}
	}
	return resultBytes
}
