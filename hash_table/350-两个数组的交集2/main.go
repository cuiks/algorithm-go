package main

func intersect(nums1 []int, nums2 []int) []int {
	record := map[int]int{}
	for _, v := range nums1 {
		record[v]++
	}
	result := make([]int, 0)
	for _, v := range nums2 {
		if record[v] > 0 {
			result = append(result, v)
			record[v]--
		}
	}
	return result
}
func main() {

}
