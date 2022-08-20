package main

import "fmt"

func main() {
	fmt.Println("Input: nums = [1,1,2]")
	nums := []int{1, 1, 2}
	fmt.Printf("Output: %v, nums = %v\n\n", removeDuplicates(nums), nums)

	fmt.Println("Input: nums = [0,0,1,1,1,2,2,3,3,4]")
	nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Printf("Output: %v, nums = %v\n\n", removeDuplicates(nums), nums)
}

func removeDuplicates(nums []int) int {
	pointer := 0

	for idx := 1; idx < len(nums); idx++ {
		if nums[pointer] != nums[idx] {
			pointer++
			nums[pointer] = nums[idx]
		}
	}

	return pointer + 1
}
