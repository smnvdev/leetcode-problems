package main

import "fmt"

func main() {
	fmt.Println("Example 1:\nInput: nums = [2 7 11 15], target = 9")
	fmt.Printf("Output: %v\n\n", twoSum([]int{2, 7, 11, 15}, 9))

	fmt.Println("Example 2:\nInput: nums = [3 2 4], target = 9")
	fmt.Printf("Output: %v\n\n", twoSum([]int{3, 2, 4}, 6))

	fmt.Println("Example 3:\nInput: nums = [3 3], target = 6")
	fmt.Printf("Output: %v\n\n", twoSum([]int{3, 3}, 6))
}

func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int, len(nums))

	for idx, num := range nums {
		if secondIdx, ok := numsMap[target-num]; ok && idx != secondIdx {
			return []int{secondIdx, idx}
		}

		numsMap[num] = idx
	}
	return nil
}
