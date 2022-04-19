package main

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	result := 0
	record := map[int]int{}
	for _, v1 := range nums1 {
		for _, v2 := range nums2 {
			record[v1+v2]++
		}
	}
	for _, v1 := range nums3 {
		for _, v2 := range nums4 {
			if count, ok := record[0-v1-v2]; ok {
				result += count
			}
		}
	}
	return result
}

func main() {

}
