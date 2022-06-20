package main

import "fmt"

func main() {
	fmt.Println("Example 1:\nInput: x = 121")
	fmt.Printf("Output: %v\n\n", isPalindrome(121))

	fmt.Println("Example 2:\nInput: x = -121")
	fmt.Printf("Output: %v\n\n", isPalindrome(121))

	fmt.Println("Example 3:\nInput: x = 10")
	fmt.Printf("Output: %v\n\n", isPalindrome(10))
}

func isPalindrome(x int) bool {
	number, reverse, lastNumber := x, 0, 0

	for number > 0 {
		lastNumber = number % 10
		reverse = reverse*10 + lastNumber
		number /= 10
	}

	return reverse == x
}
