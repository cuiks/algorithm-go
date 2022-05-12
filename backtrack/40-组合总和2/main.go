package main

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	result := make([][]int, 0)
	list := make([]int, 0)
	record := map[int]bool{}
	backtrack(candidates, target, 0, list, record, &result)
	return result
}

func backtrack(candidates []int, target int, index int, list []int, record map[int]bool, result *[][]int) {
	if target == 0 {
		tmp := make([]int, len(list))
		copy(tmp, list)
		*result = append(*result, tmp)
	}
	if target <= 0 {
		return
	}
	for i := index; i < len(candidates); i++ {
		if i > 0 && candidates[i] == candidates[i-1] && !record[candidates[i]] {
			continue
		}
		list = append(list, candidates[i])
		record[candidates[i]] = true
		backtrack(candidates, target-candidates[i], i+1, list, record, result)
		list = list[:len(list)-1]
		record[candidates[i]] = false
	}
}

func main() {

}
