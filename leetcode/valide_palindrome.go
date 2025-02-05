package leetcode

import (
	"strings"
	"unicode"
)

func IsPalindrome(s string) bool {
	var cleaned strings.Builder

	for _, char := range strings.ToLower(s) {
		if unicode.IsDigit(char) || unicode.IsLetter(char) {
			cleaned.WriteRune(char)
		}
	}

	cleanedStr := cleaned.String()
	for i := 0; i < len(cleanedStr)/2; i++ {
		if cleanedStr[i] != cleanedStr[len(cleanedStr)-1-i] {
			return false
		}
	}
	return true
}
