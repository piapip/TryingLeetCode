package main

import (
	"fmt"

	"github.com/piapip/leet_code/pkg"
)

// https://leetcode.com/problems/palindrome-linked-list/

// double pointer
func isPalindrome(head *pkg.ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	singleJump, doubleJump := head, head
	var reversedFirstHalf, tempNode *pkg.ListNode

	// we'll only reverse the first half of the link listed
	for doubleJump != nil && doubleJump.Next != nil {
		doubleJump = doubleJump.Next.Next

		// revert the order of the first half
		// from        1 -> 2 -> 3
		// to   nil <- 1 <- 2 <- ...
		reversedFirstHalf = singleJump
		singleJump = singleJump.Next
		// connect the 1 to the nil, also revert the direction
		reversedFirstHalf.Next = tempNode
		tempNode = reversedFirstHalf
	}

	secondHalf := singleJump
	// if it's an odd number then the second half shoud start on the Next
	if doubleJump != nil {
		secondHalf = singleJump.Next
	}

	for reversedFirstHalf != nil && secondHalf != nil {
		if reversedFirstHalf.Val != secondHalf.Val {
			return false
		}

		reversedFirstHalf = reversedFirstHalf.Next
		secondHalf = secondHalf.Next
	}

	return true
}

func test(input []int, expectedOutput bool) {
	linkedList := pkg.SetupNode(input)
	output := isPalindrome(linkedList)

	if output != expectedOutput {
		fmt.Printf("%+v: Failed! Output: %t, Expected output: %t\n", input, output, expectedOutput)
		return
	}

	fmt.Printf("%+v: Passed!\n", input)
}

func main() {
	test([]int{1, 2}, false)
	test([]int{1, 2, 2, 1}, true)
	test([]int{1, 2, 3, 2, 1}, true)
	test([]int{1, 2, 3, 3, 2, 1}, true)
	test([]int{1, 2, 4, 3, 2, 1}, false)
}
