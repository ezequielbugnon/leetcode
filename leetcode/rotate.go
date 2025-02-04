package leetcode

func Rotate(matrix [][]int) {
	n := len(matrix)

	for i := 0; i < n; i++ {
		for j := 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[i][n-i-1] = matrix[i][n-i-1], matrix[i][j]
		}
	}
}
