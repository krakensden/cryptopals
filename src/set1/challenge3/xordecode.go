package main

import (
	"libcryptopals"
	"fmt"
)

func main() {
	encoded, score, err := libcryptopals.SingleBitBruteForce("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	if err != nil {
		fmt.Println("Uh... the input is broken?")
		return
	}
	fmt.Println("With a score of", score, ",", encoded)
}
