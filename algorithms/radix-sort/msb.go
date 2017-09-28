package main

import (
	"fmt"
)

func isGreaterThan(a, b int, place int) bool {
	order := 1
	for i := 0; i < place-1; i++ {
		order *= 10
	}

	a = a / order
	b = b / order

	a = a % 10
	b = b % 10

	return a > b
}

func maxPlace(nums []int) (place int) {
	for _, num := range nums {
		pl := 1
		for {
			if num < 10 {
				break
			}

			num /= 10
			pl++
		}

		if pl > place {
			place = pl
		}
	}
	return
}

func valueAt(num, place int) int {
	order := 1
	for i := 0; i < place-1; i++ {
		order *= 10
	}

	num = num / order
	return num % 10
}

// RadixSort on integers against most significant digit.
func RadixSort(nums []int) []int {
	for i := 1; i <= maxPlace(nums); i++ {
		buckets := make([][]int, 10)
		for i := 0; i < 10; i++ {
			buckets[i] = make([]int, 0)
		}

		for _, num := range nums {
			// bucket := buckets[valueAt(num, i)]
			buckets[valueAt(num, i)] = append(buckets[valueAt(num, i)], num)
		}

		neworder := make([]int, 0)
		for i := 0; i < 10; i++ {
			neworder = append(neworder, buckets[i]...)
		}
		nums = neworder
	}
	return nums
}

func main() {
	fmt.Println(RadixSort([]int{2, 1, 21, 0, 41, 6, 12}))
	fmt.Println(valueAt(123, 6))
	nums := []int{12, 2, 345, 2454, 0}
	fmt.Println(maxPlace(nums))
	fmt.Println(isGreaterThan(111, 22, 4))
	fmt.Println(12 / 10)
	fmt.Println(12 / 100)
}
