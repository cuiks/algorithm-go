package main

func combine(n int, k int) [][]int {
	result := make([][]int, 0)
	list := make([]int, 0)
	backtrack(1, n+1, k, list, &result)
	return result
}

func backtrack(start, end, k int, list []int, result *[][]int) {
	if len(list) == k {
		tmp := make([]int, len(list))
		copy(tmp, list)
		*result = append(*result, tmp)
		return
	}
	for i := start; i < end; i++ {
		list = append(list, i)
		backtrack(i+1, end, k, list, result)
		list = list[:len(list)-1]
	}
}

func main() {

}
