package main

func combinationSum(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	list := make([]int, 0)
	backtrack(candidates, target, 0, list, &result)
	return result
}

func backtrack(candidates []int, target, index int, list []int, result *[][]int) {
	if target == 0 {
		tmp := make([]int, len(list))
		copy(tmp, list)
		*result = append(*result, tmp)
		return
	}
	if target < 0 {
		return
	}
	for i := index; i < len(candidates); i++ {
		list = append(list, candidates[i])
		backtrack(candidates, target-candidates[i], i, list, result)
		list = list[:len(list)-1]
	}
}

func main() {

}
