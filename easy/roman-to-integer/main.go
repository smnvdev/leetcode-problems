package main

import "fmt"

func main() {
	fmt.Println("Input: s = \"III\"")
	fmt.Printf("Output: %v\n\n", romanToInt("III"))

	fmt.Println("Input: s = \"LVIII\"")
	fmt.Printf("Output: %v\n", romanToInt("LVIII"))

	fmt.Println("Input: s = \"MCMXCIV\"")
	fmt.Printf("Output: %v\n", romanToInt("MCMXCIV"))
}

var romanNumberMap = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

func romanToInt(s string) int {
	result := romanNumberMap[s[len(s)-1]]

	for idx := len(s) - 2; idx >= 0; idx-- {
		if romanNumberMap[s[idx]] < romanNumberMap[s[idx+1]] {
			result -= romanNumberMap[s[idx]]
			continue
		}
		result += romanNumberMap[s[idx]]
	}

	return result
}
