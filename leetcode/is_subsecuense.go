package leetcode

func IsSubsequence(s string, t string) bool {
	left := len(s)
	right := len(t)
	i, j := 0, 0

	for i < left && j < right {
		if s[i] == t[j] {
			i++
		}
		j++
	}

	return i == len(s)
}
