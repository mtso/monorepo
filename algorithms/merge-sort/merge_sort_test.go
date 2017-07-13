package main

import "testing"

func BenchmarkMergeSort(b *testing.B) {
	array := []int{4, 5, 1, 3, 8, 10, 1}
	for i := 0; i < b.N; i+=1 {
		merge_sort(array, func(a, b int) bool { return b-a > 0 })
		merge_sort(array, func(a, b int) bool { return a-b > 0 })
	}
}
