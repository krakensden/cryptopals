package main

import "libcryptopals"
import "fmt"
import "flag"

// Intended to test enciphering/deciphering independent of the challenges, to
// make it easier to debug things.

func main() {
	var key *string = flag.String("key", "foo", "key to encrypt with")
	var text *string = flag.String("text", "barbazbang", "text to encipher")
	var score *bool = flag.Bool("score", false, "just print out score of text block")
	flag.Parse()

	if *score {
		fmt.Println("Score of blob is ", libcryptopals.EScore([]byte(*text)))
		return
	}
	fmt.Println(*key)
	fmt.Println(*text)
	outhex, err := libcryptopals.StringRepeatingXor(*text, *key)
	if err != nil {
		fmt.Println(err)
		return
	}
	outbytes, err := libcryptopals.Hex2Byte(outhex)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Base64 encrypted text:\n%s\n", libcryptopals.Byte2Base64(outbytes))
	libcryptopals.DecodeXorEncodedBlock(outbytes)
}
