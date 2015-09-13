package libcryptopals

import "errors"

func SliceXor(a, b []byte) ([]byte, error) {
	if len(a) != len(b) {
		return []byte{}, errors.New("mismatched input lengths")
	}
	output := make([]byte, len(a), len(a))
	for i, aval := range a {
		output[i] = aval ^ b[i]
	}
	return output, nil
}

func ByteXor(a []byte, b byte) []byte {
	output := make([]byte, len(a), len(a))
	for i, aval := range a {
		output[i] = aval ^ b
	}
	return output
}
