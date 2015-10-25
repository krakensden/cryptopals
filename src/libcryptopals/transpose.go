package libcryptopals

func TransposeBlocks(input []byte, chunk_size int) [][]byte {
	var output [][]byte = make([][]byte, chunk_size)
	for index, val := range input {
		output[index%chunk_size] = append(output[index%chunk_size], val)
	}
	return output
}
