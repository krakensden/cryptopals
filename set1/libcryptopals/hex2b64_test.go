package libcryptopals

import (
	"testing"
)

func TestHex2Base64(t *testing.T) {
	valid_inputs := map[string]string{
		"30":       "MA==", // character 0
		"0":        "AA==", // literal 0
		"68":       "aA==",
		"65":       "ZQ==",
		"6c":       "bA==",
		"6f":       "bw==",
		"00000000": "AAAAAA==",
		"30303030": "MDAwMA==",
		"20202020": "ICAgIA==",
	}
	for input, output := range valid_inputs {
		result, err := Hex2Base64(input)
		if err != nil {
			t.Errorf("Did not expect an error (%s) for %s", err.Error(), input)
		}
		if result != output {
			t.Errorf("Expected %s, got %s", output, result)
		}
	}
}

func TestHex2Byte(t *testing.T) {
	valid_inputs := map[string][]byte{
		"0":        []byte{0},
		"ff":       []byte{255},
		"4":        []byte{4},
		"40":       []byte{64},
		"A":        []byte{10},
		"20":       []byte{32},
		"204":      []byte{32, 4},
		"20202020": []byte{32, 32, 32, 32},
	}
	for input, output := range valid_inputs {
		result, err := Hex2Byte(input)
		if err != nil {
			t.Errorf("Did not expect an error (%s) for %s", err.Error(), input)
		}
		if len(result) != len(output) {
			t.Errorf("Expected %d, got %d for input %s", output, result, input)
		}
		for index, value := range output {
			if result[index] != value {
				t.Errorf("Expected %d, got %d for input %s", output, result, input)
			}
		}
	}
}
