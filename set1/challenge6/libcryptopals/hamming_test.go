package libcryptopals

import (
	"testing"
)

func TestHex2Base64(t *testing.T) {
	astr := []byte("this is a test")
	bstr := []byte("wokka wokka!!!")
	if hd, _ := HammingDistance(astr, bstr); hd != 37 {
		t.Errorf("Incorrect hamming distance, expected 37, got ", hd)
	}
	if hd, _ := HammingDistance(bstr, astr); hd != 37 {
		t.Errorf("Incorrect hamming distance, expected 37, got ", hd)
	}
	if _, err := HammingDistance(bstr, []byte("x")); err == nil {
		t.Errorf("Expected an error for strings with different lengths")
	}
}
