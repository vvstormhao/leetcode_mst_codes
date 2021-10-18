package main

import "fmt"

func remove(nums []int, target int) int {
	slowIdx := 0
	for fastIdx := 0; fastIdx < len(nums); fastIdx++ {
		if target != nums[fastIdx] {
			nums[slowIdx] = nums[fastIdx]
			slowIdx++
		}
	}

	return slowIdx
}

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	c := remove(nums, 2)
	fmt.Printf("%v\n", nums[:c])
}
