package libcryptopals

import "testing"

func TestStringRepeatingXor(t *testing.T) {
	demo := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"

	success := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"

	output, err := StringRepeatingXor(demo, "ICE")
	if err != nil {
		t.Errorf("Error in XOR process ", err)
		return
	}
	if output != success {
		t.Errorf("Decoded the string wrong")
	}
	output_parsed, err := Hex2Byte(output)
	if err != nil {
		t.Error(err)
	}
	output2, err := SliceRepeatingXor(output_parsed, []byte("ICE"))
	if len(output2) != len(demo) {
		t.Errorf("Can't reverse Xor- got %s => %s", demo, output2)
		return
	}
	for index, _ := range output2 {
		if demo[index] != output2[index] {
			t.Errorf("Strings don't match- first character @ %d (%q vs %q)\n%s\n\tvs\n%s", index, demo[index], output2[index], demo, output2)
			return
		}
	}
}

func TestReversibilityWithLongKeys(t *testing.T) {
	plaintext := "Let me not to the marriage of true minds Admit impediments. Love is not love Which alters when it alteration finds, Or bends with the remover to remove: O no! It is an ever-fixèd mark That looks on tempests and is never shaken; It is the star to every wandering bark, Whose worth's unknown, although his height be taken.  Love's not Time's fool, though rosy lips and cheeks Within his bending sickle's compass come; Love alters not with his brief hours and weeks, But bears it out even to the edge of doom.  If this be error and upon me proved, I never writ, nor no man ever loved."
	key := "hailmaryfullofgracequiznozzlepoopypants"

	output, err := StringRepeatingXor(plaintext, key)
	if err != nil {
		t.Errorf("Couldn't encipher the text", err)
	}

	output_parsed, err := Hex2Byte(output)
	if err != nil {
		t.Error(err)
	}
	output2, err := SliceRepeatingXor(output_parsed, []byte(key))
	if len(output2) != len(plaintext) {
		t.Errorf("Can't reverse Xor- got %s => %s", plaintext, output2)
		return
	}
	for index, _ := range output2 {
		if plaintext[index] != output2[index] {
			t.Errorf("Strings don't match- first character @ %d (%q vs %q)\n%s\n\tvs\n%s", index, plaintext[index], output2[index], plaintext, output2)
			return
		}
	}
}
