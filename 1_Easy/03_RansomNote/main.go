package main

import (
	"fmt"
)

// https://leetcode.com/problems/ransom-note/
func canConstruct(ransomNote string, magazine string) bool {
	mapRequirement := make(map[rune]int)

	for _, character := range magazine {
		mapRequirement[character]++
	}

	for _, character := range ransomNote {
		if mapRequirement[character] == 0 {
			return false
		}

		mapRequirement[character]--
	}

	return true
}

func test(inputNote, inputMagazine string, expectedOutput bool) {
	output := canConstruct(inputNote, inputMagazine)

	if output != expectedOutput {
		fmt.Printf("%s - %s: Failed! Output: %t, Expected output: %t\n", inputNote, inputMagazine, output, expectedOutput)
		return
	}

	fmt.Printf("%s - %s: Passed!\n", inputNote, inputMagazine)
}

func main() {
	fmt.Println("testing: ")
	test("a", "b", false)
	test("aa", "ab", false)
	test("aa", "aab", true)
	test("aab", "baa", true)
}
