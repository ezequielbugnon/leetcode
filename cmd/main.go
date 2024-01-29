package main

import (
	"fmt"

	"github.com/ezequielbugnon/leetcode/sort"
)

func main() {
	send := []int{7, 9, 15, 2, 6}
	response := sort.QuickSort(send)

	fmt.Println(response)
}
