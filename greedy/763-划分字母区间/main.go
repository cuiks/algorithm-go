package main

func partitionLabels(s string) []int {
	record := make([]int, 26)
	for i, v := range s {
		record[v-'a'] = i
	}

	start, end := 0, 0
	result := make([]int, 0)
	for i, v := range s {
		if record[v-'a'] > end {
			end = record[v-'a']
		}
		if i == end {
			result = append(result, end-start+1)
			start = end + 1
		}
	}
	return result
}

func main() {

}
