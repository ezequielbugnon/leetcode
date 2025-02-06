package leetcode

func MaxSubArray(nums []int) int {
	maxSum := 0
	currentSum := 0

	for _, num := range nums {
		currentSum = max(num, currentSum+num)
		maxSum = max(maxSum, currentSum)
	}

	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
