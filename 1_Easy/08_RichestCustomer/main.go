package main

import "fmt"

// https://leetcode.com/problems/richest-customer-wealth/
func maximumWealth(accounts [][]int) int {
	max := 0

	for _, account := range accounts {
		sum := sumOfArray(account)

		if max < sum {
			max = sum
		}
	}

	return max
}

func sumOfArray(input []int) int {
	result := 0

	for _, item := range input {
		result = result + item
	}

	return result
}

func test(input [][]int, expectedOutput int) {
	output := maximumWealth(input)

	if output != expectedOutput {
		fmt.Printf("Failed! Output: %d, Expected output: %d\n", output, expectedOutput)
		return
	}

	fmt.Printf("Passed!\n")
}

func main() {
	test([][]int{
		{1, 2, 3},
		{3, 2, 1},
	}, 6)

	test([][]int{
		{1, 5},
		{7, 3},
		{3, 5},
	}, 10)
	test([][]int{
		{2, 8, 7},
		{7, 1, 3},
		{1, 9, 5},
	}, 17)
}
