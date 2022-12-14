package main

import (
	"fmt"
	"strconv"
)

// https://leetcode.com/problems/fizz-buzz/
func fizzBuzz(n int) []string {
	result := make([]string, n)

	for i := 1; i <= n; i++ {
		item := ""

		if i%3 == 0 {
			item = item + "Fizz"
		}

		if i%5 == 0 {
			item = item + "Buzz"
		}

		if item == "" {
			item = item + strconv.Itoa(i)
		}

		result[i-1] = item
	}

	return result
}

func main() {
	fmt.Println(fizzBuzz(3))
	fmt.Println(fizzBuzz(15))
}
