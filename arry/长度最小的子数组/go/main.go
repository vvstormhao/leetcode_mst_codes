package main

import "fmt"

func minSubArry(nums []int, sum int) int {
	start, end := 0, 0
	curSum := nums[start]
	minCount := 0

	for end < len(nums) {
		if nums[start] >= sum {
			minCount = 1
			return minCount
		}

		if curSum < sum {
			end++
			if end < len(nums) {
				curSum += nums[end]
			}
		} else {
			minCount = end - start + 1
			curSum -= nums[start]
			start++
		}
	}

	return minCount
}

func main() {
	nums := []int{7, 1, 1, 1, 1, 1, 1, 1}
	c := minSubArry(nums, 4)
	fmt.Println(c)
}
