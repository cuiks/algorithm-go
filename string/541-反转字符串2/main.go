package main

func reverseStr(s string, k int) string {
	i := 0
	sByte := []byte(s)
	for i < len(s) {
		reverse(sByte, i, i+k-1)
		i += 2 * k
	}
	return string(sByte)
}

func reverse(s []byte, i, j int) {
	for j > len(s)-1 {
		j--
	}
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}

func main() {

}
