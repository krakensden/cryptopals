package libcryptopals

import "fmt"

func GuessKeySize(input []byte) int {
	// Dumbest possible guess. Presumably will need to do something 'sophisticated'
	// like 'look at more than the first two blocks' instead
	min_distance, min_size := 41, 1

	for KEYSIZE := 1; KEYSIZE < 41; KEYSIZE++ {
		single_distance, err := NormalizedDistance(KEYSIZE, 1, input)

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
	return min_distance
}

func DecodeXorEncodedBlock(input []byte) []byte {
	min_size := GuessKeySize(input)
	transposed_blocks := TransposeBlocks(input, min_size)
	key_guesses := make([]byte, min_size)

	for i, block := range transposed_blocks {
		most_likely, _ := SimpleSingleBitBruteForce(block)
		key_guesses[i] = most_likely
	}

	fmt.Println(key_guesses)
	fmt.Println("Key Guesses ", string(key_guesses))

	output, err := SliceRepeatingXor(input, key_guesses)
	if err != nil {
		fmt.Println("slicerepeatingxor ", err)
		return nil
	}
	return output
}
