package sort

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := len(arr) / 2
	left := MergeSort(arr[:pivot])
	right := MergeSort(arr[pivot:])

	return merge(left, right)
}

func merge(a, b []int) []int {
	indexA, indexB := 0, 0
	lenA, lenB := len(a), len(b)
	var result []int

	for indexA < lenA && indexB < lenB {
		if a[indexA] < b[indexB] {
			result = append(result, a[indexA])
			indexA++
		} else {
			result = append(result, b[indexB])
			indexB++
		}
	}

	result = append(result, a[indexA:]...)
	result = append(result, b[indexB:]...)

	return result

}

func recursiveMerge(a, b []int) []int {
	var result []int

	if len(a) == 0 {
		result = append(result, b...)
		return result
	}
	if len(b) == 0 {
		result = append(result, a...)
		return result
	}

	headA, tailA := a[0], a[1:]
	headB, tailB := b[0], b[1:]

	if headA < headB {
		result = append(result, headA)
		result = append(result, merge(tailA, b)...)
	} else {
		result = append(result, headB)
		result = append(result, merge(a, tailB)...)
	}

	return result
}
