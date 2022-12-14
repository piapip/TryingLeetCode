package pkg

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func SetupNode(input []int) *ListNode {
	if len(input) == 0 {
		return nil
	}

	head := &ListNode{
		Val:  input[0],
		Next: nil,
	}

	current := head

	for index, value := range input {
		if index == 0 {
			continue
		}

		newNode := &ListNode{
			Val:  value,
			Next: nil,
		}

		current.Next = newNode

		current = newNode
	}

	return head
}

func ExtractValue(head *ListNode) []int {
	if head == nil {
		return nil
	}

	result := make([]int, 0)
	nextNode := head

	for nextNode != nil {
		result = append(result, nextNode.Val)
		nextNode = nextNode.Next
	}

	return result
}

func TraverseNode(head *ListNode) {
	if head == nil {
		fmt.Println("Empty")
		return
	}

	fmt.Printf("%d", head.Val)

	nextNode := head.Next

	for nextNode != nil {
		fmt.Printf(" %d", nextNode.Val)
		nextNode = nextNode.Next
	}

	fmt.Println()
}

func convertNumber(head *ListNode) int {
	result := 0

	exponential := 1

	next := head

	for next != nil {
		result = result + exponential*next.Val
		exponential = exponential * 10
		next = next.Next
	}

	return result
}

func convertToLinkedList(input int) *ListNode {
	if input == 0 {
		return &ListNode{
			Val:  0,
			Next: nil,
		}
	}

	head := &ListNode{}
	current := head

	for {
		val := input % 10
		input = (input - val) / 10

		next := &ListNode{
			Val:  val,
			Next: nil,
		}

		current.Next = next
		current = current.Next

		if input == 0 {
			return head.Next
		}
	}
}
