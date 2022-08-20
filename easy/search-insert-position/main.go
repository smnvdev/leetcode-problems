package main

import "fmt"

func main() {
	fmt.Println("Input: nums = [1,3,5,6], target = 5")
	fmt.Printf("Output: %v\n\n", searchInsert([]int{1, 3, 5, 6}, 5))

	fmt.Println("Input: nums = [1,3,5,6], target = 2")
	fmt.Printf("Output: %v\n\n", searchInsert([]int{1, 3, 5, 6}, 2))

	fmt.Println("Input: nums = [1,3,5,6], target = 7")
	fmt.Printf("Output: %v\n\n", searchInsert([]int{1, 3, 5, 6}, 7))
}

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for right >= left {
		middle := left + (right-left)/2

		if nums[middle] == target {
			return middle
		}

		if nums[middle] > target {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}

	return left
}
