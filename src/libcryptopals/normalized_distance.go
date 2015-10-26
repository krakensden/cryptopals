package libcryptopals

//import "fmt"

func NormalizedDistance(chunk_size int, chunks_to_test int, text []byte) (int, error) {
	var normalized_distance int = 0
	for index := 0; index < chunks_to_test; index++ {
		starting := chunk_size * index
		distance, err := HammingDistance(text[starting:starting+chunk_size], text[starting+chunk_size:starting+chunk_size*2])
		//fmt.Println(starting, starting+chunk_size, starting+chunk_size, starting+chunk_size*2, text[starting:starting+chunk_size], text[starting+chunk_size:starting+chunk_size*2], distance)
		if err != nil {
			return 0, err
		}
		normalized_distance = normalized_distance + distance/chunk_size
	}
	return normalized_distance / chunks_to_test, nil
}
