package libcryptopals

import "testing"

func TestSingleBitBruteForce(t *testing.T) {
	encoded, _, err := SingleBitBruteForce("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	if err != nil {
		t.Errorf("Uh... the input is broken?")
	}
	if encoded != "With a score of 23 , Cooking MC's like a pound of bacon" {
		t.Errorf("decoding is broken")
	}
}
