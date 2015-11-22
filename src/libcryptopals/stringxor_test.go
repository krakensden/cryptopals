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
