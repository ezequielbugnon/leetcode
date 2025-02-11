package leetcode

import (
	"fmt"
	"strconv"
)

func Compress(chars []byte) int {
	n := len(chars)
	if n == 0 {
		return 0
	}
	start := 0
	index := 0

	for end := 0; end < n; end++ {
		if end == n-1 || chars[end] != chars[end+1] {
			chars[index] = chars[start]
			index++

			count := end - start + 1
			if count > 1 {
				for _, digit := range strconv.Itoa(count) {
					chars[index] = byte(digit)
					index++
				}
			}

			start = end + 1
		}
	}

	fmt.Println(string(chars))
	fmt.Println(string(chars[:index]))
	return index

}
