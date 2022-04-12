package main

func generateMatrix(n int) [][]int {
	result := make([][]int, n)
	for k := 0; k < n; k++ {
		result[k] = make([]int, n)
	}
	i, j := 0, 0
	val := 1
	index := 1 // 记录圈数
	for val <= n*n {
		if val == n*n {
			result[i][j] = val
			val++
		}
		for ; j < n-index; j++ {
			result[i][j] = val
			val++
		}
		for ; i < n-index; i++ {
			result[i][j] = val
			val++
		}
		for ; j >= index; j-- {
			result[i][j] = val
			val++
		}
		for ; i >= index; i-- {
			result[i][j] = val
			val++
		}
		i++
		j++
		index++
	}
	return result
}

func main() {

}
