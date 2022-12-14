package main

import (
	"context"
	"fmt"
	"time"

	"github.com/piapip/leet_code/pkg"
)

// https://leetcode.com/problems/add-two-numbers/solutions/?languageTags=golang
func addTwoNumbers(l1 *pkg.ListNode, l2 *pkg.ListNode) *pkg.ListNode {
	next1 := l1
	next2 := l2

	result := &pkg.ListNode{}
	current := result

	carry := 0

	for next1 != nil || next2 != nil {
		val := 0

		if next1 != nil {
			val = val + next1.Val
			next1 = next1.Next
		}

		if next2 != nil {
			val = val + next2.Val
			next2 = next2.Next
		}

		val = val + carry
		if val >= 10 {
			carry = 1
			val = val - 10
		} else {
			carry = 0
		}

		current.Next = &pkg.ListNode{
			Val:  val,
			Next: nil,
		}
		current = current.Next
	}

	if carry == 1 {
		current.Next = &pkg.ListNode{
			Val:  1,
			Next: nil,
		}
	}

	return result.Next
}

func test(input1 []int, input2 []int, expectedOutput []int) bool {
	linkedList1 := pkg.SetupNode(input1)
	linkedList2 := pkg.SetupNode(input2)
	output := addTwoNumbers(linkedList1, linkedList2)

	extractedOutput := pkg.ExtractValue(output)

	if !pkg.CompareIntSlice(extractedOutput, expectedOutput) {
		fmt.Printf("Failed! Output: %+v, Expected output: %+v\n", extractedOutput, expectedOutput)
		return false
	}

	fmt.Printf("Passed!\n")
	return true
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	results := make(chan bool)

	type args struct {
		input1 []int
		input2 []int
		want   []int
	}

	testCases := []args{
		{[]int{2, 4, 3}, []int{5, 6, 4}, []int{7, 0, 8}},
		{[]int{0}, []int{0}, []int{0}},
		{[]int{9, 9, 9, 9, 9, 9, 9}, []int{9, 9, 9, 9}, []int{8, 9, 9, 9, 0, 0, 0, 1}},
		{[]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, []int{5, 6, 4}, []int{6, 6, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}},
	}

	testCount := 0

	go func() {
		for _, ts := range testCases {
			results <- test(ts.input1, ts.input2, ts.want)
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
