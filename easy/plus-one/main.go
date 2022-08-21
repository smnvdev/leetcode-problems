package main

import "fmt"

func main() {
	fmt.Println("Input: digits = [1,2,3]")
	fmt.Printf("Output: %v\n\n", plusOne([]int{1, 2, 3}))

	fmt.Println("Input: digits = [4,3,2,1]")
	fmt.Printf("Output: %v\n\n", plusOne([]int{4, 3, 2, 1}))

	fmt.Println("Input: digits = [9]")
	fmt.Printf("Output: %v\n\n", plusOne([]int{9}))
}

func plusOne(digits []int) []int {
	overload := 0
	for idx := len(digits) - 1; idx >= 0; idx -= 1 {
		overload = (digits[idx] + 1) / 10
		digits[idx] = (digits[idx] + 1) % 10

		if overload == 0 {
			return digits
		}
	}
	return append([]int{1}, digits...)
}
