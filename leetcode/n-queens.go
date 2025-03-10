package leetcode

import "math"

func SolveNQueens(table []int, row, N int) int {
	if row == N {
		return 1
	}

	count := 0
	for col := 0; col < N; col++ {
		if isValid(table, row, col) {
			table[row] = col
			count += SolveNQueens(table, row+1, N)
		}
	}

	return count

}

func isValid(table []int, row, col int) bool {
	for r := range row {
		if table[r] == col || math.Abs(float64(table[r]-col)) == math.Abs(float64(r-row)) {
			return false
		}
	}
	return true
}
