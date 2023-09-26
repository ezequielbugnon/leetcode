package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 3
	removeElementThree(nums, val)
}

func removeElement(nums []int, val int) int {
	numsLenght := len(nums) - 1

	for i := numsLenght; i >= 0; i-- {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
		}
	}

	return len(nums)
}

func removeElementTwo(nums []int, val int) int {
	i := 0
	j := len(nums) - 1

	for i <= j {
		if nums[i] == val {
			nums[i], nums[j] = nums[j], nums[i]
			j--
		} else {
			i++
		}
	}

	fmt.Println(nums)

	return j + 1
}

func removeElementThree(nums []int, val int) int {
	j := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[j], nums[i] = nums[i], nums[j]
			j++
		}
	}
	fmt.Println(nums)
	return j
}
