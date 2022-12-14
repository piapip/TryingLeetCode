package main

import (
	"fmt"

	"github.com/piapip/leet_code/pkg"
)

// https://leetcode.com/problems/middle-of-the-linked-list/
func middleNode(head *pkg.ListNode) *pkg.ListNode {
	singleJump, doubleJump := head, head

	for doubleJump != nil && doubleJump.Next != nil {
		doubleJump = doubleJump.Next.Next

		singleJump = singleJump.Next
	}

	return singleJump
}

func test(input []int, expectedOutput int) {
	linkedList := pkg.SetupNode(input)
	output := middleNode(linkedList)

	if output == nil {
		fmt.Printf("%+v: ??? Expected output: %d\n", input, expectedOutput)
		return
	}

	if output.Val != expectedOutput {
		fmt.Printf("%+v: Failed! Output: %d, Expected output: %d\n", input, output.Val, expectedOutput)
		return
	}

	fmt.Printf("%+v: Passed!\n", input)
}

func main() {
	test([]int{1, 2}, 2)
	test([]int{1, 2, 3, 1}, 3)
	test([]int{1, 2, 3, 2, 1}, 3)
	test([]int{1, 2, 3, 3, 2, 1}, 3)
	test([]int{1, 2, 3, 4, 2, 1}, 4)
}
