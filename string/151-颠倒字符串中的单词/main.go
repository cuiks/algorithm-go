package main

func reverseWords(s string) string {
	sSlice := make([]string, 0)
	i := 0
	for i < len(s) {
		for i < len(s) && s[i] == ' ' {
			i++
		}
		if i == len(s) {
			break
		}
		j := 0
		for i+j < len(s) && s[i+j] != ' ' {
			j++
		}
		sSlice = append(sSlice, s[i:i+j])
		i += j
	}
	start, end := 0, len(sSlice)-1
	for start < end {
		sSlice[start], sSlice[end] = sSlice[end], sSlice[start]
		start++
		end--
	}
	result := ""
	for _, key := range sSlice {
		result += key + " "
	}
	return result[:len(result)-1]
}

func main() {

}
