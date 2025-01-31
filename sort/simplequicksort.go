package sort

func SimpleQuickSort(arr []int, low, high int) {

	if low < high {
		pivot := partitionSimple(arr, low, high)
		SimpleQuickSort(arr, low, pivot-1)
		SimpleQuickSort(arr, pivot+1, high)
	}
}

func partitionSimple(arr []int, low, high int) int {
	index := low - 1

	pivot := arr[high]

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			index++
			arr[j], arr[index] = arr[index], arr[j]
		}
	}

	arr[high], arr[index+1] = arr[index+1], arr[high]

	return index + 1
}
