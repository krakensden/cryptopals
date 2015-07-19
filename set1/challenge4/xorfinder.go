package main

import (
	"../challenge3/xordecode"
	"bufio"
	"fmt"
	"io"
	"os"
)

type AnalysisResult struct {
	Decoded  *string
	Original *string
	Score    int
	Error    error
}

func main() {
	results := make(chan *AnalysisResult, 10)

	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
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
			return
		}
		if len(buf) != 60 {
			fmt.Println("Count", len(buf), string(buf))
		}
		go func(tstring string) {
			decoded, score, err := xordecode.SingleBitBruteForce(tstring)
			result := &AnalysisResult{Decoded: &decoded, Score: score, Error: err, Original: &tstring}
			results <- result
		}(string(buf))
		linecount++
	}

	var best_result, cur_result *AnalysisResult
	for ; linecount > 0; linecount-- {
		cur_result = <-results
		if best_result == nil || (cur_result.Error == nil && cur_result.Score > best_result.Score) {
			best_result = cur_result
		}
		if cur_result.Error != nil {
			fmt.Println("Got an error", cur_result.Error, "from", *(cur_result.Original))
		}
	}
	if best_result != nil {
		fmt.Println("The single-byte XOR encoded string is", *(best_result.Decoded), "from", *(best_result.Original), "with a score of", best_result.Score)
	} else {
		fmt.Println("!!! processed nothing?")
	}
}
