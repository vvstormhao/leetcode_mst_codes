package main

import "fmt"

func square(nums []int) []int {
	i, j := 0, len(nums)-1
	ret := make([]int, len(nums))
	k := j

	for i <= j {
		if nums[i]*nums[i] > nums[j]*nums[j] {
			ret[k] = nums[i] * nums[i]
			i++
			k--
		} else {
			ret[k] = nums[j] * nums[j]
			j--
			k--
		}
		fmt.Println(ret)
	}

	return ret
}

func main() {
	nums := []int{-7, -3, 2, 3, 11}
	ret := square(nums)
	fmt.Println(ret)
}
