package main

import(
	"fmt"
)

// 上浮
func swim(nums []int, len int, index int) {
	// 父节点 (index - 1) / 2
	if index < 1 || len < index {
		return
	}

	parent := (index - 1) / 2
	if nums[index] > nums[parent] {
		tmp := nums[parent]
		nums[parent] = nums[index]
		nums[index] = tmp
	}
}

func sink(nums []int, len int, index int) {
	// left 2 * index + 1
	// right 2 * index + 2
	var maxIndex int
	left := 2 * index + 1
	right := 2 * index + 2
	if index >= len || left >= len {
		return
	}

	if right >= len {
		maxIndex = left
	} else {
		if nums[left] > nums[right] {
			maxIndex = left
		} else {
			maxIndex = right
		}
	}

	if nums[index] < nums[maxIndex] {
		tmp := nums[maxIndex]
		nums[maxIndex] = nums[index]
		nums[index] = tmp
		sink(nums, len, maxIndex)
	}
}

func heapInit(nums[]int) {
	n := len(nums)
	// 最后的二叉树节点 （n - 1）/2
	for index := (n-1)/2; index >= 0; index-- {
		sink(nums, n, index)
	}
}

func swap(nums []int, first, second int) {
	if first >= len(nums) || second >= len(nums) {
		return
	}

	tmp := nums[first]
	nums[first] = nums[second]
	nums[second] = tmp
}

func heapSort(nums[]int, len int) {
	if len <= 1 {
		return
	}

	swap(nums, 0, len - 1)
	fmt.Printf("after swap %v\n",nums)
	sink(nums, len - 1, 0)
	fmt.Printf("after sink %v\n",nums)
	heapSort(nums, len - 1)
}

func main() {
	data := []int{3,7,1,15,8,4,9,23}
	heapInit(data)
	fmt.Println(data)
	heapSort(data, len(data))
	fmt.Println(data)
}