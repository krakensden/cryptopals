package main

import (
	"fmt"
)

var src string = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
var dest string = "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

func Hex2Base64(input string) (output string) {
	return
}

func main() {
	if actual := Hex2Base64(src); actual != dest {
		panic("still broken")
	}
	fmt.Println(Hex2Base64(src))
}
