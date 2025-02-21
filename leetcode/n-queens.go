package leetcode

import "math"

func SolveNQueens(table []int, row, N int) bool {
	if row == N {
		println(table)
		return true
	}

	for col := 0; col < N; col++ {
		if isValid(table, row, col) {
			table[row] = col
			SolveNQueens(table, row+1, N)
		}
	}

	return false

}

func isValid(table []int, row, col int) bool {
	for i := 0; i < row; i++ {
		if table[i] == col || math.Abs(float64(table[i]-col)) == math.Abs(float64(i-row)) {
			return false
		}
	}
	return true
}
