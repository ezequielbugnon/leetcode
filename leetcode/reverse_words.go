package leetcode

import "strings"

func ReverseWords(s string) string {
	trim := strings.TrimSpace(s)
	words := strings.Fields(trim)
	start := 0
	end := len(words) - 1

	for start < end {
		words[start], words[end] = words[end], words[start]
		start++
		end--
	}

	return strings.Join(words, " ")
}
