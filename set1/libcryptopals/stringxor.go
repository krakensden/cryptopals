package libcryptopals

func SliceRepeatingXor(text, key []byte) ([]byte, error) {
	full_key := make([]byte, len(text), len(text))
	for index, _ := range text {
		full_key[index] = key[index%len(key)]
	}
	return SliceXor(text, full_key)
}

func StringRepeatingXor(text, key string) (string, error) {
	output, err := SliceRepeatingXor([]byte(text), []byte(key))
	return Byte2Hex(output), err
}
