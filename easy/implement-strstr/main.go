package main

import "fmt"

func main() {
	fmt.Println("Input: haystack = \"hello\", needle = \"ll\"")
	fmt.Printf("Output: %v\n\n", strStr("hello", "ll"))

	fmt.Println("Input: haystack = \"aaaaa\", needle = \"bba\"")
	fmt.Printf("Output: %v", strStr("aaaaa", "bba"))
}

func strStr(haystack string, needle string) int {
	if haystack == needle {
		return 0
	}
	if len(haystack) < len(needle) {
		return -1
	}

	for i := 0; i < len(haystack)-(len(needle)-1); i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}
