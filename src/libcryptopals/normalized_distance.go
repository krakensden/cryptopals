package libcryptopals

//import "fmt"

// For some reason doing no averaging- just comparing the distance of the first chunks- is
// more accurate/useful than averaging all of them
func AverageNormalizedDistance(chunk_size int, chunks_to_test int, text []byte) (float64, error) {
	var normalized_distance float64 = 0
	for index := 0; index < chunks_to_test; index++ {
		starting := chunk_size * index
		distance, err := HammingDistance(text[starting:starting+chunk_size], text[starting+chunk_size:starting+chunk_size*2])
		if err != nil {
			return 0, err
		}
		normalized_distance = normalized_distance + float64(distance)/float64(chunk_size)
	}
	return normalized_distance / float64(chunks_to_test), nil
}
