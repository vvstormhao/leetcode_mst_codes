package main

import "fmt"

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		middle := left + ((right - left) / 2)
		if nums[middle] == target {
			return middle
		}

		if nums[middle] > target {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return -1
}

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 2

	idx := search(nums, target)
	fmt.Printf("idx is %d\n", idx)
}
