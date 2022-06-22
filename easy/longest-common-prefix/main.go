package main

import "fmt"

func main() {
	fmt.Println("Input: strs = [\"flower\",\"flow\",\"flight\"]")
	fmt.Printf("Output: \"%v\"\n\n", longestCommonPrefix([]string{"flower", "flow", "flight"}))

	fmt.Println("Input: strs = [\"dog\",\"racecar\",\"car\"]")
	fmt.Printf("Output: \"%v\"\n\n", longestCommonPrefix([]string{"dog", "racecar", "car"}))
}

func longestCommonPrefix(strings []string) string {
	prefix := strings[0]

	for _, s := range strings {
		if len(s) < len(prefix) {
			prefix = prefix[:len(s)]
		}

		for prefix != s[:len(prefix)] && len(prefix) > 0 {
			prefix = prefix[:len(prefix)-1]
		}

		if len(prefix) == 0 {
			return ""
		}
	}

	return prefix
}
