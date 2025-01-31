package sort

func QuickSort(arr []int, low, high int) {
	if low < high {
		pivot := partition(arr, low, high)
		QuickSort(arr, low, pivot-1)
		QuickSort(arr, pivot+1, high)
	}

}

func partition(arr []int, low, high int) int {
	index := low - 1

	mid := (low + high) / 2

	if arr[low] > arr[mid] {
		arr[low], arr[mid] = arr[mid], arr[low]
	}

	if arr[low] > arr[high] {
		arr[low], arr[high] = arr[high], arr[low]
	}

	arr[high], arr[mid] = arr[mid], arr[high]

	pivot := arr[high]

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			index++
			arr[j], arr[index] = arr[index], arr[j]
		}
	}

	arr[index+1], arr[high] = arr[high], arr[index+1] // trocar pivot para novo pivot

	return index + 1

}
