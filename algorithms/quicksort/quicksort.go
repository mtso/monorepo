package main

import "fmt"

func partition(nums []int, start, end int) int {
	pivot := nums[end]
	part := start

	for i := start; i < end; i++ {
		if nums[i] < pivot {
			temp := nums[i]
			nums[i] = nums[part]
			nums[part] = temp
			part++
		}
	}

	nums[end] = nums[part]
	nums[part] = pivot

	return part
}

func Quicksort(nums []int, start, end int) {
	if start >= end {
		return
	}
	part := partition(nums, start, end)
	Quicksort(nums, start, part-1)
	Quicksort(nums, part, end)
}

func main() {
	nums := []int{3, -1, 1, 0, 5, 9, -4, 7}
	Quicksort(nums, 0, len(nums)-1)
	fmt.Println(nums)

	// part := partition(nums, 0, len(nums)-1)
	// fmt.Println(part, nums)
	// part1 := partition(nums, 0, part)
	// fmt.Println(part1, nums)
	// part2 := partition(nums, part, len(nums)-1)
	// fmt.Println(part2, nums)
}
