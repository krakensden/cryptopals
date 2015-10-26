package main

import (
	"bufio"
	"fmt"
	"io"
	"libcryptopals"
	"os"
)

func SlurpB64EncodedFile(filename string) ([]byte, error) {
	// ignores newlines
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer func() {
		f.Close()
	}()

	reader := bufio.NewReader(f)
	var wholebuf []byte

	for {
		buf, _, err := reader.ReadLine()
		wholebuf = append(wholebuf, buf...)
		if err == io.EOF {
			break // Done
		}
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}
	return libcryptopals.Base642Byte(wholebuf)
}

func main() {
	input, err := SlurpB64EncodedFile("src/set1/challenge6/6.txt")
	if err != nil {
		panic(err)
	}

	// Dumbest possible guess. Presumably will need to do something 'sophisticated'
	// like 'look at more than the first two blocks' instead
	min_distance, min_size := 41, 0
	for KEYSIZE := 2; KEYSIZE < 41; KEYSIZE++ {
		single_distance, err := libcryptopals.NormalizedDistance(KEYSIZE, 1, input)
		if single_distance < min_distance {
			min_size = KEYSIZE
			min_distance = single_distance
		}
		fmt.Println("KEYSIZE ", KEYSIZE, " normalized distance: ", single_distance)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("Choosing ", min_size, " as the most promising key size")

	transposed_blocks := libcryptopals.TransposeBlocks(input, min_size)
	key_guesses := make([]byte, min_size)

	for i, block := range transposed_blocks {
		most_likely, _ := libcryptopals.SimpleSingleBitBruteForce(block)
		key_guesses[i] = most_likely
	}

	fmt.Println(key_guesses)
	fmt.Println("Key Guesses ", string(key_guesses))

	output, err := libcryptopals.SliceRepeatingXor(input, key_guesses)
	fmt.Println(string(output))
}
