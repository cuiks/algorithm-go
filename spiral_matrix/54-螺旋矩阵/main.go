package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	result := make([]int, 0)
	i, j := 0, 0
	index := 1
	for len(result) < m*n {
		if len(result) == m*n-1 {
			result = append(result, matrix[i][j])
		}
		for ; j < n-index; j++ {
			result = append(result, matrix[i][j])
		}
		for ; i < m-index; i++ {
			result = append(result, matrix[i][j])
		}
		for ; j >= index; j-- {
			result = append(result, matrix[i][j])
		}
		for ; i >= index; i-- {
			result = append(result, matrix[i][j])
		}
		i++
		j++
		index++
	}
	return result[:m*n]
}

func main() {
	fmt.Println(spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
}
