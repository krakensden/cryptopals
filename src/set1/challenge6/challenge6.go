package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"libcryptopals"
	"os"
)

func SlurpB64EncodedFile(filename string) ([]byte, error) {
	// ignores newlines
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer func() {
		f.Close()
	}()

	reader := bufio.NewReader(f)
	var wholebuf []byte

	for {
		buf, _, err := reader.ReadLine()
		wholebuf = append(wholebuf, buf...)
		if err == io.EOF {
			break // Done
		}
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}
	return libcryptopals.Base642Byte(wholebuf)
}

func main() {
	var path *string = flag.String("target", "src/set1/challenge6/6.txt", "path to b64 encoded & encrypted file")
	flag.Parse()
	fmt.Println("Path: ", *path)
	input, err := SlurpB64EncodedFile(*path)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(libcryptopals.DecodeXorEncodedBlock(input)))
}
