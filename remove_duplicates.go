package main

func main() {
	nums := []int{1, 1, 2}
	removeDuplicates(nums)
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	k := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[k] = nums[i] // Move unique element to the front
			k++
		}
	}

	return k
}
