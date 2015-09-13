package main

import (
	"fmt"
	"libcryptopals"
)

func main() {
	fmt.Println("go")
	a := "1c0111001f010100061a024b53535009181c"
	b := "686974207468652062756c6c277320657965"
	abyte, _ := libcryptopals.Hex2Byte(a)
	bbyte, _ := libcryptopals.Hex2Byte(b)
	final, _ := libcryptopals.SliceXor(abyte, bbyte)
	desired, _ := libcryptopals.Hex2Byte("746865206b696420646f6e277420706c6179")

	if len(final) != len(desired) {
		fmt.Println("!!! lengths don't match")
	}
	for i, _ := range final {
		if final[i] != desired[i] {
			fmt.Println("output incorrect at offset ", i)
			return
		}
	}
	fmt.Println("A-OK everything works")
}
