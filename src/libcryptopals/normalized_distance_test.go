package libcryptopals

import "testing"

//NormalizedDistance(chunk_size int, chunks_to_test int, text []byte) (int, error) {

func TestDoubleSimpleNormalizedDistance(t *testing.T) {
	var tf = func(real_distance int, input []byte) {
		distance, err := NormalizedDistance(1, 3, input)

		if err != nil {
			t.Errorf("Errors should be impossible", err)
		}

		if distance != real_distance {
			t.Errorf("Expected %d got %d with input %q", real_distance, distance, input)
		}
	}
	tf(0, []byte{1, 1, 1, 1})
	tf(2, []byte{5, 6, 7, 8})
}

func TestSimpleNormalizedDistance(t *testing.T) {
	distance, err := NormalizedDistance(1, 1, []byte{1, 1})

	if err != nil {
		t.Errorf("Errors should be impossible", err)
	}

	if distance != 0 {
		t.Errorf("simplest possible case of difference failed")
	}
}
