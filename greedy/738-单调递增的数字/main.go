package main

import "strconv"

func monotoneIncreasingDigits(n int) int {
	s := []byte(strconv.Itoa(n))
	i := 1
	for i < len(s) && s[i-1] <= s[i] {
		i++
	}
	if i < len(s) {
		for i > 0 && s[i] < s[i-1] {
			s[i-1]--
			i--
		}
		for k := i + 1; k < len(s); k++ {
			s[k] = '9'
		}
	}
	result, _ := strconv.Atoi(string(s))
	return result
}

func main() {

}
