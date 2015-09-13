package main

import (
	//	"libcryptopals"
	"fmt"
	"os"
)

func SlurpB64EncodedFile(filename string) ([]byte, err) {
	// ignores newlines
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer func() {
		f.Close()
	}()

	reader := bufio.NewReader(f)
	linecount := 0

	for {
		buf, _, err := reader.ReadLine()
		if err == io.EOF {
			break // Done
		}
		if err != nil {
			fmt.Println(err)
			return nil
		}
	}
}

func main() {
	for KEYSIZE := 2; KEYSIZE < 41; KEYSIZE++ {
		fmt.Println(KEYSIZE)
	}
}
