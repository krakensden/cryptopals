package main

import (
	"./stringxor"
	"fmt"
)

func main() {
	demo0 := "Burning 'em, if you ain't quick and nimble"
	demo1 := "I go crazy when I hear a cymbal"

	success0 := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272"
	success1 := "a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"

	output0, err := stringxor.StringXor(demo0, "ICE")
	if err != nil {
		fmt.Println("Error in XOR process ", err)
		return
	}

	fmt.Println("First line succeeded:", output0 == success0)
	fmt.Println("Output0: ", output0)

	output1, err := stringxor.StringXor(demo1, "ICE")
	if err != nil {
		fmt.Println("Error in XOR process ", err)
		return
	}
	fmt.Println("First line succeeded:", output1 == success1)
	fmt.Println("Output1: ", output1)
}
