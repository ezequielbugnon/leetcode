package sort

import (
	"fmt"
)

func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := arr[len(arr)/2]
	var left []int
	var right []int
	var result []int
	fmt.Println(mid)

	for _, v := range arr {

		if v < mid {
			left = append(left, v)
		} else if v > mid {
			right = append(right, v)
		}

	}

	left = QuickSort(left)
	right = QuickSort(right)

	result = append(result, left...)
	result = append(result, mid)
	result = append(result, right...)

	return result
}
