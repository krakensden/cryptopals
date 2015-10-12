package libcryptopals

import "testing"

func TestXorFinder(t *testing.T) {
	best_result := XorFinderDecodeFile("../set1/challenge4/input.txt")
	if best_result == nil {
		t.Errorf("Failed to parse the input")
		return
	}
	if *(best_result.Decoded) != "Now that the party is jumping\n" {
		t.Errorf("Returned the wrong output", *best_result.Decoded)
	}
}
