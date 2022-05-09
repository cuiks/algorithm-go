package main

func diStringMatch(s string) []int {
	start, end := 0, len(s)
	result := make([]int, 0)
	for _, key := range s {
		if key == 'I' {
			result = append(result, start)
			start++
		} else {
			result = append(result, end)
			end--
		}
	}
	result = append(result, start)
	return result
}

func main() {

}
