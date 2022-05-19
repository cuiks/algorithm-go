package main

import (
	"fmt"
	"strings"
)

func solveNQueens(n int) [][]string {
	result := make([][]string, 0)
	board := make([][]string, n)
	for i := 0; i < n; i++ {
		board[i] = make([]string, n)
		for j := 0; j < n; j++ {
			board[i][j] = "."
		}
	}
	backtrack(board, 0, &result)
	return result
}

func backtrack(board [][]string, row int, result *[][]string) {
	if row == len(board) {
		tmp := make([]string, 0)
		for _, v := range board {
			tmp = append(tmp, strings.Join(v, ""))
		}
		*result = append(*result, tmp)
		return
	}

	for i := 0; i < len(board); i++ {
		if isValid(board, row, i) {
			board[row][i] = "Q"
			backtrack(board, row+1, result)
			board[row][i] = "."
		}
	}
}

func isValid(board [][]string, row, col int) bool {
	for i := 0; i < row; i++ {
		if board[i][col] == "Q" {
			return false
		}
	}
	for i, j := row, col; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == "Q" {
			return false
		}
	}
	for i, j := row, col; i >= 0 && j < len(board); i, j = i-1, j+1 {
		if board[i][j] == "Q" {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(solveNQueens(4))
}
