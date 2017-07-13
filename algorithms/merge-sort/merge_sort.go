package main

import (
	"fmt"
)

type Comparer func(int, int) bool

func merge(first, second []int, comparer Comparer) []int {
	leng := len(first) + len(second)
	index1 := 0
	index2 := 0

	merged := make([]int, leng)

	for i := 0; i < leng; i += 1 {
		if index1 >= len(first) {
			merged[i] = second[index2]
			index2 += 1
		} else if index2 >= len(second) {
			merged[i] = first[index1]
			index1 += 1
		} else {
			// Use comparer func instead to customize order.
			// first[index1] < second[index2]
			if comparer(first[index1], second[index2]) {
				merged[i] = first[index1]
				index1 += 1
			} else {
				merged[i] = second[index2]
				index2 += 1
			}
		}
	}

	return merged
}

func merge_sort(slice []int, comparer Comparer) []int {
	if len(slice) < 2 {
		return slice
	}

	middle := len(slice) / 2
	a1 := merge_sort(slice[:middle], comparer)
	a2 := merge_sort(slice[middle:], comparer)
	return merge(a1, a2, comparer)
}

func main() {
	arr := []int{5, 7, 2, 3}

	arr = merge_sort(arr, func(a, b int) bool {
		return b-a < 0
	})

	fmt.Println(arr)
}
