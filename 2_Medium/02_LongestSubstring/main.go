package main

import (
	"context"
	"fmt"
	"time"

	"github.com/piapip/leet_code/pkg"
)

// https://leetcode.com/problems/longest-substring-without-repeating-characters/

// sliding window
func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}

	startPointer := 0

	characterMap := make(map[rune]int)

	max := 0

	for i, character := range s {
		if _, ok := characterMap[character]; !ok {
			characterMap[character] = i
		} else {
			startPointer = pkg.CheckMax(characterMap[character]+1, startPointer)
			// fmt.Println("new start: ", startPointer)
			characterMap[character] = i
		}

		max = pkg.CheckMax(i-startPointer+1, max)
	}

	return max
}

func test(input string, expectedOutput int) bool {
	output := lengthOfLongestSubstring(input)

	if output != expectedOutput {
		fmt.Printf("%s: Failed! Output: %d, Expected output: %d\n", input, output, expectedOutput)
		return false
	}

	fmt.Printf("%s: Passed!\n", input)
	return true
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	results := make(chan bool)

	type args struct {
		input string
		want  int
	}

	testCases := []args{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"dvdf", 3},
		{"anviaj", 5},
	}
	testCount := 0

	go func() {
		for _, ts := range testCases {
			results <- test(ts.input, ts.want)
		}
	}()

	// do this to prevent infinite loop from breaking my laptop
	for {
		select {
		case <-results:
			testCount = testCount + 1
			if testCount == len(testCases) {
				fmt.Println("Done testing")
				return
			}
		case <-ctx.Done():
			fmt.Println("Time's up!!!")
			return
		}
	}
}
