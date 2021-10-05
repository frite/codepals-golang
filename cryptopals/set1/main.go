package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func to_hex(input []byte) string {
	return hex.EncodeToString(input)
}

func hex_to_base64(hexstring string) string{
	decodehex, _ := hex.DecodeString(hexstring)
	enc64 := base64.StdEncoding.EncodeToString(decodehex)
	return enc64

}

func fixed_xor(input, second string) []byte {
	decodeinput, _ := hex.DecodeString(input)
	decodesecond, _ := hex.DecodeString(second)
	output := make([] byte, 0)

	for i := range decodeinput {
		output = append(output, decodeinput[i] ^ decodesecond[i])
	}
	return output
}

func main(){
	hex_to_base64("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	res := fixed_xor("1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965")
	fmt.Println("%s", to_hex(res))
}


