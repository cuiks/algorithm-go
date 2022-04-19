package main

func twoSum(nums []int, target int) []int {
	record := map[int]int{}
	for i, v := range nums {
		if index, ok := record[target-v]; ok {
			return []int{i, index}
		}
		record[v] = i
	}
	return []int{}
}

func main() {

}
