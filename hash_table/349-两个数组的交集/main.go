package main

func intersection(nums1 []int, nums2 []int) []int {
	record := map[int]bool{}
	for _, v := range nums1 {
		record[v] = true
	}
	result := make([]int, 0)
	for _, v := range nums2 {
		if record[v] {
			result = append(result, v)
			delete(record, v)
		}
	}
	return result
}

func main() {

}
