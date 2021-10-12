package main

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math"
)

/*
* From http://www.data-compression.com/english.html
*  This
* http://practicalcryptography.com/cryptanalysis/letter-frequencies-various-languages/english-letter-frequencies/
* may be more useful and is more rich in data, trigrams etc.
*/
var ENGLISH_FREQUENCY = map[byte]float64 {
    'a': 0.0651738,
    'b': 0.0124248,
    'c': 0.0217339,
    'd': 0.0349835,
    'e': 0.1041442,
    'f': 0.0197881,
    'g': 0.0158610,
    'h': 0.0492888,
    'i': 0.0558094,
    'j': 0.0009033,
    'k': 0.0050529,
    'l': 0.0331490,
    'm': 0.0202124,
    'n': 0.0564513,
    'o': 0.0596302,
    'p': 0.0137645,
    'q': 0.0008606,
    'r': 0.0497563,
    's': 0.0515760,
    't': 0.0729357,
    'u': 0.0225134,
    'v': 0.0082903,
    'w': 0.0171272,
    'x': 0.0013692,
    'y': 0.0145984,
    'z': 0.0007836,
    ' ': 0.1918182,
}

func to_hex(input []byte) string {
	return hex.EncodeToString(input)
}

func from_str(input string) ([]byte, error) {
	return hex.DecodeString(input)
}

func str_from_bytes(input []byte) string {
	return string(input)
}

func hex_to_base64(input string) string{
	decodehex, _ := from_str(input)
	enc64 := base64.StdEncoding.EncodeToString(decodehex)
	return enc64

}

func fixed_xor(param1, param2 string) []byte {

	decodeinput, _ := from_str(param1)
	decodesecond, _ := from_str(param2)
	output := make([]byte, 0)

	for i := range decodeinput {
		output = append(output, decodeinput[i] ^ decodesecond[i])
	}
	return output
}

func xor(input, key []byte){
	for i := 0; i < len(key); i += 1 {
		input[i] ^= key[i]
	}

}

func calcEtaoinShrdlu(input []byte) float64 {
	// https://en.wikipedia.org/wiki/Etaoin_shrdlu
	// Taken from https://github.com/philandstuff/cryptopals-go/blob/59b2c4790a72360389aa4b84717d314da2cbc0dd/english.go#L38
	var total_per_letter [26]int
	for _, b := range input {
		if b >= 'a' && b <= 'z' {
			total_per_letter[b-'a']++
		}
	}
	coeff := float64(0)
	for i := 0; i < 26; i++ {
		coeff += math.Sqrt(ENGLISH_FREQUENCY[byte(i+'a')] * float64(total_per_letter[i]) / float64(len(input)))
	}
	return coeff
}

func single_byte_xor(input []byte) ([]byte, byte, float64){
	size := len(input)
	/*
	* WHY ON EARTH IS THIS NOT WORKING?
	results := make(map[byte][]byte)
	for i := 0; i< 256; i += 1 {
		results[i], _ := xor(input, bytes.Repeat([]byte{byte(i)}, size))
	}
	*/
	var score float64
	var plaintext []byte
	var key byte
	for i := 0; i < 256; i++ {
		tempResult := make([]byte, len(input))
		copy(tempResult, input)
		xor(tempResult, bytes.Repeat([]byte{byte(i)}, size))
		tempScore := calcEtaoinShrdlu(tempResult)
		if tempScore > score {
			score = tempScore
			plaintext = tempResult
			key = byte(i)
		}
	}
	return plaintext, key, score

}

func main(){
	hex_to_base64("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	fixed_xor("1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965")
	cipher, _ := from_str("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	ptext, _, _ := single_byte_xor(cipher)
	fmt.Println(str_from_bytes(ptext))
}


