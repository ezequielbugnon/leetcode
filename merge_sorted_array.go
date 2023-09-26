package main

import (
	"sort"
)

/*func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3 //num mergead
	nums2 := []int{2, 5, 6}
	n := 3 // num ignored
	merge(nums1, m, nums2, n)
}*/

/*func merge(nums1 []int, m int, nums2 []int, n int) {
	lenght := m + n
	inside := 0
	for len(nums1) >= lenght && lenght > 0 {
		if nums1[lenght-1] == 0 {
			nums1[lenght-1] = nums2[inside]
			inside++
		}
		lenght--
	}

	sort.Sort(sort.IntSlice(nums1))

	fmt.Println("result", nums1)
}*/

//solution more eficient
func merge(nums1 []int, m int, nums2 []int, n int) {
	p1 := m - 1
	p2 := n - 1

	for i := m + n - 1; i >= 0; i-- {
		if p1 >= 0 && (p2 < 0 || nums1[p1] >= nums2[p2]) {
			nums1[i] = nums1[p1]
			p1--
		} else {
			nums1[i] = nums2[p2]
			p2--
		}
	}
}

func mergeTwo(nums1 []int, m int, nums2 []int, n int) {
	// Copia los primeros 'm' elementos de 'nums1' a un nuevo slice
	result := make([]int, m)
	copy(result, nums1[:m])

	// Fusiona los dos slices en uno nuevo
	result = append(result, nums2[:n]...)

	// Ordena el slice fusionado
	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})

	// Copia los elementos del resultado ordenado de vuelta a 'nums1'
	copy(nums1, result)
}
