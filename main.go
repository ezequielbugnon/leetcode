package main

import (
	"fmt"
	"leetcode/leetcode"
	twopointers "leetcode/two-pointers"
)

func main() {
	fmt.Println("Hello world")
	array := []int{2, 7, 10, 15}

	fmt.Println(twopointers.SumArrayTwoPointers(array, 9))

	/*arr := []int{34, 7, 23, 32, 5, 62}
	fmt.Println("Array desordenado:", arr)
	sort.QuickSort(arr, 0, len(arr)-1)
	fmt.Println("Array ordenado:", arr)

	arrTwo := []int{10, 7, 23, 32, 55, 62}
	fmt.Println("Array desordenado simple quicksort:", arrTwo)
	sort.SimpleQuickSort(arrTwo, 0, len(arrTwo)-1)
	fmt.Println("Array ordenado simple quicksort:", arrTwo)

	arrThree := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println("Array deleteduplicate:", arrThree)
	result := twopointers.DeleteDuplicate(arrThree)
	fmt.Println("Array with deleteduplicate:", result)

	container := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println("Array water container result:", container)
	resultContainer := twopointers.WaterContainer(container)
	fmt.Println("Array water container result:", resultContainer)*/

	element := []int{3, 2, 2, 3}
	result := leetcode.RemoveElement(element, 3)
	fmt.Println(" result:", result)

	elementTwo := []int{3, 2, 2, 3}
	sort := MergeSort(elementTwo)
	fmt.Println(" result:", sort)

}

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}

	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}
