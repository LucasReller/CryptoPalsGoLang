package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

const keySizeMin int = 2
const keySizeMax int = 40
const chars string = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM1234567890"
const vowels string = "aeiouy" //these 6 letters account for 39.8% of the characters found in english sentences
const targetPercent float64 = .398

func main() {
	fmt.Println(computeEditDistance([]byte("this is a test"), []byte("wokka wokka!!!")))

	hexStrings := readInFileByLines()
	s, _ := base64.StdEncoding.DecodeString(hexStrings)
	keySize := findKeySize(string(s), 40)
	fmt.Println("Keysize:", keySize)

	keyBlocks := createBlocks(s, keySize)
	key := findKey(keyBlocks)
	fmt.Println("Key:", key)

}

func readInFileByLines() string {
	file, err := os.ReadFile("challenge6-data.txt")
	if err != nil {
		log.Fatal(err)
	}

	return string(file)
}

func findKey(arr [][]byte) string {
	var temp byte
	//var percent float64
	var result string

	for i := 0; i < len(arr); i++ {
		temp, _ = scoreEachLetter(arr[i])
		result += string(temp)
	}
	return result
}

var letterFrequencies = map[string]float64{
	"A": 8.167,
	"B": 1.492,
	"C": 2.782,
	"D": 4.253,
	"E": 12.702,
	"F": 2.228,
	"G": 2.015,
	"H": 6.094,
	"I": 6.966,
	"J": 0.153,
	"K": 0.772,
	"L": 4.025,
	"M": 2.406,
	"N": 6.749,
	"O": 7.507,
	"P": 1.929,
	"Q": 0.095,
	"R": 5.987,
	"S": 6.327,
	"T": 9.056,
	"U": 2.758,
	"V": 0.978,
	"W": 2.360,
	"X": 0.150,
	"Y": 1.974,
	"Z": 0.074,
}

var text = regexp.MustCompile("^[a-zA-Z ]$")

func isAlphabetic(str string) bool {
	return text.MatchString(str)
}

func calculateScore(buffer []byte) float64 {
	score := 0.0
	for _, b := range buffer {
		str := string(b)
		if isAlphabetic(str) {
			score += letterFrequencies[strings.ToUpper(str)]
		} else {
			score -= 10.0
		}
	}

	return score
}

func singleByteXOR(buffer []byte, key byte) []byte {
	result := make([]byte, len(buffer))
	for i := range len(buffer) {
		result[i] = buffer[i] ^ key
	}

	return result
}

func scoreEachLetter(encodedValue []byte) (byte, float64) {
	bestKey := byte(0)
	bestScore := 0.0
	for k := range 256 {
		key := byte(k)
		decrypted := singleByteXOR(encodedValue, key)

		score := calculateScore(decrypted)

		if score > bestScore {
			bestScore = score
			bestKey = key
		}
	}
	return bestKey, bestScore
}

func createBlocks(file []byte, keysize int) [][]byte {
	var result [][]byte = make([][]byte, keysize)
	for i := 0; i < len(file); i++ {
		result[i%keysize] = append(result[i%keysize], file[i])

	}

	return result
}

func findKeySize(data string, numberOfSamples int) int {
	var result int
	var prev int
	var d1 []string = make([]string, numberOfSamples)
	var d2 []string = make([]string, numberOfSamples)
	var d3 []float32 = make([]float32, numberOfSamples)
	var tempDistance float32
	var temp []float32 = make([]float32, keySizeMax+2)
	var prevDistance float32 = 100 //maybe a better way to have a large default?

	for i := keySizeMin; i < keySizeMax; i++ {
		prev = 0
		for j := 0; j < numberOfSamples; j++ {
			d1[j] = data[prev : prev+i]
			prev += i
			d2[j] = data[prev : prev+i]
			d3[j] = float32(computeEditDistance([]byte(d1[j]), []byte(d2[j]))) / float32(i)
		}
		for _, num := range d3 {
			tempDistance += num
		}
		tempDistance /= float32(len(d3) - 1)
		temp[i] = tempDistance
		if tempDistance < prevDistance {
			result = i //if distance is shorter keysize is likely correct
			prevDistance = tempDistance
		}
	}
	return result
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
		//Kernighan’s Algorithm (bitshift right and check if last bit is set until the original value is 0 because of the bit shifting)
		for temp > 0 {
			result += temp & 1 // check last bit (1 is equivalent to 0001)
			temp = temp >> 1
		}
	}
	return int(result)
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
