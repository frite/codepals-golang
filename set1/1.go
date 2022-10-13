package set_one

import (
	"encoding/base64"
	"encoding/hex"
)

func from_str(input string) ([]byte, error) {
	return hex.DecodeString(input)
}

func hex_to_base64(input string) string{
	decodehex, _ := from_str(input)
	enc64 := base64.StdEncoding.EncodeToString(decodehex)
	return enc64
}


func HexToBase64(hexString string) []byte {
	hexBytes, err := hex.DecodeString(hexString)
	if err != nil {
		panic(err)
	}

	result := make([]byte, base64.StdEncoding.EncodedLen(len(hexBytes)))
	base64.StdEncoding.Encode(result, hexBytes)
	return result
}