package set_one

import (
	"testing"
)

func TestHexToBase64(t *testing.T) {
	expected := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	base64 := hex_to_base64("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	
	if !(base64 == expected) {
		t.Errorf("Expected: %v Got: %v", expected, base64)
	}
	/*
	if !bytes.Equal(base64, expected) {
		t.Errorf("Expected: %v Got: %v", expected, base64)
	}*/
}