package xordecode

import (
	"../../challenge1/hex2b64"
	"../../challenge2/bytexor"
)

// Count the number of 'e' characters in a sequence of ascii bytes
func EScore(input []byte) int {
	var count int = 0
	vowel_map := map[byte]bool{
		'e': true,
		't': true,
		'a': true,
		'o': true,
		'i': true,
		'n': true,
		's': true,
		'h': true,
		'r': true,
		'd': true,
		'l': true,
		'u': true,
		'E': true,
		'T': true,
		'A': true,
		'O': true,
		'I': true,
		'N': true,
		'S': true,
		'H': true,
		'R': true,
		'D': true,
		'L': true,
		'U': true,
		' ': true,
	}
	for _, val := range input {
		if vowel_map[val] {
			count++
		}
	}
	return count
}

func UnPrintableScore(input []byte) int {
	var count int = 0
	for _, val := range input {
		if val < 0x20 || val > 0x7e {
			count++
		}
	}
	return count
}

func SingleBitBruteForce(input string) (string, int, error) {
	src, err := hex2b64.Hex2Byte(input)
	char_map := make(map[byte]int)
	if err != nil {
		return "", 0, err
	}
	var most_likely byte
	var most_likely_score int
	// populate with the printable characters
	for i := byte(0x0); i < 0xff; i++ {
		translated := bytexor.ByteXor(src, i)
		char_map[i] = EScore(translated)
		if UnPrintableScore(translated) == 0 && char_map[i] > most_likely_score {
			most_likely = i
			most_likely_score = char_map[i]
		}
	}
	return string(bytexor.ByteXor(src, most_likely)), most_likely_score, err
}
