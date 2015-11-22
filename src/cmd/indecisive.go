package main

import "../libcryptopals"
import "fmt"
import "flag"

// Intended to test enciphering/deciphering independent of the challenges, to
// make it easier to debug things.

func main() {
	var key *string = flag.String("key", "foo", "key to encrypt with")
	var text *string = flag.String("text", "bar", "text to encipher")
	flag.Parse()

	fmt.Println(*text)
	out, err := libcryptopals.StringRepeatingXor(*text, *key)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(out)

	var input []byte = []byte(out)
	libcryptopals.DecodeXorEncodedBlock(input)
}
