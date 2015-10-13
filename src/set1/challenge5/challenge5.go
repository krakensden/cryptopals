package main

import (
	"fmt"
	"libcryptopals"
)

func main() {
	demo0 := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"

	success0 := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"

	output0, err := libcryptopals.StringRepeatingXor(demo0, "ICE")
	if err != nil {
		fmt.Println("Error in XOR process ", err)
		return
	}
	successBytes, err := libcryptopals.Hex2Byte(success0)
	if err != nil {
		fmt.Println("Couldn't decode success0")
		return
	}
	output1, err := libcryptopals.SliceRepeatingXor(successBytes, []byte("ICE"))
	if err != nil {
		fmt.Println("Error in XOR process ", err)
		return
	}

	fmt.Println("First line succeeded:", output0 == success0)
	fmt.Println("Output0: ", output0)
	fmt.Println("len(Output0): ", len(output0), " len(dem0): ", len(demo0), " len(success0): ", len(success0))
	fmt.Println("& reverse: ", string(output1))
	fmt.Println("& reverse == success0 ", demo0 == string(output1))
}
