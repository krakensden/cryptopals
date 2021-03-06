package libcryptopals

import (
	"bufio"
	"os"
)

var word_list []string

func ReadList(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	rval := make([]string, 0, 4096)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		rval = append(rval, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		return nil, err
	}
	return rval, nil
}

// 'score' something as looking like english. Values from http://norvig.com/mayzner.html
func EScore(input []byte) int {
	var count int = 0

	vowel_map := map[byte]int{
		'A': 8,
		'a': 8,
		'C': 3,
		'c': 3,
		'B': 1,
		'b': 1,
		'E': 12,
		'e': 12,
		'D': 3,
		'd': 3,
		'G': 1,
		'g': 1,
		'F': 2,
		'f': 2,
		'I': 7,
		'i': 7,
		'H': 5,
		'h': 5,
		'K': 0,
		'k': 0,
		'J': 0,
		'j': 0,
		'M': 2,
		'm': 2,
		'L': 4,
		'l': 4,
		'O': 7,
		'o': 7,
		'N': 7,
		'n': 7,
		' ': 6,
		'Q': 0,
		'q': 0,
		'P': 2,
		'p': 2,
		'S': 6,
		's': 6,
		'R': 6,
		'r': 6,
		'U': 2,
		'u': 2,
		'T': 9,
		't': 9,
		'W': 1,
		'w': 1,
		'V': 1,
		'v': 1,
		'Y': 1,
		'y': 1,
		'X': 0,
		'x': 0,
		'Z': 0,
		'z': 0,
		'`': -1,
		'[': -1,
		']': -1,
		'<': -1,
		'>': -1,
		'(': -1,
		')': -1,
		';': -1,
	}
	for _, val := range input {
		if score, ok := vowel_map[val]; ok {
			count += score
		} else {
			// Make invalid characters have a negative value.
			// super inefficient, but maybe I don't care (?)
			found := false
			for _, valid := range valid_chars {
				if val == valid {
					found = true
					break
				}
			}
			if !found {
				count -= 5
			}
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
	src, err := Hex2Byte(input)
	if err != nil {
		return "", 0, err
	}

	most_likely, most_likely_score := SimpleSingleBitBruteForce(src)

	return string(ByteXor(src, most_likely)), most_likely_score, err
}

func SimpleSingleBitBruteForce(src []byte) (byte, int) {
	char_map := make(map[byte]int)
	var most_likely byte = 0
	var most_likely_score int = 0

	//for _, i := range valid_chars {
	for i := byte(0); i < 255; i++ {
		translated := ByteXor(src, i)
		char_map[i] = EScore(translated)
		if char_map[i] > most_likely_score {
			most_likely = i
			most_likely_score = char_map[i]
		}
	}
	return most_likely, most_likely_score
}
