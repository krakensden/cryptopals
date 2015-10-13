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

func normalizedDistance(chunk_size int, chunks_to_test int, text []byte) (int, error) {
	var normalized_distance int = 0
	for index := 0; index < chunks_to_test; index++ {
		starting := chunk_size * index
		distance, err := libcryptopals.HammingDistance(text[starting:starting+chunk_size], text[starting+chunk_size:starting+chunk_size*2])
		if err != nil {
			return 0, err
		}
		normalized_distance = normalized_distance + distance/chunk_size
	}
	return normalized_distance / chunks_to_test, nil
}

func transposeBlocks(input []byte, chunk_size int) [][]byte {
	var output [][]byte = make([][]byte, chunk_size)
	for index, val := range input {
		output[index%chunk_size] = append(output[index%chunk_size], val)
	}
	return output
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
		single_distance, err := normalizedDistance(KEYSIZE, 1, input)
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

	transposed_blocks := transposeBlocks(input, min_size)
	key_guesses := make([]byte, min_size)

	for i, block := range transposed_blocks {
		most_likely, _ := libcryptopals.SimpleSingleBitBruteForce(block)
		key_guesses[i] = most_likely
	}
	fmt.Println(key_guesses)
	fmt.Println(string(key_guesses))
	output, err := libcryptopals.SliceRepeatingXor(input, key_guesses)
	fmt.Println(string(output))
}
