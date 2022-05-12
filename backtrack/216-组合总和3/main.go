package main

func combinationSum3(k int, n int) [][]int {
	result := make([][]int, 0)
	list := make([]int, 0)
	backtrack(1, 9, n, k, list, &result)
	return result
}

func backtrack(start, end, n, k int, list []int, result *[][]int) {
	if len(list) == k && n == 0 {
		tmp := make([]int, len(list))
		copy(tmp, list)
		*result = append(*result, tmp)
		return
	}
	for i := start; i <= end; i++ {
		list = append(list, i)
		backtrack(i+1, end, n-i, k, list, result)
		list = list[:len(list)-1]
	}
}

func main() {

}
