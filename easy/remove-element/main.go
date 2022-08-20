package main

import "fmt"

func main() {
	fmt.Println("Input: nums = [3,2,2,3], val = 3")
	nums := []int{3, 2, 2, 3}
	fmt.Printf("Output: %v, nums = %v\n\n", removeElement(nums, 3), nums)

	fmt.Println("Input: nums = [0,1,2,2,3,0,4,2], val = 2")
	nums = []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Printf("Output: %v, nums = %v\n\n", removeElement(nums, 2), nums)
}

func removeElement(nums []int, val int) int {
	leftPointer, rightPointer := 0, len(nums)

	for leftPointer < rightPointer {
		if nums[leftPointer] == val {
			rightPointer--
			nums[leftPointer] = nums[rightPointer]
		} else {
			leftPointer++
		}
	}

	return rightPointer
}
