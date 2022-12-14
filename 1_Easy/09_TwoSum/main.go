package main

import (
	"fmt"

	"github.com/piapip/leet_code/pkg"
)

// https://leetcode.com/problems/two-sum/
func twoSum(nums []int, target int) []int {
	// value - index
	list := make(map[int]int)
	for i, num := range nums {
		// complement = target - current
		// if the complement value exists in the list, (that's why we store value as the key)
		// then we'll take that current index and the complement's index.
		if value, ok := list[target-num]; ok {
			return []int{value, i}
		}

		list[num] = i
	}

	return nil
}

func test(inputNums []int, inputTarget int, expectedOutput []int) {
	output := twoSum(inputNums, inputTarget)

	if !pkg.CompareIntSlice(output, expectedOutput) {
		fmt.Printf("%+v: Failed! Output: %+v, Expected output: %+v\n", inputNums, output, expectedOutput)
		return
	}

	fmt.Printf("%+v: Passed!\n", inputNums)
}

func main() {
	test([]int{2, 7, 11, 15}, 9, []int{0, 1})
	test([]int{3, 2, 4}, 6, []int{1, 2})
	test([]int{3, 3}, 6, []int{0, 1})
}
