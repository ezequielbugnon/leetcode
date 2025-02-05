package leetcode

import "fmt"

func ProductExceptSelf(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	prefix := 1
	suffix := 1

	// Primer paso: Producto de prefijos
	for i := 0; i < n; i++ {
		result[i] = prefix
		prefix *= nums[i]
		fmt.Printf("Prefijo paso %d: result = %v, prefix = %d\n", i, result, prefix)
	}

	// Segundo paso: Producto de sufijos
	for i := n - 1; i >= 0; i-- {
		result[i] *= suffix
		suffix *= nums[i]
		fmt.Printf("Sufijo paso %d: result = %v, suffix = %d\n", i, result, suffix)
	}

	return result
}
