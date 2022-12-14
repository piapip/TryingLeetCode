package main

import (
	"fmt"
	"strings"
)

// https://leetcode.com/problems/roman-to-integer/

var normalTerm = []string{"I", "V", "X", "L", "C", "D", "M"}
var specialTerm = []string{"IV", "IX", "XL", "XC", "CD", "CM"}

var termToValue = map[string]int{
	"I":  1,
	"V":  5,
	"X":  10,
	"L":  50,
	"C":  100,
	"D":  500,
	"M":  1000,
	"IV": 4,
	"IX": 9,
	"XL": 40,
	"XC": 90,
	"CD": 400,
	"CM": 900,
}

func romanToInt(s string) int {
	result := 0
	for _, term := range specialTerm {
		count := strings.Count(s, term)

		result = result + count*termToValue[term]
		s = strings.Replace(s, term, "", -1)
	}

	for _, term := range normalTerm {
		count := strings.Count(s, term)

		result = result + count*termToValue[term]
	}

	return result
}

func test(input string, expectedOutput int) {
	output := romanToInt(input)

	if output != expectedOutput {
		fmt.Printf("%s: Failed! Output: %d, Expected output: %d\n", input, output, expectedOutput)
		return
	}

	fmt.Printf("%s: Passed!\n", input)
}

func main() {
	test("III", 3)
	test("LVIII", 58)
	test("MCMXCIV", 1994)
}
