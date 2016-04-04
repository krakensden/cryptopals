package libcryptopals

import "fmt"

func GuessKeySize(input []byte) int {
	// Dumbest possible guess. Presumably will need to do something 'sophisticated'
	// like 'look at more than the first two blocks' instead
	min_distance := 41.0
	min_size := 1

	for KEYSIZE := 1; KEYSIZE < 41; KEYSIZE++ {
		if KEYSIZE*2 > len(input) {
			break
		}

		//single_distance, err := HammingDistance(input[0:KEYSIZE], input[KEYSIZE:KEYSIZE*2])
		//single_distance = single_distance / KEYSIZE
		single_distance, err := AverageNormalizedDistance(KEYSIZE, len(input)/KEYSIZE-1, input)
		if err != nil {
			panic(err)
		}

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
	return min_size
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
