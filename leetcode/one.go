package leetcode

import "strings"

func mergeAlternately(word1 string, word2 string) string {
	var result strings.Builder
	result.Grow(len(word1) + len(word2))
	p1, p2 := 0, 0

	for p1 < len(word1) && p2 < len(word2) {
		result.WriteByte(word1[p1])
		result.WriteByte(word2[p2])
		p1++
		p2++
	}

	result.WriteString(word1[p1:])
	result.WriteString(word2[p2:])

	return result.String()
}
