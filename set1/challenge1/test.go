package main

import (
	"./hex2b64"
	"fmt"
)

var src string = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
var dest string = "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

func main() {
	if actual, err := hex2b64.Hex2Base64(src); actual != dest || err != nil {
		panic(fmt.Sprintf("still broken", actual))
	}
	fmt.Println(hex2b64.Hex2Base64(src))
}
