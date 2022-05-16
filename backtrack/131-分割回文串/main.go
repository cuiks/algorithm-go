package main

func partition(s string) [][]string {
	result := make([][]string, 0)
	list := make([]string, 0)
	backtrack(s, 0, 0, list, &result)
	return result
}

func backtrack(s string, index int, end int, list []string, result *[][]string) {
	if end == len(s) {
		tmp := make([]string, len(list))
		copy(tmp, list)
		*result = append(*result, tmp)
	}
	if end >= len(s) {
		return
	}
	for i := index; i < len(s); i++ {
		if isPal(s[index : i+1]) {
			list = append(list, s[index:i+1])
		} else {
			continue
		}
		backtrack(s, i+1, i+1, list, result)
		list = list[:len(list)-1]
	}
}

func isPal(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {

}
