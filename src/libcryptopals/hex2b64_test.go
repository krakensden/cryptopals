package libcryptopals

import (
	"fmt"
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
			fmt.Println(len(output), " ", len(result), " ", index, " ", value)
			if result[index] != value {
				t.Errorf("Expected %d, got %d for input %s", output, result, input)
			}
		}
	}
}

func TestBase642Byte(t *testing.T) {
	valid_inputs := map[string][]byte{
		"MA==":     []byte{'0'},
		"AA==":     []byte{0},
		"aA==":     []byte{0x68},
		"ZA==":     []byte{'d'},
		"ZQ==":     []byte{'e'},
		"bA==":     []byte{0x6c},
		"bw==":     []byte{0x6f},
		"AAAAAA==": []byte{0x0, 0x0, 0x0, 0x0},
		"MDAwMA==": []byte{0x30, 0x30, 0x30, 0x30},
		"ICAgIA==": []byte{32, 32, 32, 32},
	}
	for input, output := range valid_inputs {
		fmt.Println(input, output)
		input := []byte(input)
		result, err := Base642Byte(input)
		if err != nil {
			t.Errorf("Did not expect an error (%s) for %v", err.Error(), input)
			return
		}
		if len(result) != len(output) {
			t.Errorf("Result (%v) length is %d, which doesn't match the expected %v with length %d", result, len(result), output, len(output))
			return
		}
		for i := range output {
			if result[i] != output[i] {
				t.Errorf("Expected %v, got %v", output, result)
				return
			}
		}
	}
}
