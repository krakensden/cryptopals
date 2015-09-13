package libcryptopals

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type XorFinderAnalysisResult struct {
	Decoded  *string
	Original *string
	Score    int
	Error    error
}

func XorFinderDecodeFile(filename string) *XorFinderAnalysisResult {
	results := make(chan *XorFinderAnalysisResult, 10)

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
		if len(buf) != 60 {
			continue
		}
		go func(tstring string) {
			decoded, score, err := SingleBitBruteForce(tstring)
			result := &XorFinderAnalysisResult{Decoded: &decoded, Score: score, Error: err, Original: &tstring}
			results <- result
		}(string(buf))
		linecount++
	}

	var best_result, cur_result *XorFinderAnalysisResult
	for ; linecount > 0; linecount-- {
		cur_result = <-results
		if best_result == nil || (cur_result.Error == nil && cur_result.Score >= best_result.Score) {
			best_result = cur_result
		}
		if cur_result.Error != nil {
			fmt.Println("Got an error", cur_result.Error, "from", *(cur_result.Original))
		}
	}
	return best_result
}
