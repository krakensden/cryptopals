package main

import (
	"./xordecode"
	"fmt"
)

func main() {
	encoded, err := xordecode.SingleBitBruteForce("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	if err != nil {
		fmt.Println("Uh... the input is broken?")
		return
	}
	fmt.Println(encoded)
}
