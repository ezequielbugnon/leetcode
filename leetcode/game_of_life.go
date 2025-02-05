package leetcode

func GameOfLife(board [][]int) {
	m, n := len(board), len(board[0])

	// Direcciones de los vecinos (8 direcciones)
	directions := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	// Función para contar los vecinos vivos
	countLiveNeighbors := func(x, y int) int {
		count := 0
		for _, dir := range directions {
			nx, ny := x+dir[0], y+dir[1]
			if nx >= 0 && nx < m && ny >= 0 && ny < n {
				if board[nx][ny] == 1 || board[nx][ny] == -1 { // Contamos las celdas vivas o muertas temporalmente
					count++
				}
			}
		}
		return count
	}

	// Fase 1: Aplicar las reglas del Juego de la Vida sin afectar la matriz original
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			liveNeighbors := countLiveNeighbors(i, j)

			if board[i][j] == 1 { // Celda viva
				if liveNeighbors < 2 || liveNeighbors > 3 {
					board[i][j] = -1 // Se muere (marcado temporalmente como -1)
				}
			} else if board[i][j] == 0 { // Celda muerta
				if liveNeighbors == 3 {
					board[i][j] = 2 // Revive (marcado temporalmente como 2)
				}
			}
		}
	}

	// Fase 2: Actualizar los valores finales
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == -1 {
				board[i][j] = 0 // Se convierte en muerta
			} else if board[i][j] == 2 {
				board[i][j] = 1 // Se convierte en viva
			}
		}
	}
}
