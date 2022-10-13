package set_one

import (
	"encoding/base64"
	"encoding/hex"
)

func HexToBase64(hexString string) []byte {
	hexBytes, err := hex.DecodeString(hexString)
	if err != nil {
		panic(err)
	}

	result := make([]byte, base64.StdEncoding.EncodedLen(len(hexBytes)))
	base64.StdEncoding.Encode(result, hexBytes)
	return result
}