package main

import "fmt"

func main() {
	fmt.Println("Input: s = \"()\"")
	fmt.Printf("Output: %v\n\n", isValid("()"))

	fmt.Println("Input: s = \"()[]{}\"")
	fmt.Printf("Output: %v\n\n", isValid("()[]{}"))

	fmt.Println("Input: s = \"(]\"")
	fmt.Printf("Output: %v\n\n", isValid(")"))
}

var bracketMap = map[byte]byte{')': '(', ']': '[', '}': '{'}

func isValid(brackets string) bool {
	var stack []byte

	for idx := 0; idx < len(brackets); idx++ {
		openBracket, ok := bracketMap[brackets[idx]]
		if !ok {
			stack = append(stack, brackets[idx])
			continue
		}

		if len(stack) == 0 || stack[len(stack)-1] != openBracket {
			return false
		}

		stack = stack[:len(stack)-1]
	}

	return len(stack) == 0
}
