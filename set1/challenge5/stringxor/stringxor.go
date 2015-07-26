package stringxor

import (
	"../../challenge1/hex2b64"
	"../../challenge2/bytexor"
)

func SliceXor(text, key []byte) ([]byte, error) {
	full_key := make([]byte, len(text), len(text))
	for index, _ := range text {
		full_key[index] = key[index%len(key)]
	}
	return bytexor.SliceXor(text, full_key)
}

func StringXor(text, key string) (string, error) {
	output, err := SliceXor([]byte(text), []byte(key))
	return hex2b64.Byte2Hex(output), err
}
