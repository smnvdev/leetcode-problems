package main

import "fmt"

func main() {
	fmt.Println("Input: a = \"11\", b = \"1\"")
	fmt.Printf("Output: %v\n\n", addBinary("11", "1"))

	fmt.Println("Input: a = \"1010\", b = \"1011\"")
	fmt.Printf("Output: %v\n\n", addBinary("1010", "1011"))
}

func addBinary(a string, b string) string {
	aIdx, bIdx := len(a)-1, len(b)-1

	var result string
	sum, overload := uint8(0), uint8(0)
	for aIdx >= 0 || bIdx >= 0 || overload > 0 {
		sum = overload

		if aIdx >= 0 {
			sum += a[aIdx] - '0'
			aIdx--
		}
		if bIdx >= 0 {
			sum += b[bIdx] - '0'
			bIdx--
		}

		overload = sum / 2
		result = string(sum%2+'0') + result
	}

	return result
}
