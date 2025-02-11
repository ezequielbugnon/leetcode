package main

import (
	"fmt"
	"leetcode/leetcode"
)

func main() {
	/*fmt.Println("Hello world")
	array := []int{2, 7, 10, 15}

	fmt.Println(twopointers.SumArrayTwoPointers(array, 9))

	arr := []int{34, 7, 23, 32, 5, 62}
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
	fmt.Println("Array water container result:", resultContainer)

	element := []int{3, 2, 2, 3}
	result := leetcode.RemoveElement(element, 3)
	fmt.Println(" result:", result)

	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	leetcode.Rotate(matrix)
	fmt.Println(" rotate:", matrix)
	nums := []int{1, 2, 3, 4}
	resultProduct := leetcode.ProductExceptSelf(nums)
	fmt.Println("Resultado final:", resultProduct)

	linkedlist := datastructures.LinkedList{}

	linkedlist.AddNode(2)
	linkedlist.AddNode(3)
	linkedlist.AddNode(7)
	linkedlist.AddNode(20)

	linkedlist.Print()
	linkedlist.Reverse()
	println()

	linkedlist.Print()

	fmt.Println(leetcode.ReverseWords("the sky is blue"))
	nums := []int{1, 2, 3, 4}
	resultProduct := leetcode.ProductExceptSelf(nums)
	fmt.Println("Resultado final:", resultProduct)*/

	test := "aabbbbbbbbbbbbbbbccc"
	bs := []byte(test)
	result := leetcode.Compress(bs)
	fmt.Println("Resultado final:", result)

}
