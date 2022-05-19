package main

func solveSudoku(board [][]byte) {
	var backtrack func() bool
	backtrack = func() bool {
		for i := 0; i < len(board); i++ {
			for j := 0; j < len(board); j++ {
				if board[i][j] != '.' {
					continue
				}
				for k := '1'; k <= '9'; k++ {
					if isValid(board, i, j, byte(k)) {
						board[i][j] = byte(k)
						if backtrack() {
							return true
						}
						board[i][j] = '.'
					}
				}
				return false
			}
		}
		return true
	}
	backtrack()
}

func isValid(board [][]byte, row, col int, k byte) bool {
	for i := 0; i < len(board); i++ {
		if board[row][i] == k {
			return false
		}
	}
	for i := 0; i < len(board); i++ {
		if board[i][col] == k {
			return false
		}
	}
	si, sj := row/3*3, col/3*3
	for i := si; i < si+3; i++ {
		for j := sj; j < sj+3; j++ {
			if board[i][j] == k {
				return false
			}
		}
	}
	return true
}

func main() {

}
