package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func to_hex(input []byte) string {
	return hex.EncodeToString(input)
}

func from_str(input string) ([]byte, error) {
	return hex.DecodeString(input)
}

func hex_to_base64(input string) string{
	decodehex, _ := from_str(input)
	enc64 := base64.StdEncoding.EncodeToString(decodehex)
	return enc64

}

func fixed_xor(param1, param2 string) []byte {

	decodeinput, _ := from_str(param1)
	decodesecond, _ := from_str(param2)
	output := make([] byte, 0)

	for i := range decodeinput {
		output = append(output, decodeinput[i] ^ decodesecond[i])
	}
	return output
}

func main(){
	hex_to_base64("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	fixed_xor("1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965")
}


