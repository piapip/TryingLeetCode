package main

import (
	"fmt"
	"sort"

	"github.com/piapip/leet_code/pkg"
)

type rowFormation struct {
	index        int
	soldierCount int
}

// https://leetcode.com/problems/the-k-weakest-rows-in-a-matrix/
func kWeakestRows(mat [][]int, k int) []int {
	weakestRows := make([]*rowFormation, 0)
	for i, row := range mat {
		soldierCount := getCivilianPosition(row)

		weakestRows = append(weakestRows, &rowFormation{
			index:        i,
			soldierCount: soldierCount,
		})
	}

	sort.Slice(weakestRows, func(i, j int) bool {
		if weakestRows[i].soldierCount == weakestRows[j].soldierCount {
			return weakestRows[i].index < weakestRows[j].index
		}

		return weakestRows[i].soldierCount < weakestRows[j].soldierCount
	})

	result := make([]int, 0)

	for i, row := range weakestRows {
		result = append(result, row.index)

		if i == k-1 {
			return result
		}
	}

	return result
}

func getCivilianPosition(input []int) int {
	if len(input) == 0 {
		return 0
	}

	start, end := 0, len(input)-1

	if end == start {
		return 0
	}

	// no civilian
	if input[end] == 1 {
		return end + 1
	}

	// no soldier
	if input[start] == 0 {
		return 0
	}

	for end != start {
		middle := (end + start - (end+start)%2) / 2
		if input[middle] == 1 {
			start = middle

			if input[start+1] == 0 {
				return start + 1
			}
		} else {
			end = middle

			if input[end-1] == 1 {
				return end
			}
		}
	}

	return start
}

func testGetCivilianPosition(input []int, expectedOutput int) {
	output := getCivilianPosition(input)

	if output != expectedOutput {
		fmt.Printf("%+v: Failed! Output: %d, Expected output: %d\n", input, output, expectedOutput)
		return
	}

	fmt.Printf("%+v: Passed!\n", input)
}

func test(inputMat [][]int, inputK int, expectedOutput []int) {
	output := kWeakestRows(inputMat, inputK)

	if !pkg.CompareIntSlice(output, expectedOutput) {
		fmt.Printf("Failed! Output: %d, Expected output: %d\n", output, expectedOutput)
		return
	}

	fmt.Printf("Passed!\n")
}

func main() {
	testGetCivilianPosition([]int{0, 0, 0}, 0)
	testGetCivilianPosition([]int{1, 1, 1}, 3)
	testGetCivilianPosition([]int{1, 1, 0}, 2)
	testGetCivilianPosition([]int{1, 1, 1, 0}, 3)
	testGetCivilianPosition([]int{1, 1, 1, 0, 0, 0, 0, 0, 0}, 3)
	testGetCivilianPosition([]int{1, 1, 1, 1, 1, 1, 1, 1, 0}, 8)
	testGetCivilianPosition([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 0}, 9)

	test([][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 1},
	}, 3, []int{2, 0, 3})

	test([][]int{
		{1, 0, 0, 0},
		{1, 1, 1, 1},
		{1, 0, 0, 0},
		{1, 0, 0, 0},
	}, 2, []int{0, 2})
}
