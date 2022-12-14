package main

import "fmt"

// https://leetcode.com/problems/number-of-steps-to-reduce-a-number-to-zero/
func numberOfSteps(num int) int {
	count := 0

	for num > 0 {
		num = reduce(num)

		count = count + 1
	}

	return count
}

func reduce(num int) int {
	if num%2 == 0 {
		return num / 2
	}

	return num - 1
}

func test(input int, expectedOutput int) {
	output := numberOfSteps(input)

	if output != expectedOutput {
		fmt.Printf("%d: Failed! Output: %d, Expected output: %d\n", input, output, expectedOutput)
		return
	}

	fmt.Printf("%d: Passed!\n", input)
}

func main() {
	test(123, 12)
	test(14, 6)
	test(8, 4)
}
