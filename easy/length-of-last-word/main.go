package main

import "fmt"

func main() {
	fmt.Println("Input: s = \"Hello World\"")
	fmt.Printf("Output: \"%v\"\n\n", lengthOfLastWord("Hello World"))

	fmt.Println("Input: s = \"   fly me   to   the moon  \"")
	fmt.Printf("Output: \"%v\"\n\n", lengthOfLastWord("   fly me   to   the moon  "))

	fmt.Println("Input: s = \"luffy is still joyboy\"")
	fmt.Printf("Output: \"%v\"\n\n", lengthOfLastWord("luffy is still joyboy"))
}

func lengthOfLastWord(s string) int {
	length := 0

	for i := len(s) - 1; i >= 0; i -= 1 {
		if s[i] == ' ' {
			if length == 0 {
				continue
			}
			return length
		}
		length += 1
	}

	return length
}
